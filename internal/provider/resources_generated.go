// Code généré par scripts/gen_terraform_provider.py — NE PAS ÉDITER À LA MAIN.
// Régénérer : python3 scripts/gen_terraform_provider.py --openapi <spec>

package provider

import "github.com/hashicorp/terraform-plugin-framework/resource"

// generatedResources : resources CRUD générées depuis l'OpenAPI (apiResource générique).
var generatedResources = []func() resource.Resource{
	func() resource.Resource {
		return &apiResource{name: "admin_ai_act_model_cards", path: "/admin/ai-act/model-cards", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/ai-act/model-cards (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_alerts", path: "/admin/alerts", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/alerts (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_blog_posts", path: "/admin/blog/posts", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/blog/posts (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_bot_conversations", path: "/admin/bot/conversations", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/bot/conversations (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_config_snapshots", path: "/admin/config/snapshots", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/config/snapshots (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_db_migrations", path: "/admin/db/migrations", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/db/migrations (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_discovery_scans", path: "/admin/discovery/scans", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/discovery/scans (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_dns_records", path: "/admin/dns/records", updateVerb: "", canDelete: true, desc: "Entité CRUD /admin/dns/records (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_dr_drills", path: "/admin/dr/drills", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/dr/drills (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_email_templates", path: "/admin/email-templates", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/email-templates (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_finetune_jobs", path: "/admin/finetune/jobs", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/finetune/jobs (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_gdpr_requests", path: "/admin/gdpr/requests", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/gdpr/requests (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_hierarchy_nodes", path: "/admin/hierarchy/nodes", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/hierarchy/nodes (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_i18n_locales", path: "/admin/i18n/locales", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/i18n/locales (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_learning_finetune_runs", path: "/admin/learning/finetune-runs", updateVerb: "", canDelete: false, desc: "Entité CRUD /admin/learning/finetune-runs (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_oidc_providers", path: "/admin/oidc-providers", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/oidc-providers (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_seo_feeds", path: "/admin/seo/feeds", updateVerb: "PUT", canDelete: true, desc: "Entité CRUD /admin/seo/feeds (create/read/update/delete via API)."}
	},
	func() resource.Resource {
		return &apiResource{name: "v1_runs", path: "/v1/runs", updateVerb: "", canDelete: false, desc: "Entité CRUD /v1/runs (create/read/update/delete via API)."}
	},
}
