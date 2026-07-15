// Code généré par scripts/generate/06_terraform_provider.py — NE PAS ÉDITER À LA MAIN.
// Régénérer : python3 scripts/generate/06_terraform_provider.py --openapi <spec>

package provider

import "github.com/hashicorp/terraform-plugin-framework/resource"

// generatedResources : resources WRITE générées depuis l'OpenAPI.
var generatedResources = []func() resource.Resource{
	func() resource.Resource {
		return &apiResource{name: "admin_ai_act_model_cards", path: "/admin/ai-act/model-cards", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "org_id", TFType: "string", Required: false, Desc: "champ org_id"}, {Name: "provider_id", TFType: "string", Required: false, Desc: "champ provider_id"}, {Name: "model_name", TFType: "string", Required: false, Desc: "champ model_name"}, {Name: "intended_use", TFType: "string", Required: false, Desc: "champ intended_use"}, {Name: "out_of_scope_use", TFType: "string", Required: false, Desc: "champ out_of_scope_use"}, {Name: "training_data", TFType: "string", Required: false, Desc: "champ training_data"}, {Name: "evaluation_data", TFType: "string", Required: false, Desc: "champ evaluation_data"}, {Name: "ethical_considerations", TFType: "string", Required: false, Desc: "champ ethical_considerations"}, {Name: "known_limitations", TFType: "string", Required: false, Desc: "champ known_limitations"}, {Name: "biases_identified", TFType: "string", Required: false, Desc: "champ biases_identified"}, {Name: "risk_level", TFType: "string", Required: false, Desc: "champ risk_level"}, {Name: "annexe_iii_category", TFType: "string", Required: false, Desc: "champ annexe_iii_category"}, {Name: "published", TFType: "bool", Required: false, Desc: "champ published"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_ai_act_v2_incidents", path: "/admin/ai-act/v2/incidents", updateVerb: "", canDelete: false, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "severity", TFType: "string", Required: false, Desc: "champ severity"}, {Name: "category", TFType: "string", Required: false, Desc: "champ category"}, {Name: "description", TFType: "string", Required: false, Desc: "champ description"}, {Name: "org_id", TFType: "string", Required: false, Desc: "champ org_id"}, {Name: "impact_assessment", TFType: "string", Required: false, Desc: "champ impact_assessment"}, {Name: "metadata", TFType: "string", Required: false, Desc: "champ metadata"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_alerts", path: "/admin/alerts", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_blog_posts", path: "/admin/blog/posts", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "id", TFType: "string", Required: false, Desc: "champ id"}, {Name: "slug", TFType: "string", Required: false, Desc: "champ slug"}, {Name: "title", TFType: "string", Required: false, Desc: "champ title"}, {Name: "content_md", TFType: "string", Required: false, Desc: "champ content_md"}, {Name: "excerpt", TFType: "string", Required: false, Desc: "champ excerpt"}, {Name: "tags", TFType: "list", Required: false, Desc: "champ tags"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_bot_conversations", path: "/admin/bot/conversations", updateVerb: "", canDelete: false, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_campaigns", path: "/admin/campaigns", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "name", TFType: "string", Required: true, Desc: "champ name"}, {Name: "channel", TFType: "string", Required: true, Desc: "champ channel"}, {Name: "target_audience", TFType: "string", Required: false, Desc: "champ target_audience"}, {Name: "content", TFType: "string", Required: true, Desc: "champ content"}, {Name: "scheduled_at", TFType: "string", Required: false, Desc: "champ scheduled_at"}, {Name: "status", TFType: "string", Required: false, Desc: "champ status"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_cloud_providers_item", path: "/admin/cloud-providers", updateVerb: "PUT", canDelete: false, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_config_email_templates_item", path: "/admin/config/email/templates", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "name", TFType: "string", Required: false, Desc: "champ name"}, {Name: "subject", TFType: "string", Required: false, Desc: "champ subject"}, {Name: "html", TFType: "string", Required: false, Desc: "champ html"}, {Name: "text", TFType: "string", Required: false, Desc: "champ text"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_config_snapshots", path: "/admin/config/snapshots", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_connectors_catalog_item", path: "/admin/connectors-catalog", updateVerb: "PUT", canDelete: false, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_connectors_mcp_servers", path: "/admin/connectors/mcp/servers", updateVerb: "", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "server_id", TFType: "string", Required: true, Desc: "champ server_id"}, {Name: "url", TFType: "string", Required: true, Desc: "champ url"}, {Name: "headers", TFType: "string", Required: false, Desc: "champ headers"}, {Name: "org_id", TFType: "string", Required: false, Desc: "champ org_id"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_credentials_item", path: "/admin/credentials", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_db_migrations", path: "/admin/db/migrations", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_discovery_scans", path: "/admin/discovery/scans", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_dns_records", path: "/admin/dns/records", updateVerb: "", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "field_type", TFType: "string", Required: false, Desc: "champ field_type"}, {Name: "force_acme_subdomain", TFType: "bool", Required: false, Desc: "champ force_acme_subdomain"}, {Name: "sub_domain", TFType: "string", Required: false, Desc: "champ sub_domain"}, {Name: "target", TFType: "string", Required: false, Desc: "champ target"}, {Name: "ttl", TFType: "int", Required: false, Desc: "champ ttl"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_dr_drills", path: "/admin/dr/drills", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_email_templates", path: "/admin/email-templates", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_events", path: "/admin/events", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "type", TFType: "string", Required: true, Desc: "champ type"}, {Name: "title", TFType: "string", Required: true, Desc: "champ title"}, {Name: "description", TFType: "string", Required: false, Desc: "champ description"}, {Name: "host_email", TFType: "string", Required: false, Desc: "champ host_email"}, {Name: "scheduled_at", TFType: "string", Required: true, Desc: "champ scheduled_at"}, {Name: "duration_min", TFType: "int", Required: false, Desc: "champ duration_min"}, {Name: "max_attendees", TFType: "int", Required: false, Desc: "champ max_attendees"}, {Name: "registration_url", TFType: "string", Required: false, Desc: "champ registration_url"}, {Name: "status", TFType: "string", Required: false, Desc: "champ status"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_finetune_jobs", path: "/admin/finetune/jobs", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_gdpr_requests", path: "/admin/gdpr/requests", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_groups", path: "/admin/groups", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "name", TFType: "string", Required: true, Desc: "champ name"}, {Name: "description", TFType: "string", Required: false, Desc: "champ description"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_hierarchy_nodes", path: "/admin/hierarchy/nodes", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_i18n_locales", path: "/admin/i18n/locales", updateVerb: "PUT", canDelete: false, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_leads_item", path: "/admin/leads", updateVerb: "PUT", canDelete: false, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_learning_finetune_runs", path: "/admin/learning/finetune-runs", updateVerb: "", canDelete: false, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "base_model", TFType: "string", Required: true, Desc: "champ base_model"}, {Name: "dataset", TFType: "string", Required: true, Desc: "champ dataset"}, {Name: "epochs", TFType: "int", Required: false, Desc: "champ epochs"}, {Name: "lr", TFType: "float", Required: false, Desc: "champ lr"}, {Name: "notes", TFType: "string", Required: false, Desc: "champ notes"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_oidc_providers", path: "/admin/oidc-providers", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "provider_id", TFType: "string", Required: false, Desc: "champ provider_id"}, {Name: "enabled", TFType: "bool", Required: false, Desc: "champ enabled"}, {Name: "config", TFType: "string", Required: false, Desc: "champ config"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_orgs_quotas_item", path: "/admin/orgs/quotas", updateVerb: "PUT", canDelete: false, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_plans", path: "/admin/plans", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_provider_pricing_item", path: "/admin/provider-pricing", updateVerb: "PUT", canDelete: false, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_seo_feeds", path: "/admin/seo/feeds", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "id", TFType: "string", Required: false, Desc: "champ id"}, {Name: "category", TFType: "string", Required: false, Desc: "champ category"}, {Name: "url", TFType: "string", Required: false, Desc: "champ url"}, {Name: "name", TFType: "string", Required: false, Desc: "champ name"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_skills", path: "/admin/skills", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "name", TFType: "string", Required: true, Desc: "champ name"}, {Name: "description", TFType: "string", Required: false, Desc: "champ description"}, {Name: "category", TFType: "string", Required: false, Desc: "champ category"}, {Name: "enabled", TFType: "bool", Required: false, Desc: "champ enabled"}, {Name: "invocation_type", TFType: "string", Required: false, Desc: "champ invocation_type"}, {Name: "payload", TFType: "string", Required: false, Desc: "champ payload"}, {Name: "org_id", TFType: "string", Required: false, Desc: "champ org_id"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_token_quotas", path: "/admin/token-quotas", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "org_agents", path: "/org/agents", updateVerb: "PUT", canDelete: true, desc: "Endpoint API AISIA — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "allowed_tools", TFType: "list", Required: false, Desc: "champ allowed_tools"}, {Name: "description", TFType: "string", Required: false, Desc: "champ description"}, {Name: "max_steps", TFType: "int", Required: false, Desc: "champ max_steps"}, {Name: "max_tokens", TFType: "int", Required: false, Desc: "champ max_tokens"}, {Name: "model", TFType: "string", Required: false, Desc: "champ model"}, {Name: "model_id", TFType: "string", Required: false, Desc: "champ model_id"}, {Name: "name", TFType: "string", Required: false, Desc: "champ name"}, {Name: "schedule_interval_s", TFType: "int", Required: false, Desc: "champ schedule_interval_s"}, {Name: "system_prompt", TFType: "string", Required: false, Desc: "champ system_prompt"}, {Name: "temperature", TFType: "float", Required: false, Desc: "champ temperature"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "org_api_keys", path: "/org/api-keys", updateVerb: "", canDelete: true, desc: "Endpoint API AISIA — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "expires_at", TFType: "string", Required: false, Desc: "champ expires_at"}, {Name: "name", TFType: "string", Required: false, Desc: "champ name"}, {Name: "rate_limit_rpd", TFType: "int", Required: false, Desc: "champ rate_limit_rpd"}, {Name: "rate_limit_rpm", TFType: "int", Required: false, Desc: "champ rate_limit_rpm"}, {Name: "scopes", TFType: "list", Required: false, Desc: "champ scopes"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "org_groups", path: "/org/groups", updateVerb: "", canDelete: false, desc: "Endpoint API AISIA — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "description", TFType: "string", Required: false, Desc: "champ description"}, {Name: "name", TFType: "string", Required: false, Desc: "champ name"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "org_support_tickets", path: "/org/support/tickets", updateVerb: "", canDelete: false, desc: "Endpoint API AISIA — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "category", TFType: "string", Required: false, Desc: "champ category"}, {Name: "description", TFType: "string", Required: false, Desc: "champ description"}, {Name: "priority", TFType: "string", Required: false, Desc: "champ priority"}, {Name: "subject", TFType: "string", Required: false, Desc: "champ subject"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "org_webhooks", path: "/org/webhooks", updateVerb: "", canDelete: true, desc: "Endpoint API AISIA — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "active", TFType: "bool", Required: false, Desc: "champ active"}, {Name: "events", TFType: "list", Required: false, Desc: "champ events"}, {Name: "secret", TFType: "string", Required: false, Desc: "champ secret"}, {Name: "url", TFType: "string", Required: false, Desc: "champ url"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "scim_v2_users", path: "/scim/v2/Users", updateVerb: "PUT", canDelete: true, desc: "Endpoint API AISIA — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "support_tickets", path: "/support/tickets", updateVerb: "", canDelete: false, desc: "Endpoint API AISIA — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "subject", TFType: "string", Required: true, Desc: "champ subject"}, {Name: "description", TFType: "string", Required: true, Desc: "champ description"}, {Name: "priority", TFType: "string", Required: false, Desc: "champ priority"}, {Name: "category", TFType: "string", Required: false, Desc: "champ category"}}}
	},
	func() resource.Resource {
		return &apiResource{name: "v1_conversations_item", path: "/v1/conversations", updateVerb: "PATCH", canDelete: true, desc: "Endpoint public API v1 — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "v1_debate", path: "/v1/debate", updateVerb: "", canDelete: false, desc: "Endpoint public API v1 — entité CRUD (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "v1_runs", path: "/v1/runs", updateVerb: "", canDelete: false, desc: "Endpoint public API v1 — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "name", TFType: "string", Required: true, Desc: "champ name"}, {Name: "prompts", TFType: "list", Required: true, Desc: "champ prompts"}, {Name: "providers", TFType: "list", Required: false, Desc: "champ providers"}, {Name: "iterations", TFType: "int", Required: false, Desc: "champ iterations"}, {Name: "concurrency", TFType: "int", Required: false, Desc: "champ concurrency"}, {Name: "timeout_s", TFType: "int", Required: false, Desc: "champ timeout_s"}}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_ai_rules_config", path: "/admin/ai-rules", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin)."}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_config_email_config", path: "/admin/config/email", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin)."}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_config_email_templates_config", path: "/admin/config/email/templates", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin)."}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_config_identity_config", path: "/admin/config/identity", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin)."}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_config_saml_config", path: "/admin/config/saml", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin)."}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_routing_policy_config", path: "/admin/routing/policy", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_cycle_run_action", path: "/admin/agent/cycle/run", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/agent/cycle/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_email_run_action", path: "/admin/agent/email/run", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/agent/email/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_orchestration_toggle_action", path: "/admin/agent/orchestration/toggle", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/agent/orchestration/toggle` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_scheduler_reload_action", path: "/admin/agent/scheduler/reload", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/agent/scheduler/reload` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_tasks_task_id_approve_action", path: "", pathTemplate: "/admin/agent/tasks/{task_id}/approve", method: "POST", readPath: "/admin/agent/tasks", desc: "Mutation admin POST `/admin/agent/tasks/{task_id}/approve` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_tasks_task_id_reject_action", path: "", pathTemplate: "/admin/agent/tasks/{task_id}/reject", method: "POST", readPath: "/admin/agent/tasks", desc: "Mutation admin POST `/admin/agent/tasks/{task_id}/reject` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agents_reload_custom_action", path: "/admin/agents/reload-custom", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/agents/reload-custom` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agents_code_trigger_action", path: "", pathTemplate: "/admin/agents/{code}/trigger", method: "POST", readPath: "/admin/agents", desc: "Mutation admin POST `/admin/agents/{code}/trigger` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_annexe_iii_classifications_action", path: "/admin/ai-act/annexe-iii-classifications", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/annexe-iii-classifications", desc: "Mutation admin POST `/admin/ai-act/annexe-iii-classifications` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_risk_assessments_action", path: "/admin/ai-act/risk-assessments", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/risk-assessments", desc: "Mutation admin POST `/admin/ai-act/risk-assessments` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_audit_purge_action", path: "/admin/ai-act/v2/audit/purge", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ai-act/v2/audit/purge` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_export_package_action", path: "/admin/ai-act/v2/export-package", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ai-act/v2/export-package` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_incidents_incident_id_report_authority_action", path: "", pathTemplate: "/admin/ai-act/v2/incidents/{incident_id}/report-authority", method: "POST", readPath: "/admin/ai-act/qms", desc: "Mutation admin POST `/admin/ai-act/v2/incidents/{incident_id}/report-authority` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_kill_switch_action", path: "/admin/ai-act/v2/kill-switch", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/v2/kill-switch", desc: "Mutation admin POST `/admin/ai-act/v2/kill-switch` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_model_cards_seed_action", path: "/admin/ai-act/v2/model-cards/seed", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ai-act/v2/model-cards/seed` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_oversight_oversight_id_approve_action", path: "", pathTemplate: "/admin/ai-act/v2/oversight/{oversight_id}/approve", method: "POST", readPath: "/admin/ai-act/qms", desc: "Mutation admin POST `/admin/ai-act/v2/oversight/{oversight_id}/approve` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_oversight_oversight_id_reject_action", path: "", pathTemplate: "/admin/ai-act/v2/oversight/{oversight_id}/reject", method: "POST", readPath: "/admin/ai-act/qms", desc: "Mutation admin POST `/admin/ai-act/v2/oversight/{oversight_id}/reject` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_watermark_add_action", path: "/admin/ai-act/v2/watermark/add", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ai-act/v2/watermark/add` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_watermark_detect_action", path: "/admin/ai-act/v2/watermark/detect", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ai-act/v2/watermark/detect` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_rules_reset_action", path: "/admin/ai-rules/reset", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ai-rules/reset` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_alerts_alert_id_ack_action", path: "", pathTemplate: "/admin/alerts/{alert_id}/ack", method: "POST", readPath: "/admin/alerts", desc: "Mutation admin POST `/admin/alerts/{alert_id}/ack` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_analytics_action", path: "/admin/analytics", pathTemplate: "", method: "POST", readPath: "/admin/analytics", desc: "Mutation admin POST `/admin/analytics` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_api_keys_rotation_action", path: "/admin/api-keys/rotation", pathTemplate: "", method: "POST", readPath: "/admin/api-keys/rotation", desc: "Mutation admin POST `/admin/api-keys/rotation` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_api_keys_validate_all_action", path: "/admin/api-keys/validate-all", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/api-keys/validate-all` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_api_keys_provider_id_mutation", path: "", pathTemplate: "/admin/api-keys/{provider_id}", method: "DELETE", readPath: "/admin/api-keys", desc: "Mutation admin DELETE `/admin/api-keys/{provider_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_auto_test_action", path: "/admin/auto-test", pathTemplate: "", method: "POST", readPath: "/admin/auto-test", desc: "Mutation admin POST `/admin/auto-test` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_autonomy_reload_action", path: "/admin/autonomy/reload", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/autonomy/reload` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_b2b_approval_request_id_approve_action", path: "", pathTemplate: "/admin/b2b/approval/{request_id}/approve", method: "POST", readPath: "/admin/b2b/approval/queue", desc: "Mutation admin POST `/admin/b2b/approval/{request_id}/approve` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_b2b_approval_request_id_reject_action", path: "", pathTemplate: "/admin/b2b/approval/{request_id}/reject", method: "POST", readPath: "/admin/b2b/approval/queue", desc: "Mutation admin POST `/admin/b2b/approval/{request_id}/reject` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_backups_btype_create_action", path: "", pathTemplate: "/admin/backups/{btype}/create", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/backups/{btype}/create` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_backups_btype_restore_action", path: "", pathTemplate: "/admin/backups/{btype}/restore", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/backups/{btype}/restore` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_billing_portal_session_action", path: "/admin/billing/portal-session", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/billing/portal-session` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_blog_posts_post_id_publish_action", path: "", pathTemplate: "/admin/blog/posts/{post_id}/publish", method: "POST", readPath: "/admin/blog/posts", desc: "Mutation admin POST `/admin/blog/posts/{post_id}/publish` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_blacklist_reset_action", path: "/admin/bot/blacklist/reset", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/bot/blacklist/reset` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_credential_rotation_candidate_id_dismiss_action", path: "", pathTemplate: "/admin/bot/credential-rotation/{candidate_id}/dismiss", method: "POST", readPath: "/admin/bot/status", desc: "Mutation admin POST `/admin/bot/credential-rotation/{candidate_id}/dismiss` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_force_task_action", path: "/admin/bot/force-task", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/bot/force-task` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_run_now_action", path: "/admin/bot/run-now", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/bot/run-now` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_start_action", path: "/admin/bot/start", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/bot/start` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_stop_action", path: "/admin/bot/stop", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/bot/stop` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cache_preload_action", path: "/admin/cache/preload", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/cache/preload` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cache_purge_action", path: "/admin/cache/purge", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/cache/purge` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_campaigns_campaign_id_email_blast_action", path: "", pathTemplate: "/admin/campaigns/{campaign_id}/email-blast", method: "POST", readPath: "/admin/campaigns/", desc: "Mutation admin POST `/admin/campaigns/{campaign_id}/email-blast` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_campaigns_campaign_id_send_email_action", path: "", pathTemplate: "/admin/campaigns/{campaign_id}/send-email", method: "POST", readPath: "/admin/campaigns/", desc: "Mutation admin POST `/admin/campaigns/{campaign_id}/send-email` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_campaigns_campaign_id_status_mutation", path: "", pathTemplate: "/admin/campaigns/{campaign_id}/status", method: "PUT", readPath: "/admin/campaigns/", desc: "Mutation admin PUT `/admin/campaigns/{campaign_id}/status` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_catalog_sync_action", path: "/admin/catalog/sync", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/catalog/sync` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_certificates_issue_action", path: "/admin/certificates/issue", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/certificates/issue` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_certificates_domain_mutation", path: "", pathTemplate: "/admin/certificates/{domain}", method: "DELETE", readPath: "/admin/certificates", desc: "Mutation admin DELETE `/admin/certificates/{domain}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_certificates_domain_renew_action", path: "", pathTemplate: "/admin/certificates/{domain}/renew", method: "POST", readPath: "/admin/certificates", desc: "Mutation admin POST `/admin/certificates/{domain}/renew` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_circuit_breakers_provider_id_reset_action", path: "", pathTemplate: "/admin/circuit-breakers/{provider_id}/reset", method: "POST", readPath: "/admin/circuit-breakers", desc: "Mutation admin POST `/admin/circuit-breakers/{provider_id}/reset` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cloud_models_provider_id_model_name_enabled_mutation", path: "", pathTemplate: "/admin/cloud-models/{provider_id}/{model_name}/enabled", method: "PUT", readPath: "/admin/cloud-models", desc: "Mutation admin PUT `/admin/cloud-models/{provider_id}/{model_name}/enabled` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cloud_providers_provider_id_provision_action", path: "", pathTemplate: "/admin/cloud-providers/{provider_id}/provision", method: "POST", readPath: "/admin/cloud-providers", desc: "Mutation admin POST `/admin/cloud-providers/{provider_id}/provision` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cloud_providers_provider_id_test_action", path: "", pathTemplate: "/admin/cloud-providers/{provider_id}/test", method: "POST", readPath: "/admin/cloud-providers", desc: "Mutation admin POST `/admin/cloud-providers/{provider_id}/test` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cluster_node_drain_action", path: "/admin/cluster/node/drain", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/cluster/node/drain` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cluster_prune_action", path: "/admin/cluster/prune", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/cluster/prune` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cluster_scale_action", path: "/admin/cluster/scale", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/cluster/scale` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_compliance_action", path: "/admin/compliance", pathTemplate: "", method: "POST", readPath: "/admin/compliance", desc: "Mutation admin POST `/admin/compliance` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_email_test_action", path: "/admin/config/email/test", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/config/email/test` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_identity_reload_action", path: "/admin/config/identity/reload", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/config/identity/reload` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_import_action", path: "/admin/config/import", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/config/import` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_catalog_import_openapi_action", path: "/admin/connectors-catalog/import-openapi", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/connectors-catalog/import-openapi` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_mcp_servers_server_id_call_action", path: "", pathTemplate: "/admin/connectors/mcp/servers/{server_id}/call", method: "POST", readPath: "/admin/connectors/manifests", desc: "Mutation admin POST `/admin/connectors/mcp/servers/{server_id}/call` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_schema_mapping_apply_action", path: "/admin/connectors/schema-mapping/apply", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/connectors/schema-mapping/apply` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_schema_mapping_validate_action", path: "/admin/connectors/schema-mapping/validate", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/connectors/schema-mapping/validate` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_run_action", path: "/admin/connectors/wizard/run", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/connectors/wizard/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_start_action", path: "/admin/connectors/wizard/start", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/connectors/wizard/start` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_session_id_activate_action", path: "", pathTemplate: "/admin/connectors/wizard/{session_id}/activate", method: "POST", readPath: "/admin/connectors/manifests", desc: "Mutation admin POST `/admin/connectors/wizard/{session_id}/activate` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_session_id_auth_action", path: "", pathTemplate: "/admin/connectors/wizard/{session_id}/auth", method: "POST", readPath: "/admin/connectors/manifests", desc: "Mutation admin POST `/admin/connectors/wizard/{session_id}/auth` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_session_id_mapping_action", path: "", pathTemplate: "/admin/connectors/wizard/{session_id}/mapping", method: "POST", readPath: "/admin/connectors/manifests", desc: "Mutation admin POST `/admin/connectors/wizard/{session_id}/mapping` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_session_id_test_action", path: "", pathTemplate: "/admin/connectors/wizard/{session_id}/test", method: "POST", readPath: "/admin/connectors/manifests", desc: "Mutation admin POST `/admin/connectors/wizard/{session_id}/test` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_connector_id_vendor_updates_action", path: "", pathTemplate: "/admin/connectors/{connector_id}/vendor-updates", method: "POST", readPath: "/admin/connectors/manifests", desc: "Mutation admin POST `/admin/connectors/{connector_id}/vendor-updates` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_crews_orphans_drain_action", path: "/admin/crews/orphans/drain", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/crews/orphans/drain` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_crews_runs_run_id_approve_action", path: "", pathTemplate: "/admin/crews/runs/{run_id}/approve", method: "POST", readPath: "/admin/crews", desc: "Mutation admin POST `/admin/crews/runs/{run_id}/approve` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_crews_name_run_action", path: "", pathTemplate: "/admin/crews/{name}/run", method: "POST", readPath: "/admin/crews", desc: "Mutation admin POST `/admin/crews/{name}/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_database_migrations_name_apply_action", path: "", pathTemplate: "/admin/database/migrations/{name}/apply", method: "POST", readPath: "/admin/database/info", desc: "Mutation admin POST `/admin/database/migrations/{name}/apply` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_cleanup_action", path: "/admin/datasets/cleanup", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/datasets/cleanup` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_download_action", path: "/admin/datasets/download", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/datasets/download` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_dataset_id_mutation", path: "", pathTemplate: "/admin/datasets/{dataset_id}", method: "DELETE", readPath: "/admin/datasets", desc: "Mutation admin DELETE `/admin/datasets/{dataset_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_dataset_id_redownload_action", path: "", pathTemplate: "/admin/datasets/{dataset_id}/redownload", method: "POST", readPath: "/admin/datasets", desc: "Mutation admin POST `/admin/datasets/{dataset_id}/redownload` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_dataset_id_reingest_action", path: "", pathTemplate: "/admin/datasets/{dataset_id}/reingest", method: "POST", readPath: "/admin/datasets", desc: "Mutation admin POST `/admin/datasets/{dataset_id}/reingest` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_deploy_worker_run_once_action", path: "/admin/deploy/worker/run-once", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/deploy/worker/run-once` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_discovery_run_action", path: "/admin/discovery/run", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/discovery/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_discovery_runs_action", path: "/admin/discovery/runs", pathTemplate: "", method: "POST", readPath: "/admin/discovery/runs", desc: "Mutation admin POST `/admin/discovery/runs` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_dr_drill_action", path: "/admin/dr/drill", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/dr/drill` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_dr_drills_run_action", path: "/admin/dr/drills/run", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/dr/drills/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_egress_test_action", path: "/admin/egress/test", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/egress/test` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_events_event_id_registrations_mark_attended_action", path: "", pathTemplate: "/admin/events/{event_id}/registrations/mark-attended", method: "POST", readPath: "/admin/events/", desc: "Mutation admin POST `/admin/events/{event_id}/registrations/mark-attended` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_feature_flags_action", path: "/admin/feature-flags", pathTemplate: "", method: "POST", readPath: "/admin/feature-flags", desc: "Mutation admin POST `/admin/feature-flags` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_finetune_action", path: "/admin/finetune", pathTemplate: "", method: "POST", readPath: "/admin/finetune", desc: "Mutation admin POST `/admin/finetune` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_delete_action", path: "/admin/gdpr/delete", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/gdpr/delete` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_export_action", path: "/admin/gdpr/export", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/gdpr/export` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_requests_delete_action", path: "/admin/gdpr/requests/delete", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/gdpr/requests/delete` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_requests_export_action", path: "/admin/gdpr/requests/export", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/gdpr/requests/export` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_groups_group_id_members_mutation", path: "", pathTemplate: "/admin/groups/{group_id}/members", method: "PUT", readPath: "/admin/groups/", desc: "Mutation admin PUT `/admin/groups/{group_id}/members` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_groups_group_id_members_user_id_mutation", path: "", pathTemplate: "/admin/groups/{group_id}/members/{user_id}", method: "DELETE", readPath: "/admin/groups/", desc: "Mutation admin DELETE `/admin/groups/{group_id}/members/{user_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_groups_group_id_permissions_mutation", path: "", pathTemplate: "/admin/groups/{group_id}/permissions", method: "PUT", readPath: "/admin/groups/", desc: "Mutation admin PUT `/admin/groups/{group_id}/permissions` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_guardrails_gid_mutation", path: "", pathTemplate: "/admin/guardrails/{gid}", method: "DELETE", readPath: "/admin/guardrails", desc: "Mutation admin DELETE `/admin/guardrails/{gid}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_healthcheck_run_action", path: "/admin/healthcheck/run", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/healthcheck/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_integrations_integration_id_test_action", path: "", pathTemplate: "/admin/integrations/{integration_id}/test", method: "POST", readPath: "/admin/integrations/status", desc: "Mutation admin POST `/admin/integrations/{integration_id}/test` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_integrations_provider_adapter_test_action", path: "", pathTemplate: "/admin/integrations/{provider}/adapter-test", method: "POST", readPath: "/admin/integrations/status", desc: "Mutation admin POST `/admin/integrations/{provider}/adapter-test` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_investors_ndas_nda_id_grant_action", path: "", pathTemplate: "/admin/investors/ndas/{nda_id}/grant", method: "POST", readPath: "/admin/investors", desc: "Mutation admin POST `/admin/investors/ndas/{nda_id}/grant` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_investors_ndas_nda_id_revoke_action", path: "", pathTemplate: "/admin/investors/ndas/{nda_id}/revoke", method: "POST", readPath: "/admin/investors", desc: "Mutation admin POST `/admin/investors/ndas/{nda_id}/revoke` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_investors_investor_id_send_nda_action", path: "", pathTemplate: "/admin/investors/{investor_id}/send-nda", method: "POST", readPath: "/admin/investors", desc: "Mutation admin POST `/admin/investors/{investor_id}/send-nda` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_knowledge_clear_action", path: "/admin/knowledge/clear", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/knowledge/clear` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_knowledge_ingest_action", path: "/admin/knowledge/ingest", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/knowledge/ingest` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_learning_finetune_action", path: "/admin/learning/finetune", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/learning/finetune` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_learning_proposals_proposal_id_apply_action", path: "", pathTemplate: "/admin/learning/proposals/{proposal_id}/apply", method: "POST", readPath: "/admin/learning/status", desc: "Mutation admin POST `/admin/learning/proposals/{proposal_id}/apply` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_learning_run_action", path: "/admin/learning/run", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/learning/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_licenses_action", path: "/admin/licenses", pathTemplate: "", method: "POST", readPath: "/admin/licenses", desc: "Mutation admin POST `/admin/licenses` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_linkedin_post_action", path: "/admin/linkedin/post", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/linkedin/post` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_local_models_model_id_mutation", path: "", pathTemplate: "/admin/local-models/{model_id}", method: "DELETE", readPath: "/admin/local-models", desc: "Mutation admin DELETE `/admin/local-models/{model_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_local_models_model_id_enabled_mutation", path: "", pathTemplate: "/admin/local-models/{model_id}/enabled", method: "PUT", readPath: "/admin/local-models", desc: "Mutation admin PUT `/admin/local-models/{model_id}/enabled` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_maintenance_reload_action", path: "/admin/maintenance/reload", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/maintenance/reload` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_maintenance_tasks_task_id_enable_action", path: "", pathTemplate: "/admin/maintenance/tasks/{task_id}/enable", method: "POST", readPath: "/admin/maintenance/status", desc: "Mutation admin POST `/admin/maintenance/tasks/{task_id}/enable` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_maintenance_tasks_task_id_run_action", path: "", pathTemplate: "/admin/maintenance/tasks/{task_id}/run", method: "POST", readPath: "/admin/maintenance/status", desc: "Mutation admin POST `/admin/maintenance/tasks/{task_id}/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_models_sync_action", path: "/admin/models/sync", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/models/sync` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_network_action", path: "/admin/network", pathTemplate: "", method: "POST", readPath: "/admin/network", desc: "Mutation admin POST `/admin/network` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ollama_models_model_name_mutation", path: "", pathTemplate: "/admin/ollama/models/{model_name}", method: "DELETE", readPath: "/admin/ollama/models", desc: "Mutation admin DELETE `/admin/ollama/models/{model_name}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ollama_pull_action", path: "/admin/ollama/pull", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ollama/pull` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_onboarding_onboarding_id_review_mutation", path: "", pathTemplate: "/admin/onboarding/{onboarding_id}/review", method: "PUT", readPath: "/admin/onboarding", desc: "Mutation admin PUT `/admin/onboarding/{onboarding_id}/review` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_backup_action", path: "/admin/ops/backup", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ops/backup` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_build_action", path: "/admin/ops/build", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ops/build` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_deploy_action", path: "/admin/ops/deploy", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ops/deploy` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_rollback_action", path: "/admin/ops/rollback", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ops/rollback` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_scale_action", path: "/admin/ops/scale", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/ops/scale` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}", method: "DELETE", readPath: "/admin/organizations/", desc: "Mutation admin DELETE `/admin/organizations/{org_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_approve_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/approve", method: "PUT", readPath: "/admin/organizations/", desc: "Mutation admin PUT `/admin/organizations/{org_id}/approve` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_branding_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/branding", method: "PUT", readPath: "/admin/organizations/", desc: "Mutation admin PUT `/admin/organizations/{org_id}/branding` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_licence_action", path: "", pathTemplate: "/admin/organizations/{org_id}/licence", method: "POST", readPath: "/admin/organizations/", desc: "Mutation admin POST `/admin/organizations/{org_id}/licence` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_llm_access_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/llm-access", method: "PUT", readPath: "/admin/organizations/", desc: "Mutation admin PUT `/admin/organizations/{org_id}/llm-access` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_members_action", path: "", pathTemplate: "/admin/organizations/{org_id}/members", method: "POST", readPath: "/admin/organizations/", desc: "Mutation admin POST `/admin/organizations/{org_id}/members` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_members_user_id_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/members/{user_id}", method: "DELETE", readPath: "/admin/organizations/", desc: "Mutation admin DELETE `/admin/organizations/{org_id}/members/{user_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_purge_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/purge", method: "DELETE", readPath: "/admin/organizations/", desc: "Mutation admin DELETE `/admin/organizations/{org_id}/purge` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_reactivate_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/reactivate", method: "PUT", readPath: "/admin/organizations/", desc: "Mutation admin PUT `/admin/organizations/{org_id}/reactivate` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_suspend_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/suspend", method: "PUT", readPath: "/admin/organizations/", desc: "Mutation admin PUT `/admin/organizations/{org_id}/suspend` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_api_keys_action", path: "", pathTemplate: "/admin/orgs/{org_id}/api-keys", method: "POST", readPath: "/admin/orgs", desc: "Mutation admin POST `/admin/orgs/{org_id}/api-keys` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_api_keys_key_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/api-keys/{key_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/api-keys/{key_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_branding_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/branding", method: "PUT", readPath: "/admin/orgs", desc: "Mutation admin PUT `/admin/orgs/{org_id}/branding` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_bundles_action", path: "", pathTemplate: "/admin/orgs/{org_id}/config/bundles", method: "POST", readPath: "/admin/orgs", desc: "Mutation admin POST `/admin/orgs/{org_id}/config/bundles` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_bundles_seed_defaults_action", path: "", pathTemplate: "/admin/orgs/{org_id}/config/bundles/seed-defaults", method: "POST", readPath: "/admin/orgs", desc: "Mutation admin POST `/admin/orgs/{org_id}/config/bundles/seed-defaults` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_bundles_bundle_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/bundles/{bundle_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/config/bundles/{bundle_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_bundles_bundle_id_assign_action", path: "", pathTemplate: "/admin/orgs/{org_id}/config/bundles/{bundle_id}/assign", method: "POST", readPath: "/admin/orgs", desc: "Mutation admin POST `/admin/orgs/{org_id}/config/bundles/{bundle_id}/assign` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_llm_access_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/llm-access", method: "PUT", readPath: "/admin/orgs", desc: "Mutation admin PUT `/admin/orgs/{org_id}/config/llm-access` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_llm_access_entry_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/llm-access/{entry_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/config/llm-access/{entry_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_quotas_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/quotas", method: "PUT", readPath: "/admin/orgs", desc: "Mutation admin PUT `/admin/orgs/{org_id}/config/quotas` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_quotas_entry_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/quotas/{entry_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/config/quotas/{entry_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_email_templates_key_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/email-templates/{key}", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/email-templates/{key}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_invitations_action", path: "", pathTemplate: "/admin/orgs/{org_id}/invitations", method: "POST", readPath: "/admin/orgs", desc: "Mutation admin POST `/admin/orgs/{org_id}/invitations` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_invitations_invitation_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/invitations/{invitation_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/invitations/{invitation_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_oidc_providers_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/oidc-providers", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/oidc-providers` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_provider_keys_provider_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/provider-keys/{provider_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/provider-keys/{provider_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_quota_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/quota", method: "PUT", readPath: "/admin/orgs", desc: "Mutation admin PUT `/admin/orgs/{org_id}/quota` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_routing_mode_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/routing-mode", method: "PUT", readPath: "/admin/orgs", desc: "Mutation admin PUT `/admin/orgs/{org_id}/routing-mode` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_security_policy_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/security-policy", method: "PUT", readPath: "/admin/orgs", desc: "Mutation admin PUT `/admin/orgs/{org_id}/security-policy` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_users_action", path: "", pathTemplate: "/admin/orgs/{org_id}/users", method: "POST", readPath: "/admin/orgs", desc: "Mutation admin POST `/admin/orgs/{org_id}/users` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_users_user_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/users/{user_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/users/{user_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_webhooks_action", path: "", pathTemplate: "/admin/orgs/{org_id}/webhooks", method: "POST", readPath: "/admin/orgs", desc: "Mutation admin POST `/admin/orgs/{org_id}/webhooks` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_webhooks_webhook_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/webhooks/{webhook_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Mutation admin DELETE `/admin/orgs/{org_id}/webhooks/{webhook_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_keys_rotate_cycle_action", path: "/admin/providers/keys/rotate-cycle", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/providers/keys/rotate-cycle` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_keys_verify_action", path: "/admin/providers/keys/verify", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/providers/keys/verify` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_mutation", path: "", pathTemplate: "/admin/providers/{provider_id}", method: "DELETE", readPath: "/admin/providers/config", desc: "Mutation admin DELETE `/admin/providers/{provider_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_auto_recover_action", path: "", pathTemplate: "/admin/providers/{provider_id}/auto-recover", method: "POST", readPath: "/admin/providers/config", desc: "Mutation admin POST `/admin/providers/{provider_id}/auto-recover` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_config_mutation", path: "", pathTemplate: "/admin/providers/{provider_id}/config", method: "PUT", readPath: "/admin/providers/config", desc: "Mutation admin PUT `/admin/providers/{provider_id}/config` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_enabled_mutation", path: "", pathTemplate: "/admin/providers/{provider_id}/enabled", method: "PUT", readPath: "/admin/providers/config", desc: "Mutation admin PUT `/admin/providers/{provider_id}/enabled` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_key_action", path: "", pathTemplate: "/admin/providers/{provider_id}/key", method: "POST", readPath: "/admin/providers/config", desc: "Mutation admin POST `/admin/providers/{provider_id}/key` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_key_rotate_action", path: "", pathTemplate: "/admin/providers/{provider_id}/key/rotate", method: "POST", readPath: "/admin/providers/config", desc: "Mutation admin POST `/admin/providers/{provider_id}/key/rotate` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_qdrant_bootstrap_action", path: "/admin/qdrant/bootstrap", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/qdrant/bootstrap` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_rate_limits_action", path: "/admin/rate-limits", pathTemplate: "", method: "POST", readPath: "/admin/rate-limits", desc: "Mutation admin POST `/admin/rate-limits` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_rbac_action", path: "/admin/rbac", pathTemplate: "", method: "POST", readPath: "/admin/rbac", desc: "Mutation admin POST `/admin/rbac` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_resellers_reseller_id_mutation", path: "", pathTemplate: "/admin/resellers/{reseller_id}", method: "DELETE", readPath: "/admin/resellers", desc: "Mutation admin DELETE `/admin/resellers/{reseller_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_scheduled_backups_action", path: "/admin/scheduled-backups", pathTemplate: "", method: "POST", readPath: "/admin/scheduled-backups", desc: "Mutation admin POST `/admin/scheduled-backups` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_self_repair_action", path: "/admin/self-repair", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/self-repair` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_comments_comment_id_approve_reply_action", path: "", pathTemplate: "/admin/seo/comments/{comment_id}/approve-reply", method: "POST", readPath: "/admin/seo/feeds", desc: "Mutation admin POST `/admin/seo/comments/{comment_id}/approve-reply` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_posts_post_id_approve_action", path: "", pathTemplate: "/admin/seo/posts/{post_id}/approve", method: "POST", readPath: "/admin/seo/feeds", desc: "Mutation admin POST `/admin/seo/posts/{post_id}/approve` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_posts_post_id_reject_action", path: "", pathTemplate: "/admin/seo/posts/{post_id}/reject", method: "POST", readPath: "/admin/seo/feeds", desc: "Mutation admin POST `/admin/seo/posts/{post_id}/reject` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_sessions_bulk_revoke_action", path: "/admin/sessions/bulk-revoke", pathTemplate: "", method: "POST", readPath: "/admin/sessions/bulk-revoke", desc: "Mutation admin POST `/admin/sessions/bulk-revoke` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_sessions_sid_revoke_action", path: "", pathTemplate: "/admin/sessions/{sid}/revoke", method: "POST", readPath: "/admin/sessions", desc: "Mutation admin POST `/admin/sessions/{sid}/revoke` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_stripe_checkout_action", path: "/admin/stripe/checkout", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/stripe/checkout` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_tests_run_action", path: "/admin/tests/run", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/tests/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_tests_suite_run_action", path: "", pathTemplate: "/admin/tests/{suite}/run", method: "POST", readPath: "/admin/tests/history", desc: "Mutation admin POST `/admin/tests/{suite}/run` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_mutation", path: "", pathTemplate: "/admin/users/{user_id}", method: "DELETE", readPath: "/admin/users", desc: "Mutation admin DELETE `/admin/users/{user_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_active_mutation", path: "", pathTemplate: "/admin/users/{user_id}/active", method: "PUT", readPath: "/admin/users", desc: "Mutation admin PUT `/admin/users/{user_id}/active` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_assign_org_action", path: "", pathTemplate: "/admin/users/{user_id}/assign-org", method: "POST", readPath: "/admin/users", desc: "Mutation admin POST `/admin/users/{user_id}/assign-org` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_display_name_mutation", path: "", pathTemplate: "/admin/users/{user_id}/display-name", method: "PUT", readPath: "/admin/users", desc: "Mutation admin PUT `/admin/users/{user_id}/display-name` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_force_logout_action", path: "", pathTemplate: "/admin/users/{user_id}/force-logout", method: "POST", readPath: "/admin/users", desc: "Mutation admin POST `/admin/users/{user_id}/force-logout` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_force_mfa_action", path: "", pathTemplate: "/admin/users/{user_id}/force-mfa", method: "POST", readPath: "/admin/users", desc: "Mutation admin POST `/admin/users/{user_id}/force-mfa` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_lock_action", path: "", pathTemplate: "/admin/users/{user_id}/lock", method: "POST", readPath: "/admin/users", desc: "Mutation admin POST `/admin/users/{user_id}/lock` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_permissions_mutation", path: "", pathTemplate: "/admin/users/{user_id}/permissions", method: "PUT", readPath: "/admin/users", desc: "Mutation admin PUT `/admin/users/{user_id}/permissions` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_profile_mutation", path: "", pathTemplate: "/admin/users/{user_id}/profile", method: "PUT", readPath: "/admin/users", desc: "Mutation admin PUT `/admin/users/{user_id}/profile` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_quota_mutation", path: "", pathTemplate: "/admin/users/{user_id}/quota", method: "PUT", readPath: "/admin/users", desc: "Mutation admin PUT `/admin/users/{user_id}/quota` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_reset_password_action", path: "", pathTemplate: "/admin/users/{user_id}/reset-password", method: "POST", readPath: "/admin/users", desc: "Mutation admin POST `/admin/users/{user_id}/reset-password` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_role_mutation", path: "", pathTemplate: "/admin/users/{user_id}/role", method: "PUT", readPath: "/admin/users", desc: "Mutation admin PUT `/admin/users/{user_id}/role` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_unlock_action", path: "", pathTemplate: "/admin/users/{user_id}/unlock", method: "POST", readPath: "/admin/users", desc: "Mutation admin POST `/admin/users/{user_id}/unlock` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_vault_action", path: "/admin/vault", pathTemplate: "", method: "POST", readPath: "/admin/vault", desc: "Mutation admin POST `/admin/vault` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhook_alert_action", path: "/admin/webhook/alert", pathTemplate: "", method: "POST", readPath: "", desc: "Mutation admin POST `/admin/webhook/alert` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_resend_action", path: "/admin/webhooks/resend", pathTemplate: "", method: "POST", readPath: "/admin/webhooks/resend", desc: "Mutation admin POST `/admin/webhooks/resend` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_webhook_id_mutation", path: "", pathTemplate: "/admin/webhooks/{webhook_id}", method: "DELETE", readPath: "/admin/webhooks", desc: "Mutation admin DELETE `/admin/webhooks/{webhook_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_webhook_id_test_action", path: "", pathTemplate: "/admin/webhooks/{webhook_id}/test", method: "POST", readPath: "/admin/webhooks", desc: "Mutation admin POST `/admin/webhooks/{webhook_id}/test` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_wh_id_mutation", path: "", pathTemplate: "/admin/webhooks/{wh_id}", method: "DELETE", readPath: "/admin/webhooks", desc: "Mutation admin DELETE `/admin/webhooks/{wh_id}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_wid_mutation", path: "", pathTemplate: "/admin/webhooks/{wid}", method: "DELETE", readPath: "/admin/webhooks", desc: "Mutation admin DELETE `/admin/webhooks/{wid}` (idempotent via Terraform, GOV-01)."}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_wid_test_action", path: "", pathTemplate: "/admin/webhooks/{wid}/test", method: "POST", readPath: "/admin/webhooks", desc: "Mutation admin POST `/admin/webhooks/{wid}/test` (idempotent via Terraform, GOV-01)."}
	},
}
