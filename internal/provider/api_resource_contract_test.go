package provider

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGeneratedAPIResourcesHaveStrictRuntimeContract(t *testing.T) {
	expected := map[string]bool{
		"admin_ai_act_model_cards": false,
		"admin_campaigns":          false,
		"admin_events":             false,
		"admin_groups":             false,
		"admin_skills":             false,
		"scim_v2_users":            false,
	}

	count := 0
	for _, factory := range generatedResources {
		generated, ok := factory().(*apiResource)
		if !ok {
			continue
		}
		count++
		if _, wanted := expected[generated.name]; !wanted {
			t.Errorf("unexpected generic apiResource %s", generated.name)
		} else {
			expected[generated.name] = true
		}
		if generated.path == "" || generated.itemPathTemplate == "" || generated.idParam == "" {
			t.Errorf("%s has incomplete collection/item metadata", generated.name)
		}
		if generated.updateVerb == "" || !generated.canDelete {
			t.Errorf("%s is not a complete update/delete lifecycle", generated.name)
		}
		resolved, err := generated.itemPath("id/with space")
		if err != nil {
			t.Errorf("%s item path cannot be resolved: %v", generated.name, err)
			continue
		}
		if !strings.Contains(resolved, "id%2Fwith%20space") {
			t.Errorf("%s item identifier is not path escaped: %s", generated.name, resolved)
		}
		if strings.ContainsAny(resolved, "{}") {
			t.Errorf("%s leaves an unresolved placeholder: %s", generated.name, resolved)
		}
	}

	if count != len(expected) {
		t.Fatalf("expected %d strict apiResource values, got %d", len(expected), count)
	}
	for name, found := range expected {
		if !found {
			t.Errorf("strict apiResource %s not generated", name)
		}
	}
}

type campaignDefaultsPlanModel struct {
	ID     types.String `tfsdk:"id"`
	JSON   types.String `tfsdk:"json"`
	Name   types.String `tfsdk:"name"`
	Status types.String `tfsdk:"status"`
}

func TestCRUDCreateHydratesServerDefaultsIntoOptionalComputedState(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case request.Method == http.MethodPost && request.URL.Path == "/admin/campaigns/":
			var body map[string]any
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Errorf("invalid create body: %v", err)
			}
			if _, sent := body["status"]; sent {
				t.Errorf("unknown server default was sent by provider: %#v", body)
			}
			_, _ = w.Write([]byte(`{"id":"campaign-1"}`))
		case request.Method == http.MethodGet && request.URL.Path == "/admin/campaigns/campaign-1":
			_, _ = w.Write([]byte(`{"id":"campaign-1","name":"Launch","status":"draft"}`))
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()

	generated := apiResource{
		data:             &providerData{endpoint: server.URL, token: "test", http: server.Client()},
		name:             "admin_campaigns",
		path:             "/admin/campaigns/",
		itemPathTemplate: "/admin/campaigns/{campaign_id}",
		idParam:          "campaign_id",
		updateVerb:       "PUT",
		canDelete:        true,
		typedAttrs: []apiAttr{
			{Name: "name", TFType: "string", Required: true},
			{Name: "status", TFType: "string"},
		},
		createFields: []string{"name", "status"},
		updateFields: []string{"name"},
	}
	var schemaResponse resource.SchemaResponse
	generated.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	status, ok := schemaResponse.Schema.Attributes["status"].(schema.StringAttribute)
	if !ok || !status.Optional || !status.Computed {
		t.Fatalf("server-default field must be Optional+Computed: %#v", schemaResponse.Schema.Attributes["status"])
	}
	plan := tfsdk.Plan{Schema: schemaResponse.Schema}
	diagnostics := plan.Set(t.Context(), campaignDefaultsPlanModel{
		ID:     types.StringUnknown(),
		JSON:   types.StringUnknown(),
		Name:   types.StringValue("Launch"),
		Status: types.StringUnknown(),
	})
	if diagnostics.HasError() {
		t.Fatalf("cannot build campaign plan: %v", diagnostics)
	}
	response := resource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	generated.Create(t.Context(), resource.CreateRequest{Plan: plan}, &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("campaign create failed: %v", response.Diagnostics)
	}
	var state campaignDefaultsPlanModel
	if diagnostics := response.State.Get(t.Context(), &state); diagnostics.HasError() {
		t.Fatalf("cannot decode campaign state: %v", diagnostics)
	}
	if state.ID.ValueString() != "campaign-1" || state.Status.ValueString() != "draft" || state.Status.IsUnknown() {
		t.Fatalf("server default not hydrated into stable state: %#v", state)
	}
}

func TestAPIResourceIdentifierIsFailClosed(t *testing.T) {
	if got := idFromResponseForParam(map[string]any{"group_id": "g-42", "id": "fallback"}, "group_id"); got != "g-42" {
		t.Fatalf("expected path parameter identifier, got %q", got)
	}
	if got := idFromResponseForParam(map[string]any{"id": "fallback"}, "group_id"); got != "fallback" {
		t.Fatalf("expected compatible identifier alias, got %q", got)
	}
	if got := idFromResponseForParam(map[string]any{"status": "created"}, "group_id"); got != "" {
		t.Fatalf("missing identifier must stay empty, got %q", got)
	}

	resource := apiResource{itemPathTemplate: "/admin/groups/{group_id}", idParam: "group_id"}
	if _, err := resource.itemPath(""); err == nil {
		t.Fatal("empty item identifier must fail closed")
	}
	resource.itemPathTemplate = "/admin/groups/{other_id}"
	if _, err := resource.itemPath("g-42"); err == nil {
		t.Fatal("mismatched item placeholder must fail closed")
	}
}

func TestEveryGeneratedActionHasOneResolvableExactMutation(t *testing.T) {
	allowedMethods := map[string]bool{"POST": true, "PUT": true, "PATCH": true, "DELETE": true}
	count := 0
	for _, factory := range generatedResources {
		generated, ok := factory().(*actionApiResource)
		if !ok {
			continue
		}
		count++
		if !allowedMethods[generated.httpMethod()] {
			t.Errorf("%s has unsupported method %q", generated.name, generated.httpMethod())
		}
		if (generated.path == "") == (generated.pathTemplate == "") {
			t.Errorf("%s must define exactly one literal or templated path", generated.name)
			continue
		}

		fields := map[string]any{}
		for _, match := range _tplVar.FindAllStringSubmatch(generated.pathTemplate, -1) {
			fields[match[1]] = "id/with space"
		}
		for _, parameter := range generated.queryParams {
			fields[parameter.Name] = "query/value"
		}
		resolved, err := generated.resolvePath(fields)
		if err != nil {
			t.Errorf("%s path cannot be resolved: %v", generated.name, err)
			continue
		}
		resolved, err = generated.resolveQuery(resolved, fields)
		if err != nil {
			t.Errorf("%s query cannot be resolved: %v", generated.name, err)
			continue
		}
		if strings.ContainsAny(resolved, "{}") || strings.Contains(resolved, "id/with space") {
			t.Errorf("%s has an unresolved or unescaped path: %s", generated.name, resolved)
		}
	}
	if count == 0 {
		t.Fatal("no generated action resource found")
	}
}

func TestActionRequiredQueryParamsAreFailClosed(t *testing.T) {
	resource := actionApiResource{queryParams: []actionQueryParam{{Name: "user_id", Required: true}}}
	if _, err := resource.resolveQuery("/admin/gdpr/export", map[string]any{}); err == nil {
		t.Fatal("missing required query parameter must fail closed")
	}
	got, err := resource.resolveQuery(
		"/admin/gdpr/export",
		map[string]any{"user_id": "user/42"},
	)
	if err != nil {
		t.Fatalf("required query parameter cannot be resolved: %v", err)
	}
	if got != "/admin/gdpr/export?user_id=user%2F42" {
		t.Fatalf("unexpected encoded query path: %s", got)
	}
}

func TestActionTransportFieldsAreRemovedFromHTTPRequestBody(t *testing.T) {
	action := actionApiResource{
		pathTemplate: "/admin/blog/posts/{post_id}",
		hasJSONBody:  true,
		bodyFields:   []string{"title", "status"},
		queryParams:  []actionQueryParam{{Name: "dry_run", Required: false}},
	}
	body := map[string]any{
		"post_id": "post/42",
		"dry_run": true,
		"title":   "Titre",
		"status":  "draft",
	}

	got, ok := action.requestBody(body).(map[string]any)
	if !ok {
		t.Fatalf("expected object body, got %T", action.requestBody(body))
	}
	if _, found := got["post_id"]; found {
		t.Fatalf("path parameter leaked into body: %#v", got)
	}
	if _, found := got["dry_run"]; found {
		t.Fatalf("query parameter leaked into body: %#v", got)
	}
	if !reflect.DeepEqual(got, map[string]any{"title": "Titre", "status": "draft"}) {
		t.Fatalf("unexpected filtered body: %#v", got)
	}
}

func TestActionKeepsTransportFieldWhenOpenAPIBodyAlsoDeclaresIt(t *testing.T) {
	action := actionApiResource{
		pathTemplate: "/admin/config/email/templates/{code}",
		hasJSONBody:  true,
		bodyFields:   []string{"code", "subject"},
	}
	body := map[string]any{"code": "welcome", "subject": "Bienvenue"}
	got := action.requestBody(body).(map[string]any)
	if got["code"] != "welcome" {
		t.Fatalf("OpenAPI body field shared with path must be preserved: %#v", got)
	}
}

func TestLegacyActionFieldAliasesPreserveSchemaAndExactWireContract(t *testing.T) {
	action := actionApiResource{
		fieldAliases: map[string]string{
			"name": "code",
			"html": "body_html",
			"text": "body_text",
		},
	}
	fields := map[string]any{
		"name":    "welcome",
		"subject": "Bienvenue",
		"html":    "<p>Bonjour</p>",
		"text":    "Bonjour",
	}
	body := map[string]any{}
	for key, value := range fields {
		body[key] = value
	}
	action.applyFieldAliases(fields, body)

	for _, legacy := range []string{"name", "html", "text"} {
		if _, found := body[legacy]; found {
			t.Fatalf("legacy field %s leaked into exact API body: %#v", legacy, body)
		}
	}
	if fields["code"] != "welcome" || body["body_html"] != "<p>Bonjour</p>" || body["body_text"] != "Bonjour" {
		t.Fatalf("legacy aliases not translated: fields=%#v body=%#v", fields, body)
	}
}

func TestActionWithoutOpenAPIBodySendsNil(t *testing.T) {
	action := actionApiResource{hasJSONBody: false}
	if got := action.requestBody(map[string]any{"task_id": "42"}); got != nil {
		t.Fatalf("endpoint without requestBody must receive nil, got %#v", got)
	}
}

func TestLegacyTypedResourcesRemainTypedImmutableActions(t *testing.T) {
	expected := map[string]bool{
		"admin_ai_act_v2_incidents":         false,
		"admin_blog_posts":                  false,
		"admin_config_email_templates_item": false,
		"admin_connectors_mcp_servers":      false,
		"admin_dns_records":                 false,
		"admin_integrations":                false,
		"admin_learning_finetune_runs":      false,
		"admin_maintenance_tasks":           false,
		"admin_mcp_servers":                 false,
		"admin_oidc_providers":              false,
		"admin_seo_feeds":                   false,
		"admin_vault":                       false,
		"org_agents":                        false,
		"org_api_keys":                      false,
		"org_groups":                        false,
		"org_support_tickets":               false,
		"org_webhooks":                      false,
		"support_tickets":                   false,
		"v1_runs":                           false,
	}
	for _, factory := range generatedResources {
		action, ok := factory().(*actionApiResource)
		if !ok {
			continue
		}
		if _, wanted := expected[action.name]; !wanted {
			continue
		}
		if len(action.typedAttrs) == 0 {
			t.Errorf("legacy typed resource %s lost its typed schema", action.name)
		}
		for _, attribute := range action.typedAttrs {
			if attribute.Name == "id" || attribute.Name == "json" || attribute.Name == "body" {
				t.Errorf("legacy typed resource %s silently collides with reserved attribute %s", action.name, attribute.Name)
			}
		}
		if action.name == "admin_integrations" {
			foundRequestID := false
			for _, attribute := range action.typedAttrs {
				if attribute.Name == "request_id" && attribute.wireName() == "id" && attribute.Required {
					foundRequestID = true
				}
			}
			if !foundRequestID {
				t.Error("admin_integrations must expose required OpenAPI id as required request_id")
			}
		}
		var response resource.SchemaResponse
		action.Schema(t.Context(), resource.SchemaRequest{}, &response)
		body, ok := response.Schema.Attributes["body"].(schema.StringAttribute)
		if !ok || len(body.PlanModifiers) == 0 {
			t.Errorf("action %s body must require replacement", action.name)
		}
		expected[action.name] = true
	}
	for name, found := range expected {
		if !found {
			t.Errorf("legacy typed action %s not found", name)
		}
	}
}

type reservedActionPlanModel struct {
	ID        types.String `tfsdk:"id"`
	Body      types.String `tfsdk:"body"`
	JSON      types.String `tfsdk:"json"`
	RequestID types.String `tfsdk:"request_id"`
}

func TestReservedActionAttributeAliasUsesExactWireName(t *testing.T) {
	action := actionApiResource{
		name:         "admin_integrations",
		path:         "/admin/integrations",
		method:       "POST",
		hasJSONBody:  true,
		bodyRequired: true,
		typedAttrs: []apiAttr{{
			Name:     "request_id",
			WireName: "id",
			TFType:   "string",
			Required: true,
		}},
	}
	var schemaResponse resource.SchemaResponse
	action.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("reserved alias schema failed: %v", schemaResponse.Diagnostics)
	}
	requestID, ok := schemaResponse.Schema.Attributes["request_id"].(schema.StringAttribute)
	if !ok || !requestID.Required {
		t.Fatalf("request_id alias is not required: %#v", schemaResponse.Schema.Attributes["request_id"])
	}
	body, ok := schemaResponse.Schema.Attributes["body"].(schema.StringAttribute)
	if !ok || body.Required {
		t.Fatalf("typed required id must make opaque body optional: %#v", schemaResponse.Schema.Attributes["body"])
	}

	plan := tfsdk.Plan{Schema: schemaResponse.Schema}
	diagnostics := plan.Set(t.Context(), reservedActionPlanModel{
		ID:        types.StringUnknown(),
		Body:      types.StringNull(),
		JSON:      types.StringUnknown(),
		RequestID: types.StringValue("integration/42"),
	})
	if diagnostics.HasError() {
		t.Fatalf("cannot build reserved alias plan: %v", diagnostics)
	}
	wire, diagnostics := (&apiResource{typedAttrs: action.typedAttrs}).bodyFromTyped(t.Context(), plan)
	if diagnostics.HasError() {
		t.Fatalf("cannot build reserved alias wire body: %v", diagnostics)
	}
	if wire["id"] != "integration/42" {
		t.Fatalf("request_id was not translated to exact OpenAPI id: %#v", wire)
	}
	if _, leaked := wire["request_id"]; leaked {
		t.Fatalf("Terraform alias leaked to HTTP body: %#v", wire)
	}
}

func TestOnlyDurableGeneratedTypesSupportImport(t *testing.T) {
	var _ resource.ResourceWithImportState = &apiResource{}
	var _ resource.ResourceWithImportState = &singletonApiResource{}
	if _, ok := any(&actionApiResource{}).(resource.ResourceWithImportState); ok {
		t.Fatal("ephemeral action resources must not advertise import")
	}
}

func TestGeneratedSingletonDeleteFailsClosedWithoutPretendingRemoteReset(t *testing.T) {
	count := 0
	for _, factory := range generatedResources {
		singleton, ok := factory().(*singletonApiResource)
		if !ok {
			continue
		}
		count++
		var response resource.DeleteResponse
		singleton.Delete(t.Context(), resource.DeleteRequest{}, &response)
		if !response.Diagnostics.HasError() {
			t.Errorf("singleton %s must fail closed when no DELETE/reset contract exists", singleton.name)
			continue
		}
		if !strings.Contains(response.Diagnostics.Errors()[0].Detail(), "destroy=false") {
			t.Errorf("singleton %s delete diagnostic lacks explicit safe removal guidance: %v", singleton.name, response.Diagnostics)
		}
	}
	if count != 11 {
		t.Fatalf("expected 11 exact GET+PUT/PATCH singletons, got %d", count)
	}
}

func TestActionUpdateFailsClosedInsteadOfReplayingMutation(t *testing.T) {
	action := actionApiResource{method: "POST", path: "/admin/feature-flags"}
	var response resource.UpdateResponse
	action.Update(t.Context(), resource.UpdateRequest{}, &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("action Update must fail closed instead of replaying POST")
	}
}

func TestGeneratedCRUDSeparatesCreateAndUpdateContracts(t *testing.T) {
	expected := map[string]struct {
		createOnly string
		updateOnly string
	}{
		"admin_ai_act_model_cards": {createOnly: "org_id", updateOnly: "last_reviewed_at"},
		"admin_campaigns":          {createOnly: "channel"},
		"admin_events":             {createOnly: "host_email", updateOnly: "recording_url"},
		"admin_groups":             {createOnly: "org_id"},
		"admin_skills":             {createOnly: "org_id"},
	}

	for _, factory := range generatedResources {
		generated, ok := factory().(*apiResource)
		if !ok {
			continue
		}
		contract, wanted := expected[generated.name]
		if !wanted {
			continue
		}
		if !containsString(generated.createFields, contract.createOnly) || containsString(generated.updateFields, contract.createOnly) {
			t.Errorf("%s create-only field %s has invalid wire contract", generated.name, contract.createOnly)
		}
		if contract.updateOnly != "" && (!containsString(generated.updateFields, contract.updateOnly) || containsString(generated.createFields, contract.updateOnly)) {
			t.Errorf("%s update-only field %s has invalid wire contract", generated.name, contract.updateOnly)
		}

		var schemaResponse resource.SchemaResponse
		generated.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
		attribute, ok := schemaResponse.Schema.Attributes[contract.createOnly].(schema.StringAttribute)
		if !ok || len(attribute.PlanModifiers) == 0 {
			t.Errorf("%s.%s must require replacement", generated.name, contract.createOnly)
		}
		delete(expected, generated.name)
	}
	for name := range expected {
		t.Errorf("CRUD resource %s not found", name)
	}
}

type actionStateModel struct {
	ID   types.String `tfsdk:"id"`
	Body types.String `tfsdk:"body"`
	JSON types.String `tfsdk:"json"`
}

func actionStateForTest(t *testing.T, action *actionApiResource, identifier string) tfsdk.State {
	t.Helper()
	var schemaResponse resource.SchemaResponse
	action.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	state := tfsdk.State{Schema: schemaResponse.Schema}
	diagnostics := state.Set(t.Context(), actionStateModel{
		ID:   types.StringValue(identifier),
		Body: types.StringValue("{}"),
		JSON: types.StringValue("{}"),
	})
	if diagnostics.HasError() {
		t.Fatalf("cannot build action state: %v", diagnostics)
	}
	return state
}

func TestLegacyActionStateIsBlockedBeforeDuplicatePOST(t *testing.T) {
	action := actionApiResource{name: "admin_blog_posts", method: "POST", path: "/admin/blog/posts"}
	legacyState := actionStateForTest(t, &action, "legacy-entity-id")
	var response resource.ModifyPlanResponse
	action.ModifyPlan(
		t.Context(),
		resource.ModifyPlanRequest{State: legacyState},
		&response,
	)
	if !response.Diagnostics.HasError() {
		t.Fatal("legacy entity state must fail closed before planning a duplicate POST")
	}
	if !strings.Contains(response.Diagnostics.Errors()[0].Detail(), "dupliquerait") {
		t.Fatalf("migration diagnostic is not explicit: %v", response.Diagnostics)
	}

	currentState := actionStateForTest(t, &action, action.actionID())
	response = resource.ModifyPlanResponse{}
	action.ModifyPlan(
		t.Context(),
		resource.ModifyPlanRequest{State: currentState},
		&response,
	)
	if response.Diagnostics.HasError() {
		t.Fatalf("current action state unexpectedly blocked: %v", response.Diagnostics)
	}
}
