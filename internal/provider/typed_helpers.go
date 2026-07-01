package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

// knownStr renvoie la valeur d'un types.String s'il est connu et non-null, sinon "".
// Utilisé pour omettre les champs optionnels absents du payload JSON (tag omitempty).
func knownStr(s types.String) string {
	if s.IsNull() || s.IsUnknown() {
		return ""
	}
	return s.ValueString()
}

// coalesceStr garantit une valeur connue pour un attribut Optional+Computed :
// priorité à la réponse API, sinon la valeur courante connue, sinon "".
func coalesceStr(apiVal string, cur types.String) types.String {
	if apiVal != "" {
		return types.StringValue(apiVal)
	}
	if !cur.IsNull() && !cur.IsUnknown() {
		return cur
	}
	return types.StringValue("")
}
