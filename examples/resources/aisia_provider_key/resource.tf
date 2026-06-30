# Clé OpenAI propre à l'organisation ACME (isolation multi-tenant KEY-2).
resource "aisia_provider_key" "acme_openai" {
  org_id        = aisia_organization.acme.id
  provider_id   = "openai"
  key_value     = var.openai_key
  account_email = "billing@acme.example"
}

variable "openai_key" {
  type      = string
  sensitive = true
}
