package provider

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func TestAPIDiagnosticsNeverExposeResponseOrQuerySecrets(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`{"detail":{"input":"cleartext-secret"}}`))
	}))
	defer server.Close()

	client := providerData{endpoint: server.URL, token: "bearer-secret", http: server.Client()}
	_, err := client.apiDo(
		context.Background(),
		http.MethodPost,
		"/admin/test?api_key=query-secret",
		map[string]any{"api_key": "body-secret"},
		nil,
	)
	if err == nil {
		t.Fatal("expected API error")
	}
	message := err.Error()
	for _, secret := range []string{"cleartext-secret", "query-secret", "body-secret", "bearer-secret", "api_key="} {
		if strings.Contains(message, secret) {
			t.Fatalf("diagnostic leaked %q: %s", secret, message)
		}
	}
	if !strings.Contains(message, "HTTP 422") || !strings.Contains(message, "expurg") {
		t.Fatalf("diagnostic lost safe status context: %s", message)
	}
}

func TestAPIDiagnosticsNeverExposeTransportErrorOrURL(t *testing.T) {
	client := providerData{
		endpoint: "https://api.invalid",
		token:    "bearer-secret",
		http: &http.Client{Transport: roundTripperFunc(func(request *http.Request) (*http.Response, error) {
			return nil, errors.New("dial refused for " + request.URL.String() + " transport-secret")
		})},
	}
	_, err := client.apiDo(
		context.Background(),
		http.MethodPost,
		"/admin/test?api_key=query-secret",
		map[string]any{"api_key": "body-secret"},
		nil,
	)
	if err == nil {
		t.Fatal("expected transport error")
	}
	message := err.Error()
	for _, secret := range []string{"transport-secret", "query-secret", "body-secret", "bearer-secret", "api_key="} {
		if strings.Contains(message, secret) {
			t.Fatalf("transport diagnostic leaked %q: %s", secret, message)
		}
	}
	if !strings.Contains(message, "POST /admin/test") || !strings.Contains(message, "transport") {
		t.Fatalf("safe transport context missing: %s", message)
	}
}

func TestAPIDiagnosticsNeverExposeRedirectURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		http.Redirect(w, request, "/target?token=redirect-secret", http.StatusFound)
	}))
	defer server.Close()

	client := server.Client()
	client.CheckRedirect = func(request *http.Request, _ []*http.Request) error {
		return errors.New("redirect refused " + request.URL.String())
	}
	provider := providerData{endpoint: server.URL, token: "bearer-secret", http: client}
	_, err := provider.apiDo(context.Background(), http.MethodPost, "/admin/start", nil, nil)
	if err == nil {
		t.Fatal("expected redirect error")
	}
	message := err.Error()
	for _, secret := range []string{"redirect-secret", "bearer-secret", "token="} {
		if strings.Contains(message, secret) {
			t.Fatalf("redirect diagnostic leaked %q: %s", secret, message)
		}
	}
}
