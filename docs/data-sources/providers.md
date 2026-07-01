# aisia_providers (Data Source)

Catalogue des providers LLM enregistrés (`providers.yaml`). Lecture seule.

## Example Usage
```hcl
data "aisia_providers" "all" {}

output "provider_ids" { value = data.aisia_providers.all.ids }
output "providers"    { value = jsondecode(data.aisia_providers.all.json) }
```

## Attributes Reference
- `json` (String) — réponse brute de l'API (JSON), à exploiter via `jsondecode`.
- `count` (Number) — nombre de providers.
- `ids` (List of String) — identifiants des providers.
