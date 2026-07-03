# =============================================================================
# Exemple complet AISIA — cycle de vie d'un tenant B2B
#
# Deploie en une seule passe :
#   - 1 organisation cliente (tenant SaaS)
#   - 1 administrateur de l'organisation (mot de passe auto-genere)
#   - 1 utilisateur standard
#   - 1 cle provider OpenAI isolee (KEY-2 multi-tenant)
#   - 1 cle d'API programmatique pour le CI/CD
#   - 1 guardrail de contenu (conformite)
#   - 1 webhook de monitoring
#
# Prerequis :
#   export AISIA_TOKEN="votre-jeton-super-admin"
#   terraform init && terraform apply
# =============================================================================

terraform {
  required_providers {
    aisia = {
      source  = "aisia-foundation/aisia"
      version = "~> 6.9"
    }
  }
  required_version = ">= 1.5"
}

provider "aisia" {
  # token et endpoint lus depuis AISIA_TOKEN / AISIA_ENDPOINT
}

# ---------------------------------------------------------------------------
# Variables
# ---------------------------------------------------------------------------
variable "openai_api_key" {
  type        = string
  sensitive   = true
  description = "Cle API OpenAI propre au client (sera isolee dans leur org)."
}

variable "webhook_secret" {
  type        = string
  sensitive   = true
  description = "Secret HMAC pour la signature des webhooks."
  default     = null
}

variable "client_name" {
  type        = string
  description = "Nom de l'organisation cliente."
  default     = "ACME Corp"
}

variable "client_slug" {
  type        = string
  description = "Slug (sous-domaine) de l'organisation."
  default     = "acme"
}

# ---------------------------------------------------------------------------
# 1. Organisation (tenant)
# ---------------------------------------------------------------------------
resource "aisia_organization" "client" {
  name          = var.client_name
  slug          = var.client_slug
  contract_type = "saas"
}

# ---------------------------------------------------------------------------
# 2. Utilisateurs
# ---------------------------------------------------------------------------
resource "aisia_user" "admin" {
  email        = "admin@${var.client_slug}.example"
  display_name = "Admin ${var.client_name}"
  role         = "org_admin"
  org_id       = aisia_organization.client.id
  # password non fourni → AISIA en genere un (expose une seule fois dans generated_password)
}

resource "aisia_user" "developer" {
  email        = "dev@${var.client_slug}.example"
  display_name = "Developer ${var.client_name}"
  role         = "org_user"
  org_id       = aisia_organization.client.id
}

# ---------------------------------------------------------------------------
# 3. Cle provider LLM (isolee par org — isolation KEY-2)
# ---------------------------------------------------------------------------
resource "aisia_provider_key" "openai" {
  org_id        = aisia_organization.client.id
  provider_id   = "openai"
  key_value     = var.openai_api_key
  account_email = "billing@${var.client_slug}.example"
}

# ---------------------------------------------------------------------------
# 4. Cle API programmatique (CI/CD)
# ---------------------------------------------------------------------------
resource "aisia_api_key" "ci" {
  org_id         = aisia_organization.client.id
  name           = "${var.client_slug}-ci-pipeline"
  scopes         = ["invoke", "compare"]
  rate_limit_rpm = 60
  rate_limit_rpd = 5000
}

# ---------------------------------------------------------------------------
# 5. Guardrail de conformite (global — s'applique a toutes les orgs)
# ---------------------------------------------------------------------------
resource "aisia_guardrail" "no_pii_card" {
  name    = "no-credit-card"
  pattern = "\\b(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13})\\b"
  action  = "block"
  enabled = true
}

# ---------------------------------------------------------------------------
# 6. Webhook de monitoring (optionnel — cree seulement si URL fournie)
# ---------------------------------------------------------------------------
resource "aisia_webhook" "monitoring" {
  name   = "${var.client_slug}-monitoring"
  url    = "https://monitoring.${var.client_slug}.example/aisia/events"
  events = ["invoke.failed", "org.created", "org.deleted"]
  secret = var.webhook_secret
}

# ---------------------------------------------------------------------------
# Sorties (sensibles — a stocker dans votre vault)
# ---------------------------------------------------------------------------
output "organization_id" {
  value       = aisia_organization.client.id
  description = "ID de l'organisation cliente."
}

output "admin_password" {
  value       = aisia_user.admin.generated_password
  sensitive   = true
  description = "Mot de passe auto-genere pour l'admin (one-shot — sauvegardez-le)."
}

output "ci_api_key" {
  value       = aisia_api_key.ci.key
  sensitive   = true
  description = "Cle API CI (one-shot — sauvegardez-la immediatement)."
}

output "ci_key_prefix" {
  value       = aisia_api_key.ci.key_prefix
  description = "Prefixe public de la cle CI (pour identification sans exposer la cle)."
}
