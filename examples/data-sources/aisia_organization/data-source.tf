# Lire une organisation existante par son ID
# Utile pour referencer une org creee hors de ce module Terraform.
data "aisia_organization" "existing" {
  id = "org-uuid-existant"
}

output "org_name" {
  value = data.aisia_organization.existing.name
}
