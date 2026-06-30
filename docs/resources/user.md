---
page_title: "aisia_user Resource - aisia"
description: |-
  Un utilisateur AISIA.
---

# aisia_user

Crée et gère un **utilisateur** AISIA (rôle + rattachement org). Si `password`
n'est pas fourni, AISIA en génère un, exposé **une seule fois** dans
`generated_password` (sensible).

## Exemple

```terraform
resource "aisia_user" "alice" {
  email        = "alice@acme.example"
  display_name = "Alice Martin"
  role         = "org_admin"
  org_id       = aisia_organization.acme.id
}
```

## Schema

### Required
- `email` (String, ForceNew) Email.

### Optional
- `display_name` (String) Nom affiché.
- `role` (String) `super_admin`|`org_owner`|`org_admin`|`org_user`|`b2c_free`|`investor` (défaut `org_user`).
- `org_id` (String) Org de rattachement.
- `password` (String, Sensitive) Mot de passe (vide = auto-généré).

### Read-Only
- `id` (String) Identifiant.
- `generated_password` (String, Sensitive) Mot de passe auto-généré (one-shot).
