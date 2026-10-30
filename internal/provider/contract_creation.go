package provider

import (
	"context"
	"fmt"
	"net/url"
)

func (r *actionApiResource) isContractCreation() bool {
	return r.httpMethod() == "POST" && r.actionDocPath() == "/admin/contracts"
}

func validContractCreationKey(key string) bool {
	if len(key) < 8 || len(key) > 128 {
		return false
	}
	for _, character := range []byte(key) {
		if !((character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') ||
			(character >= '0' && character <= '9') || character == '_' || character == '.' || character == ':' || character == '-') {
			return false
		}
	}
	return true
}

// Un ancien serveur ignore les en-têtes inconnus : l'envoi d'une clé seule ne
// suffit donc pas à rendre sûre la recréation d'une ressource Terraform tainted.
func (r *actionApiResource) checkContractCreationProtocol(ctx context.Context, fields map[string]any) error {
	org, ok := fields["org_id"].(string)
	if !ok || org == "" {
		return fmt.Errorf("org_id explicite requis pour qualifier la création idempotente")
	}
	var response struct {
		Protocol struct {
			Version string `json:"version"`
			Create  struct {
				Header         string `json:"header"`
				Scope          string `json:"scope"`
				Persistent     bool   `json:"persistent"`
				ConflictStatus int    `json:"conflict_status"`
			} `json:"create_idempotency"`
		} `json:"mutation_protocol"`
	}
	_, err := r.data.apiDo(ctx, "GET", "/admin/contracts?org_id="+url.QueryEscape(org), nil, &response)
	if err != nil {
		return fmt.Errorf("qualification du protocole de création impossible : %w", err)
	}
	p := response.Protocol
	if p.Version != "contracts-recovery-v1" || p.Create.Header != "Idempotency-Key" || p.Create.Scope != "organization_and_canonical_payload" || !p.Create.Persistent || p.Create.ConflictStatus != 409 {
		return fmt.Errorf("création refusée : l'API doit annoncer contracts-recovery-v1 avec idempotence persistante ; aucun POST envoyé")
	}
	return nil
}
