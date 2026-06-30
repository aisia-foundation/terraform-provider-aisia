---
page_title: "aisia_organization Resource - aisia"
description: |-
  Une organisation (tenant) AISIA.
---

# aisia_organization

Crée et gère une **organisation** AISIA (tenant isolé). Chaque org a ses propres
utilisateurs, clés providers, quotas et environnement, selon son `contract_type`
(SaaS partagé / BaaS dédié / PaaS on-prem).

## Exemple

```terraform
resource "aisia_organization" "acme" {
  name          = "ACME Corp"
  slug          = "acme"
  contract_type = "saas"
}
```

## Schema

### Required
- `name` (String) Nom de l'organisation.

### Optional
- `slug` (String) Slug / sous-domaine.
- `contract_type` (String) `saas` | `baas` | `paas` (défaut `saas`).

### Read-Only
- `id` (String) Identifiant généré.
