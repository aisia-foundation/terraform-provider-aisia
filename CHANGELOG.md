# Changelog — terraform-provider-aisia

Versioning **couplé à AISIA** (la version du provider = la version d'AISIA).

## [6.9.63] — 2026-06-30

### Changed
- Alignement de version sur AISIA v6.9.63 (versioning couplé). Sprint durcissement sécurité AISIA
  (authz get_root_admin/scoping, helper anti-SSRF, verify=True, agents non-autonomes). Provider
  inchangé fonctionnellement : resources/data sources stables (organizations, users, api_keys,
  provider_keys). Repo passé public pour CI GitHub Actions (docs sanitizées, IP-safe).

## [6.9.61] — 2026-06-30

### Changed
- Alignement de version sur AISIA v6.9.61 (versioning couplé). Provider inchangé
  fonctionnellement depuis 6.9.30 : resources/data sources stables (organizations,
  users, api_keys, provider_keys). Republication pour rétablir la parité registry
  (le provider était figé à 6.9.30 alors qu'AISIA progressait jusqu'à 6.9.61).

## [6.9.30] — 2026-06-24

### Added
- Scaffold initial du provider (terraform-plugin-framework, protocol v6).
- Provider config : `endpoint` (env `AISIA_ENDPOINT`) + `token` (env `AISIA_TOKEN`, sensible).
- Resource `aisia_organization` (CRUD via /admin/organizations) + data source.
- Resources `aisia_user` (CRUD /admin/users, mot de passe auto-généré one-shot) + `aisia_api_key` (clé programmatique org, one-shot, immuable).
- Resource `aisia_provider_key` (clé provider par org, KEY-2, write-only, import `org_id/provider_id`).
- Client HTTP Bearer, GoReleaser + signature GPG, CI (test + release), exemples, docs registry manifest.
- Couplage de version automatisé via `scripts/bump_version.py`.
