// Code généré par scripts/gen_terraform_provider.py — NE PAS ÉDITER À LA MAIN.
// Régénérer : python3 scripts/gen_terraform_provider.py --openapi <spec>

package provider

import "github.com/hashicorp/terraform-plugin-framework/datasource"

// generatedDataSources : data sources LECTURE (GET) générées depuis l'OpenAPI.
var generatedDataSources = []func() datasource.DataSource{
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_activity_log", path: "/admin/activity-log", listKey: "admin_activity_log", desc: "Lecture de /admin/activity-log."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_email_stats", path: "/admin/agent/email/stats", listKey: "admin_agent_email_stats", desc: "Lecture de /admin/agent/email/stats."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_reports", path: "/admin/agent/reports", listKey: "admin_agent_reports", desc: "Lecture de /admin/agent/reports."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_scheduler_status", path: "/admin/agent/scheduler/status", listKey: "admin_agent_scheduler_status", desc: "Lecture de /admin/agent/scheduler/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_status", path: "/admin/agent/status", listKey: "admin_agent_status", desc: "Lecture de /admin/agent/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agent_tasks", path: "/admin/agent/tasks", listKey: "admin_agent_tasks", desc: "Lecture de /admin/agent/tasks."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_agents_tasks", path: "/admin/agents/tasks", listKey: "admin_agents_tasks", desc: "Lecture de /admin/agents/tasks."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_annexe_iii_classifications", path: "/admin/ai-act/annexe-iii-classifications", listKey: "admin_ai_act_annexe_iii_classifications", desc: "Lecture de /admin/ai-act/annexe-iii-classifications."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_model_cards", path: "/admin/ai-act/model-cards", listKey: "admin_ai_act_model_cards", desc: "Lecture de /admin/ai-act/model-cards."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_risk_assessments", path: "/admin/ai-act/risk-assessments", listKey: "admin_ai_act_risk_assessments", desc: "Lecture de /admin/ai-act/risk-assessments."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_annex_iv", path: "/admin/ai-act/v2/annex-iv", listKey: "admin_ai_act_v2_annex_iv", desc: "Lecture de /admin/ai-act/v2/annex-iv."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_audit", path: "/admin/ai-act/v2/audit", listKey: "admin_ai_act_v2_audit", desc: "Lecture de /admin/ai-act/v2/audit."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_kill_switch", path: "/admin/ai-act/v2/kill-switch", listKey: "admin_ai_act_v2_kill_switch", desc: "Lecture de /admin/ai-act/v2/kill-switch."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_oversight_queue", path: "/admin/ai-act/v2/oversight/queue", listKey: "admin_ai_act_v2_oversight_queue", desc: "Lecture de /admin/ai-act/v2/oversight/queue."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_oversight_stats", path: "/admin/ai-act/v2/oversight/stats", listKey: "admin_ai_act_v2_oversight_stats", desc: "Lecture de /admin/ai-act/v2/oversight/stats."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_providers_compliance", path: "/admin/ai-act/v2/providers-compliance", listKey: "admin_ai_act_v2_providers_compliance", desc: "Lecture de /admin/ai-act/v2/providers-compliance."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_transparency_disclosure", path: "/admin/ai-act/v2/transparency/disclosure", listKey: "admin_ai_act_v2_transparency_disclosure", desc: "Lecture de /admin/ai-act/v2/transparency/disclosure."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_act_v2_watermark_image_status", path: "/admin/ai-act/v2/watermark/image-status", listKey: "admin_ai_act_v2_watermark_image_status", desc: "Lecture de /admin/ai-act/v2/watermark/image-status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ai_rules", path: "/admin/ai-rules", listKey: "admin_ai_rules", desc: "Lecture de /admin/ai-rules."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_analytics", path: "/admin/analytics", listKey: "admin_analytics", desc: "Lecture de /admin/analytics."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_api_keys", path: "/admin/api-keys", listKey: "admin_api_keys", desc: "Lecture de /admin/api-keys."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_api_keys_expiring", path: "/admin/api-keys/expiring", listKey: "admin_api_keys_expiring", desc: "Lecture de /admin/api-keys/expiring."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_api_keys_rotation", path: "/admin/api-keys/rotation", listKey: "admin_api_keys_rotation", desc: "Lecture de /admin/api-keys/rotation."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_audit_log", path: "/admin/audit-log", listKey: "admin_audit_log", desc: "Lecture de /admin/audit-log."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_audit", path: "/admin/audit/", listKey: "admin_audit", desc: "Lecture de /admin/audit/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_autonomy_status", path: "/admin/autonomy/status", listKey: "admin_autonomy_status", desc: "Lecture de /admin/autonomy/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_b2b_approval_queue", path: "/admin/b2b/approval/queue", listKey: "admin_b2b_approval_queue", desc: "Lecture de /admin/b2b/approval/queue."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_backup_status", path: "/admin/backup/status", listKey: "admin_backup_status", desc: "Lecture de /admin/backup/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing", path: "/admin/billing", listKey: "admin_billing", desc: "Lecture de /admin/billing."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_baas_tiers", path: "/admin/billing/baas-tiers", listKey: "admin_billing_baas_tiers", desc: "Lecture de /admin/billing/baas-tiers."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_cost_by_org", path: "/admin/billing/cost-by-org", listKey: "admin_billing_cost_by_org", desc: "Lecture de /admin/billing/cost-by-org."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_cost_by_provider", path: "/admin/billing/cost-by-provider", listKey: "admin_billing_cost_by_provider", desc: "Lecture de /admin/billing/cost-by-provider."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_cost_trend", path: "/admin/billing/cost-trend", listKey: "admin_billing_cost_trend", desc: "Lecture de /admin/billing/cost-trend."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_organizations", path: "/admin/billing/organizations", listKey: "admin_billing_organizations", desc: "Lecture de /admin/billing/organizations."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_billing_stats", path: "/admin/billing/stats", listKey: "admin_billing_stats", desc: "Lecture de /admin/billing/stats."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_blog_articles", path: "/admin/blog/articles", listKey: "admin_blog_articles", desc: "Lecture de /admin/blog/articles."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_blog_posts", path: "/admin/blog/posts", listKey: "admin_blog_posts", desc: "Lecture de /admin/blog/posts."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_auto_suggest_providers", path: "/admin/bot/auto-suggest-providers", listKey: "admin_bot_auto_suggest_providers", desc: "Lecture de /admin/bot/auto-suggest-providers."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_blacklist", path: "/admin/bot/blacklist", listKey: "admin_bot_blacklist", desc: "Lecture de /admin/bot/blacklist."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_conversations", path: "/admin/bot/conversations", listKey: "admin_bot_conversations", desc: "Lecture de /admin/bot/conversations."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_credential_rotation", path: "/admin/bot/credential-rotation", listKey: "admin_bot_credential_rotation", desc: "Lecture de /admin/bot/credential-rotation."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_scoring_usage", path: "/admin/bot/scoring-usage", listKey: "admin_bot_scoring_usage", desc: "Lecture de /admin/bot/scoring-usage."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_bot_status", path: "/admin/bot/status", listKey: "admin_bot_status", desc: "Lecture de /admin/bot/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cache", path: "/admin/cache", listKey: "admin_cache", desc: "Lecture de /admin/cache."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cache_stats", path: "/admin/cache/stats", listKey: "admin_cache_stats", desc: "Lecture de /admin/cache/stats."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_campaigns", path: "/admin/campaigns/", listKey: "admin_campaigns", desc: "Lecture de /admin/campaigns/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_certificates", path: "/admin/certificates", listKey: "admin_certificates", desc: "Lecture de /admin/certificates."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_circuit_breakers", path: "/admin/circuit-breakers", listKey: "admin_circuit_breakers", desc: "Lecture de /admin/circuit-breakers."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cloud_providers", path: "/admin/cloud-providers", listKey: "admin_cloud_providers", desc: "Lecture de /admin/cloud-providers."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cloud_providers_terraform_modules", path: "/admin/cloud-providers/terraform-modules", listKey: "admin_cloud_providers_terraform_modules", desc: "Lecture de /admin/cloud-providers/terraform-modules."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cluster_nodes", path: "/admin/cluster/nodes", listKey: "admin_cluster_nodes", desc: "Lecture de /admin/cluster/nodes."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cluster_services", path: "/admin/cluster/services", listKey: "admin_cluster_services", desc: "Lecture de /admin/cluster/services."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_compliance", path: "/admin/compliance", listKey: "admin_compliance", desc: "Lecture de /admin/compliance."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_config_email", path: "/admin/config/email", listKey: "admin_config_email", desc: "Lecture de /admin/config/email."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_config_email_templates", path: "/admin/config/email/templates", listKey: "admin_config_email_templates", desc: "Lecture de /admin/config/email/templates."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_config_history", path: "/admin/config/history", listKey: "admin_config_history", desc: "Lecture de /admin/config/history."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_config_identity", path: "/admin/config/identity", listKey: "admin_config_identity", desc: "Lecture de /admin/config/identity."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_connectors_catalog", path: "/admin/connectors-catalog", listKey: "admin_connectors_catalog", desc: "Lecture de /admin/connectors-catalog."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_connectors_manifests", path: "/admin/connectors/manifests", listKey: "admin_connectors_manifests", desc: "Lecture de /admin/connectors/manifests."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_contracts", path: "/admin/contracts", listKey: "admin_contracts", desc: "Lecture de /admin/contracts."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_cost_dashboard", path: "/admin/cost-dashboard", listKey: "admin_cost_dashboard", desc: "Lecture de /admin/cost-dashboard."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_credentials", path: "/admin/credentials", listKey: "admin_credentials", desc: "Lecture de /admin/credentials."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_crews", path: "/admin/crews", listKey: "admin_crews", desc: "Lecture de /admin/crews."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_crews_orphans_count", path: "/admin/crews/orphans/count", listKey: "admin_crews_orphans_count", desc: "Lecture de /admin/crews/orphans/count."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_crews_runs", path: "/admin/crews/runs", listKey: "admin_crews_runs", desc: "Lecture de /admin/crews/runs."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_database_info", path: "/admin/database/info", listKey: "admin_database_info", desc: "Lecture de /admin/database/info."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_database_migrations", path: "/admin/database/migrations", listKey: "admin_database_migrations", desc: "Lecture de /admin/database/migrations."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_datasets", path: "/admin/datasets", listKey: "admin_datasets", desc: "Lecture de /admin/datasets."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_db_ddl_status", path: "/admin/db/ddl-status", listKey: "admin_db_ddl_status", desc: "Lecture de /admin/db/ddl-status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_deployable_services", path: "/admin/deployable-services", listKey: "admin_deployable_services", desc: "Lecture de /admin/deployable-services."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_diagnostics", path: "/admin/diagnostics", listKey: "admin_diagnostics", desc: "Lecture de /admin/diagnostics."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_discovery_runs", path: "/admin/discovery/runs", listKey: "admin_discovery_runs", desc: "Lecture de /admin/discovery/runs."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_dns_zone", path: "/admin/dns/zone", listKey: "admin_dns_zone", desc: "Lecture de /admin/dns/zone."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_dr_drills", path: "/admin/dr/drills", listKey: "admin_dr_drills", desc: "Lecture de /admin/dr/drills."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_egress_audit", path: "/admin/egress/audit", listKey: "admin_egress_audit", desc: "Lecture de /admin/egress/audit."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_egress_policies", path: "/admin/egress/policies", listKey: "admin_egress_policies", desc: "Lecture de /admin/egress/policies."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_email_templates", path: "/admin/email-templates", listKey: "admin_email_templates", desc: "Lecture de /admin/email-templates."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_events", path: "/admin/events/", listKey: "admin_events", desc: "Lecture de /admin/events/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_feature_flags", path: "/admin/feature-flags", listKey: "admin_feature_flags", desc: "Lecture de /admin/feature-flags."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_gdpr_requests", path: "/admin/gdpr/requests", listKey: "admin_gdpr_requests", desc: "Lecture de /admin/gdpr/requests."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_gdpr_requests_audit_log", path: "/admin/gdpr/requests/audit-log", listKey: "admin_gdpr_requests_audit_log", desc: "Lecture de /admin/gdpr/requests/audit-log."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_golden_path", path: "/admin/golden-path", listKey: "admin_golden_path", desc: "Lecture de /admin/golden-path."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_grafana_dashboards", path: "/admin/grafana/dashboards", listKey: "admin_grafana_dashboards", desc: "Lecture de /admin/grafana/dashboards."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_groups", path: "/admin/groups/", listKey: "admin_groups", desc: "Lecture de /admin/groups/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_guardrails", path: "/admin/guardrails", listKey: "admin_guardrails", desc: "Lecture de /admin/guardrails."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_hierarchy", path: "/admin/hierarchy", listKey: "admin_hierarchy", desc: "Lecture de /admin/hierarchy."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_i18n_locales", path: "/admin/i18n/locales", listKey: "admin_i18n_locales", desc: "Lecture de /admin/i18n/locales."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_integrations_status", path: "/admin/integrations/status", listKey: "admin_integrations_status", desc: "Lecture de /admin/integrations/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_investor_me", path: "/admin/investor/me", listKey: "admin_investor_me", desc: "Lecture de /admin/investor/me."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_investors", path: "/admin/investors", listKey: "admin_investors", desc: "Lecture de /admin/investors."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_investors_ndas", path: "/admin/investors/ndas", listKey: "admin_investors_ndas", desc: "Lecture de /admin/investors/ndas."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_knowledge", path: "/admin/knowledge", listKey: "admin_knowledge", desc: "Lecture de /admin/knowledge."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_leads", path: "/admin/leads", listKey: "admin_leads", desc: "Lecture de /admin/leads."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_learning_cycles", path: "/admin/learning/cycles", listKey: "admin_learning_cycles", desc: "Lecture de /admin/learning/cycles."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_learning_finetune_runs", path: "/admin/learning/finetune-runs", listKey: "admin_learning_finetune_runs", desc: "Lecture de /admin/learning/finetune-runs."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_learning_proposals", path: "/admin/learning/proposals", listKey: "admin_learning_proposals", desc: "Lecture de /admin/learning/proposals."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_learning_status", path: "/admin/learning/status", listKey: "admin_learning_status", desc: "Lecture de /admin/learning/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_licenses", path: "/admin/licenses", listKey: "admin_licenses", desc: "Lecture de /admin/licenses."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_logs", path: "/admin/logs", listKey: "admin_logs", desc: "Lecture de /admin/logs."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_logs_recent", path: "/admin/logs/recent", listKey: "admin_logs_recent", desc: "Lecture de /admin/logs/recent."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_maintenance_status", path: "/admin/maintenance/status", listKey: "admin_maintenance_status", desc: "Lecture de /admin/maintenance/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_mcp_status", path: "/admin/mcp/status", listKey: "admin_mcp_status", desc: "Lecture de /admin/mcp/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_me", path: "/admin/me", listKey: "admin_me", desc: "Lecture de /admin/me."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_multitenant_subdomains", path: "/admin/multitenant/subdomains", listKey: "admin_multitenant_subdomains", desc: "Lecture de /admin/multitenant/subdomains."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_network", path: "/admin/network", listKey: "admin_network", desc: "Lecture de /admin/network."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_oidc_providers", path: "/admin/oidc-providers", listKey: "admin_oidc_providers", desc: "Lecture de /admin/oidc-providers."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ollama_models", path: "/admin/ollama-models", listKey: "admin_ollama_models", desc: "Lecture de /admin/ollama-models."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ollama_models", path: "/admin/ollama/models", listKey: "admin_ollama_models", desc: "Lecture de /admin/ollama/models."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_onboarding", path: "/admin/onboarding", listKey: "admin_onboarding", desc: "Lecture de /admin/onboarding."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_onboarding_pending", path: "/admin/onboarding/pending", listKey: "admin_onboarding_pending", desc: "Lecture de /admin/onboarding/pending."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ops_backups", path: "/admin/ops/backups", listKey: "admin_ops_backups", desc: "Lecture de /admin/ops/backups."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ops_build_history", path: "/admin/ops/build-history", listKey: "admin_ops_build_history", desc: "Lecture de /admin/ops/build-history."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_ops_services", path: "/admin/ops/services", listKey: "admin_ops_services", desc: "Lecture de /admin/ops/services."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_organizations", path: "/admin/organizations/", listKey: "admin_organizations", desc: "Lecture de /admin/organizations/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_orgs", path: "/admin/orgs", listKey: "admin_orgs", desc: "Lecture de /admin/orgs."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_orgs_quotas", path: "/admin/orgs/quotas", listKey: "admin_orgs_quotas", desc: "Lecture de /admin/orgs/quotas."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_plans", path: "/admin/plans/", listKey: "admin_plans", desc: "Lecture de /admin/plans/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_prometheus_alerts", path: "/admin/prometheus/alerts", listKey: "admin_prometheus_alerts", desc: "Lecture de /admin/prometheus/alerts."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_prometheus_query", path: "/admin/prometheus/query", listKey: "admin_prometheus_query", desc: "Lecture de /admin/prometheus/query."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_prometheus_targets", path: "/admin/prometheus/targets", listKey: "admin_prometheus_targets", desc: "Lecture de /admin/prometheus/targets."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_providers_capabilities", path: "/admin/providers/capabilities", listKey: "admin_providers_capabilities", desc: "Lecture de /admin/providers/capabilities."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_providers_config", path: "/admin/providers/config", listKey: "admin_providers_config", desc: "Lecture de /admin/providers/config."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_providers_dashboard", path: "/admin/providers/dashboard", listKey: "admin_providers_dashboard", desc: "Lecture de /admin/providers/dashboard."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_providers_keys_status", path: "/admin/providers/keys/status", listKey: "admin_providers_keys_status", desc: "Lecture de /admin/providers/keys/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_qdrant_collections", path: "/admin/qdrant/collections", listKey: "admin_qdrant_collections", desc: "Lecture de /admin/qdrant/collections."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_qdrant_info", path: "/admin/qdrant/info", listKey: "admin_qdrant_info", desc: "Lecture de /admin/qdrant/info."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_rate_limits", path: "/admin/rate-limits", listKey: "admin_rate_limits", desc: "Lecture de /admin/rate-limits."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_rbac", path: "/admin/rbac", listKey: "admin_rbac", desc: "Lecture de /admin/rbac."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_redis_info", path: "/admin/redis/info", listKey: "admin_redis_info", desc: "Lecture de /admin/redis/info."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_resellers", path: "/admin/resellers", listKey: "admin_resellers", desc: "Lecture de /admin/resellers."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_resellers", path: "/admin/resellers/", listKey: "admin_resellers", desc: "Lecture de /admin/resellers/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_scheduled_backups", path: "/admin/scheduled-backups", listKey: "admin_scheduled_backups", desc: "Lecture de /admin/scheduled-backups."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_comments", path: "/admin/seo/comments", listKey: "admin_seo_comments", desc: "Lecture de /admin/seo/comments."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_feeds", path: "/admin/seo/feeds", listKey: "admin_seo_feeds", desc: "Lecture de /admin/seo/feeds."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_posts", path: "/admin/seo/posts", listKey: "admin_seo_posts", desc: "Lecture de /admin/seo/posts."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_stats", path: "/admin/seo/stats", listKey: "admin_seo_stats", desc: "Lecture de /admin/seo/stats."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_seo_stats_global", path: "/admin/seo/stats/global", listKey: "admin_seo_stats_global", desc: "Lecture de /admin/seo/stats/global."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_sessions", path: "/admin/sessions", listKey: "admin_sessions", desc: "Lecture de /admin/sessions."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_sessions_bulk_revoke", path: "/admin/sessions/bulk-revoke", listKey: "admin_sessions_bulk_revoke", desc: "Lecture de /admin/sessions/bulk-revoke."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_swarm_services", path: "/admin/swarm/services", listKey: "admin_swarm_services", desc: "Lecture de /admin/swarm/services."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_tests_history", path: "/admin/tests/history", listKey: "admin_tests_history", desc: "Lecture de /admin/tests/history."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_usage", path: "/admin/usage/", listKey: "admin_usage", desc: "Lecture de /admin/usage/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_users", path: "/admin/users", listKey: "admin_users", desc: "Lecture de /admin/users."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_users", path: "/admin/users/", listKey: "admin_users", desc: "Lecture de /admin/users/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_vault", path: "/admin/vault", listKey: "admin_vault", desc: "Lecture de /admin/vault."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_verticals", path: "/admin/verticals", listKey: "admin_verticals", desc: "Lecture de /admin/verticals."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_verticals", path: "/admin/verticals/", listKey: "admin_verticals", desc: "Lecture de /admin/verticals/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_webhooks", path: "/admin/webhooks", listKey: "admin_webhooks", desc: "Lecture de /admin/webhooks."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_webhooks_resend", path: "/admin/webhooks/resend", listKey: "admin_webhooks_resend", desc: "Lecture de /admin/webhooks/resend."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "admin_yousign_status", path: "/admin/yousign/status", listKey: "admin_yousign_status", desc: "Lecture de /admin/yousign/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "api_investor_docs_access", path: "/api/investor-docs/access", listKey: "api_investor_docs_access", desc: "Lecture de /api/investor-docs/access."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_me_consents", path: "/auth/me/consents", listKey: "auth_me_consents", desc: "Lecture de /auth/me/consents."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_mfa_status", path: "/auth/mfa/status", listKey: "auth_mfa_status", desc: "Lecture de /auth/mfa/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "auth_oidc_providers", path: "/auth/oidc/providers", listKey: "auth_oidc_providers", desc: "Lecture de /auth/oidc/providers."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "billing_checkout_llm_catalog", path: "/billing/checkout/llm-catalog", listKey: "billing_checkout_llm_catalog", desc: "Lecture de /billing/checkout/llm-catalog."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "billing_checkout_status", path: "/billing/checkout/status", listKey: "billing_checkout_status", desc: "Lecture de /billing/checkout/status."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_deployable_services", path: "/org/deployable-services", listKey: "org_deployable_services", desc: "Lecture de /org/deployable-services."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_deployments", path: "/org/deployments", listKey: "org_deployments", desc: "Lecture de /org/deployments."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "org_members", path: "/org/members", listKey: "org_members", desc: "Lecture de /org/members."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "support_tickets", path: "/support/tickets/", listKey: "support_tickets", desc: "Lecture de /support/tickets/."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_conversations", path: "/v1/conversations", listKey: "v1_conversations", desc: "Lecture de /v1/conversations."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_local_models", path: "/v1/local-models", listKey: "v1_local_models", desc: "Lecture de /v1/local-models."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_me_providers", path: "/v1/me/providers", listKey: "v1_me_providers", desc: "Lecture de /v1/me/providers."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_me_quotas", path: "/v1/me/quotas", listKey: "v1_me_quotas", desc: "Lecture de /v1/me/quotas."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_models", path: "/v1/models", listKey: "v1_models", desc: "Lecture de /v1/models."}
	},
	func() datasource.DataSource {
		return &catalogDataSource{name: "v1_public_stats", path: "/v1/public/stats", listKey: "v1_public_stats", desc: "Lecture de /v1/public/stats."}
	},
}
