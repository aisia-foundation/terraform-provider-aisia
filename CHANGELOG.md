# Changelog — terraform-provider-aisia

Versioning **couplé à AISIA** (la version du provider = la version d'AISIA).

## [6.12.81] — 2026-08-05

### Changed
- Alignement de version sur AISIA v6.12.81 (versioning couplé). Déjà publié publiquement
  sur registry.terraform.io (vérifié : `aisia-foundation/aisia` v6.12.81 disponible,
  installable via `~> 6.12`). Provider inchangé fonctionnellement. ⚠️ Contrairement aux
  modules `terraform-registry/*`, ce provider ne porte pas de variable `image_tag` — le
  couplage de version ne pose donc pas le même risque de « tag d'image inexistant » que
  celui corrigé cette session sur les modules per-cloud.

## [6.12.80] — 2026-08-05

### Changed
- Alignement de version sur AISIA v6.12.80 (versioning couplé, release AISIA LIVE certifiée
  DEPLOY-REPORT all-green). Entrée rétroactive (bump réel non documenté au moment du commit
  `38058f47f`). Provider inchangé fonctionnellement.

## [6.12.79] — 2026-08-04

### Changed
- Alignement de version sur AISIA v6.12.79 (versioning couplé). Entrée rétroactive (bump
  réel non documenté au moment du commit `0ac97ec9d`). Provider inchangé fonctionnellement.

## [6.12.78] — 2026-08-04

### Changed
- Alignement de version sur AISIA v6.12.78 (versioning couplé, docs/index.md +
  guides/getting-started.md + README régénérés). Rattrape aussi le saut `v6.12.77`
  (VERSION bumpé sans entrée CHANGELOG par le commit `ad31e4ac8`, jamais publié).
  Corrige aussi la contrainte de version obsolète `~> 6.9` dans `examples/provider/
  provider.tf` et `examples/complete/main.tf` (→ `~> 6.12`, alignée sur README/docs).
  Provider inchangé fonctionnellement (resources/data sources stables) — préparation
  publication, pas de build/signature ici.

## [6.12.76] — 2026-08-02

### Changed
- Alignement de version sur AISIA v6.12.76 (versioning couplé, docs/index.md +
  guides/getting-started.md + README régénérés). Provider inchangé fonctionnellement
  (resources/data sources stables) — préparation publication, pas de build/signature ici.

## [6.9.63] — 2026-06-30

### Changed
- Alignement de version sur AISIA v6.9.63 (versioning couplé). Sprint durcissement sécurité
  (authz get_root_admin/scoping, helper anti-SSRF, verify=True, agents non-autonomes) + fixes
  audit admin. Provider inchangé fonctionnellement (resources/data sources stables).

## [6.9.61] — 2026-06-29

### Changed
- Alignement de version sur AISIA v6.9.61 (versioning couplé). Embarque le batch Voie A
  (cache sémantique TTL, /metrics bot+agent exposés, AsyncMock corrigés, E2E B2B). Provider
  inchangé fonctionnellement (resources/data sources stables).
- Modules cloud (azure/ovh/scaleway/aws/gcp) : `image_tag` default → v6.9.61 ; versions.tf
  présents partout ; parité K8s+Swarm sur GCP. Module `terraform-aisia-cluster` reste SemVer 1.0.1.

## [6.9.60] — 2026-06-29

### Changed
- Alignement de version sur AISIA v6.9.60 (versioning couplé). Aucun changement fonctionnel
  du provider : resources/data sources inchangés (organizations, users, api_keys, provider_keys).
- Embarque les correctifs v6.9.49-59 : SEC-1 (providers 401 anonyme), AI Act Art.50,
  OOM api, leak Qdrant crew, Redis Sentinel 56s, DDL CrateDB bootstrap.

## [6.9.48] — 2026-06-29

### Changed
- Alignement de version sur AISIA v6.9.48 (versioning couplé). Aucun changement fonctionnel
  du provider depuis 6.9.30 : resources/data sources inchangés. Entrée ajoutée pour rétablir
  la continuité du changelog (IND-6) figé à 6.9.30 alors qu'AISIA avançait.

## [6.9.30] — 2026-06-24

### Added
- Scaffold initial du provider (terraform-plugin-framework, protocol v6).
- Provider config : `endpoint` (env `AISIA_ENDPOINT`) + `token` (env `AISIA_TOKEN`, sensible).
- Resource `aisia_organization` (CRUD via /admin/organizations) + data source.
- Resources `aisia_user` (CRUD /admin/users, mot de passe auto-généré one-shot) + `aisia_api_key` (clé programmatique org, one-shot, immuable).
- Resource `aisia_provider_key` (clé provider par org, KEY-2, write-only, import `org_id/provider_id`).
- Client HTTP Bearer, GoReleaser + signature GPG, CI (test + release), exemples, docs registry manifest.
- Couplage de version automatisé via `scripts/bump_version.py`.
