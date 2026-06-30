---
page_title: "aisia_api_key Resource - aisia"
description: |-
  Clé d'accès programmatique (aisia_sk_…) scoped à une organisation.
---

# aisia_api_key

Crée une **clé d'accès programmatique** AISIA (`aisia_sk_…`) pour une organisation
(appels API automatisés). **Immuable** (toute modification → recréation). La
valeur brute n'est exposée **qu'à la création** (`key`, sensible).

## Exemple

```terraform
resource "aisia_api_key" "acme_ci" {
  org_id         = aisia_organization.acme.id
  name           = "ci-pipeline"
  scopes         = ["invoke", "compare"]
  rate_limit_rpm = 120
}
```

## Schema

### Required
- `org_id` (String, ForceNew) Org propriétaire.
- `name` (String, ForceNew) Nom de la clé.

### Optional
- `scopes` (List of String, ForceNew) Défaut `["invoke"]`.
- `rate_limit_rpm` / `rate_limit_rpd` (Number) Limites de débit.

### Read-Only
- `id` (String) Identifiant.
- `key` (String, Sensitive) Valeur brute (one-shot).
- `key_prefix` (String) Préfixe public.
