resource "aisia_user" "alice" {
  email        = "alice@acme.example"
  display_name = "Alice Martin"
  role         = "org_admin"
  org_id       = aisia_organization.acme.id
  # password non fourni → AISIA en génère un, exposé une fois :
}

output "alice_generated_password" {
  value     = aisia_user.alice.generated_password
  sensitive = true
}
