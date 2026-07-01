# aisia_agents (Data Source)

Catalogue des agents AISIA (`identity.yaml`). **Lecture seule** — les agents ne sont pas
gérables en CRUD Terraform : leur définition est versionnée (identity.yaml) et leur
exécution est orchestrée au runtime (cycle/run, HITL). Pour piloter les agents, utiliser
l'API `/admin/agent/*` (actions impératives), pas Terraform.

## Example Usage
```hcl
data "aisia_agents" "all" {}
output "agents" { value = jsonencode(data.aisia_agents.all.json) }
output "codes"  { value = data.aisia_agents.all.ids }
```

## Attributes Reference
- `json` (String) — réponse brute (JSON) : `code`, `name`, `model`, `role`.
- `count` (Number) — nombre d'agents.
- `ids` (List of String) — codes des agents.
