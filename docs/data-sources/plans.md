# aisia_plans (Data Source)

Plans d'abonnement AISIA (Free / SaaS / BaaS / PaaS) et leurs quotas. Lecture seule.

## Example Usage
```hcl
data "aisia_plans" "all" {}
output "plans" { value = jsondecode(data.aisia_plans.all.json) }
```

## Attributes Reference
- `json` (String) — réponse brute (JSON) : `name`, `plan_type`, `price_eur`, `active`, quotas…
- `count` (Number) — nombre de plans.
- `ids` (List of String) — slugs des plans.
