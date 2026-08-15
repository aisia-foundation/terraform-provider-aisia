package provider

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func emailSingletonForTest() singletonApiResource {
	return singletonApiResource{
		name:       "admin_config_email_config",
		path:       "/admin/config/email",
		updateVerb: "PUT",
		bodyFields: []string{
			"from_email", "from_name", "smtp_host", "smtp_password",
			"smtp_port", "smtp_user", "use_tls",
		},
		readFieldAliases: map[string]string{
			"smtp_host": "host",
			"smtp_port": "port",
			"smtp_user": "user",
		},
		sensitiveFields: []string{"smtp_password"},
	}
}

func TestSingletonReadReconcilesPublicDriftAndPreservesWriteOnlySecret(t *testing.T) {
	generated := emailSingletonForTest()
	current := types.StringValue(`{"smtp_host":"smtp.old.test","smtp_port":587,"smtp_password":"write-only-secret","from_email":"ops@example.test"}`)
	observed := map[string]any{
		"host":         "smtp.new.test",
		"port":         float64(587),
		"password_set": true,
		"from_email":   "ops@example.test",
		"readonly":     "must-not-enter-body",
	}

	reconciled, err := generated.reconcileBody(current, observed)
	if err != nil {
		t.Fatalf("cannot reconcile singleton: %v", err)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(reconciled.ValueString()), &body); err != nil {
		t.Fatalf("invalid reconciled body: %v", err)
	}
	if body["smtp_host"] != "smtp.new.test" || body["smtp_password"] != "write-only-secret" {
		t.Fatalf("public drift or write-only secret contract lost: %#v", body)
	}
	if _, leaked := body["readonly"]; leaked {
		t.Fatalf("read-only response field entered PUT body: %#v", body)
	}

	imported, err := generated.reconcileBody(types.StringUnknown(), observed)
	if err != nil {
		t.Fatalf("cannot hydrate imported singleton: %v", err)
	}
	var importedBody map[string]any
	if err := json.Unmarshal([]byte(imported.ValueString()), &importedBody); err != nil {
		t.Fatalf("invalid imported singleton body: %v", err)
	}
	if importedBody["smtp_host"] != "smtp.new.test" {
		t.Fatalf("import did not hydrate public writable fields: %#v", importedBody)
	}
	if _, leaked := importedBody["smtp_password"]; leaked {
		t.Fatalf("import invented or copied a write-only secret: %#v", importedBody)
	}
	if _, leaked := importedBody["readonly"]; leaked {
		t.Fatalf("import copied a read-only field: %#v", importedBody)
	}
}

func TestSingletonUpdateSendsOnlyExactPUTContractAndConfirmsState(t *testing.T) {
	const configured = `{"smtp_host":"smtp.example.test","smtp_port":587,"smtp_user":"aisia","smtp_password":"write-only-secret","from_email":"ops@example.test","from_name":"AISIA","use_tls":true}`
	putCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch request.Method {
		case http.MethodPut:
			putCount++
			var body map[string]any
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Errorf("invalid singleton PUT: %v", err)
			}
			if len(body) != 7 || body["smtp_password"] != "write-only-secret" {
				t.Errorf("singleton PUT diverged from exact OpenAPI contract: %#v", body)
			}
			_, _ = w.Write([]byte(`{"saved":true,"smtp_password":"must-not-reach-json"}`))
		case http.MethodGet:
			_, _ = w.Write([]byte(`{"host":"smtp.example.test","port":587,"user":"aisia","from_email":"ops@example.test","from_name":"AISIA","use_tls":true,"password_set":true,"readonly":"ignored"}`))
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()

	generated := emailSingletonForTest()
	generated.data = &providerData{endpoint: server.URL, token: "test", http: server.Client()}
	var schemaResponse resource.SchemaResponse
	generated.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	state := tfsdk.State{Schema: schemaResponse.Schema}
	diagnostics := state.Set(t.Context(), singletonApiResourceModel{
		ID:   types.StringValue(generated.singletonID()),
		Body: types.StringValue(configured),
		JSON: types.StringValue("{}"),
	})
	if diagnostics.HasError() {
		t.Fatalf("cannot build singleton state: %v", diagnostics)
	}
	plan := tfsdk.Plan{Schema: schemaResponse.Schema}
	diagnostics = plan.Set(t.Context(), singletonApiResourceModel{
		ID:   types.StringUnknown(),
		Body: types.StringValue(configured),
		JSON: types.StringUnknown(),
	})
	if diagnostics.HasError() {
		t.Fatalf("cannot build singleton plan: %v", diagnostics)
	}
	response := resource.UpdateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	generated.Update(
		t.Context(),
		resource.UpdateRequest{Plan: plan, State: state},
		&response,
	)
	if response.Diagnostics.HasError() {
		t.Fatalf("singleton update failed: %v", response.Diagnostics)
	}
	if putCount != 1 {
		t.Fatalf("expected exactly one PUT, got %d", putCount)
	}
	var updated singletonApiResourceModel
	if diagnostics := response.State.Get(t.Context(), &updated); diagnostics.HasError() {
		t.Fatalf("cannot decode singleton state: %v", diagnostics)
	}
	if updated.Body.ValueString() != configured {
		t.Fatalf("known configured body was rewritten after apply: %s", updated.Body.ValueString())
	}
	if strings.Contains(updated.JSON.ValueString(), "write-only-secret") || strings.Contains(updated.JSON.ValueString(), "must-not-reach-json") {
		t.Fatalf("singleton computed JSON leaked a secret: %s", updated.JSON.ValueString())
	}
}

func TestSingletonRejectsUnknownPUTFieldBeforeHTTP(t *testing.T) {
	generated := emailSingletonForTest()
	if _, err := generated.decodeBody(types.StringValue(`{"smtp_host":"ok","unknown":"no"}`)); err == nil {
		t.Fatal("singleton must reject fields outside exact PUT schema")
	}
}

func TestGeneratedMulticloudPlanReconcilesNestedPlanObject(t *testing.T) {
	var plan *singletonApiResource
	for _, factory := range generatedResources {
		candidate, ok := factory().(*singletonApiResource)
		if ok && candidate.path == "/admin/multicloud/pack/plan" {
			plan = candidate
			break
		}
	}
	if plan == nil {
		t.Fatal("generated multicloud plan singleton not found")
	}
	if plan.responseObjectKey != "plan" {
		t.Fatalf("multicloud response must be read from nested plan object: %q", plan.responseObjectKey)
	}

	current := types.StringValue(`{"aws":"S","gcp":"M"}`)
	observed := map[string]any{
		"plan":   map[string]any{"aws": "L", "gcp": "M", "azure": "S"},
		"clouds": []any{"aws", "gcp", "azure"},
	}
	reconciled, err := plan.reconcileBody(current, observed)
	if err != nil {
		t.Fatalf("cannot reconcile nested multicloud plan: %v", err)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(reconciled.ValueString()), &body); err != nil {
		t.Fatalf("invalid reconciled multicloud body: %v", err)
	}
	if body["aws"] != "L" || body["gcp"] != "M" {
		t.Fatalf("nested multicloud drift was not reconciled: %#v", body)
	}
	if _, leaked := body["azure"]; leaked {
		t.Fatalf("unconfigured cloud entered the desired subset: %#v", body)
	}
}
