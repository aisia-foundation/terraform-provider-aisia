terraform {
  required_providers {
    aisia = {
      source  = "AISIA-fr/aisia"
      version = "~> 6.9"
    }
  }
}

provider "aisia" {
  endpoint = "https://api.aisia.fr" # ou env AISIA_ENDPOINT
  token    = var.aisia_token        # préférez env AISIA_TOKEN (sensible)
}

variable "aisia_token" {
  type      = string
  sensitive = true
  default   = null
}
