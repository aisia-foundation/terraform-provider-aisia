# aisia_provider (Resource)

Provider LLM AISIA géré en DB (overlay du `providers.yaml`, v6.9.64 Wave 2). Permet de
créer/modifier/supprimer un provider en IaC. Le `providers.yaml` reste les défauts versionnés ;
la DB l'emporte sur un même `id`.

## Example Usage
```hcl
resource "aisia_provider" "openrouter" {
  id          = "openrouter"
  config_json = jsonencode({
    name         = "OpenRouter"
    adapter_name = "openai"
    base_url     = "https://openrouter.ai/api/v1"
    model        = "openai/gpt-4o"
    status       = "active"
  })
}
```

## Argument Reference
- `id` (Required, immuable) — identifiant du provider.
- `config_json` (Required) — définition complète du provider en JSON (`jsonencode({...})`).

## Attributes Reference
- `name` (Computed) — nom dérivé de la config.
- `status` (Computed) — statut (active/inactive).
