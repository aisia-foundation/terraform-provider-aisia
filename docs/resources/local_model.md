# aisia_local_model (Resource)

Modèle local AISIA géré en DB (overlay du `local_models.yaml`, v6.9.64 Wave 2). Le YAML
reste les défauts versionnés ; la DB l'emporte sur un même `model_id`.

## Example Usage
```hcl
resource "aisia_local_model" "llama70b" {
  model_id    = "ollama-llama33-70b"
  config_json = jsonencode({
    display_name = "Llama 3.3 70B (Ollama)"
    runtime      = "ollama"
    endpoint     = "http://localhost:11434"
    model_name   = "llama3.3:70b"
    priority     = 120
    status       = "active"
  })
}
```

## Argument Reference
- `model_id` (Required, immuable) — identifiant du modèle local.
- `config_json` (Required) — définition complète en JSON (`jsonencode({...})`).

## Attributes Reference
- `display_name` (Computed) · `status` (Computed).
