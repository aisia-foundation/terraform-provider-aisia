package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTFAttrPropagatesSensitive(t *testing.T) {
	attr, ok := tfAttr(apiAttr{
		Name:      "secret",
		TFType:    "string",
		Required:  true,
		Sensitive: true,
	}).(schema.StringAttribute)
	if !ok {
		t.Fatalf("expected schema.StringAttribute, got %T", attr)
	}
	if !attr.Sensitive {
		t.Fatal("generated secret attribute must be marked Sensitive")
	}
}

func TestGeneratedResourcesMarkObviousSecretsSensitive(t *testing.T) {
	expected := map[string]string{
		"admin_connectors_mcp_servers":                   "headers",
		"admin_credentials_item":                         "api_key",
		"admin_integrations":                             "config",
		"admin_integrations_integration_id_put_mutation": "config",
		"admin_maintenance_tasks":                        "config",
		"admin_mcp_servers":                              "headers",
		"admin_oidc_providers":                           "config",
		"admin_oidc_providers_provider_id_put_mutation":  "config",
		"admin_vault":                                    "value",
		"org_webhooks":                                   "secret",
		"admin_orgs_org_id_api_keys_action":              "key",
		"admin_orgs_org_id_invitations_action":           "accept_url",
		"org_api_keys":                                   "key",
	}

	for _, factory := range generatedResources {
		var name string
		var fields []string
		sensitiveTyped := map[string]bool{}
		switch generated := factory().(type) {
		case *apiResource:
			name, fields = generated.name, generated.sensitiveFields
			for _, attr := range generated.typedAttrs {
				sensitiveTyped[attr.Name] = attr.Sensitive
			}
		case *singletonApiResource:
			name, fields = generated.name, generated.sensitiveFields
		case *actionApiResource:
			name, fields = generated.name, generated.sensitiveFields
		default:
			continue
		}
		field, wanted := expected[name]
		if !wanted {
			continue
		}
		found := sensitiveTyped[field]
		if !found {
			for _, candidate := range fields {
				if candidate == field {
					found = true
					break
				}
			}
		}
		if !found {
			t.Errorf("%s.%s must be generated as Sensitive", name, field)
		}
		delete(expected, name)
	}

	for resource, field := range expected {
		t.Errorf("generated resource %s.%s not found", resource, field)
	}
}

func TestGeneratedOpaqueBodiesMarkObviousSecretsSensitive(t *testing.T) {
	expected := map[string]string{
		"admin_cloud_providers_item":                     "config",
		"admin_credentials_item":                         "api_key",
		"admin_seo_app_credentials_item":                 "client_secret",
		"admin_config_email_config":                      "smtp_password",
		"admin_connectors_wizard_session_id_auth_action": "credentials",
		"admin_orgs_org_id_users_action":                 "password",
		"admin_orgs_org_id_webhooks_action":              "secret",
		"admin_providers_provider_id_config_mutation":    "config",
		"admin_providers_provider_id_key_action":         "api_key",
	}

	for _, factory := range generatedResources {
		var name string
		var fields []string
		switch generated := factory().(type) {
		case *apiResource:
			name, fields = generated.name, generated.sensitiveFields
		case *singletonApiResource:
			name, fields = generated.name, generated.sensitiveFields
		case *actionApiResource:
			name, fields = generated.name, generated.sensitiveFields
		default:
			continue
		}
		field, wanted := expected[name]
		if !wanted {
			continue
		}
		found := false
		for _, candidate := range fields {
			if candidate == field {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%s body must mark %s as sensitive", name, field)
		}
		delete(expected, name)
	}

	for name, field := range expected {
		t.Errorf("generated opaque resource %s.%s not found", name, field)
	}
}

func TestOpaqueBodySchemaAndResponseAreProtected(t *testing.T) {
	blob := apiResource{sensitiveFields: []string{"api_key"}}
	var blobSchema resource.SchemaResponse
	blob.Schema(context.Background(), resource.SchemaRequest{}, &blobSchema)
	for _, name := range []string{"body", "json"} {
		attr, ok := blobSchema.Schema.Attributes[name].(schema.StringAttribute)
		if !ok || !attr.Sensitive {
			t.Errorf("body-blob %s must be Sensitive, got %#v", name, attr)
		}
	}

	action := actionApiResource{sensitiveFields: []string{"api_key"}}
	var actionSchema resource.SchemaResponse
	action.Schema(context.Background(), resource.SchemaRequest{}, &actionSchema)
	for _, name := range []string{"body", "json"} {
		attr, ok := actionSchema.Schema.Attributes[name].(schema.StringAttribute)
		if !ok || !attr.Sensitive {
			t.Errorf("action %s must be Sensitive, got %#v", name, attr)
		}
	}

	singleton := singletonApiResource{sensitiveFields: []string{"smtp_password"}}
	var singletonSchema resource.SchemaResponse
	singleton.Schema(context.Background(), resource.SchemaRequest{}, &singletonSchema)
	for _, name := range []string{"body", "json"} {
		attr, ok := singletonSchema.Schema.Attributes[name].(schema.StringAttribute)
		if !ok || !attr.Sensitive {
			t.Errorf("singleton %s must be Sensitive, got %#v", name, attr)
		}
	}

	got := marshalStateJSON(map[string]any{
		"api_key": "cleartext",
		"key":     "one-shot-cleartext",
		"safe":    "kept",
	}, nil)
	if strings.Contains(got, "cleartext") || !strings.Contains(got, `"safe":"kept"`) {
		t.Fatalf("opaque response not safely redacted: %s", got)
	}
}

func TestSensitiveResponseIsRedactedBeforeStateJSON(t *testing.T) {
	resource := apiResource{typedAttrs: []apiAttr{{
		Name:      "value",
		TFType:    "string",
		Sensitive: true,
	}}}
	got := resource.responseJSON(map[string]any{
		"value":  "vault-cleartext",
		"nested": map[string]any{"value": "nested-cleartext", "safe": "kept"},
	})

	if strings.Contains(got, "cleartext") {
		t.Fatalf("state JSON re-exposes a secret: %s", got)
	}
	if !strings.Contains(got, `"value":"[REDACTED]"`) || !strings.Contains(got, `"safe":"kept"`) {
		t.Fatalf("unexpected redacted JSON: %s", got)
	}
}

func TestOneShotSecretsHaveDedicatedSensitiveOutputs(t *testing.T) {
	expected := map[string]string{
		"admin_orgs_org_id_api_keys_action":         "key",
		"admin_orgs_org_id_invitations_action":      "accept_url",
		"admin_organizations_org_id_licence_action": "licence_key",
		"org_api_keys":                       "key",
		"org_deployments_dep_id_plan_action": "confirm_token",
	}
	for _, factory := range generatedResources {
		action, ok := factory().(*actionApiResource)
		if !ok {
			continue
		}
		output, wanted := expected[action.name]
		if !wanted {
			continue
		}
		found := false
		for _, candidate := range action.secretOutputs {
			found = found || candidate == output
		}
		if !found {
			t.Errorf("%s lost one-shot output %s", action.name, output)
			continue
		}
		var response resource.SchemaResponse
		action.Schema(t.Context(), resource.SchemaRequest{}, &response)
		attribute, ok := response.Schema.Attributes[output].(schema.StringAttribute)
		if !ok || !attribute.Computed || !attribute.Sensitive {
			t.Errorf("%s.%s must be Computed+Sensitive, got %#v", action.name, output, attribute)
		}
		delete(expected, action.name)
	}
	for name, output := range expected {
		t.Errorf("generated action %s with output %s not found", name, output)
	}
}

type oneShotActionPlan struct {
	ID   types.String `tfsdk:"id"`
	Body types.String `tfsdk:"body"`
	JSON types.String `tfsdk:"json"`
	Key  types.String `tfsdk:"key"`
}

func TestOneShotSecretIsUsableButAbsentFromJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"key-1","key":"one-shot-cleartext"}`))
	}))
	defer server.Close()

	action := actionApiResource{
		data:            &providerData{endpoint: server.URL, token: "test", http: server.Client()},
		name:            "test_key",
		path:            "/create",
		method:          "POST",
		sensitiveFields: []string{"key"},
		secretOutputs:   []string{"key"},
	}
	var schemaResponse resource.SchemaResponse
	action.Schema(t.Context(), resource.SchemaRequest{}, &schemaResponse)
	plan := tfsdk.Plan{Schema: schemaResponse.Schema}
	diagnostics := plan.Set(t.Context(), oneShotActionPlan{
		ID:   types.StringUnknown(),
		Body: types.StringNull(),
		JSON: types.StringUnknown(),
		Key:  types.StringUnknown(),
	})
	if diagnostics.HasError() {
		t.Fatalf("cannot build action plan: %v", diagnostics)
	}
	jsonState, outputs, diagnostics, mutationSucceeded, err := action.runMutation(t.Context(), plan)
	if err != nil || diagnostics.HasError() {
		t.Fatalf("one-shot mutation failed: err=%v diagnostics=%v", err, diagnostics)
	}
	if !mutationSucceeded {
		t.Fatal("successful HTTP mutation was not reported")
	}
	if outputs["key"] != "one-shot-cleartext" {
		t.Fatalf("one-shot output lost: %#v", outputs)
	}
	if strings.Contains(jsonState, "one-shot-cleartext") || !strings.Contains(jsonState, "[REDACTED]") {
		t.Fatalf("json state was not redacted: %s", jsonState)
	}
}
