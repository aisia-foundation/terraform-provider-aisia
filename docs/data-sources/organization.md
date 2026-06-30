---
page_title: "aisia_organization Data Source - aisia"
description: |-
  Lit une organisation AISIA existante.
---

# aisia_organization (Data Source)

Lit une organisation AISIA existante par son `id`.

## Exemple

```terraform
data "aisia_organization" "acme" {
  id = "org-123"
}
```

## Schema

### Required
- `id` (String) ID de l'organisation.

### Read-Only
- `name`, `slug`, `contract_type` (String).
