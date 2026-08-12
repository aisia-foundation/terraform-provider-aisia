package provider

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

var configJSONServerMetadata = map[string]struct{}{
	"persisted": {}, "verified": {}, "runtime_applied": {},
	"runtime_consistency": {}, "cache_invalidation": {},
	"created_at": {}, "updated_at": {}, "deleted_at": {},
}

func configJSONKeySensitive(key string) bool {
	normalized := strings.ToLower(strings.TrimSpace(key))
	if _, exact := sensitiveFieldSet(nil)[normalized]; exact {
		return true
	}
	return strings.HasSuffix(normalized, "_password") ||
		strings.HasSuffix(normalized, "_secret") ||
		strings.HasSuffix(normalized, "_token") ||
		strings.HasSuffix(normalized, "_api_key") ||
		strings.HasSuffix(normalized, "_private_key")
}

func decodeConfigJSONObject(value types.String, label string) (map[string]any, error) {
	if value.IsNull() || value.IsUnknown() || strings.TrimSpace(value.ValueString()) == "" {
		return nil, fmt.Errorf("%s requis", label)
	}
	var object map[string]any
	if err := json.Unmarshal([]byte(value.ValueString()), &object); err != nil {
		return nil, fmt.Errorf("%s invalide (objet JSON attendu): %w", label, err)
	}
	return object, nil
}

func sanitizeImportedConfig(value any) any {
	switch typed := value.(type) {
	case map[string]any:
		out := make(map[string]any)
		for key, item := range typed {
			if configJSONKeySensitive(key) {
				continue
			}
			out[key] = sanitizeImportedConfig(item)
		}
		return out
	case []any:
		out := make([]any, len(typed))
		for index, item := range typed {
			out[index] = sanitizeImportedConfig(item)
		}
		return out
	default:
		return value
	}
}

func reconcileConfiguredConfig(key string, configured any, observed any) (any, bool) {
	if configJSONKeySensitive(key) {
		return configured, false
	}
	desiredObject, desiredObjectOK := configured.(map[string]any)
	remoteObject, remoteObjectOK := observed.(map[string]any)
	if desiredObjectOK && remoteObjectOK {
		result := make(map[string]any, len(desiredObject))
		changed := false
		for childKey, desired := range desiredObject {
			result[childKey] = desired
			if remote, found := remoteObject[childKey]; found {
				reconciled, nestedChanged := reconcileConfiguredConfig(childKey, desired, remote)
				result[childKey] = reconciled
				changed = changed || nestedChanged
			}
		}
		return result, changed
	}
	if reflect.DeepEqual(configured, observed) {
		return configured, false
	}
	return observed, true
}

// reconcileConfigJSON met à jour uniquement les clés que Terraform gérait déjà.
// Les secrets write-only/masqués et les clés absentes du GET sont préservés. Lors
// d'un import, seules les clés publiques de la définition distante sont adoptées.
func reconcileConfigJSON(current types.String, remote map[string]any, identityFields ...string) (types.String, error) {
	excluded := make(map[string]struct{}, len(identityFields)+len(configJSONServerMetadata))
	for key := range configJSONServerMetadata {
		excluded[key] = struct{}{}
	}
	for _, key := range identityFields {
		excluded[key] = struct{}{}
	}
	if current.IsNull() || current.IsUnknown() {
		public := make(map[string]any)
		for key, value := range remote {
			if _, skip := excluded[key]; skip || configJSONKeySensitive(key) {
				continue
			}
			public[key] = sanitizeImportedConfig(value)
		}
		raw, err := json.Marshal(public)
		if err != nil {
			return types.StringNull(), err
		}
		return types.StringValue(string(raw)), nil
	}
	desired, err := decodeConfigJSONObject(current, "config_json")
	if err != nil {
		return types.StringNull(), err
	}
	changed := false
	for key, configured := range desired {
		if _, skip := excluded[key]; skip {
			continue
		}
		observed, found := remote[key]
		if !found {
			continue
		}
		reconciled, fieldChanged := reconcileConfiguredConfig(key, configured, observed)
		if fieldChanged {
			desired[key] = reconciled
			changed = true
		}
	}
	if !changed {
		return current, nil
	}
	raw, err := json.Marshal(desired)
	if err != nil {
		return types.StringNull(), err
	}
	return types.StringValue(string(raw)), nil
}
