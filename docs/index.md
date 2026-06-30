---
page_title: "AISIA Provider"
description: |-
  Gérez les ressources de la plateforme AISIA (orchestration IA souveraine) en Terraform.
---

# AISIA Provider

Le provider AISIA permet de gérer **en Infrastructure-as-Code** les ressources de
la plateforme [AISIA](https://aisia.fr) — orchestration d'IA souveraine, local-first,
multi-providers : organisations (tenants), clés providers LLM par organisation,
utilisateurs et clés d'API programmatiques.

Couplé au module [`terraform-aisia-cluster`](https://app.terraform.io/app/AISIA/registry/modules/private/AISIA/aisia/kubernetes)
(qui **déploie** AISIA), ce provider **gère** AISIA → cycle de vie complet en Terraform.

## Exemple

```terraform
provider "aisia" {
  endpoint = "https://api.aisia.fr" # ou AISIA_ENDPOINT
  # token via AISIA_TOKEN (sensible)
}
```

## Schema

### Optional

- `endpoint` (String) URL de l'API AISIA. Défaut `https://api.aisia.fr`. Env `AISIA_ENDPOINT`.
- `token` (String, Sensitive) Jeton admin AISIA (Bearer). Env `AISIA_TOKEN`.
