# aisia_local_models (Data Source)

Catalogue des modèles locaux (`local_models.yaml`). Lecture seule.

## Example Usage
```hcl
data "aisia_local_models" "all" {}
output "local_models" { value = jsondecode(data.aisia_local_models.all.json) }
```

## Attributes Reference
- `json` (String) — réponse brute (JSON).
- `count` (Number) — nombre de modèles locaux.
- `ids` (List of String) — identifiants (`model_id`/`name`).
