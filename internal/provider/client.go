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

// apiDo exécute une requête contre l'API AISIA avec auth Bearer.
// body (si non-nil) est sérialisé en JSON ; out (si non-nil) reçoit la réponse.
func (d *providerData) apiDo(ctx context.Context, method, path string, body any, out any) (int, error) {
	token, err := d.tokenForPath(path)
	if err != nil {
		return 0, err
	}
	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return 0, fmt.Errorf("AISIA API %s %s : corps non sérialisable (%T)", method, safeAPIPath(path), err)
		}
		reader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, d.endpoint+path, reader)
	if err != nil {
		return 0, fmt.Errorf("AISIA API %s %s : requête invalide (%T)", method, safeAPIPath(path), err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := d.http.Do(req)
	if err != nil {
		return 0, fmt.Errorf("AISIA API %s %s : erreur de transport (%T)", method, safeAPIPath(path), err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return resp.StatusCode, fmt.Errorf("AISIA API %s %s → HTTP %d (réponse volontairement expurgée)", method, safeAPIPath(path), resp.StatusCode)
	}
	if out != nil && len(raw) > 0 {
		ct := strings.ToLower(resp.Header.Get("Content-Type"))
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
