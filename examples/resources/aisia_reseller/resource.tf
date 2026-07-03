# Revendeur partenaire avec 30% de part de revenu et domaine personnalise
resource "aisia_reseller" "partner_acme" {
  name          = "ACME Integrator"
  contact_email = "partnership@acme-integrator.example"
  domain        = "ai.acme-integrator.example"
  rev_share_pct = 30
  notes         = "Partenaire Gold — integration ERP secteur manufacturing"
}

output "reseller_id" {
  value       = aisia_reseller.partner_acme.id
  description = "ID du revendeur ACME (a passer aux organisations enfants)."
}
