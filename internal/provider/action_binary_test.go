package provider

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type binaryActionPlanModel struct {
	ID           types.String `tfsdk:"id"`
	Body         types.String `tfsdk:"body"`
	JSON         types.String `tfsdk:"json"`
	SourceFile   types.String `tfsdk:"source_file"`
	SourceSHA256 types.String `tfsdk:"source_sha256"`
}

func binaryActionForTest() *actionApiResource {
	return &actionApiResource{
		name: "admin_contracts_contract_id_document_action", method: "POST",
		pathTemplate:      "/admin/contracts/{contract_id}/document",
		binaryContentType: "application/pdf", bodyRequired: true,
		queryParams: []actionQueryParam{{Name: "filename"}},
	}
}

func binaryActionPlan(t *testing.T, action *actionApiResource, model binaryActionPlanModel) tfsdk.Plan {
	t.Helper()
	var response resource.SchemaResponse
	action.Schema(t.Context(), resource.SchemaRequest{}, &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("binary schema: %v", response.Diagnostics)
	}
	plan := tfsdk.Plan{Schema: response.Schema}
	if diagnostics := plan.Set(t.Context(), model); diagnostics.HasError() {
		t.Fatalf("binary plan: %v", diagnostics)
	}
	return plan
}

func binaryFile(t *testing.T, content []byte) (string, string) {
	t.Helper()
	name := filepath.Join(t.TempDir(), "document.pdf")
	if err := os.WriteFile(name, content, 0600); err != nil {
		t.Fatal(err)
	}
	hash := sha256.Sum256(content)
	return name, hex.EncodeToString(hash[:])
}

func TestBinaryActionSendsExactPDFAndKeepsOnlyMetadataInState(t *testing.T) {
	// Octets non UTF-8 : une sérialisation JSON/base64 doit faire échouer le test.
	pdf := append([]byte("%PDF-1.7\n"), 0xff, 0x00, 0x80)
	filename, hash := binaryFile(t, pdf)
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		calls++
		body, err := io.ReadAll(request.Body)
		if err != nil || !bytes.Equal(body, pdf) {
			t.Errorf("body is not the exact PDF bytes: length=%d error=%v", len(body), err)
		}
		if request.Method != "POST" || request.URL.EscapedPath() != "/admin/contracts/c%2F1%20x/document" || request.URL.Query().Get("filename") != "c space.pdf" {
			t.Errorf("incorrect method/path/query: %s %s", request.Method, request.URL)
		}
		if request.Header.Get("Content-Type") != "application/pdf" || request.Header.Get("Authorization") != "Bearer synthetic-test-token" {
			t.Error("binary content type or Bearer authentication missing")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"contract":{"id":"c/1 x"},"persisted":true,"verified":true}`))
	}))
	defer server.Close()
	action := binaryActionForTest()
	action.data = &providerData{endpoint: server.URL, token: "synthetic-test-token", http: server.Client()}
	plan := binaryActionPlan(t, action, binaryActionPlanModel{
		ID: types.StringUnknown(), JSON: types.StringUnknown(),
		Body:       types.StringValue(`{"contract_id":"c/1 x","filename":"c space.pdf"}`),
		SourceFile: types.StringValue(filename), SourceSHA256: types.StringValue(hash),
	})
	response := resource.CreateResponse{State: tfsdk.State{Schema: plan.Schema}}
	action.Create(t.Context(), resource.CreateRequest{Plan: plan}, &response)
	if response.Diagnostics.HasError() || calls != 1 {
		t.Fatalf("binary create failed or replayed: calls=%d diagnostics=%v", calls, response.Diagnostics)
	}
	var state binaryActionPlanModel
	if diagnostics := response.State.Get(t.Context(), &state); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	if state.ID.ValueString() != "_action" || state.SourceSHA256.ValueString() != hash || strings.Contains(state.JSON.ValueString(), "%PDF-") {
		t.Fatalf("unexpected binary state: id=%s json=%s", state.ID.ValueString(), state.JSON.ValueString())
	}
	for _, name := range []string{"source_file", "source_sha256"} {
		attr := plan.Schema.(schema.Schema).Attributes[name].(schema.StringAttribute)
		if !attr.Required || len(attr.PlanModifiers) == 0 {
			t.Errorf("%s must be required and require replacement", name)
		}
	}
}

func TestBinaryActionRejectsInvalidInputBeforeNetwork(t *testing.T) {
	for _, scenario := range []string{"missing", "directory", "oversize", "non_pdf", "changed", "bad_hash", "unknown", "extra_body", "array_body", "multipart"} {
		t.Run(scenario, func(t *testing.T) {
			pdf := []byte("%PDF-1.7\n%%EOF\n")
			filename, hash := binaryFile(t, pdf)
			action := binaryActionForTest()
			calls := 0
			action.data = &providerData{endpoint: "http://local.invalid", token: "synthetic", http: &http.Client{Transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
				calls++
				return nil, fmt.Errorf("unexpected network call")
			})}}
			model := binaryActionPlanModel{ID: types.StringUnknown(), JSON: types.StringUnknown(), Body: types.StringValue(`{"contract_id":"c1"}`), SourceFile: types.StringValue(filename), SourceSHA256: types.StringValue(hash)}
			switch scenario {
			case "missing":
				model.SourceFile = types.StringValue(filename + ".missing")
			case "directory":
				model.SourceFile = types.StringValue(filepath.Dir(filename))
			case "oversize":
				if err := os.Truncate(filename, maxActionPDFBytes+1); err != nil {
					t.Fatal(err)
				}
			case "non_pdf":
				if err := os.WriteFile(filename, []byte("not a PDF"), 0600); err != nil {
					t.Fatal(err)
				}
			case "changed":
				if err := os.WriteFile(filename, []byte("%PDF-1.7\nchanged\n"), 0600); err != nil {
					t.Fatal(err)
				}
			case "bad_hash":
				model.SourceSHA256 = types.StringValue("invalid")
			case "unknown":
				model.SourceFile = types.StringUnknown()
			case "extra_body":
				model.Body = types.StringValue(`{"contract_id":"c1","document":"ignored JSON"}`)
			case "array_body":
				model.Body = types.StringValue(`[]`)
			case "multipart":
				action.binaryContentType = "multipart/form-data"
			}
			plan := binaryActionPlan(t, action, model)
			_, _, diagnostics, succeeded, err := action.runMutation(t.Context(), plan)
			if err == nil || diagnostics.HasError() || succeeded || calls != 0 {
				t.Fatalf("expected local refusal: error=%v diagnostics=%v succeeded=%v calls=%d", err, diagnostics, succeeded, calls)
			}
			if strings.Contains(err.Error(), filename) || strings.Contains(err.Error(), string(pdf)) {
				t.Fatal("local diagnostic disclosed path or document")
			}
		})
	}
}

func TestBinaryClientRejectsPDFResponseInsteadOfClaimingJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/pdf")
		_, _ = w.Write([]byte("%PDF-1.7\nprivate-document"))
	}))
	defer server.Close()
	client := providerData{endpoint: server.URL, token: "synthetic", http: server.Client()}
	var out any
	_, err := client.apiDo(t.Context(), "GET", "/admin/contracts/c1/document", nil, &out)
	if err == nil || out != nil || !strings.Contains(err.Error(), "PDF non représentable") || strings.Contains(err.Error(), "private-document") {
		t.Fatalf("binary response falsely accepted or leaked: error=%v out=%#v", err, out)
	}
}

func TestBinaryAuditFailureIsPartialAndNeverRetriedDuringCreate(t *testing.T) {
	for _, confirmed := range []bool{true, false} {
		t.Run(fmt.Sprintf("persisted_%v", confirmed), func(t *testing.T) {
			filename, hash := binaryFile(t, []byte("%PDF-1.7\n%%EOF\n"))
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				calls++
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusServiceUnavailable)
				if confirmed {
					_, _ = w.Write([]byte(`{"detail":{"code":"audit_not_confirmed","persisted":true,"verified":false,"retry_safe":false,"message":"private-response"}}`))
				} else {
					_, _ = w.Write([]byte(`{"detail":{"code":"unknown","persisted":true,"message":"private-response"}}`))
				}
			}))
			defer server.Close()
			action := binaryActionForTest()
			action.data = &providerData{endpoint: server.URL, token: "synthetic", http: server.Client()}
			plan := binaryActionPlan(t, action, binaryActionPlanModel{ID: types.StringUnknown(), JSON: types.StringUnknown(), Body: types.StringValue(`{"contract_id":"c1"}`), SourceFile: types.StringValue(filename), SourceSHA256: types.StringValue(hash)})
			response := resource.CreateResponse{State: tfsdk.State{Schema: plan.Schema}}
			action.Create(t.Context(), resource.CreateRequest{Plan: plan}, &response)
			if !response.Diagnostics.HasError() || calls != 1 {
				t.Fatalf("partial create must fail visibly with exactly one request: calls=%d diagnostics=%v", calls, response.Diagnostics)
			}
			message := response.Diagnostics.Errors()[0].Detail()
			if strings.Contains(message, "private-response") {
				t.Fatal("error response leaked")
			}
			if confirmed {
				var state binaryActionPlanModel
				if diagnostics := response.State.Get(t.Context(), &state); diagnostics.HasError() {
					t.Fatal(diagnostics)
				}
				if state.ID.ValueString() != "_action" || !strings.Contains(state.JSON.ValueString(), `"verified":false`) || !strings.Contains(message, "tainted") {
					t.Fatalf("persisted partial state or explicit recovery context missing: %s %s", state.JSON.ValueString(), message)
				}
			} else if !response.State.Raw.IsNull() {
				t.Fatal("unknown error was falsely classified as persisted")
			}
		})
	}
}
