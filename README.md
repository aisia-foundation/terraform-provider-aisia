<!--
  GÉNÉRÉ — ne pas éditer à la main.
  Source: scripts/generate/09_publications.py
  Régénérer: python3 scripts/aisia.py regen
  Gate deploy: python3 scripts/release/deploy.py <ver> --mode docs
-->

> **Prod live vérifiée** : **v6.12.69** (2026-07-19) — chiffres : `project_facts.json` · régénéré par `09_publications.py`.

# Terraform Provider AISIA

[![Registry](https://img.shields.io/badge/registry-aisia--foundation%2Faisia-7B42BC)](https://registry.terraform.io/providers/aisia-foundation/aisia/latest)
[![License: MPL-2.0](https://img.shields.io/badge/license-MPL--2.0-blue)](./LICENSE)

Provider Terraform officiel pour **AISIA** — organisations, clés providers par org,
utilisateurs et clés API en **Infrastructure-as-Code** (`api.aisia.fr`).

## Cœur d'AISIA (identité produit)

AISIA est le **chef d'orchestre IA local-first** : une requête entre, le meilleur modèle (local ou cloud) exécute, la réponse sort traçable et gouvernée.

**Fonction première** : orchestrer chaque requête IA en **local-first** (Ollama sur cluster)
puis cloud si nécessaire — via `BanditRouter`, pas un simple reverse-proxy.

**Différenciation** : orchestration local-first — pas un proxy LLM stateless.

| vs proxy LLM | AISIA |
|--------------|-------|
| 1 provider fixe | **88** providers déclarés |
| Catalogue modèles | **3275** modèles catalogue · **115** locaux déclarés · **58** locaux actifs |
| Stateless | Qdrant + audit AI Act + multi-tenant |
| SaaS opaque | Déployable Swarm/K8s — **v6.12.69** LIVE |

Documentation : [README racine](../../../README.md) ·
[Product Identity](../../../specification/03-Project-State/Product-Identity-AISIA.md)

```mermaid
flowchart LR
  App[Application] --> AISIA[AISIA orchestration]
  AISIA --> Local[Ollama local]
  AISIA --> Cloud[Providers cloud]
```


---

## Ce que ce provider vous permet de faire

- Gérer vos **organisations** (tenants), **clés providers** isolées par org, **utilisateurs** et **clés d'API**.
- **Multi-tenant** : isolation par organisation, quotas, déploiement self-service.
- **IaC** : ce provider (gérer AISIA) + module public [`aisia-foundation/cluster/aisia`](https://registry.terraform.io/modules/aisia-foundation/cluster/aisia/latest) (déployer AISIA).
- **Guide** : [getting-started](docs/guides/getting-started.md) — parcours déployer + gérer en Terraform.

> Module (déployer) **+** provider (gouverner) = cycle de vie complet en Terraform.

---

## Démarrage rapide

```hcl
terraform {
  required_providers {
    aisia = {
      source  = "aisia-foundation/aisia"
      version = "~> 6.12"
    }
  }
}

provider "aisia" {
  # endpoint = "https://api.aisia.fr"
  # token via AISIA_TOKEN
}

resource "aisia_organization" "acme" {
  name          = "ACME Corp"
  slug          = "acme"
  contract_type = "saas"
}
```

## Authentification

| Variable | Rôle |
|----------|------|
| `AISIA_ENDPOINT` | URL API (défaut `https://api.aisia.fr`) |
| `AISIA_TOKEN` | Jeton admin Bearer — **sensible** |

## Versioning

Provider **couplé à AISIA** : `aisia 6.12.69` cible la plateforme **v6.12.69**.

<!-- TF-REGISTRY-STATUS -->
## Statut publication registry (honnête)

> Mesuré à la régénération docs · version repo **v6.12.69** (`VERSION` modules + provider).

| Artefact | Repo | Public registry.terraform.io |
|----------|------|------------------------------|
| Provider `aisia-foundation/aisia` | `6.12.69` | ⚠️ non mesuré (provider: <urlopen error [SSL: CERTIFICATE_VERIFY_FAILED] certificate verify failed: unable to get local issuer certificate (_ssl.c:1032)>) |
| Module `terraform-aisia-cluster` (`cluster/aisia`) | `6.12.69` | ⚠️ non mesuré (offline) |
| Module `terraform-aisia-swarm` (`swarm/aisia`) | `6.12.69` | ⚠️ non mesuré (offline) |
| Module `terraform-aws-aisia` (`aisia/aws`) | `6.12.69` | ⚠️ non mesuré (offline) |
| Module `terraform-azure-aisia` (`aisia/azure`) | `6.12.69` | ⚠️ non mesuré (offline) |
| Module `terraform-google-aisia` (`aisia/google`) | `6.12.69` | ⚠️ non mesuré (offline) |
| Module `terraform-ovh-aisia` (`aisia/ovh`) | `6.12.69` | ⚠️ non mesuré (offline) |
| Module `terraform-scaleway-aisia` (`aisia/scaleway`) | `6.12.69` | ⚠️ non mesuré (offline) |

HCP privé (`app.terraform.io/AISIA`) : non interrogé ici (token fondateur). Ne pas écrire « 100 % registry » si une ligne public est absente ou en écart.


## Développement

```bash
make build && make validate && make docs
```

## Licence

[MPL-2.0](./LICENSE)
