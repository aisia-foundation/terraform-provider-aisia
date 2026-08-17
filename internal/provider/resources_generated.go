// Code généré par scripts/generate/06_terraform_provider.py — NE PAS ÉDITER À LA MAIN.
// Régénérer : python3 scripts/generate/06_terraform_provider.py --openapi <spec>

package provider

import "github.com/hashicorp/terraform-plugin-framework/resource"

// generatedResources : resources WRITE générées depuis l'OpenAPI.
var generatedResources = []func() resource.Resource{
	func() resource.Resource {
		return &apiResource{name: "admin_ai_act_model_cards", path: "/admin/ai-act/model-cards", itemPathTemplate: "/admin/ai-act/model-cards/{card_id}", idParam: "card_id", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "annexe_iii_category", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ annexe_iii_category"}, {Name: "biases_identified", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ biases_identified"}, {Name: "ethical_considerations", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ ethical_considerations"}, {Name: "evaluation_data", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ evaluation_data"}, {Name: "intended_use", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ intended_use"}, {Name: "known_limitations", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ known_limitations"}, {Name: "model_name", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ model_name"}, {Name: "org_id", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ org_id"}, {Name: "out_of_scope_use", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ out_of_scope_use"}, {Name: "provider_id", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ provider_id"}, {Name: "published", TFType: "bool", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ published"}, {Name: "risk_level", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ risk_level"}, {Name: "training_data", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ training_data"}, {Name: "last_reviewed_at", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ last_reviewed_at"}}, createFields: []string{"annexe_iii_category", "biases_identified", "ethical_considerations", "evaluation_data", "intended_use", "known_limitations", "model_name", "org_id", "out_of_scope_use", "provider_id", "published", "risk_level", "training_data"}, updateFields: []string{"annexe_iii_category", "biases_identified", "ethical_considerations", "evaluation_data", "intended_use", "known_limitations", "last_reviewed_at", "model_name", "out_of_scope_use", "provider_id", "published", "risk_level", "training_data"}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_campaigns", path: "/admin/campaigns/", itemPathTemplate: "/admin/campaigns/{campaign_id}", idParam: "campaign_id", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "channel", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ channel"}, {Name: "content", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ content"}, {Name: "name", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ name"}, {Name: "scheduled_at", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ scheduled_at"}, {Name: "status", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ status"}, {Name: "target_audience", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ target_audience"}}, createFields: []string{"channel", "content", "name", "scheduled_at", "status", "target_audience"}, updateFields: []string{"content", "name", "scheduled_at", "target_audience"}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_events", path: "/admin/events/", itemPathTemplate: "/admin/events/{event_id}", idParam: "event_id", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "description", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ description"}, {Name: "duration_min", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ duration_min"}, {Name: "host_email", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ host_email"}, {Name: "max_attendees", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ max_attendees"}, {Name: "registration_url", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ registration_url"}, {Name: "scheduled_at", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ scheduled_at"}, {Name: "status", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ status"}, {Name: "title", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ title"}, {Name: "type", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ type"}, {Name: "recording_url", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ recording_url"}}, createFields: []string{"description", "duration_min", "host_email", "max_attendees", "registration_url", "scheduled_at", "status", "title", "type"}, updateFields: []string{"description", "max_attendees", "recording_url", "registration_url", "scheduled_at", "status", "title"}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_groups", path: "/admin/groups/", itemPathTemplate: "/admin/groups/{group_id}", idParam: "group_id", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "description", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ description"}, {Name: "name", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ name"}, {Name: "org_id", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ org_id"}}, createFields: []string{"description", "name", "org_id"}, updateFields: []string{"description", "name"}, updateRequiredFields: []string{"name"}}
	},
	func() resource.Resource {
		return &apiResource{name: "admin_skills", path: "/admin/skills", itemPathTemplate: "/admin/skills/{skill_id}", idParam: "skill_id", updateVerb: "PUT", canDelete: true, desc: "Endpoint d'administration (Bearer requis) — entité CRUD (create/read/update/delete via API).", typedAttrs: []apiAttr{{Name: "category", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ category"}, {Name: "description", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ description"}, {Name: "enabled", TFType: "bool", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ enabled"}, {Name: "invocation_type", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ invocation_type"}, {Name: "name", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ name"}, {Name: "org_id", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ org_id"}, {Name: "payload", TFType: "string", Required: false, Sensitive: false, JSONEncoded: true, Desc: "champ payload"}}, createFields: []string{"category", "description", "enabled", "invocation_type", "name", "org_id", "payload"}, updateFields: []string{"category", "description", "enabled", "invocation_type", "name", "payload"}}
	},
	func() resource.Resource {
		return &apiResource{name: "scim_v2_users", path: "/scim/v2/Users", itemPathTemplate: "/scim/v2/Users/{user_id}", idParam: "user_id", updateVerb: "PUT", canDelete: true, desc: "Endpoint API AISIA — entité CRUD (create/read/update/delete via API).", bodyFields: []string{"active", "displayName", "emails", "name", "schemas", "userName"}, createBodyRequiredFields: []string{"userName"}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_ai_rules_config", path: "/admin/ai-rules", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"rules"}, deepSubsetFields: []string{"rules"}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_billing_b2c_tier_limits_config", path: "/admin/billing/b2c-tier-limits", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"limits"}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_config_email_config", path: "/admin/config/email", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"from_email", "from_name", "smtp_host", "smtp_password", "smtp_port", "smtp_user", "use_tls"}, readFieldAliases: map[string]string{"smtp_host": "host", "smtp_port": "port", "smtp_user": "user"}, sensitiveFields: []string{"smtp_password"}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_config_identity_config", path: "/admin/config/identity", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"content", "yaml"}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_config_saml_config", path: "/admin/config/saml", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"clock_skew_s", "enabled", "idp_cert", "idp_entity_id", "idp_metadata_url", "idp_sso_url", "sp_acs_url", "sp_entity_id"}, writeOnlyFields: []string{"idp_cert"}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_multicloud_pack_plan_config", path: "/admin/multicloud/pack/plan", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"aws", "azure", "gcp", "ovh", "scaleway"}, responseObjectKey: "plan"}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "admin_routing_policy_config", path: "/admin/routing/policy", updateVerb: "PUT", desc: "Endpoint d'administration (Bearer requis) — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"cost_weight", "epsilon", "exploration_bonus", "latency_weight", "strategy"}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "org_account_config", path: "/org/account", updateVerb: "PUT", desc: "Endpoint API AISIA — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"display_name"}, responseObjectKey: "account"}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "org_branding_config", path: "/org/branding", updateVerb: "PUT", desc: "Endpoint API AISIA — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"accent_color", "company_name", "custom_domain", "favicon_url", "logo_url", "primary_color", "tagline"}, responseObjectKey: "branding"}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "org_budget_config", path: "/org/budget", updateVerb: "PUT", desc: "Endpoint API AISIA — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"alert_pct", "monthly_cap_eur"}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "org_oidc_config_config", path: "/org/oidc-config", updateVerb: "PUT", desc: "Endpoint API AISIA — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"client_id", "client_secret", "discovery_url", "domain", "email_domain", "enabled", "provider", "tenant_id"}, sensitiveFields: []string{"client_secret"}}
	},
	func() resource.Resource {
		return &singletonApiResource{name: "org_settings_config", path: "/org/settings", updateVerb: "PUT", desc: "Endpoint API AISIA — configuration singleton (GET+PUT/PATCH sur le même chemin).", bodyFields: []string{"name"}, responseObjectKey: "settings"}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_cycle_run_action", path: "/admin/agent/cycle/run", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/agent/cycle/run` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_email_run_action", path: "/admin/agent/email/run", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/agent/email/run` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_orchestration_toggle_action", path: "/admin/agent/orchestration/toggle", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/agent/orchestration/toggle` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"agent_code", "enabled"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_scheduler_reload_action", path: "/admin/agent/scheduler/reload", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/agent/scheduler/reload` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_tasks_task_id_approve_action", path: "", pathTemplate: "/admin/agent/tasks/{task_id}/approve", method: "POST", readPath: "/admin/agent/tasks", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/agent/tasks/{task_id}/approve` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agent_tasks_task_id_reject_action", path: "", pathTemplate: "/admin/agent/tasks/{task_id}/reject", method: "POST", readPath: "/admin/agent/tasks", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/agent/tasks/{task_id}/reject` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"reason"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agents_reload_custom_action", path: "/admin/agents/reload-custom", pathTemplate: "", method: "POST", readPath: "/admin/agents", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/agents/reload-custom` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_agents_code_trigger_action", path: "", pathTemplate: "/admin/agents/{code}/trigger", method: "POST", readPath: "/admin/agents", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/agents/{code}/trigger` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_annexe_iii_classifications_action", path: "/admin/ai-act/annexe-iii-classifications", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/annexe-iii-classifications", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/annexe-iii-classifications` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"annexe_iii_area", "classification_reason", "org_id", "use_case_description", "use_case_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_model_cards_purge_stubs_action", path: "/admin/ai-act/model-cards/purge-stubs", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/model-cards", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/model-cards/purge-stubs` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "dry_run", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_model_cards_review_bulk_action", path: "/admin/ai-act/model-cards/review-bulk", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/model-cards", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/model-cards/review-bulk` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"all_published", "card_ids", "note"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_model_cards_card_id_review_action", path: "", pathTemplate: "/admin/ai-act/model-cards/{card_id}/review", method: "POST", readPath: "/admin/ai-act/model-cards", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/model-cards/{card_id}/review` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_risk_assessments_action", path: "/admin/ai-act/risk-assessments", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/risk-assessments", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/risk-assessments` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"accuracy_score", "bias_score", "drift_score", "fairness_score", "findings", "mitigation_actions", "model_card_id", "next_review_at", "org_id", "overall_risk_level", "privacy_score", "status"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_audit_purge_action", path: "/admin/ai-act/v2/audit/purge", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/v2/audit", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/audit/purge` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_export_package_action", path: "/admin/ai-act/v2/export-package", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/export-package` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_incidents", path: "/admin/ai-act/v2/incidents", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/v2/incidents", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/incidents` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"category", "description", "impact_assessment", "metadata", "org_id", "severity"}, typedAttrs: []apiAttr{{Name: "category", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ category"}, {Name: "description", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ description"}, {Name: "impact_assessment", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ impact_assessment"}, {Name: "metadata", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ metadata"}, {Name: "org_id", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ org_id"}, {Name: "severity", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ severity"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_incidents_incident_id_report_authority_action", path: "", pathTemplate: "/admin/ai-act/v2/incidents/{incident_id}/report-authority", method: "POST", readPath: "/admin/ai-act/v2/incidents", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/incidents/{incident_id}/report-authority` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"authority_reference", "reference"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_kill_switch_action", path: "/admin/ai-act/v2/kill-switch", pathTemplate: "", method: "POST", readPath: "/admin/ai-act/v2/kill-switch", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/kill-switch` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"active", "by_user_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_model_cards_seed_action", path: "/admin/ai-act/v2/model-cards/seed", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/model-cards/seed` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "overwrite", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_oversight_oversight_id_approve_action", path: "", pathTemplate: "/admin/ai-act/v2/oversight/{oversight_id}/approve", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/oversight/{oversight_id}/approve` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"by_user_id", "comment"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_oversight_oversight_id_reject_action", path: "", pathTemplate: "/admin/ai-act/v2/oversight/{oversight_id}/reject", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/oversight/{oversight_id}/reject` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"by_user_id", "reason"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_watermark_add_action", path: "/admin/ai-act/v2/watermark/add", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/watermark/add` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"text", "visible"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_act_v2_watermark_detect_action", path: "/admin/ai-act/v2/watermark/detect", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-act/v2/watermark/detect` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"text"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ai_rules_reset_action", path: "/admin/ai-rules/reset", pathTemplate: "", method: "POST", readPath: "/admin/ai-rules", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ai-rules/reset` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_alerts", path: "/admin/alerts", pathTemplate: "", method: "POST", readPath: "/admin/alerts", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/alerts` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_alerts_alert_id_ack_action", path: "", pathTemplate: "/admin/alerts/{alert_id}/ack", method: "POST", readPath: "/admin/alerts", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/alerts/{alert_id}/ack` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_alerts_item_id_mutation", path: "", pathTemplate: "/admin/alerts/{item_id}", method: "DELETE", readPath: "/admin/alerts", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/alerts/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_alerts_item_id_put_mutation", path: "", pathTemplate: "/admin/alerts/{item_id}", method: "PUT", readPath: "/admin/alerts", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/alerts/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_analytics_action", path: "/admin/analytics", pathTemplate: "", method: "POST", readPath: "/admin/analytics", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/analytics` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_api_keys_rotation_action", path: "/admin/api-keys/rotation", pathTemplate: "", method: "POST", readPath: "/admin/api-keys/rotation", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/api-keys/rotation` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_api_keys_validate_all_action", path: "/admin/api-keys/validate-all", pathTemplate: "", method: "POST", readPath: "/admin/api-keys", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/api-keys/validate-all` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_api_keys_provider_id_mutation", path: "", pathTemplate: "/admin/api-keys/{provider_id}", method: "DELETE", readPath: "/admin/api-keys", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/api-keys/{provider_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_api_keys_provider_id_put_mutation", path: "", pathTemplate: "/admin/api-keys/{provider_id}", method: "PUT", readPath: "/admin/api-keys", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/api-keys/{provider_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"account_email", "key_value"}, sensitiveFields: []string{"key_value"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_auto_test_action", path: "/admin/auto-test", pathTemplate: "", method: "POST", readPath: "/admin/auto-test", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/auto-test` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_autonomy_reload_action", path: "/admin/autonomy/reload", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/autonomy/reload` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_b2b_approval_request_id_approve_action", path: "", pathTemplate: "/admin/b2b/approval/{request_id}/approve", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/b2b/approval/{request_id}/approve` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "confirm_rehome", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_b2b_approval_request_id_reject_action", path: "", pathTemplate: "/admin/b2b/approval/{request_id}/reject", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/b2b/approval/{request_id}/reject` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_backups_btype_create_action", path: "", pathTemplate: "/admin/backups/{btype}/create", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/backups/{btype}/create` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_backups_btype_restore_action", path: "", pathTemplate: "/admin/backups/{btype}/restore", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/backups/{btype}/restore` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"backup_file", "backup_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_blog_posts", path: "/admin/blog/posts", pathTemplate: "", method: "POST", readPath: "/admin/blog/posts", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/blog/posts` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"content_md", "excerpt", "id", "slug", "tags", "title"}, typedAttrs: []apiAttr{{Name: "content_md", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ content_md"}, {Name: "excerpt", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ excerpt"}, {Name: "request_id", TFType: "string", WireName: "id", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ id"}, {Name: "slug", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ slug"}, {Name: "tags", TFType: "list", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ tags"}, {Name: "title", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ title"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_blog_posts_post_id_mutation", path: "", pathTemplate: "/admin/blog/posts/{post_id}", method: "DELETE", readPath: "/admin/blog/posts", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/blog/posts/{post_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_blog_posts_post_id_put_mutation", path: "", pathTemplate: "/admin/blog/posts/{post_id}", method: "PUT", readPath: "/admin/blog/posts", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/blog/posts/{post_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"content_md", "excerpt", "slug", "status", "tags", "title"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_blog_posts_post_id_publish_action", path: "", pathTemplate: "/admin/blog/posts/{post_id}/publish", method: "POST", readPath: "/admin/blog/posts", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/blog/posts/{post_id}/publish` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_blacklist_reset_action", path: "/admin/bot/blacklist/reset", pathTemplate: "", method: "POST", readPath: "/admin/bot/blacklist", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/bot/blacklist/reset` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"provider_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_conversations", path: "/admin/bot/conversations", pathTemplate: "", method: "POST", readPath: "/admin/bot/conversations", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/bot/conversations` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_conversations_item_id_mutation", path: "", pathTemplate: "/admin/bot/conversations/{item_id}", method: "DELETE", readPath: "/admin/bot/conversations", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/bot/conversations/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_conversations_item_id_put_mutation", path: "", pathTemplate: "/admin/bot/conversations/{item_id}", method: "PUT", readPath: "/admin/bot/conversations", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/bot/conversations/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_credential_rotation_candidate_id_dismiss_action", path: "", pathTemplate: "/admin/bot/credential-rotation/{candidate_id}/dismiss", method: "POST", readPath: "/admin/bot/credential-rotation", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/bot/credential-rotation/{candidate_id}/dismiss` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"days"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_force_task_action", path: "/admin/bot/force-task", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/bot/force-task` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"kind", "task"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_run_now_action", path: "/admin/bot/run-now", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/bot/run-now` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_start_action", path: "/admin/bot/start", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/bot/start` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_bot_stop_action", path: "/admin/bot/stop", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/bot/stop` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cache_preload_action", path: "/admin/cache/preload", pathTemplate: "", method: "POST", readPath: "/admin/cache", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/cache/preload` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cache_purge_action", path: "/admin/cache/purge", pathTemplate: "", method: "POST", readPath: "/admin/cache", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/cache/purge` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "pattern", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_campaigns_campaign_id_email_blast_action", path: "", pathTemplate: "/admin/campaigns/{campaign_id}/email-blast", method: "POST", readPath: "/admin/campaigns/", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/campaigns/{campaign_id}/email-blast` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"recipients", "subject"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_campaigns_campaign_id_send_email_action", path: "", pathTemplate: "/admin/campaigns/{campaign_id}/send-email", method: "POST", readPath: "/admin/campaigns/", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/campaigns/{campaign_id}/send-email` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"recipients", "subject"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_campaigns_campaign_id_status_mutation", path: "", pathTemplate: "/admin/campaigns/{campaign_id}/status", method: "PUT", readPath: "/admin/campaigns/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/campaigns/{campaign_id}/status` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"status"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_catalog_sync_action", path: "/admin/catalog/sync", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/catalog/sync` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_certificates_issue_action", path: "/admin/certificates/issue", pathTemplate: "", method: "POST", readPath: "/admin/certificates", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/certificates/issue` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"allow_traefik_zone", "domain", "sans"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_certificates_domain_mutation", path: "", pathTemplate: "/admin/certificates/{domain}", method: "DELETE", readPath: "/admin/certificates", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/certificates/{domain}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_certificates_domain_renew_action", path: "", pathTemplate: "/admin/certificates/{domain}/renew", method: "POST", readPath: "/admin/certificates", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/certificates/{domain}/renew` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_circuit_breakers_provider_id_reset_action", path: "", pathTemplate: "/admin/circuit-breakers/{provider_id}/reset", method: "POST", readPath: "/admin/circuit-breakers", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/circuit-breakers/{provider_id}/reset` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cloud_models_provider_id_model_name_enabled_mutation", path: "", pathTemplate: "/admin/cloud-models/{provider_id}/{model_name}/enabled", method: "PUT", readPath: "/admin/cloud-models", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/cloud-models/{provider_id}/{model_name}/enabled` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"enabled", "is_default"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cloud_providers_item", path: "", pathTemplate: "/admin/cloud-providers/{provider_id}", method: "PUT", readPath: "/admin/cloud-providers", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/cloud-providers/{provider_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"config", "enabled"}, sensitiveFields: []string{"config"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cloud_providers_provider_id_provision_action", path: "", pathTemplate: "/admin/cloud-providers/{provider_id}/provision", method: "POST", readPath: "/admin/cloud-providers", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/cloud-providers/{provider_id}/provision` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cloud_providers_provider_id_test_action", path: "", pathTemplate: "/admin/cloud-providers/{provider_id}/test", method: "POST", readPath: "/admin/cloud-providers", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/cloud-providers/{provider_id}/test` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cluster_node_drain_action", path: "/admin/cluster/node/drain", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/cluster/node/drain` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"node"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cluster_prune_action", path: "/admin/cluster/prune", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/cluster/prune` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"target"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_cluster_scale_action", path: "/admin/cluster/scale", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/cluster/scale` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"replicas", "service"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_compliance_action", path: "/admin/compliance", pathTemplate: "", method: "POST", readPath: "/admin/compliance", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/compliance` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_email_action", path: "/admin/config/email", pathTemplate: "", method: "POST", readPath: "/admin/config/email", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/config/email` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"from_email", "from_name", "smtp_host", "smtp_password", "smtp_port", "smtp_user", "use_tls"}, sensitiveFields: []string{"smtp_password"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_email_templates_code_mutation", path: "", pathTemplate: "/admin/config/email/templates/{code}", method: "DELETE", readPath: "/admin/config/email/templates", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/config/email/templates/{code}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_email_templates_item", path: "", pathTemplate: "/admin/config/email/templates/{code}", method: "PUT", readPath: "/admin/config/email/templates", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/config/email/templates/{code}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"body_html", "body_text", "code", "enabled", "locale", "subject"}, typedAttrs: []apiAttr{{Name: "name", TFType: "string", Required: false, Sensitive: false, Desc: "champ name"}, {Name: "subject", TFType: "string", Required: false, Sensitive: false, Desc: "champ subject"}, {Name: "html", TFType: "string", Required: false, Sensitive: false, Desc: "champ html"}, {Name: "text", TFType: "string", Required: false, Sensitive: false, Desc: "champ text"}}, fieldAliases: map[string]string{"html": "body_html", "name": "code", "text": "body_text"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_email_test_action", path: "/admin/config/email/test", pathTemplate: "", method: "POST", readPath: "/admin/config/email", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/config/email/test` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"to"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_identity_reload_action", path: "/admin/config/identity/reload", pathTemplate: "", method: "POST", readPath: "/admin/config/identity", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/config/identity/reload` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_import_action", path: "/admin/config/import", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/config/import` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"groups", "guardrails"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_snapshots", path: "/admin/config/snapshots", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/config/snapshots` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_snapshots_item_id_mutation", path: "", pathTemplate: "/admin/config/snapshots/{item_id}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/config/snapshots/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_config_snapshots_item_id_put_mutation", path: "", pathTemplate: "/admin/config/snapshots/{item_id}", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/config/snapshots/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_catalog_import_openapi_action", path: "/admin/connectors-catalog/import-openapi", pathTemplate: "", method: "POST", readPath: "/admin/connectors-catalog", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors-catalog/import-openapi` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"service_id", "spec"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_catalog_service_id_mutation", path: "", pathTemplate: "/admin/connectors-catalog/{service_id}", method: "DELETE", readPath: "/admin/connectors-catalog", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/connectors-catalog/{service_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_catalog_service_id_action", path: "", pathTemplate: "/admin/connectors-catalog/{service_id}", method: "POST", readPath: "/admin/connectors-catalog", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors-catalog/{service_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"auth_type", "base_url", "category", "description", "display_name", "enabled_global", "endpoints", "oauth_authorize_url", "oauth_token_url", "pricing_tier", "scopes_default"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_catalog_item", path: "", pathTemplate: "/admin/connectors-catalog/{service_id}", method: "PUT", readPath: "/admin/connectors-catalog", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/connectors-catalog/{service_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"enabled_global"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_catalog_service_id_metadata_mutation", path: "", pathTemplate: "/admin/connectors-catalog/{service_id}/metadata", method: "PUT", readPath: "/admin/connectors-catalog", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/connectors-catalog/{service_id}/metadata` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"auth_type", "base_url", "category", "description", "display_name", "enabled_global", "endpoints", "oauth_authorize_url", "oauth_token_url", "pricing_tier", "scopes_default"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_mcp_servers", path: "/admin/connectors/mcp/servers", pathTemplate: "", method: "POST", readPath: "/admin/connectors/mcp/servers", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/mcp/servers` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"headers", "org_id", "server_id", "url"}, typedAttrs: []apiAttr{{Name: "headers", TFType: "string", Required: false, Sensitive: true, JSONEncoded: true, Desc: "champ headers"}, {Name: "org_id", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ org_id"}, {Name: "server_id", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ server_id"}, {Name: "url", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ url"}}, sensitiveFields: []string{"headers"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_mcp_servers_server_id_mutation", path: "", pathTemplate: "/admin/connectors/mcp/servers/{server_id}", method: "DELETE", readPath: "/admin/connectors/mcp/servers", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/connectors/mcp/servers/{server_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "org_id", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_mcp_servers_server_id_call_action", path: "", pathTemplate: "/admin/connectors/mcp/servers/{server_id}/call", method: "POST", readPath: "/admin/connectors/mcp/servers", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/mcp/servers/{server_id}/call` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "org_id", Required: false}}, bodyFields: []string{"arguments", "tool_name"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_schema_mapping_apply_action", path: "/admin/connectors/schema-mapping/apply", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/schema-mapping/apply` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"mapping_spec", "record", "records"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_schema_mapping_validate_action", path: "/admin/connectors/schema-mapping/validate", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/schema-mapping/validate` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"mapping_spec"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_run_action", path: "/admin/connectors/wizard/run", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/wizard/run` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"connector_id", "dry_run", "org_id", "spec"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_start_action", path: "/admin/connectors/wizard/start", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/wizard/start` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"connector_id", "dry_run", "org_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_session_id_activate_action", path: "", pathTemplate: "/admin/connectors/wizard/{session_id}/activate", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/wizard/{session_id}/activate` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_session_id_auth_action", path: "", pathTemplate: "/admin/connectors/wizard/{session_id}/auth", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/wizard/{session_id}/auth` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"auth_method", "credentials"}, sensitiveFields: []string{"credentials"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_session_id_mapping_action", path: "", pathTemplate: "/admin/connectors/wizard/{session_id}/mapping", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/wizard/{session_id}/mapping` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"mapping_spec", "sample_record"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_wizard_session_id_test_action", path: "", pathTemplate: "/admin/connectors/wizard/{session_id}/test", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/wizard/{session_id}/test` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_connectors_connector_id_vendor_updates_action", path: "", pathTemplate: "/admin/connectors/{connector_id}/vendor-updates", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/connectors/{connector_id}/vendor-updates` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"previous_snapshot"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_credentials_provider_mutation", path: "", pathTemplate: "/admin/credentials/{provider}", method: "DELETE", readPath: "/admin/credentials", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/credentials/{provider}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_credentials_item", path: "", pathTemplate: "/admin/credentials/{provider}", method: "PUT", readPath: "/admin/credentials", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/credentials/{provider}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"api_key", "fields"}, sensitiveFields: []string{"api_key", "fields"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_crews_orphans_drain_action", path: "/admin/crews/orphans/drain", pathTemplate: "", method: "POST", readPath: "/admin/crews", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/crews/orphans/drain` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"crew_name", "decision", "dry_run", "max_drain"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_crews_runs_run_id_approve_action", path: "", pathTemplate: "/admin/crews/runs/{run_id}/approve", method: "POST", readPath: "/admin/crews/runs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/crews/runs/{run_id}/approve` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"decision", "step_index"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_crews_name_run_action", path: "", pathTemplate: "/admin/crews/{name}/run", method: "POST", readPath: "/admin/crews", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/crews/{name}/run` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"payload"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_database_migrations_name_apply_action", path: "", pathTemplate: "/admin/database/migrations/{name}/apply", method: "POST", readPath: "/admin/database/migrations", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/database/migrations/{name}/apply` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_cleanup_action", path: "/admin/datasets/cleanup", pathTemplate: "", method: "POST", readPath: "/admin/datasets", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/datasets/cleanup` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_download_action", path: "/admin/datasets/download", pathTemplate: "", method: "POST", readPath: "/admin/datasets", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/datasets/download` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"dataset_id", "name", "source"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_dataset_id_mutation", path: "", pathTemplate: "/admin/datasets/{dataset_id}", method: "DELETE", readPath: "/admin/datasets", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/datasets/{dataset_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_dataset_id_redownload_action", path: "", pathTemplate: "/admin/datasets/{dataset_id}/redownload", method: "POST", readPath: "/admin/datasets", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/datasets/{dataset_id}/redownload` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_datasets_dataset_id_reingest_action", path: "", pathTemplate: "/admin/datasets/{dataset_id}/reingest", method: "POST", readPath: "/admin/datasets", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/datasets/{dataset_id}/reingest` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_db_migrations", path: "/admin/db/migrations", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/db/migrations` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_db_migrations_item_id_mutation", path: "", pathTemplate: "/admin/db/migrations/{item_id}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/db/migrations/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_db_migrations_item_id_put_mutation", path: "", pathTemplate: "/admin/db/migrations/{item_id}", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/db/migrations/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_deploy_worker_run_once_action", path: "/admin/deploy/worker/run-once", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/deploy/worker/run-once` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_discovery_run_action", path: "/admin/discovery/run", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/discovery/run` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_discovery_runs_action", path: "/admin/discovery/runs", pathTemplate: "", method: "POST", readPath: "/admin/discovery/runs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/discovery/runs` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_discovery_scans", path: "/admin/discovery/scans", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/discovery/scans` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_discovery_scans_item_id_mutation", path: "", pathTemplate: "/admin/discovery/scans/{item_id}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/discovery/scans/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_discovery_scans_item_id_put_mutation", path: "", pathTemplate: "/admin/discovery/scans/{item_id}", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/discovery/scans/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_dns_records", path: "/admin/dns/records", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/dns/records` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"field_type", "force_acme_subdomain", "sub_domain", "target", "ttl"}, typedAttrs: []apiAttr{{Name: "field_type", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ field_type"}, {Name: "force_acme_subdomain", TFType: "bool", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ force_acme_subdomain"}, {Name: "sub_domain", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ sub_domain"}, {Name: "target", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ target"}, {Name: "ttl", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ ttl"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_dns_records_record_id_mutation", path: "", pathTemplate: "/admin/dns/records/{record_id}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/dns/records/{record_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_dr_drill_action", path: "/admin/dr/drill", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/dr/drill` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "type", Required: true}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_dr_drills", path: "/admin/dr/drills", pathTemplate: "", method: "POST", readPath: "/admin/dr/drills", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/dr/drills` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_dr_drills_run_action", path: "/admin/dr/drills/run", pathTemplate: "", method: "POST", readPath: "/admin/dr/drills", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/dr/drills/run` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"dry_run", "notes", "scenario"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_dr_drills_item_id_mutation", path: "", pathTemplate: "/admin/dr/drills/{item_id}", method: "DELETE", readPath: "/admin/dr/drills", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/dr/drills/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_dr_drills_item_id_put_mutation", path: "", pathTemplate: "/admin/dr/drills/{item_id}", method: "PUT", readPath: "/admin/dr/drills", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/dr/drills/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_egress_test_action", path: "/admin/egress/test", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/egress/test` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_email_templates", path: "/admin/email-templates", pathTemplate: "", method: "POST", readPath: "/admin/email-templates", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/email-templates` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_email_templates_item_id_mutation", path: "", pathTemplate: "/admin/email-templates/{item_id}", method: "DELETE", readPath: "/admin/email-templates", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/email-templates/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_email_templates_item_id_put_mutation", path: "", pathTemplate: "/admin/email-templates/{item_id}", method: "PUT", readPath: "/admin/email-templates", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/email-templates/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_events_event_id_registrations_mark_attended_action", path: "", pathTemplate: "/admin/events/{event_id}/registrations/mark-attended", method: "POST", readPath: "/admin/events/", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/events/{event_id}/registrations/mark-attended` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"email"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_feature_flags", path: "/admin/feature-flags", pathTemplate: "", method: "POST", readPath: "/admin/feature-flags", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/feature-flags` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_feature_flags_item_id_mutation", path: "", pathTemplate: "/admin/feature-flags/{item_id}", method: "DELETE", readPath: "/admin/feature-flags", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/feature-flags/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_feature_flags_item_id_put_mutation", path: "", pathTemplate: "/admin/feature-flags/{item_id}", method: "PUT", readPath: "/admin/feature-flags", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/feature-flags/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_finetune_action", path: "/admin/finetune", pathTemplate: "", method: "POST", readPath: "/admin/finetune", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/finetune` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"base_model", "dataset", "epochs"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_finetune_jobs", path: "/admin/finetune/jobs", pathTemplate: "", method: "POST", readPath: "/admin/finetune", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/finetune/jobs` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_finetune_jobs_item_id_mutation", path: "", pathTemplate: "/admin/finetune/jobs/{item_id}", method: "DELETE", readPath: "/admin/finetune", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/finetune/jobs/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_finetune_jobs_item_id_put_mutation", path: "", pathTemplate: "/admin/finetune/jobs/{item_id}", method: "PUT", readPath: "/admin/finetune", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/finetune/jobs/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_delete_action", path: "/admin/gdpr/delete", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/gdpr/delete` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "user_id", Required: true}}, bodyFields: []string{"confirm_user_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_export_action", path: "/admin/gdpr/export", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/gdpr/export` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "user_id", Required: true}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_requests", path: "/admin/gdpr/requests", pathTemplate: "", method: "POST", readPath: "/admin/gdpr/requests", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/gdpr/requests` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_requests_delete_action", path: "/admin/gdpr/requests/delete", pathTemplate: "", method: "POST", readPath: "/admin/gdpr/requests", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/gdpr/requests/delete` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"confirm", "reason", "user_email"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_requests_export_action", path: "/admin/gdpr/requests/export", pathTemplate: "", method: "POST", readPath: "/admin/gdpr/requests", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/gdpr/requests/export` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"reason", "user_email"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_requests_item_id_mutation", path: "", pathTemplate: "/admin/gdpr/requests/{item_id}", method: "DELETE", readPath: "/admin/gdpr/requests", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/gdpr/requests/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_gdpr_requests_item_id_put_mutation", path: "", pathTemplate: "/admin/gdpr/requests/{item_id}", method: "PUT", readPath: "/admin/gdpr/requests", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/gdpr/requests/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_groups_group_id_members_action", path: "", pathTemplate: "/admin/groups/{group_id}/members", method: "POST", readPath: "/admin/groups/", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/groups/{group_id}/members` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"user_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_groups_group_id_members_mutation", path: "", pathTemplate: "/admin/groups/{group_id}/members", method: "PUT", readPath: "/admin/groups/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/groups/{group_id}/members` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"user_ids"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_groups_group_id_members_user_id_mutation", path: "", pathTemplate: "/admin/groups/{group_id}/members/{user_id}", method: "DELETE", readPath: "/admin/groups/", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/groups/{group_id}/members/{user_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_groups_group_id_permissions_mutation", path: "", pathTemplate: "/admin/groups/{group_id}/permissions", method: "PUT", readPath: "/admin/groups/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/groups/{group_id}/permissions` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_guardrails_mutation", path: "/admin/guardrails", pathTemplate: "", method: "PUT", readPath: "/admin/guardrails", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/guardrails` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true, bodyFields: []string{"allowed_origins", "blocked_patterns", "enabled", "max_request_size", "rate_limit_rpm"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_guardrails_gid_mutation", path: "", pathTemplate: "/admin/guardrails/{gid}", method: "DELETE", readPath: "/admin/guardrails", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/guardrails/{gid}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_guardrails_gid_put_mutation", path: "", pathTemplate: "/admin/guardrails/{gid}", method: "PUT", readPath: "/admin/guardrails", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/guardrails/{gid}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"action", "category", "enabled", "name", "pattern"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_healthcheck_run_action", path: "/admin/healthcheck/run", pathTemplate: "", method: "POST", readPath: "/admin/healthcheck", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/healthcheck/run` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_hierarchy_nodes", path: "/admin/hierarchy/nodes", pathTemplate: "", method: "POST", readPath: "/admin/hierarchy", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/hierarchy/nodes` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_hierarchy_nodes_item_id_mutation", path: "", pathTemplate: "/admin/hierarchy/nodes/{item_id}", method: "DELETE", readPath: "/admin/hierarchy", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/hierarchy/nodes/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_hierarchy_nodes_item_id_put_mutation", path: "", pathTemplate: "/admin/hierarchy/nodes/{item_id}", method: "PUT", readPath: "/admin/hierarchy", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/hierarchy/nodes/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_i18n_locales", path: "/admin/i18n/locales", pathTemplate: "", method: "POST", readPath: "/admin/i18n/locales", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/i18n/locales` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_i18n_locales_item_id_mutation", path: "", pathTemplate: "/admin/i18n/locales/{item_id}", method: "DELETE", readPath: "/admin/i18n/locales", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/i18n/locales/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_i18n_locales_lang_mutation", path: "", pathTemplate: "/admin/i18n/locales/{lang}", method: "PUT", readPath: "/admin/i18n/locales", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/i18n/locales/{lang}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_integrations", path: "/admin/integrations", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/integrations` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"api_version", "base_url", "category", "config", "connector_class", "connector_version", "display_name", "enabled", "id"}, typedAttrs: []apiAttr{{Name: "api_version", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ api_version"}, {Name: "base_url", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ base_url"}, {Name: "category", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ category"}, {Name: "config", TFType: "string", Required: false, Sensitive: true, JSONEncoded: true, Desc: "champ config"}, {Name: "connector_class", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ connector_class"}, {Name: "connector_version", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ connector_version"}, {Name: "display_name", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ display_name"}, {Name: "enabled", TFType: "bool", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ enabled"}, {Name: "request_id", TFType: "string", WireName: "id", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ id"}}, sensitiveFields: []string{"config"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_integrations_integration_id_mutation", path: "", pathTemplate: "/admin/integrations/{integration_id}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/integrations/{integration_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_integrations_integration_id_put_mutation", path: "", pathTemplate: "/admin/integrations/{integration_id}", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/integrations/{integration_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"api_version", "base_url", "category", "config", "connector_class", "connector_version", "display_name", "enabled", "id"}, sensitiveFields: []string{"config"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_integrations_integration_id_enable_action", path: "", pathTemplate: "/admin/integrations/{integration_id}/enable", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/integrations/{integration_id}/enable` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"enabled"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_integrations_integration_id_test_action", path: "", pathTemplate: "/admin/integrations/{integration_id}/test", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/integrations/{integration_id}/test` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"body", "event_type", "severity", "title"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_integrations_provider_adapter_test_action", path: "", pathTemplate: "/admin/integrations/{provider}/adapter-test", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/integrations/{provider}/adapter-test` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_investors_ndas_nda_id_grant_action", path: "", pathTemplate: "/admin/investors/ndas/{nda_id}/grant", method: "POST", readPath: "/admin/investors/ndas", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/investors/ndas/{nda_id}/grant` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"tech_doc"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_investors_ndas_nda_id_revoke_action", path: "", pathTemplate: "/admin/investors/ndas/{nda_id}/revoke", method: "POST", readPath: "/admin/investors/ndas", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/investors/ndas/{nda_id}/revoke` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_investors_investor_id_send_nda_action", path: "", pathTemplate: "/admin/investors/{investor_id}/send-nda", method: "POST", readPath: "/admin/investors", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/investors/{investor_id}/send-nda` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"investor_company", "investor_email", "investor_name"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_knowledge_clear_action", path: "/admin/knowledge/clear", pathTemplate: "", method: "POST", readPath: "/admin/knowledge", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/knowledge/clear` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_knowledge_ingest_action", path: "/admin/knowledge/ingest", pathTemplate: "", method: "POST", readPath: "/admin/knowledge", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/knowledge/ingest` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"source", "text", "url"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_leads_item", path: "", pathTemplate: "/admin/leads/{lead_id}", method: "PUT", readPath: "/admin/leads", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/leads/{lead_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"assigned_to", "notes", "status"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_learning_finetune_action", path: "/admin/learning/finetune", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/learning/finetune` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"batch_size", "dataset_id", "epochs", "learning_rate", "lora_alpha", "lora_rank", "model_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_learning_finetune_runs", path: "/admin/learning/finetune-runs", pathTemplate: "", method: "POST", readPath: "/admin/learning/finetune-runs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/learning/finetune-runs` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"base_model", "dataset", "epochs", "lr", "notes"}, typedAttrs: []apiAttr{{Name: "base_model", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ base_model"}, {Name: "dataset", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ dataset"}, {Name: "epochs", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ epochs"}, {Name: "lr", TFType: "float", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ lr"}, {Name: "notes", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ notes"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_learning_proposals_proposal_id_apply_action", path: "", pathTemplate: "/admin/learning/proposals/{proposal_id}/apply", method: "POST", readPath: "/admin/learning/proposals", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/learning/proposals/{proposal_id}/apply` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_learning_public_sources_harvest_action", path: "/admin/learning/public-sources/harvest", pathTemplate: "", method: "POST", readPath: "/admin/learning/public-sources", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/learning/public-sources/harvest` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "lane", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_learning_run_action", path: "/admin/learning/run", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/learning/run` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_licenses_action", path: "/admin/licenses", pathTemplate: "", method: "POST", readPath: "/admin/licenses", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/licenses` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_linkedin_disconnect_mutation", path: "/admin/linkedin/disconnect", pathTemplate: "", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/linkedin/disconnect` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_linkedin_post_action", path: "/admin/linkedin/post", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/linkedin/post` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"text"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_local_models_model_id_mutation", path: "", pathTemplate: "/admin/local-models/{model_id}", method: "DELETE", readPath: "/admin/local-models", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/local-models/{model_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_local_models_model_id_put_mutation", path: "", pathTemplate: "/admin/local-models/{model_id}", method: "PUT", readPath: "/admin/local-models", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/local-models/{model_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"enabled", "priority"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_local_models_model_id_enabled_mutation", path: "", pathTemplate: "/admin/local-models/{model_id}/enabled", method: "PUT", readPath: "/admin/local-models", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/local-models/{model_id}/enabled` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "enabled", Required: false}}, bodyFields: []string{"enabled"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_maintenance_reload_action", path: "/admin/maintenance/reload", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/maintenance/reload` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_maintenance_tasks", path: "/admin/maintenance/tasks", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/maintenance/tasks` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"category", "config", "description", "enabled", "interval_seconds", "name", "schedule", "task_id"}, typedAttrs: []apiAttr{{Name: "category", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ category"}, {Name: "config", TFType: "string", Required: false, Sensitive: true, JSONEncoded: true, Desc: "champ config"}, {Name: "description", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ description"}, {Name: "enabled", TFType: "bool", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ enabled"}, {Name: "interval_seconds", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ interval_seconds"}, {Name: "name", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ name"}, {Name: "schedule", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ schedule"}, {Name: "task_id", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ task_id"}}, sensitiveFields: []string{"config"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_maintenance_tasks_task_id_mutation", path: "", pathTemplate: "/admin/maintenance/tasks/{task_id}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/maintenance/tasks/{task_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_maintenance_tasks_task_id_put_mutation", path: "", pathTemplate: "/admin/maintenance/tasks/{task_id}", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/maintenance/tasks/{task_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"category", "config", "description", "enabled", "interval_seconds", "name", "schedule"}, sensitiveFields: []string{"config"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_maintenance_tasks_task_id_enable_action", path: "", pathTemplate: "/admin/maintenance/tasks/{task_id}/enable", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/maintenance/tasks/{task_id}/enable` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "enabled", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_maintenance_tasks_task_id_run_action", path: "", pathTemplate: "/admin/maintenance/tasks/{task_id}/run", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/maintenance/tasks/{task_id}/run` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "async_run", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_mcp_servers", path: "/admin/mcp-servers", pathTemplate: "", method: "POST", readPath: "/admin/mcp-servers", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/mcp-servers` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"headers", "server_id", "url"}, typedAttrs: []apiAttr{{Name: "headers", TFType: "string", Required: false, Sensitive: true, JSONEncoded: true, Desc: "champ headers"}, {Name: "server_id", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ server_id"}, {Name: "url", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ url"}}, sensitiveFields: []string{"headers"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_mcp_servers_server_id_mutation", path: "", pathTemplate: "/admin/mcp-servers/{server_id}", method: "DELETE", readPath: "/admin/mcp-servers", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/mcp-servers/{server_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_models_sync_action", path: "/admin/models/sync", pathTemplate: "", method: "POST", readPath: "/admin/models", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/models/sync` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_multicloud_pack_dry_run_action", path: "/admin/multicloud/pack/dry-run", pathTemplate: "", method: "POST", readPath: "/admin/multicloud/pack", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/multicloud/pack/dry-run` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"cloud", "tier"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_multicloud_pack_dry_run_plan_action", path: "/admin/multicloud/pack/dry-run-plan", pathTemplate: "", method: "POST", readPath: "/admin/multicloud/pack", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/multicloud/pack/dry-run-plan` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_multitenant_org_domains_verify_all_action", path: "/admin/multitenant/org-domains/verify-all", pathTemplate: "", method: "POST", readPath: "/admin/multitenant/org-domains", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/multitenant/org-domains/verify-all` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_multitenant_org_domains_org_id_verify_action", path: "", pathTemplate: "/admin/multitenant/org-domains/{org_id}/verify", method: "POST", readPath: "/admin/multitenant/org-domains", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/multitenant/org-domains/{org_id}/verify` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_network_action", path: "/admin/network", pathTemplate: "", method: "POST", readPath: "/admin/network", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/network` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_oidc_providers", path: "/admin/oidc-providers", pathTemplate: "", method: "POST", readPath: "/admin/oidc-providers", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/oidc-providers` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true, bodyFields: []string{"config", "enabled", "provider_id"}, typedAttrs: []apiAttr{{Name: "config", TFType: "string", Required: false, Sensitive: true, JSONEncoded: true, Desc: "champ config"}, {Name: "enabled", TFType: "bool", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ enabled"}, {Name: "provider_id", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ provider_id"}}, sensitiveFields: []string{"config"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_oidc_providers_provider_id_mutation", path: "", pathTemplate: "/admin/oidc-providers/{provider_id}", method: "DELETE", readPath: "/admin/oidc-providers", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/oidc-providers/{provider_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_oidc_providers_provider_id_put_mutation", path: "", pathTemplate: "/admin/oidc-providers/{provider_id}", method: "PUT", readPath: "/admin/oidc-providers", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/oidc-providers/{provider_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"config", "enabled"}, sensitiveFields: []string{"config"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ollama_models_model_name_mutation", path: "", pathTemplate: "/admin/ollama/models/{model_name}", method: "DELETE", readPath: "/admin/ollama/models", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/ollama/models/{model_name}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "tier", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ollama_pull_action", path: "/admin/ollama/pull", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ollama/pull` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"model", "tier"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_onboarding_onboarding_id_review_mutation", path: "", pathTemplate: "/admin/onboarding/{onboarding_id}/review", method: "PUT", readPath: "/admin/onboarding", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/onboarding/{onboarding_id}/review` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"approved", "rejection_reason"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_backup_action", path: "/admin/ops/backup", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ops/backup` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true, bodyFields: []string{"backup_type", "scope"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_build_action", path: "/admin/ops/build", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ops/build` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"arch", "service", "version"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_deploy_action", path: "/admin/ops/deploy", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ops/deploy` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true, bodyFields: []string{"confirm", "force", "image", "image_tag", "parallelism", "service"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_rollback_action", path: "/admin/ops/rollback", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ops/rollback` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true, bodyFields: []string{"confirm", "service"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_ops_scale_action", path: "/admin/ops/scale", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/ops/scale` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true, bodyFields: []string{"replicas", "service"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}", method: "DELETE", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/organizations/{org_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_put_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}", method: "PUT", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/organizations/{org_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"contract_type", "deploy_channel", "description", "max_requests_day", "max_tokens_day", "max_users", "name", "plan", "slug", "support_tier"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_approve_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/approve", method: "PUT", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/organizations/{org_id}/approve` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_branding_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/branding", method: "PUT", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/organizations/{org_id}/branding` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"accent_color", "company_name", "custom_domain", "favicon_url", "logo_url", "primary_color", "tagline"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_custom_domain_provision_action", path: "", pathTemplate: "/admin/organizations/{org_id}/custom-domain/provision", method: "POST", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/organizations/{org_id}/custom-domain/provision` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_licence_action", path: "", pathTemplate: "/admin/organizations/{org_id}/licence", method: "POST", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/organizations/{org_id}/licence` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"contract_type", "expires_days", "features", "max_nodes", "max_users"}, sensitiveFields: []string{"licence_key"}, secretOutputs: []string{"licence_key"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_llm_access_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/llm-access", method: "PUT", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/organizations/{org_id}/llm-access` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"allowed", "group_id", "provider_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_members_action", path: "", pathTemplate: "/admin/organizations/{org_id}/members", method: "POST", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/organizations/{org_id}/members` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"role", "user_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_members_user_id_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/members/{user_id}", method: "DELETE", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/organizations/{org_id}/members/{user_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_members_user_id_put_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/members/{user_id}", method: "PUT", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/organizations/{org_id}/members/{user_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"role"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_purge_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/purge", method: "DELETE", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/organizations/{org_id}/purge` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_reactivate_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/reactivate", method: "PUT", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/organizations/{org_id}/reactivate` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_organizations_org_id_suspend_mutation", path: "", pathTemplate: "/admin/organizations/{org_id}/suspend", method: "PUT", readPath: "/admin/organizations/", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/organizations/{org_id}/suspend` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"reason"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_quotas_item", path: "", pathTemplate: "/admin/orgs/quotas/{org_id}", method: "PUT", readPath: "/admin/orgs/quotas", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/quotas/{org_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"daily_requests_cap", "monthly_tokens_cap", "notes", "org_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"description", "max_requests_day", "max_tokens_day", "max_users", "name", "plan", "status", "support_tier"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_api_keys_action", path: "", pathTemplate: "/admin/orgs/{org_id}/api-keys", method: "POST", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/orgs/{org_id}/api-keys` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"allowed_models", "expires_at", "max_budget_eur", "name", "rate_limit_rpd", "rate_limit_rpm", "rate_limit_tpm", "scopes"}, sensitiveFields: []string{"key"}, secretOutputs: []string{"key"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_api_keys_key_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/api-keys/{key_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/api-keys/{key_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_branding_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/branding", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/branding` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"brand_footer_text", "brand_logo_url", "brand_primary_color", "brand_secondary_color"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_bundles_action", path: "", pathTemplate: "/admin/orgs/{org_id}/config/bundles", method: "POST", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/orgs/{org_id}/config/bundles` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"bundle_key", "description", "includes", "name"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_bundles_seed_defaults_action", path: "", pathTemplate: "/admin/orgs/{org_id}/config/bundles/seed-defaults", method: "POST", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/orgs/{org_id}/config/bundles/seed-defaults` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_bundles_bundle_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/bundles/{bundle_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/config/bundles/{bundle_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_bundles_bundle_id_assign_action", path: "", pathTemplate: "/admin/orgs/{org_id}/config/bundles/{bundle_id}/assign", method: "POST", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/orgs/{org_id}/config/bundles/{bundle_id}/assign` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_llm_access_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/llm-access", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/config/llm-access` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"allowed", "model", "priority", "provider", "scope_type", "scope_value"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_llm_access_entry_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/llm-access/{entry_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/config/llm-access/{entry_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_quotas_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/quotas", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/config/quotas` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"cost_usd_day", "requests_day", "scope_type", "scope_value", "tokens_day", "tokens_month"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_config_quotas_entry_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/config/quotas/{entry_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/config/quotas/{entry_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_email_templates_key_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/email-templates/{key}", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/email-templates/{key}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_email_templates_key_put_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/email-templates/{key}", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/email-templates/{key}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"body_html", "body_text", "locale", "subject"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_invitations_action", path: "", pathTemplate: "/admin/orgs/{org_id}/invitations", method: "POST", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/orgs/{org_id}/invitations` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"email", "role"}, sensitiveFields: []string{"accept_url"}, secretOutputs: []string{"accept_url"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_invitations_invitation_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/invitations/{invitation_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/invitations/{invitation_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_oidc_providers_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/oidc-providers", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/oidc-providers` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_oidc_providers_put_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/oidc-providers", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/oidc-providers` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"client_id", "client_secret", "discovery_url", "domain", "email_domain", "enabled", "provider", "tenant_id"}, sensitiveFields: []string{"client_secret"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_provider_keys_provider_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/provider-keys/{provider_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/provider-keys/{provider_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_provider_keys_provider_id_put_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/provider-keys/{provider_id}", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/provider-keys/{provider_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"account_email", "key_value", "owner", "rotation_days"}, sensitiveFields: []string{"key_value"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_quota_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/quota", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/quota` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"max_requests_day", "max_requests_minute", "max_tokens_day", "max_users"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_routing_mode_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/routing-mode", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/routing-mode` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"mode"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_security_policy_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/security-policy", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/security-policy` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"allow_signup", "enforce_2fa", "ip_allowlist", "password_min_length", "session_ttl_minutes"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_users_action", path: "", pathTemplate: "/admin/orgs/{org_id}/users", method: "POST", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/orgs/{org_id}/users` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"display_name", "email", "password", "role"}, sensitiveFields: []string{"password"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_users_user_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/users/{user_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/users/{user_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_users_user_id_put_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/users/{user_id}", method: "PUT", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/orgs/{org_id}/users/{user_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"active", "display_name", "role"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_webhooks_action", path: "", pathTemplate: "/admin/orgs/{org_id}/webhooks", method: "POST", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/orgs/{org_id}/webhooks` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"events", "name", "secret", "url"}, sensitiveFields: []string{"secret"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_orgs_org_id_webhooks_webhook_id_mutation", path: "", pathTemplate: "/admin/orgs/{org_id}/webhooks/{webhook_id}", method: "DELETE", readPath: "/admin/orgs", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/orgs/{org_id}/webhooks/{webhook_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_plans", path: "/admin/plans/", pathTemplate: "", method: "POST", readPath: "/admin/plans/", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/plans/` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"billing_period", "max_requests_day", "max_tokens_day", "max_users", "name", "plan_type", "price_eur", "slug"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_plans_slug_mutation", path: "", pathTemplate: "/admin/plans/{slug}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/plans/{slug}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_plans_slug_put_mutation", path: "", pathTemplate: "/admin/plans/{slug}", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/plans/{slug}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"active", "billing_period", "custom_quote", "max_requests_day", "max_requests_minute", "max_tokens_day", "max_users", "name", "plan_type", "price_eur", "tokens_overage_eur_per_1k"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_provider_pricing_item", path: "", pathTemplate: "/admin/provider-pricing/{provider_id}", method: "PUT", readPath: "/admin/provider-pricing", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/provider-pricing/{provider_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"eur_per_1k_in", "eur_per_1k_out"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_keys_rotate_cycle_action", path: "/admin/providers/keys/rotate-cycle", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/providers/keys/rotate-cycle` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_keys_verify_action", path: "/admin/providers/keys/verify", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/providers/keys/verify` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_mutation", path: "", pathTemplate: "/admin/providers/{provider_id}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/providers/{provider_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_put_mutation", path: "", pathTemplate: "/admin/providers/{provider_id}", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/providers/{provider_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true, bodyFields: []string{"adapter_name", "api_key_env", "base_url", "display_name", "enabled", "id", "models", "name", "priority"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_auto_recover_action", path: "", pathTemplate: "/admin/providers/{provider_id}/auto-recover", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/providers/{provider_id}/auto-recover` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_config_mutation", path: "", pathTemplate: "/admin/providers/{provider_id}/config", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/providers/{provider_id}/config` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"config"}, sensitiveFields: []string{"config"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_enabled_mutation", path: "", pathTemplate: "/admin/providers/{provider_id}/enabled", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/providers/{provider_id}/enabled` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"enabled"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_key_action", path: "", pathTemplate: "/admin/providers/{provider_id}/key", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/providers/{provider_id}/key` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"api_key", "key"}, sensitiveFields: []string{"api_key"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_providers_provider_id_key_rotate_action", path: "", pathTemplate: "/admin/providers/{provider_id}/key/rotate", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/providers/{provider_id}/key/rotate` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_proxy_usage_retention_run_action", path: "/admin/proxy-usage/retention/run", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/proxy-usage/retention/run` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_proxy_usage_users_principal_id_mutation", path: "", pathTemplate: "/admin/proxy-usage/users/{principal_id}", method: "DELETE", readPath: "/admin/proxy-usage/users", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/proxy-usage/users/{principal_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "org_id", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_qdrant_bootstrap_action", path: "/admin/qdrant/bootstrap", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/qdrant/bootstrap` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_rate_limits_action", path: "/admin/rate-limits", pathTemplate: "", method: "POST", readPath: "/admin/rate-limits", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/rate-limits` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_rbac_action", path: "/admin/rbac", pathTemplate: "", method: "POST", readPath: "/admin/rbac", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/rbac` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_resellers_reseller_id_mutation", path: "", pathTemplate: "/admin/resellers/{reseller_id}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/resellers/{reseller_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_resellers_reseller_id_put_mutation", path: "", pathTemplate: "/admin/resellers/{reseller_id}", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/resellers/{reseller_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"branding_config_id", "contact_email", "domain", "name", "notes", "rev_share_pct", "status"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_scheduled_backups_action", path: "/admin/scheduled-backups", pathTemplate: "", method: "POST", readPath: "/admin/scheduled-backups", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/scheduled-backups` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_self_repair_action", path: "/admin/self-repair", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/self-repair` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_app_credentials_item", path: "", pathTemplate: "/admin/seo/app-credentials/{network}", method: "PUT", readPath: "/admin/seo/app-credentials", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/seo/app-credentials/{network}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"client_id", "client_secret"}, sensitiveFields: []string{"client_secret"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_comments_comment_id_approve_reply_action", path: "", pathTemplate: "/admin/seo/comments/{comment_id}/approve-reply", method: "POST", readPath: "/admin/seo/comments", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/seo/comments/{comment_id}/approve-reply` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"content"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_feeds", path: "/admin/seo/feeds", pathTemplate: "", method: "POST", readPath: "/admin/seo/feeds", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/seo/feeds` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"category", "id", "name", "url"}, typedAttrs: []apiAttr{{Name: "category", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ category"}, {Name: "request_id", TFType: "string", WireName: "id", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ id"}, {Name: "name", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ name"}, {Name: "url", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ url"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_feeds_feed_id_mutation", path: "", pathTemplate: "/admin/seo/feeds/{feed_id}", method: "DELETE", readPath: "/admin/seo/feeds", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/seo/feeds/{feed_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_feeds_feed_id_put_mutation", path: "", pathTemplate: "/admin/seo/feeds/{feed_id}", method: "PUT", readPath: "/admin/seo/feeds", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/seo/feeds/{feed_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"category", "enabled", "name"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_linkedin_force_publish_action", path: "/admin/seo/linkedin/force-publish", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/seo/linkedin/force-publish` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_posts_post_id_approve_action", path: "", pathTemplate: "/admin/seo/posts/{post_id}/approve", method: "POST", readPath: "/admin/seo/posts", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/seo/posts/{post_id}/approve` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_posts_post_id_reject_action", path: "", pathTemplate: "/admin/seo/posts/{post_id}/reject", method: "POST", readPath: "/admin/seo/posts", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/seo/posts/{post_id}/reject` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"reason"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_seo_scheduler_action", path: "/admin/seo/scheduler", pathTemplate: "", method: "POST", readPath: "/admin/seo/scheduler", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/seo/scheduler` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_sessions_bulk_revoke_action", path: "/admin/sessions/bulk-revoke", pathTemplate: "", method: "POST", readPath: "/admin/sessions/bulk-revoke", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/sessions/bulk-revoke` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"confirm", "exclude_session_ids", "session_ids", "user_email"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_sessions_sid_revoke_action", path: "", pathTemplate: "/admin/sessions/{sid}/revoke", method: "POST", readPath: "/admin/sessions", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/sessions/{sid}/revoke` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_tests_run_action", path: "/admin/tests/run", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/tests/run` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"profile", "spec_file"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_tests_suite_run_action", path: "", pathTemplate: "/admin/tests/{suite}/run", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/tests/{suite}/run` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_token_quotas", path: "/admin/token-quotas", pathTemplate: "", method: "POST", readPath: "/admin/token-quotas", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/token-quotas` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_token_quotas_item_id_mutation", path: "", pathTemplate: "/admin/token-quotas/{item_id}", method: "DELETE", readPath: "/admin/token-quotas", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/token-quotas/{item_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_token_quotas_item_id_put_mutation", path: "", pathTemplate: "/admin/token-quotas/{item_id}", method: "PUT", readPath: "/admin/token-quotas", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/token-quotas/{item_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_mutation", path: "", pathTemplate: "/admin/users/{user_id}", method: "DELETE", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/users/{user_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_put_mutation", path: "", pathTemplate: "/admin/users/{user_id}", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/users/{user_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"active", "display_name", "email", "role", "user_type"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_active_mutation", path: "", pathTemplate: "/admin/users/{user_id}/active", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/users/{user_id}/active` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"active"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_assign_org_action", path: "", pathTemplate: "/admin/users/{user_id}/assign-org", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/users/{user_id}/assign-org` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"group_id", "org_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_disable_mfa_action", path: "", pathTemplate: "/admin/users/{user_id}/disable-mfa", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/users/{user_id}/disable-mfa` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_display_name_mutation", path: "", pathTemplate: "/admin/users/{user_id}/display-name", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/users/{user_id}/display-name` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"display_name"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_force_logout_action", path: "", pathTemplate: "/admin/users/{user_id}/force-logout", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/users/{user_id}/force-logout` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_force_mfa_action", path: "", pathTemplate: "/admin/users/{user_id}/force-mfa", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/users/{user_id}/force-mfa` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_lock_action", path: "", pathTemplate: "/admin/users/{user_id}/lock", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/users/{user_id}/lock` (resource action Terraform).", hasJSONBody: true, bodyRequired: false, allowAdditionalBodyFields: false, bodyFields: []string{"reason"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_permissions_mutation", path: "", pathTemplate: "/admin/users/{user_id}/permissions", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/users/{user_id}/permissions` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_profile_mutation", path: "", pathTemplate: "/admin/users/{user_id}/profile", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/users/{user_id}/profile` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true, bodyFields: []string{"address", "city", "company", "country", "first_name", "last_name", "phone", "subscription_type", "token_quota_daily", "token_quota_monthly", "token_quota_weekly"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_quota_mutation", path: "", pathTemplate: "/admin/users/{user_id}/quota", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/users/{user_id}/quota` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"daily", "monthly", "weekly"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_reset_password_action", path: "", pathTemplate: "/admin/users/{user_id}/reset-password", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/users/{user_id}/reset-password` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_role_mutation", path: "", pathTemplate: "/admin/users/{user_id}/role", method: "PUT", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/users/{user_id}/role` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"role"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_users_user_id_unlock_action", path: "", pathTemplate: "/admin/users/{user_id}/unlock", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/users/{user_id}/unlock` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_vault", path: "/admin/vault", pathTemplate: "", method: "POST", readPath: "/admin/vault", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/vault` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"org_id", "secret_name", "value"}, typedAttrs: []apiAttr{{Name: "org_id", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ org_id"}, {Name: "secret_name", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ secret_name"}, {Name: "value", TFType: "string", Required: true, Sensitive: true, JSONEncoded: false, Desc: "champ value"}}, sensitiveFields: []string{"value"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_vault_secret_name_mutation", path: "", pathTemplate: "/admin/vault/{secret_name}", method: "DELETE", readPath: "/admin/vault", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/vault/{secret_name}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, queryParams: []actionQueryParam{{Name: "org_id", Required: false}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhook_alert_action", path: "/admin/webhook/alert", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/webhook/alert` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"alerts", "status"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_resend_action", path: "/admin/webhooks/resend", pathTemplate: "", method: "POST", readPath: "/admin/webhooks/resend", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/webhooks/resend` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_webhook_id_mutation", path: "", pathTemplate: "/admin/webhooks/{webhook_id}", method: "DELETE", readPath: "/admin/webhooks", desc: "Endpoint d'administration (Bearer requis) — mutation DELETE exacte `/admin/webhooks/{webhook_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_webhook_id_put_mutation", path: "", pathTemplate: "/admin/webhooks/{webhook_id}", method: "PUT", readPath: "/admin/webhooks", desc: "Endpoint d'administration (Bearer requis) — mutation PUT exacte `/admin/webhooks/{webhook_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"active", "events", "name", "secret", "url"}, sensitiveFields: []string{"secret"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "admin_webhooks_webhook_id_test_action", path: "", pathTemplate: "/admin/webhooks/{webhook_id}/test", method: "POST", readPath: "/admin/webhooks", desc: "Endpoint d'administration (Bearer requis) — mutation POST exacte `/admin/webhooks/{webhook_id}/test` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_agents", path: "/org/agents", pathTemplate: "", method: "POST", readPath: "/org/agents", desc: "Endpoint API AISIA — mutation POST exacte `/org/agents` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"allowed_tools", "description", "max_steps", "max_tokens", "model", "model_id", "name", "schedule_interval_s", "system_prompt", "temperature"}, typedAttrs: []apiAttr{{Name: "allowed_tools", TFType: "list", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ allowed_tools"}, {Name: "description", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ description"}, {Name: "max_steps", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ max_steps"}, {Name: "max_tokens", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ max_tokens"}, {Name: "model", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ model"}, {Name: "model_id", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ model_id"}, {Name: "name", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ name"}, {Name: "schedule_interval_s", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ schedule_interval_s"}, {Name: "system_prompt", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ system_prompt"}, {Name: "temperature", TFType: "float", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ temperature"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_agents_code_mutation", path: "", pathTemplate: "/org/agents/{code}", method: "DELETE", readPath: "/org/agents", desc: "Endpoint API AISIA — mutation DELETE exacte `/org/agents/{code}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_agents_code_put_mutation", path: "", pathTemplate: "/org/agents/{code}", method: "PUT", readPath: "/org/agents", desc: "Endpoint API AISIA — mutation PUT exacte `/org/agents/{code}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"model", "model_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_agents_code_run_action", path: "", pathTemplate: "/org/agents/{code}/run", method: "POST", readPath: "/org/agents", desc: "Endpoint API AISIA — mutation POST exacte `/org/agents/{code}/run` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"input", "prompt", "task"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_api_keys", path: "/org/api-keys", pathTemplate: "", method: "POST", readPath: "/org/api-keys", desc: "Endpoint API AISIA — mutation POST exacte `/org/api-keys` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"allowed_models", "expires_at", "max_budget_eur", "name", "rate_limit_rpd", "rate_limit_rpm", "rate_limit_tpm", "scopes"}, typedAttrs: []apiAttr{{Name: "allowed_models", TFType: "list", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ allowed_models"}, {Name: "expires_at", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ expires_at"}, {Name: "max_budget_eur", TFType: "float", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ max_budget_eur"}, {Name: "name", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ name"}, {Name: "rate_limit_rpd", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ rate_limit_rpd"}, {Name: "rate_limit_rpm", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ rate_limit_rpm"}, {Name: "rate_limit_tpm", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ rate_limit_tpm"}, {Name: "scopes", TFType: "list", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ scopes"}}, sensitiveFields: []string{"key"}, secretOutputs: []string{"key"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_api_keys_key_id_mutation", path: "", pathTemplate: "/org/api-keys/{key_id}", method: "DELETE", readPath: "/org/api-keys", desc: "Endpoint API AISIA — mutation DELETE exacte `/org/api-keys/{key_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_custom_domain_provision_action", path: "/org/custom-domain/provision", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint API AISIA — mutation POST exacte `/org/custom-domain/provision` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_deployments_action", path: "/org/deployments", pathTemplate: "", method: "POST", readPath: "/org/deployments", desc: "Endpoint API AISIA — mutation POST exacte `/org/deployments` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"profile", "provider", "region", "service_key", "target_type"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_deployments_dep_id_apply_action", path: "", pathTemplate: "/org/deployments/{dep_id}/apply", method: "POST", readPath: "/org/deployments", desc: "Endpoint API AISIA — mutation POST exacte `/org/deployments/{dep_id}/apply` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"confirm_token"}, sensitiveFields: []string{"confirm_token"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_deployments_dep_id_destroy_action", path: "", pathTemplate: "/org/deployments/{dep_id}/destroy", method: "POST", readPath: "/org/deployments", desc: "Endpoint API AISIA — mutation POST exacte `/org/deployments/{dep_id}/destroy` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_deployments_dep_id_plan_action", path: "", pathTemplate: "/org/deployments/{dep_id}/plan", method: "POST", readPath: "/org/deployments", desc: "Endpoint API AISIA — mutation POST exacte `/org/deployments/{dep_id}/plan` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false, sensitiveFields: []string{"confirm_token"}, secretOutputs: []string{"confirm_token"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_groups", path: "/org/groups", pathTemplate: "", method: "POST", readPath: "/org/groups", desc: "Endpoint API AISIA — mutation POST exacte `/org/groups` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"description", "name"}, typedAttrs: []apiAttr{{Name: "description", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ description"}, {Name: "name", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ name"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_groups_group_id_mutation", path: "", pathTemplate: "/org/groups/{group_id}", method: "DELETE", readPath: "/org/groups", desc: "Endpoint API AISIA — mutation DELETE exacte `/org/groups/{group_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_groups_group_id_members_action", path: "", pathTemplate: "/org/groups/{group_id}/members", method: "POST", readPath: "/org/groups", desc: "Endpoint API AISIA — mutation POST exacte `/org/groups/{group_id}/members` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"user_id"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_groups_group_id_members_user_id_mutation", path: "", pathTemplate: "/org/groups/{group_id}/members/{user_id}", method: "DELETE", readPath: "/org/groups", desc: "Endpoint API AISIA — mutation DELETE exacte `/org/groups/{group_id}/members/{user_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_groups_group_id_permissions_mutation", path: "", pathTemplate: "/org/groups/{group_id}/permissions", method: "PUT", readPath: "/org/groups", desc: "Endpoint API AISIA — mutation PUT exacte `/org/groups/{group_id}/permissions` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"permissions"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_invitations", path: "/org/invitations", pathTemplate: "", method: "POST", readPath: "/org/invitations", desc: "Endpoint API AISIA — mutation POST exacte `/org/invitations` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"email", "role"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_invitations_invitation_id_mutation", path: "", pathTemplate: "/org/invitations/{invitation_id}", method: "DELETE", readPath: "/org/invitations", desc: "Endpoint API AISIA — mutation DELETE exacte `/org/invitations/{invitation_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_knowledge_action", path: "/org/knowledge", pathTemplate: "", method: "POST", readPath: "/org/knowledge", desc: "Endpoint API AISIA — mutation POST exacte `/org/knowledge` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"source", "text", "url"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_notifications_notif_id_read_action", path: "", pathTemplate: "/org/notifications/{notif_id}/read", method: "POST", readPath: "/org/notifications", desc: "Endpoint API AISIA — mutation POST exacte `/org/notifications/{notif_id}/read` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_seo_networks_network_mutation", path: "", pathTemplate: "/org/seo/networks/{network}", method: "DELETE", readPath: "/org/seo/networks", desc: "Endpoint API AISIA — mutation DELETE exacte `/org/seo/networks/{network}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_support_tickets", path: "/org/support/tickets", pathTemplate: "", method: "POST", readPath: "/org/support/tickets", desc: "Endpoint API AISIA — mutation POST exacte `/org/support/tickets` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"category", "description", "priority", "subject"}, typedAttrs: []apiAttr{{Name: "category", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ category"}, {Name: "description", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ description"}, {Name: "priority", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ priority"}, {Name: "subject", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ subject"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_support_tickets_ticket_id_messages_action", path: "", pathTemplate: "/org/support/tickets/{ticket_id}/messages", method: "POST", readPath: "/org/support/tickets", desc: "Endpoint API AISIA — mutation POST exacte `/org/support/tickets/{ticket_id}/messages` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"content", "message"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_webhooks", path: "/org/webhooks", pathTemplate: "", method: "POST", readPath: "/org/webhooks", desc: "Endpoint API AISIA — mutation POST exacte `/org/webhooks` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"active", "events", "secret", "url"}, typedAttrs: []apiAttr{{Name: "active", TFType: "bool", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ active"}, {Name: "events", TFType: "list", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ events"}, {Name: "secret", TFType: "string", Required: false, Sensitive: true, JSONEncoded: false, Desc: "champ secret"}, {Name: "url", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ url"}}, sensitiveFields: []string{"secret"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "org_webhooks_wh_id_mutation", path: "", pathTemplate: "/org/webhooks/{wh_id}", method: "DELETE", readPath: "/org/webhooks", desc: "Endpoint API AISIA — mutation DELETE exacte `/org/webhooks/{wh_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "scim_v2_groups_action", path: "/scim/v2/Groups", pathTemplate: "", method: "POST", readPath: "/scim/v2/Groups", desc: "Endpoint API AISIA — mutation POST exacte `/scim/v2/Groups` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"displayName", "members", "schemas"}, bodyRequiredFields: []string{"displayName"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "scim_v2_users_user_id_patch_mutation", path: "", pathTemplate: "/scim/v2/Users/{user_id}", method: "PATCH", readPath: "/scim/v2/Users", desc: "Endpoint API AISIA — mutation PATCH exacte `/scim/v2/Users/{user_id}` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"Operations", "schemas"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "support_tickets", path: "/support/tickets/", pathTemplate: "", method: "POST", readPath: "/support/tickets/", desc: "Endpoint API AISIA — mutation POST exacte `/support/tickets/` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"category", "description", "priority", "subject"}, typedAttrs: []apiAttr{{Name: "category", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ category"}, {Name: "description", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ description"}, {Name: "priority", TFType: "string", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ priority"}, {Name: "subject", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ subject"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "support_tickets_ticket_id_messages_action", path: "", pathTemplate: "/support/tickets/{ticket_id}/messages", method: "POST", readPath: "/support/tickets/", desc: "Endpoint API AISIA — mutation POST exacte `/support/tickets/{ticket_id}/messages` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"content"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "support_tickets_ticket_id_status_mutation", path: "", pathTemplate: "/support/tickets/{ticket_id}/status", method: "PUT", readPath: "/support/tickets/", desc: "Endpoint API AISIA — mutation PUT exacte `/support/tickets/{ticket_id}/status` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"assigned_to", "status"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_agent_action", path: "/v1/agent", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/agent` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"provider_id", "steps_max", "task"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_analytics_conversion_action", path: "/v1/analytics/conversion", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/analytics/conversion` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"analytics_consent", "error_code", "event_name", "form_name", "surface", "target"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_compare_action", path: "/v1/compare", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/compare` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_compare_llms_action", path: "/v1/compare-llms", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/compare-llms` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_comparer_llms_action", path: "/v1/comparer-llms", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/comparer-llms` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_conversations_mutation", path: "/v1/conversations", pathTemplate: "", method: "DELETE", readPath: "/v1/conversations", desc: "Endpoint public API v1 — mutation DELETE exacte `/v1/conversations` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_conversations_export_action", path: "/v1/conversations/export", pathTemplate: "", method: "POST", readPath: "/v1/conversations", desc: "Endpoint public API v1 — mutation POST exacte `/v1/conversations/export` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_conversations_conv_id_mutation", path: "", pathTemplate: "/v1/conversations/{conv_id}", method: "DELETE", readPath: "/v1/conversations", desc: "Endpoint public API v1 — mutation DELETE exacte `/v1/conversations/{conv_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_conversations_item", path: "", pathTemplate: "/v1/conversations/{conv_id}", method: "PATCH", readPath: "/v1/conversations", desc: "Endpoint public API v1 — mutation PATCH exacte `/v1/conversations/{conv_id}` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_debate", path: "/v1/debate", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/debate` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_debate_debate_id_stop_action", path: "", pathTemplate: "/v1/debate/{debate_id}/stop", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/debate/{debate_id}/stop` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_diagnostic_lead_action", path: "/v1/diagnostic/lead", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/diagnostic/lead` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"consent", "email", "org", "score", "tier", "weakest_axis"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_embeddings_action", path: "/v1/embeddings", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/embeddings` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_mcp_action", path: "/v1/mcp", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/mcp` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_media_audio_action", path: "/v1/media/audio", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/media/audio` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_media_image_action", path: "/v1/media/image", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/media/image` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_media_video_action", path: "/v1/media/video", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/media/video` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_messages_message_id_feedback_action", path: "", pathTemplate: "/v1/messages/{message_id}/feedback", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/messages/{message_id}/feedback` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"comment", "model", "provider_id", "rating"}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_rerank_action", path: "/v1/rerank", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/rerank` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: true}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_runs", path: "/v1/runs", pathTemplate: "", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/runs` (resource action Terraform).", hasJSONBody: true, bodyRequired: true, allowAdditionalBodyFields: false, bodyFields: []string{"concurrency", "iterations", "name", "prompts", "providers", "timeout_s"}, typedAttrs: []apiAttr{{Name: "concurrency", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ concurrency"}, {Name: "iterations", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ iterations"}, {Name: "name", TFType: "string", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ name"}, {Name: "prompts", TFType: "list", Required: true, Sensitive: false, JSONEncoded: false, Desc: "champ prompts"}, {Name: "providers", TFType: "list", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ providers"}, {Name: "timeout_s", TFType: "int", Required: false, Sensitive: false, JSONEncoded: false, Desc: "champ timeout_s"}}}
	},
	func() resource.Resource {
		return &actionApiResource{name: "v1_runs_run_id_stop_action", path: "", pathTemplate: "/v1/runs/{run_id}/stop", method: "POST", readPath: "", desc: "Endpoint public API v1 — mutation POST exacte `/v1/runs/{run_id}/stop` (resource action Terraform).", hasJSONBody: false, bodyRequired: false, allowAdditionalBodyFields: false}
	},
}
