# Terraform Provider AISIA

[![Registry](https://img.shields.io/badge/registry-AISIA%2Faisia-7B42BC)](https://app.terraform.io/app/AISIA/registry/providers/private/AISIA/aisia)
[![License: MPL-2.0](https://img.shields.io/badge/license-MPL--2.0-blue)](./LICENSE)

Provider Terraform officiel pour **AISIA** — gérez vos ressources AISIA
(organisations, clés providers par org, utilisateurs, clés d'API) **en
Infrastructure-as-Code**, contre l'API AISIA (`api.aisia.fr`).

---

## C'est quoi AISIA ?

**AISIA** est une plateforme d'**orchestration d'IA souveraine, local-first et
multi-providers**. Elle place une couche d'orchestration intelligente entre vos
applications et les modèles d'IA (cloud **et** locaux), pour router chaque requête
vers le meilleur modèle au meilleur coût, sans verrouillage fournisseur, en
gardant la maîtrise de vos données.

### Les problèmes qu'AISIA résout

| Problème | Réponse AISIA |
|---|---|
| **Coûts LLM** qui explosent | Routage *cost-aware* + modèles **locaux** → réduction significative des coûts |
| **Vendor lock-in** (un seul fournisseur) | Providers cloud unifiés derrière une API OpenAI-compatible + bascule automatique |
| **Souveraineté / RGPD / EU AI Act** | Déploiement **on-prem / cloud souverain**, données maîtrisées, journal d'audit AI Act |
| **Fiabilité** | Multi-provider, fallback IA locale, haute disponibilité |
| **Multi-tenant** (B2B/B2C, SaaS/BaaS/PaaS) | Isolation par organisation, clés par org, quotas, déploiement self-service |

### Ce que ce provider vous permet de faire

- Gérer vos **organisations** (tenants), leurs **clés providers** isolées par org, leurs **utilisateurs** et leurs **clés d'API** — en Infrastructure-as-Code.
- **Multi-tenant** : isolation par organisation, clés providers **par organisation**, quotas.
- **IaC** : ce provider (gérer les ressources AISIA) + le module [`terraform-aisia-cluster`](https://app.terraform.io/app/AISIA/registry/modules/private/AISIA/aisia/kubernetes) (déployer la plateforme).

> Module (déployer AISIA) **+** provider (gérer AISIA) = cycle de vie complet en Terraform.
>
> ℹ️ AISIA est une solution **propriétaire** (brevet / dépôt INPI). Cette documentation
> décrit **l'usage** de l'outillage IaC, pas l'architecture interne ni la conception du produit.

---

## Démarrage rapide

```hcl
terraform {
  required_providers {
    aisia = {
      source  = "app.terraform.io/AISIA/aisia" # registry privé HCP
      version = "~> 6.9"
    }
  }
}

provider "aisia" {
  # endpoint = "https://api.aisia.fr"  # défaut, ou env AISIA_ENDPOINT
  # token    = "..."                    # préférez l'env AISIA_TOKEN (sensible)
}

# Une organisation cliente isolée
resource "aisia_organization" "acme" {
  name          = "ACME Corp"
  slug          = "acme"
  contract_type = "saas"
}

# Sa clé OpenAI dédiée (isolée des autres orgs — KEY-2)
resource "aisia_provider_key" "acme_openai" {
  org_id      = aisia_organization.acme.id
  provider_id = "openai"
  key_value   = var.openai_key
}

# Un admin pour cette org (mot de passe auto-généré)
resource "aisia_user" "alice" {
  email  = "alice@acme.example"
  role   = "org_admin"
  org_id = aisia_organization.acme.id
}

# Une clé d'API programmatique pour leur CI
resource "aisia_api_key" "acme_ci" {
  org_id = aisia_organization.acme.id
  name   = "ci-pipeline"
  scopes = ["invoke", "compare"]
}
```

## Authentification

| Variable | Rôle |
|---|---|
| `AISIA_ENDPOINT` | URL API (défaut `https://api.aisia.fr`) |
| `AISIA_TOKEN` | Jeton admin AISIA (Bearer) — **sensible**, ne jamais commiter |

## Ressources & data sources

| Type | Description |
|---|---|
| [`aisia_organization`](docs/resources/organization.md) | Organisation (tenant) — CRUD |
| [`aisia_provider_key`](docs/resources/provider_key.md) | Clé provider LLM **par org** (isolation), write-only |
| [`aisia_user`](docs/resources/user.md) | Utilisateur (rôle, org, mot de passe one-shot) |
| [`aisia_api_key`](docs/resources/api_key.md) | Clé d'accès programmatique (`aisia_sk_…`), one-shot |
| [`aisia_organization` (data)](docs/data-sources/organization.md) | Lecture d'une org existante |

## Versioning — couplé à AISIA

La version du provider **suit la version d'AISIA** : `aisia 6.9.30` cible la
plateforme AISIA v6.9.30. À chaque release AISIA, une version identique du
provider est publiée (`scripts/bump_version.py` automatise le couplage).

## Développement

```bash
make build      # compile
make validate   # go vet + build
make docs       # tfplugindocs (régénère docs/)
make testacc    # tests d'acceptation (TF_ACC=1, instance AISIA requise)
```

## Sécurité

Voir [SECURITY.md](./SECURITY.md). Binaires signés GPG (clé `96D2AD35E6565778`).

## Licence

[MPL-2.0](./LICENSE).
