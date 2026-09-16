package provider

import "testing"

func TestExtractListRecognizesBusinessEnvelopes(t *testing.T) {
	tests := []struct {
		name    string
		payload map[string]any
		want    int
	}{
		{name: "plans", payload: map[string]any{"plans": []any{map[string]any{"id": "starter"}}}, want: 1},
		{name: "services", payload: map[string]any{"services": []any{map[string]any{"id": "api"}, map[string]any{"id": "bot"}}}, want: 2},
		{name: "custom list key", payload: map[string]any{"records": []any{map[string]any{"id": "one"}}}, want: 1},
		{name: "single unknown array", payload: map[string]any{"future_key": []any{map[string]any{"id": "future"}}}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractList(tt.payload, "not_the_envelope")
			if len(got) != tt.want {
				t.Fatalf("extractList() length = %d, want %d", len(got), tt.want)
			}
		})
	}
}

func TestExtractListRejectsAmbiguousUnknownArrays(t *testing.T) {
	payload := map[string]any{
		"first":  []any{map[string]any{"id": "one"}},
		"second": []any{map[string]any{"id": "two"}},
	}

	if got := extractList(payload, "missing"); got != nil {
		t.Fatalf("extractList() = %#v for ambiguous envelope, want nil", got)
	}
}
