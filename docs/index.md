---
page_title: "Provider: aisia"
description: |-
  Le provider Terraform AISIA gere l'ensemble des ressources de la plateforme AISIA
  (organisations, utilisateurs, cles API, cles providers LLM, guardrails, webhooks,
  revendeurs, modeles locaux) via l'API admin aisia.fr.
---

# Provider AISIA

Le provider **aisia** est le provider Terraform officiel de la plateforme [AISIA](https://aisia.fr).
Il vous permet de gerer votre instance AISIA — organisations (tenants), utilisateurs, cles API,
cles providers LLM, guardrails de contenu, webhooks, revendeurs et modeles locaux — entierement
en **Infrastructure-as-Code**, sans passer par l'interface graphique.

Source : [github.com/aisia-foundation/terraform-provider-aisia](https://github.com/aisia-foundation/terraform-provider-aisia)
Registry : [registry.terraform.io/providers/aisia-foundation/aisia](https://registry.terraform.io/providers/aisia-foundation/aisia)

---

## Qu'est-ce qu'AISIA ?

**AISIA** est une plateforme d'**orchestration d'IA souveraine, local-first et multi-providers**.
Elle place une couche d'orchestration intelligente entre vos applications et les modeles d'IA
(cloud **et** locaux) pour router chaque requete vers le meilleur modele au meilleur cout,
sans verrouillage fournisseur, en gardant la maitrise de vos donnees.

| Probleme | Reponse AISIA |
|---|---|
| Couts LLM qui explosent | Routage cost-aware + modeles locaux → reduction significative |
| Vendor lock-in | Providers cloud unifies derriere une API OpenAI-compatible + bascule auto |
| Souverainete / RGPD / EU AI Act | Deploiement on-prem ou cloud souverain, audit journal AI Act |
| Fiabilite | Multi-provider, fallback IA locale, haute disponibilite |
| Multi-tenant B2B/B2C | Isolation par organisation, cles par org, quotas, self-service |

---

## Authentification

Le provider lit ses parametres depuis les **variables d'environnement** (recommande) ou
depuis le bloc `provider {}`.

| Variable d'environnement | Description |
|---|---|
| `AISIA_TOKEN` | Cle API AISIA `ak_live_…` (Bearer) — **sensible, ne jamais commiter** |
| `AISIA_ENDPOINT` | URL de l'API (defaut : `https://api.aisia.fr`) |

Obtenez votre cle dans l'interface AISIA : **Admin > API Keys > Creer une cle**.
La cle a le prefixe **`ak_live_`** et sert de jeton `Authorization: Bearer ak_live_…`.
Selectionnez le scope `super_admin` (ou `admin`) pour la gestion IaC complete.

```terraform
# Methode recommandee : variables d'environnement
# export AISIA_TOKEN="votre-jeton-admin"
# export AISIA_ENDPOINT="https://api.aisia.fr"   # optionnel

provider "aisia" {}
```

```terraform
# Methode alternative : configuration explicite (eviter pour le token en clair)
provider "aisia" {
  endpoint = "https://api.aisia.fr"
  token    = var.aisia_token  # var sensitive, injecte depuis Vault/CI
}

variable "aisia_token" {
  type      = string
  sensitive = true
}
```

---

## Demarrage rapide

Voici un exemple complet : une organisation cliente isolee, son admin,
sa cle provider OpenAI dediee (isolation multi-tenant KEY-2) et
une cle d'API programmatique pour son pipeline CI.

```terraform
terraform {
  required_providers {
    aisia = {
      source  = "aisia-foundation/aisia"
      version = "~> 6.9"
    }
  }
}

# Variables sensibles — injector depuis l'environnement CI ou un vault
variable "aisia_token"  { type = string; sensitive = true }
variable "openai_key"   { type = string; sensitive = true }

provider "aisia" {
  token = var.aisia_token
}

# 1. Creer une organisation (tenant)
resource "aisia_organization" "acme" {
  name          = "ACME Corp"
  slug          = "acme"
  contract_type = "saas"
}

# 2. Attacher la cle OpenAI d'ACME (isolee des autres orgs)
resource "aisia_provider_key" "acme_openai" {
  org_id        = aisia_organization.acme.id
  provider_id   = "openai"
  key_value     = var.openai_key
  account_email = "billing@acme.example"
}

# 3. Creer un admin pour cette organisation (mot de passe auto-genere)
resource "aisia_user" "alice" {
  email        = "alice@acme.example"
  display_name = "Alice Martin"
  role         = "org_admin"
  org_id       = aisia_organization.acme.id
}

# 4. Cle d'API programmatique pour le CI d'ACME
resource "aisia_api_key" "acme_ci" {
  org_id         = aisia_organization.acme.id
  name           = "ci-pipeline"
  scopes         = ["invoke", "compare"]
  rate_limit_rpm = 120
}

# Sorties sensibles (one-shot — sauvegardez-les immediatement)
output "alice_password" {
  value     = aisia_user.alice.generated_password
  sensitive = true
}
output "ci_api_key" {
  value     = aisia_api_key.acme_ci.key
  sensitive = true
}
```

---

## Ressources disponibles

| Ressource | Description |
|---|---|
| [`aisia_organization`](resources/organization) | Organisation (tenant) AISIA — CRUD complet |
| [`aisia_user`](resources/user) | Utilisateur avec role, org et mot de passe one-shot |
| [`aisia_api_key`](resources/api_key) | Cle d'API programmatique `ak_live_…` (one-shot ; `Authorization: Bearer` ou `X-API-Key`) |
| [`aisia_provider_key`](resources/provider_key) | Cle provider LLM par org (isolation KEY-2, write-only) |
| [`aisia_guardrail`](resources/guardrail) | Regle de filtrage de contenu (guardrail) |
| [`aisia_webhook`](resources/webhook) | Webhook sortant avec signature HMAC |
| [`aisia_reseller`](resources/reseller) | Revendeur avec part de revenu et branding |
| [`aisia_local_model`](resources/local_model) | Modele local (overlay du local_models.yaml) |

## Data sources disponibles (selection)

Les data sources en lecture seule permettent d'**interroger l'etat** de la plateforme
sans creer de ressources. Ils retournent toujours un champ `json` (brut) et `item_count`.

| Data source | Endpoint AISIA | Usage typique |
|---|---|---|
| [`aisia_organization`](data-sources/organization) | `/admin/organizations/:id` | Lire une org existante par ID |
| [`aisia_local_models`](data-sources/local_models) | `/local-models` | Catalogue modeles locaux disponibles |
| [`aisia_admin_api_keys`](data-sources/admin_api_keys) | `/admin/api-keys` | Lister toutes les cles API |
| [`aisia_admin_providers`](data-sources/admin_cloud_providers) | `/admin/providers` | Etat des providers LLM |
| [`aisia_admin_analytics`](data-sources/admin_analytics) | `/admin/analytics` | Metriques d'usage |
| [`aisia_admin_billing_stats`](data-sources/admin_billing_stats) | `/admin/billing/stats` | Stats de facturation |
| [`aisia_org_me`](data-sources/org_me) | `/org/me` | Profil de l'organisation courante |

> **181 data sources** sont disponibles — un par endpoint de l'API AISIA. Ils permettent
> de lire n'importe quelle donnee de la plateforme et de l'utiliser dans d'autres ressources
> Terraform via `jsondecode(data.aisia_xxx.yyy.json)`.

---

## Cycle de vie complet

Ce provider couvre la **gestion** de la plateforme. Pour le **deploiement** de la
plateforme elle-meme (cluster Kubernetes ou Docker Swarm), utilisez les modules
Terraform officiels AISIA :

| Module | Usage |
|---|---|
| [`aisia-foundation/aisia/kubernetes`](https://registry.terraform.io/modules/aisia-foundation/aisia/kubernetes) | Deploie AISIA sur un cluster K8s existant |
| [`aisia-foundation/aisia/aws`](https://registry.terraform.io/modules/aisia-foundation/aisia/aws) | Provisionne EKS + deploie AISIA (AWS) |
| [`aisia-foundation/aisia/azure`](https://registry.terraform.io/modules/aisia-foundation/aisia/azure) | Provisionne AKS + deploie AISIA (Azure) |
| [`aisia-foundation/aisia/google`](https://registry.terraform.io/modules/aisia-foundation/aisia/google) | Provisionne GKE + deploie AISIA (GCP) |
| [`aisia-foundation/aisia/ovh`](https://registry.terraform.io/modules/aisia-foundation/aisia/ovh) | Provisionne OVHcloud + deploie AISIA |
| [`aisia-foundation/aisia/scaleway`](https://registry.terraform.io/modules/aisia-foundation/aisia/scaleway) | Provisionne Scaleway Kapsule + deploie AISIA |

```
module (deployer AISIA) + provider (gerer AISIA) = cycle de vie complet en Terraform
```

---

## Versionnage

La version du provider **suit la version d'AISIA** : `6.9.81` cible la plateforme
AISIA v6.9.81. L'argument `version = "~> 6.9"` dans `required_providers` accepte
toutes les versions `6.9.x` (patch compatible).

---

## Schema

### Optional

- `endpoint` (String) URL de base de l'API AISIA (defaut `https://api.aisia.fr`, ou env `AISIA_ENDPOINT`).
- `token` (String, Sensitive) Jeton admin AISIA (Bearer). Preferez l'env `AISIA_TOKEN`. Sensible.
