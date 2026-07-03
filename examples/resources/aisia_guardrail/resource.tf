# Guardrail : bloquer les numeros de carte bancaire dans les prompts
resource "aisia_guardrail" "no_credit_card" {
  name    = "no-credit-card"
  pattern = "\\b(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13})\\b"
  action  = "block"
  enabled = true
}

# Guardrail : redacter les numeros de telephone dans les reponses LLM
resource "aisia_guardrail" "redact_phone_fr" {
  name    = "redact-phone-fr"
  pattern = "\\b0[1-9][0-9]{8}\\b"
  action  = "redact"
  enabled = true
}
