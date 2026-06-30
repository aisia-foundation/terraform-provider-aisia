---
page_title: "aisia_provider_key Resource - aisia"
description: |-
  Clé API d'un provider LLM, spécifique à une organisation (isolation multi-tenant).
---

# aisia_provider_key

Définit la clé API d'un **provider LLM pour une organisation donnée** (isolation
KEY-2 : la clé d'org A n'est jamais partagée avec org B). AISIA **valide la clé par
un appel réel** au provider avant stockage. Résolution à l'usage : clé d'org >
clé globale AISIA > secret. Valeur **write-only** (l'API la masque en lecture).

## Exemple

```terraform
resource "aisia_provider_key" "acme_openai" {
  org_id        = aisia_organization.acme.id
  provider_id   = "openai"
  key_value     = var.openai_key       # sensible
  account_email = "billing@acme.example"
}
```

## Schema

### Required
- `org_id` (String, ForceNew) Org propriétaire.
- `provider_id` (String, ForceNew) Provider (`openai`, `cohere`, `groq`, …).
- `key_value` (String, Sensitive) Valeur de la clé.

### Optional
- `account_email` (String) Email du compte SaaS propriétaire (traçabilité).

### Read-Only
- `id` (String) `org_id/provider_id`.
