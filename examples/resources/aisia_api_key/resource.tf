resource "aisia_api_key" "acme_ci" {
  org_id         = aisia_organization.acme.id
  name           = "ci-pipeline"
  scopes         = ["invoke", "compare"]
  rate_limit_rpm = 120
}

output "acme_ci_key" {
  value     = aisia_api_key.acme_ci.key # one-shot, sensible
  sensitive = true
}
