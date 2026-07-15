// Code généré par scripts/generate/06_terraform_provider.py — NE PAS ÉDITER À LA MAIN.
// Régénérer : python3 scripts/generate/06_terraform_provider.py --openapi <spec>

package provider

import "github.com/hashicorp/terraform-plugin-framework/datasource"

// generatedDataSources : data sources LECTURE (GET) générées depuis l'OpenAPI.
var generatedDataSources = []func() datasource.DataSource{
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_activity_log", path: "/admin/activity-log", listKey: "admin_activity_log", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/activity-log`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_email_stats", path: "/admin/agent/email/stats", listKey: "admin_agent_email_stats", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/agent/email/stats`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_reports", path: "/admin/agent/reports", listKey: "admin_agent_reports", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/agent/reports`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_scheduler_status", path: "/admin/agent/scheduler/status", listKey: "admin_agent_scheduler_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/agent/scheduler/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_status", path: "/admin/agent/status", listKey: "admin_agent_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/agent/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_tasks", path: "/admin/agent/tasks", listKey: "admin_agent_tasks", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/agent/tasks`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agents_catalog", path: "/admin/agents-catalog", listKey: "admin_agents_catalog", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/agents-catalog`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agents_tasks", path: "/admin/agents/tasks", listKey: "admin_agents_tasks", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/agents/tasks`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_annexe_iii_classifications", path: "/admin/ai-act/annexe-iii-classifications", listKey: "admin_ai_act_annexe_iii_classifications", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/annexe-iii-classifications`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_model_cards", path: "/admin/ai-act/model-cards", listKey: "admin_ai_act_model_cards", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/model-cards`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_model_cards_registry", path: "/admin/ai-act/model-cards/registry", listKey: "admin_ai_act_model_cards_registry", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/model-cards/registry`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_qms", path: "/admin/ai-act/qms", listKey: "admin_ai_act_qms", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/qms`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_risk_assessments", path: "/admin/ai-act/risk-assessments", listKey: "admin_ai_act_risk_assessments", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/risk-assessments`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_annex_iv", path: "/admin/ai-act/v2/annex-iv", listKey: "admin_ai_act_v2_annex_iv", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/annex-iv`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_audit", path: "/admin/ai-act/v2/audit", listKey: "admin_ai_act_v2_audit", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/audit`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_enforcement_preview", path: "/admin/ai-act/v2/enforcement/preview", listKey: "admin_ai_act_v2_enforcement_preview", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/enforcement/preview`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_incidents", path: "/admin/ai-act/v2/incidents", listKey: "admin_ai_act_v2_incidents", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/incidents`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_kill_switch", path: "/admin/ai-act/v2/kill-switch", listKey: "admin_ai_act_v2_kill_switch", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/kill-switch`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_model_cards_stubs", path: "/admin/ai-act/v2/model-cards-stubs", listKey: "admin_ai_act_v2_model_cards_stubs", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/model-cards-stubs`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_oversight_queue", path: "/admin/ai-act/v2/oversight/queue", listKey: "admin_ai_act_v2_oversight_queue", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/oversight/queue`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_oversight_stats", path: "/admin/ai-act/v2/oversight/stats", listKey: "admin_ai_act_v2_oversight_stats", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/oversight/stats`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_providers_compliance", path: "/admin/ai-act/v2/providers-compliance", listKey: "admin_ai_act_v2_providers_compliance", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/providers-compliance`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_transparency_disclosure", path: "/admin/ai-act/v2/transparency/disclosure", listKey: "admin_ai_act_v2_transparency_disclosure", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/transparency/disclosure`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_watermark_image_status", path: "/admin/ai-act/v2/watermark/image-status", listKey: "admin_ai_act_v2_watermark_image_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-act/v2/watermark/image-status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_rules", path: "/admin/ai-rules", listKey: "admin_ai_rules", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ai-rules`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_alerts", path: "/admin/alerts", listKey: "admin_alerts", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/alerts`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_analytics", path: "/admin/analytics", listKey: "admin_analytics", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/analytics`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_api_keys", path: "/admin/api-keys", listKey: "admin_api_keys", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/api-keys`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_api_keys_expiring", path: "/admin/api-keys/expiring", listKey: "admin_api_keys_expiring", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/api-keys/expiring`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_api_keys_rotation", path: "/admin/api-keys/rotation", listKey: "admin_api_keys_rotation", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/api-keys/rotation`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_audit_log", path: "/admin/audit-log", listKey: "admin_audit_log", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/audit-log`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_audit", path: "/admin/audit/", listKey: "admin_audit", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/audit/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_auto_test", path: "/admin/auto-test", listKey: "admin_auto_test", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/auto-test`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_autonomy_status", path: "/admin/autonomy/status", listKey: "admin_autonomy_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/autonomy/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_b2b_approval_queue", path: "/admin/b2b/approval/queue", listKey: "admin_b2b_approval_queue", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/b2b/approval/queue`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_backup_status", path: "/admin/backup/status", listKey: "admin_backup_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/backup/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing", path: "/admin/billing", listKey: "admin_billing", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/billing`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_baas_tiers", path: "/admin/billing/baas-tiers", listKey: "admin_billing_baas_tiers", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/billing/baas-tiers`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_cost_by_org", path: "/admin/billing/cost-by-org", listKey: "admin_billing_cost_by_org", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/billing/cost-by-org`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_cost_by_provider", path: "/admin/billing/cost-by-provider", listKey: "admin_billing_cost_by_provider", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/billing/cost-by-provider`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_cost_trend", path: "/admin/billing/cost-trend", listKey: "admin_billing_cost_trend", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/billing/cost-trend`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_organizations", path: "/admin/billing/organizations", listKey: "admin_billing_organizations", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/billing/organizations`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_stats", path: "/admin/billing/stats", listKey: "admin_billing_stats", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/billing/stats`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_blog_articles", path: "/admin/blog/articles", listKey: "admin_blog_articles", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/blog/articles`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_blog_posts", path: "/admin/blog/posts", listKey: "admin_blog_posts", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/blog/posts`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_auto_suggest_providers", path: "/admin/bot/auto-suggest-providers", listKey: "admin_bot_auto_suggest_providers", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/bot/auto-suggest-providers`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_blacklist", path: "/admin/bot/blacklist", listKey: "admin_bot_blacklist", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/bot/blacklist`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_conversations", path: "/admin/bot/conversations", listKey: "admin_bot_conversations", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/bot/conversations`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_credential_rotation", path: "/admin/bot/credential-rotation", listKey: "admin_bot_credential_rotation", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/bot/credential-rotation`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_scoring_usage", path: "/admin/bot/scoring-usage", listKey: "admin_bot_scoring_usage", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/bot/scoring-usage`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_status", path: "/admin/bot/status", listKey: "admin_bot_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/bot/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cache", path: "/admin/cache", listKey: "admin_cache", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/cache`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cache_stats", path: "/admin/cache/stats", listKey: "admin_cache_stats", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/cache/stats`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_campaigns", path: "/admin/campaigns/", listKey: "admin_campaigns", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/campaigns/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_catalog_sync_status", path: "/admin/catalog/sync/status", listKey: "admin_catalog_sync_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/catalog/sync/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_certificates", path: "/admin/certificates", listKey: "admin_certificates", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/certificates`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_circuit_breakers", path: "/admin/circuit-breakers", listKey: "admin_circuit_breakers", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/circuit-breakers`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cloud_models", path: "/admin/cloud-models", listKey: "admin_cloud_models", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/cloud-models`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cloud_providers", path: "/admin/cloud-providers", listKey: "admin_cloud_providers", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/cloud-providers`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cloud_providers_terraform_modules", path: "/admin/cloud-providers/terraform-modules", listKey: "admin_cloud_providers_terraform_modules", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/cloud-providers/terraform-modules`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cluster_nodes", path: "/admin/cluster/nodes", listKey: "admin_cluster_nodes", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/cluster/nodes`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cluster_services", path: "/admin/cluster/services", listKey: "admin_cluster_services", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/cluster/services`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_compliance", path: "/admin/compliance", listKey: "admin_compliance", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/compliance`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_config_email", path: "/admin/config/email", listKey: "admin_config_email", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/config/email`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_config_email_templates", path: "/admin/config/email/templates", listKey: "admin_config_email_templates", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/config/email/templates`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_config_history", path: "/admin/config/history", listKey: "admin_config_history", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/config/history`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_config_identity", path: "/admin/config/identity", listKey: "admin_config_identity", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/config/identity`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_config_saml", path: "/admin/config/saml", listKey: "admin_config_saml", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/config/saml`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_connectors_catalog", path: "/admin/connectors-catalog", listKey: "admin_connectors_catalog", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/connectors-catalog`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_connectors_catalog_ecosystem", path: "/admin/connectors-catalog/ecosystem", listKey: "admin_connectors_catalog_ecosystem", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/connectors-catalog/ecosystem`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_connectors_auth_methods", path: "/admin/connectors/auth-methods", listKey: "admin_connectors_auth_methods", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/connectors/auth-methods`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_connectors_manifests", path: "/admin/connectors/manifests", listKey: "admin_connectors_manifests", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/connectors/manifests`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_connectors_mcp_servers", path: "/admin/connectors/mcp/servers", listKey: "admin_connectors_mcp_servers", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/connectors/mcp/servers`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_contracts", path: "/admin/contracts", listKey: "admin_contracts", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/contracts`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cost_dashboard", path: "/admin/cost-dashboard", listKey: "admin_cost_dashboard", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/cost-dashboard`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_credentials", path: "/admin/credentials", listKey: "admin_credentials", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/credentials`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_crews", path: "/admin/crews", listKey: "admin_crews", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/crews`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_crews_orphans_count", path: "/admin/crews/orphans/count", listKey: "admin_crews_orphans_count", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/crews/orphans/count`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_crews_runs", path: "/admin/crews/runs", listKey: "admin_crews_runs", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/crews/runs`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_database_info", path: "/admin/database/info", listKey: "admin_database_info", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/database/info`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_database_migrations", path: "/admin/database/migrations", listKey: "admin_database_migrations", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/database/migrations`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_datasets", path: "/admin/datasets", listKey: "admin_datasets", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/datasets`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_db_ddl_status", path: "/admin/db/ddl-status", listKey: "admin_db_ddl_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/db/ddl-status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_deployable_services", path: "/admin/deployable-services", listKey: "admin_deployable_services", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/deployable-services`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_diagnostics", path: "/admin/diagnostics", listKey: "admin_diagnostics", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/diagnostics`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_discovery_runs", path: "/admin/discovery/runs", listKey: "admin_discovery_runs", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/discovery/runs`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_dns_zone", path: "/admin/dns/zone", listKey: "admin_dns_zone", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/dns/zone`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_dr_drills", path: "/admin/dr/drills", listKey: "admin_dr_drills", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/dr/drills`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_egress_audit", path: "/admin/egress/audit", listKey: "admin_egress_audit", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/egress/audit`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_egress_policies", path: "/admin/egress/policies", listKey: "admin_egress_policies", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/egress/policies`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_email_templates", path: "/admin/email-templates", listKey: "admin_email_templates", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/email-templates`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_events", path: "/admin/events/", listKey: "admin_events", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/events/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_feature_flags", path: "/admin/feature-flags", listKey: "admin_feature_flags", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/feature-flags`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_finetune", path: "/admin/finetune", listKey: "admin_finetune", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/finetune`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_finops_cost_by_org", path: "/admin/finops/cost-by-org", listKey: "admin_finops_cost_by_org", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/finops/cost-by-org`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_finops_cost_by_provider", path: "/admin/finops/cost-by-provider", listKey: "admin_finops_cost_by_provider", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/finops/cost-by-provider`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_finops_cost_trend", path: "/admin/finops/cost-trend", listKey: "admin_finops_cost_trend", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/finops/cost-trend`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_gdpr_requests", path: "/admin/gdpr/requests", listKey: "admin_gdpr_requests", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/gdpr/requests`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_gdpr_requests_audit_log", path: "/admin/gdpr/requests/audit-log", listKey: "admin_gdpr_requests_audit_log", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/gdpr/requests/audit-log`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_golden_path", path: "/admin/golden-path", listKey: "admin_golden_path", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/golden-path`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_grafana_dashboards", path: "/admin/grafana/dashboards", listKey: "admin_grafana_dashboards", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/grafana/dashboards`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_groups", path: "/admin/groups/", listKey: "admin_groups", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/groups/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_guardrails", path: "/admin/guardrails", listKey: "admin_guardrails", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/guardrails`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_healthcheck", path: "/admin/healthcheck", listKey: "admin_healthcheck", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/healthcheck`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_hierarchy", path: "/admin/hierarchy", listKey: "admin_hierarchy", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/hierarchy`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_i18n_locales", path: "/admin/i18n/locales", listKey: "admin_i18n_locales", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/i18n/locales`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_integrations_status", path: "/admin/integrations/status", listKey: "admin_integrations_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/integrations/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_investor_me", path: "/admin/investor/me", listKey: "admin_investor_me", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/investor/me`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_investors", path: "/admin/investors", listKey: "admin_investors", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/investors`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_investors_ndas", path: "/admin/investors/ndas", listKey: "admin_investors_ndas", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/investors/ndas`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_knowledge", path: "/admin/knowledge", listKey: "admin_knowledge", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/knowledge`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_leads", path: "/admin/leads", listKey: "admin_leads", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/leads`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_learning_cycles", path: "/admin/learning/cycles", listKey: "admin_learning_cycles", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/learning/cycles`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_learning_finetune_runs", path: "/admin/learning/finetune-runs", listKey: "admin_learning_finetune_runs", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/learning/finetune-runs`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_learning_proposals", path: "/admin/learning/proposals", listKey: "admin_learning_proposals", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/learning/proposals`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_learning_status", path: "/admin/learning/status", listKey: "admin_learning_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/learning/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_learning_summary", path: "/admin/learning/summary", listKey: "admin_learning_summary", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/learning/summary`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_licenses", path: "/admin/licenses", listKey: "admin_licenses", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/licenses`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_linkedin_connect", path: "/admin/linkedin/connect", listKey: "admin_linkedin_connect", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/linkedin/connect`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_linkedin_status", path: "/admin/linkedin/status", listKey: "admin_linkedin_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/linkedin/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_logs", path: "/admin/logs", listKey: "admin_logs", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/logs`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_logs_recent", path: "/admin/logs/recent", listKey: "admin_logs_recent", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/logs/recent`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_maintenance_status", path: "/admin/maintenance/status", listKey: "admin_maintenance_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/maintenance/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_mcp_status", path: "/admin/mcp/status", listKey: "admin_mcp_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/mcp/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_me", path: "/admin/me", listKey: "admin_me", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/me`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_metrics_dashboards", path: "/admin/metrics/dashboards", listKey: "admin_metrics_dashboards", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/metrics/dashboards`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_metrics_query", path: "/admin/metrics/query", listKey: "admin_metrics_query", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/metrics/query`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_models", path: "/admin/models", listKey: "admin_models", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/models`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_multitenant_subdomains", path: "/admin/multitenant/subdomains", listKey: "admin_multitenant_subdomains", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/multitenant/subdomains`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_network", path: "/admin/network", listKey: "admin_network", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/network`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_oidc_providers", path: "/admin/oidc-providers", listKey: "admin_oidc_providers", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/oidc-providers`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ollama_models", path: "/admin/ollama/models", listKey: "admin_ollama_models", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ollama/models`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_onboarding", path: "/admin/onboarding", listKey: "admin_onboarding", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/onboarding`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_onboarding_pending", path: "/admin/onboarding/pending", listKey: "admin_onboarding_pending", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/onboarding/pending`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ops_backups", path: "/admin/ops/backups", listKey: "admin_ops_backups", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ops/backups`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ops_build_history", path: "/admin/ops/build-history", listKey: "admin_ops_build_history", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ops/build-history`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ops_services", path: "/admin/ops/services", listKey: "admin_ops_services", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/ops/services`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_organizations", path: "/admin/organizations/", listKey: "admin_organizations", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/organizations/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_orgs", path: "/admin/orgs", listKey: "admin_orgs", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/orgs`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_orgs_quotas", path: "/admin/orgs/quotas", listKey: "admin_orgs_quotas", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/orgs/quotas`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_orgs_routing_mode", path: "/admin/orgs/routing-mode", listKey: "admin_orgs_routing_mode", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/orgs/routing-mode`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_plans", path: "/admin/plans/", listKey: "admin_plans", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/plans/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_prometheus_alerts", path: "/admin/prometheus/alerts", listKey: "admin_prometheus_alerts", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/prometheus/alerts`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_prometheus_query", path: "/admin/prometheus/query", listKey: "admin_prometheus_query", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/prometheus/query`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_prometheus_targets", path: "/admin/prometheus/targets", listKey: "admin_prometheus_targets", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/prometheus/targets`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_provider_pricing", path: "/admin/provider-pricing", listKey: "admin_provider_pricing", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/provider-pricing`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_providers_capabilities", path: "/admin/providers/capabilities", listKey: "admin_providers_capabilities", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/providers/capabilities`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_providers_config", path: "/admin/providers/config", listKey: "admin_providers_config", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/providers/config`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_providers_dashboard", path: "/admin/providers/dashboard", listKey: "admin_providers_dashboard", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/providers/dashboard`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_providers_keys_status", path: "/admin/providers/keys/status", listKey: "admin_providers_keys_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/providers/keys/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_qdrant_collections", path: "/admin/qdrant/collections", listKey: "admin_qdrant_collections", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/qdrant/collections`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_qdrant_info", path: "/admin/qdrant/info", listKey: "admin_qdrant_info", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/qdrant/info`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_rate_limits", path: "/admin/rate-limits", listKey: "admin_rate_limits", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/rate-limits`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_rbac", path: "/admin/rbac", listKey: "admin_rbac", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/rbac`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_rbac_roles", path: "/admin/rbac/roles", listKey: "admin_rbac_roles", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/rbac/roles`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_redis_info", path: "/admin/redis/info", listKey: "admin_redis_info", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/redis/info`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_resellers", path: "/admin/resellers/", listKey: "admin_resellers", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/resellers/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_routing_policy", path: "/admin/routing/policy", listKey: "admin_routing_policy", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/routing/policy`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_scheduled_backups", path: "/admin/scheduled-backups", listKey: "admin_scheduled_backups", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/scheduled-backups`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_security", path: "/admin/security", listKey: "admin_security", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/security`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_self_audit", path: "/admin/self-audit", listKey: "admin_self_audit", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/self-audit`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_comments", path: "/admin/seo/comments", listKey: "admin_seo_comments", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/seo/comments`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_feeds", path: "/admin/seo/feeds", listKey: "admin_seo_feeds", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/seo/feeds`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_posts", path: "/admin/seo/posts", listKey: "admin_seo_posts", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/seo/posts`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_stats", path: "/admin/seo/stats", listKey: "admin_seo_stats", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/seo/stats`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_stats_global", path: "/admin/seo/stats/global", listKey: "admin_seo_stats_global", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/seo/stats/global`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_service_inventory", path: "/admin/service-inventory", listKey: "admin_service_inventory", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/service-inventory`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_sessions", path: "/admin/sessions", listKey: "admin_sessions", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/sessions`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_sessions_bulk_revoke", path: "/admin/sessions/bulk-revoke", listKey: "admin_sessions_bulk_revoke", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/sessions/bulk-revoke`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_slo", path: "/admin/slo", listKey: "admin_slo", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/slo`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_support", path: "/admin/support", listKey: "admin_support", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/support`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_swarm_services", path: "/admin/swarm/services", listKey: "admin_swarm_services", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/swarm/services`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_tests_history", path: "/admin/tests/history", listKey: "admin_tests_history", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/tests/history`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_truth_pack", path: "/admin/truth/pack", listKey: "admin_truth_pack", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/truth/pack`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_truth_ui", path: "/admin/truth/ui", listKey: "admin_truth_ui", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/truth/ui`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_usage", path: "/admin/usage/", listKey: "admin_usage", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/usage/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_users", path: "/admin/users/", listKey: "admin_users", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/users/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_vault", path: "/admin/vault", listKey: "admin_vault", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/vault`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_verticals", path: "/admin/verticals/", listKey: "admin_verticals", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/verticals/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_webhooks", path: "/admin/webhooks", listKey: "admin_webhooks", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/webhooks`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_webhooks_resend", path: "/admin/webhooks/resend", listKey: "admin_webhooks_resend", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/webhooks/resend`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_yousign_status", path: "/admin/yousign/status", listKey: "admin_yousign_status", desc: "Endpoint d'administration (Bearer requis) — lecture GET de `/admin/yousign/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "api_investor_docs_access", path: "/api/investor-docs/access", listKey: "api_investor_docs_access", desc: "Endpoint API AISIA — lecture GET de `/api/investor-docs/access`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "api_plans", path: "/api/plans", listKey: "api_plans", desc: "Endpoint API AISIA — lecture GET de `/api/plans`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "api_privacy_json", path: "/api/privacy.json", listKey: "api_privacy_json", desc: "Endpoint API AISIA — lecture GET de `/api/privacy.json`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "api_tenant_context", path: "/api/tenant-context", listKey: "api_tenant_context", desc: "Endpoint API AISIA — lecture GET de `/api/tenant-context`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_me", path: "/auth/me", listKey: "auth_me", desc: "Endpoint API AISIA — lecture GET de `/auth/me`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_me_consents", path: "/auth/me/consents", listKey: "auth_me_consents", desc: "Endpoint API AISIA — lecture GET de `/auth/me/consents`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_me_profile", path: "/auth/me/profile", listKey: "auth_me_profile", desc: "Endpoint API AISIA — lecture GET de `/auth/me/profile`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_mfa_status", path: "/auth/mfa/status", listKey: "auth_mfa_status", desc: "Endpoint API AISIA — lecture GET de `/auth/mfa/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_oidc_providers", path: "/auth/oidc/providers", listKey: "auth_oidc_providers", desc: "Endpoint API AISIA — lecture GET de `/auth/oidc/providers`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_saml_metadata", path: "/auth/saml/metadata", listKey: "auth_saml_metadata", desc: "Endpoint API AISIA — lecture GET de `/auth/saml/metadata`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_saml_status", path: "/auth/saml/status", listKey: "auth_saml_status", desc: "Endpoint API AISIA — lecture GET de `/auth/saml/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_verify_email", path: "/auth/verify-email", listKey: "auth_verify_email", desc: "Endpoint API AISIA — lecture GET de `/auth/verify-email`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "billing_checkout_llm_catalog", path: "/billing/checkout/llm-catalog", listKey: "billing_checkout_llm_catalog", desc: "Endpoint API AISIA — lecture GET de `/billing/checkout/llm-catalog`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "billing_checkout_status", path: "/billing/checkout/status", listKey: "billing_checkout_status", desc: "Endpoint API AISIA — lecture GET de `/billing/checkout/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "events_upcoming", path: "/events/upcoming", listKey: "events_upcoming", desc: "Endpoint API AISIA — lecture GET de `/events/upcoming`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_account", path: "/org/account", listKey: "org_account", desc: "Endpoint API AISIA — lecture GET de `/org/account`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_account_mfa", path: "/org/account/mfa", listKey: "org_account_mfa", desc: "Endpoint API AISIA — lecture GET de `/org/account/mfa`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_agents", path: "/org/agents", listKey: "org_agents", desc: "Endpoint API AISIA — lecture GET de `/org/agents`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_api_keys", path: "/org/api-keys", listKey: "org_api_keys", desc: "Endpoint API AISIA — lecture GET de `/org/api-keys`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_audit_log", path: "/org/audit-log", listKey: "org_audit_log", desc: "Endpoint API AISIA — lecture GET de `/org/audit-log`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_billing", path: "/org/billing", listKey: "org_billing", desc: "Endpoint API AISIA — lecture GET de `/org/billing`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_branding", path: "/org/branding", listKey: "org_branding", desc: "Endpoint API AISIA — lecture GET de `/org/branding`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_budget", path: "/org/budget", listKey: "org_budget", desc: "Endpoint API AISIA — lecture GET de `/org/budget`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_deployable_services", path: "/org/deployable-services", listKey: "org_deployable_services", desc: "Endpoint API AISIA — lecture GET de `/org/deployable-services`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_deployments", path: "/org/deployments", listKey: "org_deployments", desc: "Endpoint API AISIA — lecture GET de `/org/deployments`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_groups", path: "/org/groups", listKey: "org_groups", desc: "Endpoint API AISIA — lecture GET de `/org/groups`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_integrations", path: "/org/integrations", listKey: "org_integrations", desc: "Endpoint API AISIA — lecture GET de `/org/integrations`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_invoices", path: "/org/invoices", listKey: "org_invoices", desc: "Endpoint API AISIA — lecture GET de `/org/invoices`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_knowledge", path: "/org/knowledge", listKey: "org_knowledge", desc: "Endpoint API AISIA — lecture GET de `/org/knowledge`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_me", path: "/org/me", listKey: "org_me", desc: "Endpoint API AISIA — lecture GET de `/org/me`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_members", path: "/org/members", listKey: "org_members", desc: "Endpoint API AISIA — lecture GET de `/org/members`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_notifications", path: "/org/notifications", listKey: "org_notifications", desc: "Endpoint API AISIA — lecture GET de `/org/notifications`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_oidc_config", path: "/org/oidc-config", listKey: "org_oidc_config", desc: "Endpoint API AISIA — lecture GET de `/org/oidc-config`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_quotas", path: "/org/quotas", listKey: "org_quotas", desc: "Endpoint API AISIA — lecture GET de `/org/quotas`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_roles", path: "/org/roles", listKey: "org_roles", desc: "Endpoint API AISIA — lecture GET de `/org/roles`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_sessions", path: "/org/sessions", listKey: "org_sessions", desc: "Endpoint API AISIA — lecture GET de `/org/sessions`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_settings", path: "/org/settings", listKey: "org_settings", desc: "Endpoint API AISIA — lecture GET de `/org/settings`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_support_tickets", path: "/org/support/tickets", listKey: "org_support_tickets", desc: "Endpoint API AISIA — lecture GET de `/org/support/tickets`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_usage", path: "/org/usage", listKey: "org_usage", desc: "Endpoint API AISIA — lecture GET de `/org/usage`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_webhooks", path: "/org/webhooks", listKey: "org_webhooks", desc: "Endpoint API AISIA — lecture GET de `/org/webhooks`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "privacy", path: "/privacy", listKey: "privacy", desc: "Endpoint API AISIA — lecture GET de `/privacy`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "scim_v2_groups", path: "/scim/v2/Groups", listKey: "scim_v2_groups", desc: "Endpoint API AISIA — lecture GET de `/scim/v2/Groups`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "scim_v2_users", path: "/scim/v2/Users", listKey: "scim_v2_users", desc: "Endpoint API AISIA — lecture GET de `/scim/v2/Users`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "support_tickets", path: "/support/tickets/", listKey: "support_tickets", desc: "Endpoint API AISIA — lecture GET de `/support/tickets/`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_capabilities", path: "/v1/capabilities", listKey: "v1_capabilities", desc: "Endpoint public API v1 — lecture GET de `/v1/capabilities`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_conversations", path: "/v1/conversations", listKey: "v1_conversations", desc: "Endpoint public API v1 — lecture GET de `/v1/conversations`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_debates", path: "/v1/debates", listKey: "v1_debates", desc: "Endpoint public API v1 — lecture GET de `/v1/debates`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_info", path: "/v1/info", listKey: "v1_info", desc: "Endpoint public API v1 — lecture GET de `/v1/info`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_local_models", path: "/v1/local-models", listKey: "v1_local_models", desc: "Endpoint public API v1 — lecture GET de `/v1/local-models`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_me_feedback", path: "/v1/me/feedback", listKey: "v1_me_feedback", desc: "Endpoint public API v1 — lecture GET de `/v1/me/feedback`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_me_providers", path: "/v1/me/providers", listKey: "v1_me_providers", desc: "Endpoint public API v1 — lecture GET de `/v1/me/providers`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_me_quotas", path: "/v1/me/quotas", listKey: "v1_me_quotas", desc: "Endpoint public API v1 — lecture GET de `/v1/me/quotas`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_models", path: "/v1/models", listKey: "v1_models", desc: "Endpoint public API v1 — lecture GET de `/v1/models`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_providers_ai_act_compliance", path: "/v1/providers/ai-act-compliance", listKey: "v1_providers_ai_act_compliance", desc: "Endpoint public API v1 — lecture GET de `/v1/providers/ai-act-compliance`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_public_stats", path: "/v1/public/stats", listKey: "v1_public_stats", desc: "Endpoint public API v1 — lecture GET de `/v1/public/stats`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_status", path: "/v1/status", listKey: "v1_status", desc: "Endpoint public API v1 — lecture GET de `/v1/status`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_status_incidents", path: "/v1/status/incidents", listKey: "v1_status_incidents", desc: "Endpoint public API v1 — lecture GET de `/v1/status/incidents`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_trust_certifications", path: "/v1/trust/certifications", listKey: "v1_trust_certifications", desc: "Endpoint public API v1 — lecture GET de `/v1/trust/certifications`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_truth_capabilities", path: "/v1/truth/capabilities", listKey: "v1_truth_capabilities", desc: "Endpoint public API v1 — lecture GET de `/v1/truth/capabilities`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_truth_me", path: "/v1/truth/me", listKey: "v1_truth_me", desc: "Endpoint public API v1 — lecture GET de `/v1/truth/me`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_truth_pricing", path: "/v1/truth/pricing", listKey: "v1_truth_pricing", desc: "Endpoint public API v1 — lecture GET de `/v1/truth/pricing`."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_truth_summary", path: "/v1/truth/summary", listKey: "v1_truth_summary", desc: "Endpoint public API v1 — lecture GET de `/v1/truth/summary`."}
	},
}
