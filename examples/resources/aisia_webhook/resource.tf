variable "webhook_secret" {
  type        = string
  sensitive   = true
  description = "Secret HMAC-SHA256 pour la signature des livraisons webhook."
}

# Webhook : notifications Slack sur les erreurs d'invocation
resource "aisia_webhook" "slack_alerts" {
  name   = "slack-error-alerts"
  url    = "https://hooks.slack.com/services/T.../B.../xxx"
  events = ["invoke.failed", "provider.circuit_open"]
  secret = var.webhook_secret
}

# Webhook : audit du cycle de vie des organisations (RGPD)
resource "aisia_webhook" "org_audit" {
  name   = "org-lifecycle-audit"
  url    = "https://audit.acme.example/aisia"
  events = ["org.created", "org.deleted", "user.created", "user.deleted"]
  secret = var.webhook_secret
}
