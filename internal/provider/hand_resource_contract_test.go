package provider

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func userPlanForTest(t *testing.T, schemaResponse resource.SchemaResponse) tfsdk.Plan {
	t.Helper()
	plan := tfsdk.Plan{Schema: schemaResponse.Schema}
	diagnostics := plan.Set(t.Context(), userModel{
		ID:                types.StringUnknown(),
		Email:             types.StringValue("alice@example.test"),
		DisplayName:       types.StringUnknown(),
		Role:              types.StringUnknown(),
		OrgID:             types.StringNull(),
		Password:          types.StringNull(),
		GeneratedPassword: types.StringUnknown(),
	})
	if diagnostics.HasError() {
		t.Fatalf("cannot build user plan: %v", diagnostics)
	}
	return plan
}

func TestUserResourceUsesExactCreateResponseAndCollectionRead(t *testing.T) {
	var updateBody map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case request.Method == http.MethodPost && request.URL.Path == "/admin/users":
			var body map[string]any
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Errorf("invalid create body: %v", err)
			}
			if _, leaked := body["password"]; leaked {
				t.Errorf("unsupported password leaked into OpenAPI body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"user_id":"u/42","email":"alice+canonical@example.test","role":"org_user","org_id":"org/create","active":true,"user_type":"investor","temp_password":"one-shot-password"}`))
		case request.Method == http.MethodGet && request.URL.Path == "/admin/users":
			_, _ = w.Write([]byte(`{"users":[{"id":"u/42","email":"alice@example.test","display_name":"Alice","role":"org_admin","org_id":"org/1"}]}`))
		case request.Method == http.MethodPut && request.URL.EscapedPath() == "/admin/users/u%2F42":
			if err := json.NewDecoder(request.Body).Decode(&updateBody); err != nil {
				t.Errorf("invalid user update body: %v", err)
			}
			_, _ = w.Write([]byte(`{"updated":true}`))
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()

	user := userResource{data: &providerData{endpoint: server.URL, token: "test", http: server.Client()}}
	var schemaResponse resource.SchemaResponse
	user.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	createResponse := resource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	user.Create(
		t.Context(),
		resource.CreateRequest{Plan: userPlanForTest(t, schemaResponse)},
		&createResponse,
	)
	if createResponse.Diagnostics.HasError() {
		t.Fatalf("user create failed: %v", createResponse.Diagnostics)
	}
	var created userModel
	if diagnostics := createResponse.State.Get(t.Context(), &created); diagnostics.HasError() {
		t.Fatalf("cannot decode created user state: %v", diagnostics)
	}
	if created.ID.ValueString() != "u/42" || created.GeneratedPassword.ValueString() != "one-shot-password" {
		t.Fatalf("exact response contract lost: %#v", created)
	}
	if created.Email.ValueString() != "alice+canonical@example.test" ||
		created.DisplayName.ValueString() != "alice" ||
		created.Role.ValueString() != "org_user" ||
		created.OrgID.ValueString() != "org/create" ||
		!created.Active.ValueBool() ||
		created.UserType.ValueString() != "investor" {
		t.Fatalf("server defaults not made known after create: %#v", created)
	}

	readResponse := resource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	user.Read(
		t.Context(),
		resource.ReadRequest{State: createResponse.State},
		&readResponse,
	)
	if readResponse.Diagnostics.HasError() {
		t.Fatalf("user collection read failed: %v", readResponse.Diagnostics)
	}
	var refreshed userModel
	if diagnostics := readResponse.State.Get(t.Context(), &refreshed); diagnostics.HasError() {
		t.Fatalf("cannot decode refreshed user state: %v", diagnostics)
	}
	if refreshed.DisplayName.ValueString() != "Alice" || refreshed.Role.ValueString() != "org_admin" || refreshed.OrgID.ValueString() != "org/1" {
		t.Fatalf("collection drift was not hydrated: %#v", refreshed)
	}

	updateModel := refreshed
	updateModel.ID = types.StringUnknown()
	updateModel.DisplayName = types.StringValue("Alice Investisseur")
	updateModel.Role = types.StringUnknown()
	updateModel.GeneratedPassword = types.StringUnknown()
	updatePlan := tfsdk.Plan{Schema: schemaResponse.Schema}
	diagnostics := updatePlan.Set(t.Context(), updateModel)
	if diagnostics.HasError() {
		t.Fatalf("cannot build user update plan: %v", diagnostics)
	}
	updateResponse := resource.UpdateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	user.Update(
		t.Context(),
		resource.UpdateRequest{Plan: updatePlan, State: readResponse.State},
		&updateResponse,
	)
	if updateResponse.Diagnostics.HasError() {
		t.Fatalf("user update failed: %v", updateResponse.Diagnostics)
	}
	if updateBody["display_name"] != "Alice Investisseur" || updateBody["role"] != "org_admin" {
		t.Fatalf("user update did not preserve optional computed state: %#v", updateBody)
	}
	var updated userModel
	if diagnostics := updateResponse.State.Get(t.Context(), &updated); diagnostics.HasError() {
		t.Fatalf("cannot decode updated user state: %v", diagnostics)
	}
	if updated.ID.ValueString() != "u/42" || updated.GeneratedPassword.ValueString() != "one-shot-password" || updated.Role.IsUnknown() {
		t.Fatalf("user update lost id/one-shot/computed state: %#v", updated)
	}
}

func TestUserCreatePersistsRemoteIDWhenOneShotSecretIsMissing(t *testing.T) {
	postCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if request.Method != http.MethodPost || request.URL.Path != "/admin/users" {
			http.NotFound(w, request)
			return
		}
		postCount++
		_, _ = w.Write([]byte(`{"user_id":"u/without-secret","email":"alice@example.test","role":"org_user"}`))
	}))
	defer server.Close()

	user := userResource{data: &providerData{endpoint: server.URL, token: "test", http: server.Client()}}
	var schemaResponse resource.SchemaResponse
	user.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	createResponse := resource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	user.Create(
		t.Context(),
		resource.CreateRequest{Plan: userPlanForTest(t, schemaResponse)},
		&createResponse,
	)
	if !createResponse.Diagnostics.HasError() {
		t.Fatal("missing one-shot secret must be surfaced as an error")
	}
	if postCount != 1 {
		t.Fatalf("user mutation was replayed: POST count=%d", postCount)
	}
	var created userModel
	if diagnostics := createResponse.State.Get(t.Context(), &created); diagnostics.HasError() {
		t.Fatalf("cannot decode recoverable user state: %v", diagnostics)
	}
	if created.ID.ValueString() != "u/without-secret" || !created.GeneratedPassword.IsNull() {
		t.Fatalf("remote ID was not preserved fail-closed: %#v", created)
	}
}

func TestOrganizationSlugAdapterMeetsBackendLengthContract(t *testing.T) {
	for _, name := range []string{"AI", "Été", "A very long organization name that exceeds the maximum accepted slug length by a wide margin"} {
		slug := deriveOrgSlug(name)
		if len(slug) < 3 || len(slug) > 64 {
			t.Errorf("derived slug %q for %q violates 3..64 contract", slug, name)
		}
	}
}

func TestAPIPathSegmentEscapesSlashAndSpace(t *testing.T) {
	if got := apiPathSegment("id/with space"); got != "id%2Fwith%20space" {
		t.Fatalf("unexpected escaped path segment: %s", got)
	}
}

func webhookPlanForTest(t *testing.T, schemaResponse resource.SchemaResponse, active types.Bool) tfsdk.Plan {
	t.Helper()
	events, diagnostics := types.ListValueFrom(t.Context(), types.StringType, []string{"invoke.completed"})
	if diagnostics.HasError() {
		t.Fatalf("cannot build webhook events: %v", diagnostics)
	}
	plan := tfsdk.Plan{Schema: schemaResponse.Schema}
	diagnostics = plan.Set(t.Context(), webhookModel{
		ID:     types.StringUnknown(),
		Name:   types.StringValue("Webhook investisseurs"),
		URL:    types.StringValue("https://hooks.example.test/aisia"),
		Events: events,
		Secret: types.StringUnknown(),
		OrgID:  types.StringValue("org/1"),
		Active: active,
	})
	if diagnostics.HasError() {
		t.Fatalf("cannot build webhook plan: %v", diagnostics)
	}
	return plan
}

func TestWebhookGeneratedSecretIsStrongAndSchemaSensitive(t *testing.T) {
	first, err := generateWebhookSecret()
	if err != nil {
		t.Fatalf("cannot generate webhook secret: %v", err)
	}
	second, err := generateWebhookSecret()
	if err != nil {
		t.Fatalf("cannot generate second webhook secret: %v", err)
	}
	if first == second || len(first) != 64 || len(second) != 64 {
		t.Fatalf("webhook secrets are not distinct 256-bit hex values")
	}
	if _, err := hex.DecodeString(first); err != nil {
		t.Fatalf("webhook secret is not valid hex: %v", err)
	}

	var response resource.SchemaResponse
	(&webhookResource{}).Schema(t.Context(), resource.SchemaRequest{}, &response)
	secret, ok := response.Schema.Attributes["secret"].(schema.StringAttribute)
	if !ok || !secret.Optional || !secret.Computed || !secret.Sensitive {
		t.Fatalf("webhook secret must be Optional+Computed+Sensitive: %#v", response.Schema.Attributes["secret"])
	}
}

func TestWebhookCreateReadUpdateAndImportPreserveSecretContract(t *testing.T) {
	var createdSecret string
	putCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case request.Method == http.MethodPost && request.URL.Path == "/admin/webhooks":
			var body map[string]any
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Errorf("invalid webhook create body: %v", err)
			}
			createdSecret, _ = body["secret"].(string)
			if len(createdSecret) != 64 {
				t.Errorf("provider-generated secret missing from create body")
			}
			_, _ = w.Write([]byte(`{"id":"wh/42","active":true}`))
		case request.Method == http.MethodPut && request.URL.EscapedPath() == "/admin/webhooks/wh%2F42":
			putCount++
			var body map[string]any
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Errorf("invalid webhook update body: %v", err)
			}
			if body["secret"] != createdSecret {
				t.Errorf("webhook update did not preserve the generated secret")
			}
			_, _ = w.Write([]byte(`{"id":"wh/42","active":false}`))
		case request.Method == http.MethodGet && request.URL.Path == "/admin/webhooks":
			_, _ = w.Write([]byte(`[{"id":"wh/42","org_id":"org/1","name":"Webhook investisseurs","url":"https://hooks.example.test/aisia","events":["invoke.completed"],"active":false}]`))
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()

	webhook := webhookResource{data: &providerData{endpoint: server.URL, token: "test", http: server.Client()}}
	var schemaResponse resource.SchemaResponse
	webhook.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	createResponse := resource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	webhook.Create(
		t.Context(),
		resource.CreateRequest{Plan: webhookPlanForTest(t, schemaResponse, types.BoolValue(false))},
		&createResponse,
	)
	if createResponse.Diagnostics.HasError() {
		t.Fatalf("webhook create failed: %v", createResponse.Diagnostics)
	}
	var created webhookModel
	if diagnostics := createResponse.State.Get(t.Context(), &created); diagnostics.HasError() {
		t.Fatalf("cannot decode created webhook state: %v", diagnostics)
	}
	if created.ID.ValueString() != "wh/42" || created.Secret.ValueString() != createdSecret || created.Active.ValueBool() {
		t.Fatalf("webhook create state contract lost: %#v", created)
	}

	readResponse := resource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	webhook.Read(t.Context(), resource.ReadRequest{State: createResponse.State}, &readResponse)
	if readResponse.Diagnostics.HasError() {
		t.Fatalf("webhook read failed: %v", readResponse.Diagnostics)
	}
	var refreshed webhookModel
	if diagnostics := readResponse.State.Get(t.Context(), &refreshed); diagnostics.HasError() {
		t.Fatalf("cannot decode refreshed webhook state: %v", diagnostics)
	}
	if refreshed.Secret.ValueString() != createdSecret {
		t.Fatal("webhook read discarded the write-only secret")
	}

	updatePlan := tfsdk.Plan{Schema: schemaResponse.Schema}
	diagnostics := updatePlan.Set(t.Context(), refreshed)
	if diagnostics.HasError() {
		t.Fatalf("cannot build webhook update plan: %v", diagnostics)
	}
	updateResponse := resource.UpdateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	webhook.Update(
		t.Context(),
		resource.UpdateRequest{Plan: updatePlan, State: readResponse.State},
		&updateResponse,
	)
	if updateResponse.Diagnostics.HasError() {
		t.Fatalf("webhook update failed: %v", updateResponse.Diagnostics)
	}
	if putCount != 2 {
		t.Fatalf("expected create activation + explicit update, got %d PUT calls", putCount)
	}

	importResponse := resource.ImportStateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	diagnostics = importResponse.State.Set(t.Context(), webhookModel{
		ID:     types.StringUnknown(),
		Name:   types.StringUnknown(),
		URL:    types.StringUnknown(),
		Events: types.ListUnknown(types.StringType),
		Secret: types.StringUnknown(),
		OrgID:  types.StringUnknown(),
		Active: types.BoolUnknown(),
	})
	if diagnostics.HasError() {
		t.Fatalf("cannot initialize webhook import state: %v", diagnostics)
	}
	webhook.ImportState(t.Context(), resource.ImportStateRequest{ID: "wh/42"}, &importResponse)
	if importResponse.Diagnostics.HasError() {
		t.Fatalf("webhook import failed: %v", importResponse.Diagnostics)
	}
	var importedID types.String
	diagnostics = importResponse.State.GetAttribute(t.Context(), path.Root("id"), &importedID)
	if diagnostics.HasError() || importedID.ValueString() != "wh/42" {
		t.Fatalf("webhook import lost id: %v %#v", diagnostics, importedID)
	}
	importReadResponse := resource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	webhook.Read(t.Context(), resource.ReadRequest{State: importResponse.State}, &importReadResponse)
	if importReadResponse.Diagnostics.HasError() {
		t.Fatalf("webhook import refresh failed: %v", importReadResponse.Diagnostics)
	}
	var imported webhookModel
	if diagnostics := importReadResponse.State.Get(t.Context(), &imported); diagnostics.HasError() {
		t.Fatalf("cannot decode imported webhook state: %v", diagnostics)
	}
	if imported.OrgID.ValueString() != "org/1" || !imported.Secret.IsNull() {
		t.Fatalf("webhook import did not hydrate org or safely omit secret: %#v", imported)
	}
}

func TestWebhookRejectsExplicitEmptySecretWithoutRemoteMutation(t *testing.T) {
	called := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		called = true
		http.Error(w, "unexpected", http.StatusInternalServerError)
	}))
	defer server.Close()

	webhook := webhookResource{data: &providerData{endpoint: server.URL, token: "test", http: server.Client()}}
	var schemaResponse resource.SchemaResponse
	webhook.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	plan := webhookPlanForTest(t, schemaResponse, types.BoolNull())
	diagnostics := plan.SetAttribute(t.Context(), path.Root("secret"), types.StringValue("  "))
	if diagnostics.HasError() {
		t.Fatalf("cannot configure explicit empty secret: %v", diagnostics)
	}
	response := resource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	webhook.Create(t.Context(), resource.CreateRequest{Plan: plan}, &response)
	if !response.Diagnostics.HasError() || called {
		t.Fatalf("empty configured secret must fail before HTTP: diagnostics=%v called=%v", response.Diagnostics, called)
	}
}

func TestWebhookUpdateAfterImportNeverPersistsUnknownComputedValues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if request.Method != http.MethodPut || request.URL.EscapedPath() != "/admin/webhooks/wh%2Fimported" {
			http.NotFound(w, request)
			return
		}
		var body map[string]any
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Errorf("invalid imported webhook update body: %v", err)
		}
		if _, leaked := body["secret"]; leaked {
			t.Errorf("unknown imported secret must be omitted, got %#v", body)
		}
		if active, ok := body["active"].(bool); !ok || active {
			t.Errorf("known active=false state was not preserved: %#v", body)
		}
		_, _ = w.Write([]byte(`{"id":"wh/imported","active":false}`))
	}))
	defer server.Close()

	webhook := webhookResource{data: &providerData{endpoint: server.URL, token: "test", http: server.Client()}}
	var schemaResponse resource.SchemaResponse
	webhook.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	events, diagnostics := types.ListValueFrom(t.Context(), types.StringType, []string{"invoke.completed"})
	if diagnostics.HasError() {
		t.Fatalf("cannot build webhook events: %v", diagnostics)
	}
	stateModel := webhookModel{
		ID:     types.StringValue("wh/imported"),
		Name:   types.StringValue("Imported"),
		URL:    types.StringValue("https://hooks.example.test/imported"),
		Events: events,
		Secret: types.StringNull(),
		OrgID:  types.StringValue("org/1"),
		Active: types.BoolValue(false),
	}
	state := tfsdk.State{Schema: schemaResponse.Schema}
	diagnostics = state.Set(t.Context(), stateModel)
	if diagnostics.HasError() {
		t.Fatalf("cannot build imported webhook state: %v", diagnostics)
	}
	planModel := stateModel
	planModel.Name = types.StringValue("Imported renamed")
	planModel.Secret = types.StringUnknown()
	planModel.Active = types.BoolUnknown()
	plan := tfsdk.Plan{Schema: schemaResponse.Schema}
	diagnostics = plan.Set(t.Context(), planModel)
	if diagnostics.HasError() {
		t.Fatalf("cannot build imported webhook plan: %v", diagnostics)
	}
	response := resource.UpdateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	webhook.Update(
		t.Context(),
		resource.UpdateRequest{Plan: plan, State: state},
		&response,
	)
	if response.Diagnostics.HasError() {
		t.Fatalf("imported webhook update failed: %v", response.Diagnostics)
	}
	var updated webhookModel
	if diagnostics := response.State.Get(t.Context(), &updated); diagnostics.HasError() {
		t.Fatalf("cannot decode imported webhook update state: %v", diagnostics)
	}
	if !updated.Secret.IsNull() || updated.Active.IsUnknown() || updated.Active.ValueBool() {
		t.Fatalf("imported webhook update kept unknown computed values: %#v", updated)
	}
}

func TestWebhookPartialActivationFailureKeepsRecoverableSensitiveState(t *testing.T) {
	var postedSecret string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if request.Method == http.MethodPost {
			var body map[string]any
			_ = json.NewDecoder(request.Body).Decode(&body)
			postedSecret, _ = body["secret"].(string)
			_, _ = w.Write([]byte(`{"id":"wh-partial","active":true}`))
			return
		}
		if request.Method == http.MethodPut {
			http.Error(w, `{"detail":"activation failed"}`, http.StatusInternalServerError)
			return
		}
		http.NotFound(w, request)
	}))
	defer server.Close()

	webhook := webhookResource{data: &providerData{endpoint: server.URL, token: "test", http: server.Client()}}
	var schemaResponse resource.SchemaResponse
	webhook.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	response := resource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	webhook.Create(
		t.Context(),
		resource.CreateRequest{Plan: webhookPlanForTest(t, schemaResponse, types.BoolValue(false))},
		&response,
	)
	if !response.Diagnostics.HasError() {
		t.Fatal("failed supplemental activation must surface an error")
	}
	if strings.Contains(response.Diagnostics.Errors()[0].Detail(), postedSecret) {
		t.Fatal("generated webhook secret leaked through diagnostics")
	}
	var state webhookModel
	if diagnostics := response.State.Get(t.Context(), &state); diagnostics.HasError() {
		t.Fatalf("partial create did not retain recoverable state: %v", diagnostics)
	}
	if state.ID.ValueString() != "wh-partial" || state.Secret.ValueString() != postedSecret || postedSecret == "" {
		t.Fatalf("partial create state is not recoverable: %#v", state)
	}
}
