# Lister tous les modeles locaux disponibles dans AISIA
data "aisia_local_models" "all" {}

# Exploiter le catalogue via jsondecode
locals {
  models = jsondecode(data.aisia_local_models.all.json)
}

output "local_model_count" {
  value = data.aisia_local_models.all.item_count
}

output "model_ids" {
  value = data.aisia_local_models.all.ids
}
