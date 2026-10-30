package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Une écriture métier peut précéder une panne d'audit. Garder cette distinction
// permet à l'action de conserver son state tout en rendant l'erreur visible.
type apiPersistedMutationError struct{ method, path string }

func (e *apiPersistedMutationError) Error() string {
	return fmt.Sprintf("AISIA API %s %s → HTTP 503 : écriture persistée, audit non confirmé ; relisez la ressource et traitez le state avant tout nouvel apply (Terraform peut marquer la création tainted)", e.method, safeAPIPath(e.path))
}

// apiDo exécute une requête contre l'API AISIA avec auth Bearer.
// body (si non-nil) est sérialisé en JSON ; out (si non-nil) reçoit la réponse.
func (d *providerData) apiDo(ctx context.Context, method, path string, body any, out any) (int, error) {
	return d.apiDoWithKey(ctx, method, path, body, "", out)
}

// Une clé n'est envoyée qu'au protocole de création de contrat qualifié.
func (d *providerData) apiDoWithKey(ctx context.Context, method, path string, body any, key string, out any) (int, error) {
	if key != "" && (method != http.MethodPost || path != "/admin/contracts" || !validContractCreationKey(key)) {
		return 0, fmt.Errorf("clé d'idempotence refusée pour cette mutation")
	}
	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return 0, fmt.Errorf("AISIA API %s %s : corps non sérialisable (%T)", method, safeAPIPath(path), err)
		}
		reader = bytes.NewReader(b)
	}
	return d.apiDoReaderWithKey(ctx, method, path, reader, "application/json", key, out)
}

// apiDoBinary conserve les octets du document ; aucune enveloppe JSON/base64.
func (d *providerData) apiDoBinary(ctx context.Context, method, path string, body []byte, contentType string, out any) (int, error) {
	if contentType != "application/pdf" {
		return 0, fmt.Errorf("AISIA API : type de corps binaire non supporté")
	}
	return d.apiDoReader(ctx, method, path, bytes.NewReader(body), contentType, out)
}

func (d *providerData) apiDoReader(ctx context.Context, method, path string, reader io.Reader, contentType string, out any) (int, error) {
	return d.apiDoReaderWithKey(ctx, method, path, reader, contentType, "", out)
}

func (d *providerData) apiDoReaderWithKey(ctx context.Context, method, path string, reader io.Reader, contentType, key string, out any) (int, error) {
	token, err := d.tokenForPath(path)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, method, d.endpoint+path, reader)
	if err != nil {
		return 0, fmt.Errorf("AISIA API %s %s : requête invalide (%T)", method, safeAPIPath(path), err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", contentType)
	if key != "" {
		req.Header.Set("Idempotency-Key", key)
	}

	// Une redirection 307/308 peut rejouer le POST et ses credentials vers un
	// autre chemin. Copier le client évite de changer sa politique concurremment.
	client := *d.http
	client.CheckRedirect = func(_ *http.Request, _ []*http.Request) error {
		return http.ErrUseLastResponse
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("AISIA API %s %s : erreur de transport (%T)", method, safeAPIPath(path), err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var failure struct {
			Detail map[string]any `json:"detail"`
		}
		cleanPath := safeAPIPath(path)
		contractPath := cleanPath == "/admin/contracts" || strings.HasPrefix(cleanPath, "/admin/contracts/")
		if resp.StatusCode == http.StatusServiceUnavailable && method != http.MethodGet && contractPath && json.Unmarshal(raw, &failure) == nil && failure.Detail["code"] == "audit_not_confirmed" && failure.Detail["persisted"] == true && failure.Detail["verified"] == false && failure.Detail["retry_safe"] == false {
			if target, ok := out.(*any); ok {
				*target = map[string]any{"code": "audit_not_confirmed", "persisted": true, "verified": false, "retry_safe": false}
			}
			return resp.StatusCode, &apiPersistedMutationError{method: method, path: path}
		}
		return resp.StatusCode, fmt.Errorf("AISIA API %s %s → HTTP %d (réponse volontairement expurgée)", method, safeAPIPath(path), resp.StatusCode)
	}
	ct := strings.ToLower(resp.Header.Get("Content-Type"))
	if out != nil && (strings.Contains(ct, "application/pdf") || bytes.HasPrefix(raw, []byte("%PDF-"))) {
		return resp.StatusCode, fmt.Errorf("AISIA API %s %s : réponse PDF non représentable dans un state JSON ; utilisez le téléchargement dédié", method, safeAPIPath(path))
	}
	if out != nil && len(raw) > 0 {
		looksXML := strings.Contains(ct, "xml") || bytes.HasPrefix(bytes.TrimSpace(raw), []byte("<"))
		if looksXML {
			wrapped := map[string]any{
				"raw":          string(raw),
				"content_type": "xml",
				"path":         safeAPIPath(path),
			}
			if target, ok := out.(*any); ok {
				*target = wrapped
				return resp.StatusCode, nil
			}
			encoded, err := json.Marshal(wrapped)
			if err != nil {
				return resp.StatusCode, fmt.Errorf("AISIA API %s %s : wrap XML impossible (%T)", method, safeAPIPath(path), err)
			}
			if err := json.Unmarshal(encoded, out); err != nil {
				return resp.StatusCode, fmt.Errorf("AISIA API %s %s : wrap XML non décodable (%T)", method, safeAPIPath(path), err)
			}
			return resp.StatusCode, nil
		}
		if err := json.Unmarshal(raw, out); err != nil {
			return resp.StatusCode, fmt.Errorf("AISIA API %s %s : réponse JSON invalide (%T, contenu expurgé)", method, safeAPIPath(path), err)
		}
	}
	return resp.StatusCode, nil
}

func (d *providerData) tokenForPath(path string) (string, error) {
	cleanPath := safeAPIPath(path)
	if strings.HasPrefix(cleanPath, "/scim/") {
		if strings.TrimSpace(d.scimToken) == "" {
			return "", fmt.Errorf("AISIA API %s : jeton SCIM manquant (configurez scim_token ou AISIA_SCIM_TOKEN)", cleanPath)
		}
		return d.scimToken, nil
	}
	if strings.TrimSpace(d.token) == "" {
		return "", fmt.Errorf("AISIA API %s : jeton admin manquant (configurez token ou AISIA_TOKEN)", cleanPath)
	}
	return d.token, nil
}

func safeAPIPath(path string) string {
	displayPath := path
	if queryAt := strings.IndexByte(displayPath, '?'); queryAt >= 0 {
		displayPath = displayPath[:queryAt]
	}
	if fragmentAt := strings.IndexByte(displayPath, '#'); fragmentAt >= 0 {
		displayPath = displayPath[:fragmentAt]
	}
	return displayPath
}

func apiPathSegment(value string) string { return url.PathEscape(value) }
