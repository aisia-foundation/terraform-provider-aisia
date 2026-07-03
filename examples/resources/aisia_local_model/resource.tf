# Modele haute capacite sur noeud GPU (priorite haute, routage local-first)
resource "aisia_local_model" "llama3_70b_gpu" {
  model_id = "llama3.3-70b-local"
  config_json = jsonencode({
    runtime    = "ollama"
    model_name = "llama3.3:70b"
    endpoint   = "http://gpu-node:11434"
    priority   = 120
    status     = "active"
    tier       = "gpu"
  })
}

# Modele leger sur noeud edge (priorite basse, fallback)
resource "aisia_local_model" "mistral_7b_edge" {
  model_id = "mistral-7b-edge"
  config_json = jsonencode({
    runtime    = "ollama"
    model_name = "mistral:7b"
    endpoint   = "http://edge-node:11434"
    priority   = 60
    status     = "active"
    tier       = "edge"
  })
}

output "gpu_model_status" {
  value = aisia_local_model.llama3_70b_gpu.status
}
