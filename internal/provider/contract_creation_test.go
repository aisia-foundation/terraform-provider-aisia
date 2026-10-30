package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type contractCreationModel struct {
	ID   types.String `tfsdk:"id"`
	Body types.String `tfsdk:"body"`
	JSON types.String `tfsdk:"json"`
	Key  types.String `tfsdk:"idempotency_key"`
}

const creationProtocol = `{"mutation_protocol":{"version":"contracts-recovery-v1","create_idempotency":{"header":"Idempotency-Key","scope":"organization_and_canonical_payload","persistent":true,"conflict_status":409}}}`

func contractPlan(t *testing.T, action *actionApiResource, key types.String, body string) tfsdk.Plan {
	t.Helper()
	var response resource.SchemaResponse
	action.Schema(t.Context(), resource.SchemaRequest{}, &response)
	if response.Diagnostics.HasError() {
		t.Fatal(response.Diagnostics)
	}
	plan := tfsdk.Plan{Schema: response.Schema}
	if diagnostics := plan.Set(t.Context(), contractCreationModel{ID: types.StringUnknown(), Body: types.StringValue(body), JSON: types.StringUnknown(), Key: key}); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	return plan
}

func TestContractCreationRefusesUnprovenServer(t *testing.T) {
	for _, payload := range []string{`{}`, strings.Replace(creationProtocol, `true`, `false`, 1), strings.Replace(creationProtocol, `409`, `200`, 1), strings.Replace(creationProtocol, `organization_and_canonical_payload`, `request_only`, 1)} {
		t.Run(payload, func(t *testing.T) {
			posts := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method == "POST" {
					posts++
				}
				if req.URL.Query().Get("org_id") != "synthetic/org +x" || req.Header.Get("Idempotency-Key") != "" {
					t.Error("incorrect preflight scope or header")
				}
				_, _ = w.Write([]byte(payload))
			}))
			defer server.Close()
			action := &actionApiResource{path: "/admin/contracts", method: "POST", hasJSONBody: true, allowAdditionalBodyFields: true, data: &providerData{endpoint: server.URL, token: "synthetic", http: server.Client()}}
			plan := contractPlan(t, action, types.StringValue("stable-key-1"), `{"org_id":"synthetic/org +x","title":"test"}`)
			_, _, diagnostics, succeeded, err := action.runMutation(t.Context(), plan)
			if err == nil || succeeded || diagnostics.HasError() || posts != 0 {
				t.Fatalf("unproven server accepted: %v %v %d", err, diagnostics, posts)
			}
		})
	}
}

func TestContractCreationKeepsKeyAcrossPartialCreateAndConflict(t *testing.T) {
	claims := map[string]string{}
	posts, writes := 0, 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == "GET" {
			_, _ = w.Write([]byte(creationProtocol))
			return
		}
		posts++
		key := req.Header.Get("Idempotency-Key")
		if key != "stable-key-1" {
			t.Errorf("missing stable key: %q", key)
		}
		var body map[string]any
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if _, present := body["idempotency_key"]; present {
			t.Error("transport key leaked into payload")
		}
		fingerprint, _ := json.Marshal(body)
		if old, found := claims[key]; found {
			if old != string(fingerprint) {
				w.WriteHeader(409)
				_, _ = w.Write([]byte(`{"detail":"private"}`))
				return
			}
		} else {
			claims[key] = string(fingerprint)
			writes++
		}
		if posts == 1 {
			w.WriteHeader(503)
			_, _ = w.Write([]byte(`{"detail":{"code":"audit_not_confirmed","persisted":true,"verified":false,"retry_safe":false}}`))
			return
		}
		_, _ = w.Write([]byte(`{"contract":{"id":"stable-contract-id"},"persisted":true,"verified":true}`))
	}))
	defer server.Close()
	action := &actionApiResource{path: "/admin/contracts", method: "POST", hasJSONBody: true, allowAdditionalBodyFields: true, data: &providerData{endpoint: server.URL, token: "synthetic", http: server.Client()}}
	for attempt := 0; attempt < 3; attempt++ {
		body := `{"org_id":"synthetic","title":"test"}`
		if attempt == 2 {
			body = `{"org_id":"synthetic","title":"changed"}`
		}
		plan := contractPlan(t, action, types.StringValue("stable-key-1"), body)
		response := resource.CreateResponse{State: tfsdk.State{Schema: plan.Schema}}
		action.Create(t.Context(), resource.CreateRequest{Plan: plan}, &response)
		if response.Diagnostics.HasError() != (attempt != 1) {
			t.Fatalf("attempt %d: %v", attempt, response.Diagnostics)
		}
		if attempt == 0 {
			var state contractCreationModel
			if diagnostics := response.State.Get(t.Context(), &state); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			if state.Key.ValueString() != "stable-key-1" || !strings.Contains(state.JSON.ValueString(), `"verified":false`) {
				t.Fatal("partial state/key lost")
			}
		}
	}
	if posts != 3 || writes != 1 {
		t.Fatalf("expected one business write: posts=%d writes=%d", posts, writes)
	}
}

func TestContractCreationInvalidKeyIsRejectedBeforeNetwork(t *testing.T) {
	for _, key := range []types.String{types.StringNull(), types.StringUnknown(), types.StringValue("short"), types.StringValue("stable-key/01"), types.StringValue("stable=key01"), types.StringValue("line\r\ninjection"), types.StringValue("clé-éventuelle"), types.StringValue(strings.Repeat("a", 129))} {
		calls := 0
		client := &providerData{endpoint: "http://local.invalid", token: "synthetic", http: &http.Client{Transport: roundTripperFunc(func(*http.Request) (*http.Response, error) { calls++; return nil, fmt.Errorf("unexpected HTTP") })}}
		action := &actionApiResource{path: "/admin/contracts", method: "POST", hasJSONBody: true, allowAdditionalBodyFields: true, data: client}
		plan := contractPlan(t, action, key, `{"org_id":"synthetic"}`)
		_, _, _, success, err := action.runMutation(t.Context(), plan)
		if err == nil || success || calls != 0 {
			t.Fatalf("invalid key reached HTTP: %v calls=%d", err, calls)
		}
	}
}

func TestIdempotencyHeaderIsLimitedToContractCreation(t *testing.T) {
	client := &providerData{}
	for _, route := range []string{"/admin/contracts/123", "/admin/contracts?x=1", "/admin/scim/tokens"} {
		if _, err := client.apiDoWithKey(t.Context(), "POST", route, nil, "stable-key-1", nil); err == nil {
			t.Fatalf("accepted route %s", route)
		}
	}
}

func TestAPIRejectsRedirectsWithoutReplayingCredentialsOrMutation(t *testing.T) {
	for _, status := range []int{301, 302, 303, 307, 308} {
		for _, method := range []string{"GET", "POST"} {
			t.Run(fmt.Sprintf("%s_%d", method, status), func(t *testing.T) {
				calls := 0
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
					calls++
					if request.URL.Path != "/admin/contracts" {
						t.Error("credentials/mutation followed a redirect")
					}
					w.Header().Set("Location", "/admin/other-sensitive-action")
					w.WriteHeader(status)
				}))
				defer server.Close()
				client := &providerData{endpoint: server.URL, token: "synthetic-token", http: &http.Client{Timeout: 30 * time.Second}}
				var out any
				var err error
				if method == "POST" {
					_, err = client.apiDoWithKey(t.Context(), method, "/admin/contracts", map[string]any{"org_id": "synthetic"}, "stable-key-1", &out)
				} else {
					_, err = client.apiDo(t.Context(), method, "/admin/contracts", nil, &out)
				}
				if calls != 1 || err == nil || out != nil || strings.Contains(err.Error(), "synthetic-token") {
					t.Fatalf("redirect falsely succeeded or leaked: calls=%d error=%v out=%v", calls, err, out)
				}
				if client.http.CheckRedirect != nil {
					t.Fatal("shared HTTP client policy was mutated")
				}
			})
		}
	}
}
