package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &apiKeyResource{}
	_ resource.ResourceWithConfigure   = &apiKeyResource{}
	_ resource.ResourceWithImportState = &apiKeyResource{}
)

func NewApiKeyResource() resource.Resource { return &apiKeyResource{} }

type apiKeyResource struct {
	data *providerData
}

// aisia_api_key — clé d'accès programmatique (aisia_sk_...) scoped à une org.
// Immuable : tout changement => recréation. La valeur brute n'est renvoyée QU'UNE
// fois à la création (stockée en sortie sensible `key`).
type apiKeyModel struct {
	ID            types.String  `tfsdk:"id"`
	OrgID         types.String  `tfsdk:"org_id"`
	Name          types.String  `tfsdk:"name"`
	Scopes        types.List    `tfsdk:"scopes"`
	RateRPM       types.Int64   `tfsdk:"rate_limit_rpm"`
	RateRPD       types.Int64   `tfsdk:"rate_limit_rpd"`
	RateTPM       types.Int64   `tfsdk:"rate_limit_tpm"`
	MaxBudget     types.Float64 `tfsdk:"max_budget_eur"`
	AllowedModels types.List    `tfsdk:"allowed_models"`
	ExpiresAt     types.String  `tfsdk:"expires_at"`
	Key           types.String  `tfsdk:"key"`
	KeyPrefix     types.String  `tfsdk:"key_prefix"`
}

func (r *apiKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_key"
}

func (r *apiKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	rr := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{
		MarkdownDescription: "Clé d'accès programmatique AISIA (`aisia_sk_…`) scoped à une organisation. Immuable. La valeur brute n'est exposée qu'à la création (`key`, sensible).",
		Attributes: map[string]schema.Attribute{
			"id":             schema.StringAttribute{MarkdownDescription: "ID de la clé.", Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"org_id":         schema.StringAttribute{MarkdownDescription: "Org propriétaire.", Required: true, PlanModifiers: rr},
			"name":           schema.StringAttribute{MarkdownDescription: "Nom de la clé.", Optional: true, Computed: true, PlanModifiers: rr},
			"scopes":         schema.ListAttribute{MarkdownDescription: "Scopes (défaut [\"invoke\"]).", ElementType: types.StringType, Optional: true, Computed: true, PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()}},
			"rate_limit_rpm": schema.Int64Attribute{MarkdownDescription: "Limite req/min.", Optional: true, Computed: true, PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()}},
			"rate_limit_rpd": schema.Int64Attribute{MarkdownDescription: "Limite req/jour.", Optional: true, Computed: true, PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()}},
			"rate_limit_tpm": schema.Int64Attribute{MarkdownDescription: "Limite tokens/minute.", Optional: true, Computed: true, PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()}},
			"max_budget_eur": schema.Float64Attribute{MarkdownDescription: "Budget maximal en euros.", Optional: true, Computed: true, PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()}},
			"allowed_models": schema.ListAttribute{MarkdownDescription: "Modèles autorisés pour cette clé.", ElementType: types.StringType, Optional: true, Computed: true, PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()}},
			"expires_at":     schema.StringAttribute{MarkdownDescription: "Date d'expiration optionnelle.", Optional: true, Computed: true, PlanModifiers: rr},
			"key":            schema.StringAttribute{MarkdownDescription: "Valeur brute (one-shot, sensible).", Computed: true, Sensitive: true},
			"key_prefix":     schema.StringAttribute{MarkdownDescription: "Préfixe public de la clé.", Computed: true},
		},
	}
}

func (r *apiKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d, ok := req.ProviderData.(*providerData)
	if !ok {
		resp.Diagnostics.AddError("ProviderData inattendu", fmt.Sprintf("%T", req.ProviderData))
		return
	}
	r.data = d
}

func (r *apiKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan apiKeyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Name.IsNull() || plan.Name.IsUnknown() || strings.TrimSpace(plan.Name.ValueString()) == "" {
		resp.Diagnostics.AddError("Nom de clé API manquant", "`name` est requis lors de la création d'une clé API AISIA.")
		return
	}
	var scopes []string
	if !plan.Scopes.IsNull() && !plan.Scopes.IsUnknown() {
		resp.Diagnostics.Append(plan.Scopes.ElementsAs(ctx, &scopes, false)...)
	}
	if len(scopes) == 0 {
		scopes = []string{"invoke"}
	}
	payload := map[string]any{"name": plan.Name.ValueString(), "scopes": scopes}
	if !plan.RateRPM.IsNull() && !plan.RateRPM.IsUnknown() {
		payload["rate_limit_rpm"] = plan.RateRPM.ValueInt64()
	}
	if !plan.RateRPD.IsNull() && !plan.RateRPD.IsUnknown() {
		payload["rate_limit_rpd"] = plan.RateRPD.ValueInt64()
	}
	if !plan.RateTPM.IsNull() && !plan.RateTPM.IsUnknown() {
		payload["rate_limit_tpm"] = plan.RateTPM.ValueInt64()
	}
	if !plan.MaxBudget.IsNull() && !plan.MaxBudget.IsUnknown() {
		payload["max_budget_eur"] = plan.MaxBudget.ValueFloat64()
	}
	if !plan.AllowedModels.IsNull() && !plan.AllowedModels.IsUnknown() {
		var models []string
		resp.Diagnostics.Append(plan.AllowedModels.ElementsAs(ctx, &models, false)...)
		payload["allowed_models"] = models
	}
	if !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		payload["expires_at"] = plan.ExpiresAt.ValueString()
	}
	var out map[string]any
	if _, err := r.data.apiDo(ctx, "POST", "/admin/orgs/"+apiPathSegment(plan.OrgID.ValueString())+"/api-keys", payload, &out); err != nil {
		resp.Diagnostics.AddError("Création clé API échouée", err.Error())
		return
	}
	identifier := asString(out, "id", "key_id")
	key := asString(out, "key", "full_key", "api_key")
	if identifier != "" {
		plan.ID = types.StringValue(identifier)
	}
	if key != "" {
		plan.Key = types.StringValue(key)
	} else {
		plan.Key = types.StringNull()
	}
	plan.KeyPrefix = types.StringValue(asString(out, "key_prefix", "prefix"))
	if value, ok := out["rate_limit_rpm"].(float64); ok {
		plan.RateRPM = types.Int64Value(int64(value))
	} else if plan.RateRPM.IsUnknown() {
		plan.RateRPM = types.Int64Null()
	}
	if value, ok := out["rate_limit_rpd"].(float64); ok {
		plan.RateRPD = types.Int64Value(int64(value))
	} else if plan.RateRPD.IsUnknown() {
		plan.RateRPD = types.Int64Null()
	}
	if value, ok := out["rate_limit_tpm"].(float64); ok {
		plan.RateTPM = types.Int64Value(int64(value))
	} else if plan.RateTPM.IsUnknown() {
		plan.RateTPM = types.Int64Null()
	}
	if value, ok := out["max_budget_eur"].(float64); ok {
		plan.MaxBudget = types.Float64Value(value)
	} else if plan.MaxBudget.IsUnknown() {
		plan.MaxBudget = types.Float64Null()
	}
	if values, ok := out["allowed_models"].([]any); ok {
		models := make([]string, 0, len(values))
		for _, value := range values {
			models = append(models, fmt.Sprintf("%v", value))
		}
		list, diagnostics := types.ListValueFrom(ctx, types.StringType, models)
		resp.Diagnostics.Append(diagnostics...)
		if !diagnostics.HasError() {
			plan.AllowedModels = list
		}
	} else if plan.AllowedModels.IsUnknown() {
		plan.AllowedModels = types.ListNull(types.StringType)
	}
	if value := asString(out, "expires_at"); value != "" {
		plan.ExpiresAt = types.StringValue(value)
	} else if plan.ExpiresAt.IsUnknown() {
		plan.ExpiresAt = types.StringNull()
	}
	if sc, ok := types.ListValueFrom(ctx, types.StringType, scopes); ok.HasError() {
		resp.Diagnostics.Append(ok...)
	} else {
		plan.Scopes = sc
	}
	// Dès qu'un identifiant est disponible, conserver le state avant de valider
	// la sortie one-shot : un POST 2xx ne doit jamais être rejoué aveuglément.
	if identifier != "" {
		resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	}
	if identifier == "" || key == "" {
		detail := "La réponse 2xx ne contient pas la clé one-shot requise."
		if identifier == "" {
			detail = "La réponse 2xx ne contient aucun identifiant exploitable ; la création distante doit être auditée avant toute relance."
		} else {
			detail += " L'identifiant a été conservé dans un state récupérable afin d'empêcher un second POST."
		}
		resp.Diagnostics.AddError("Création clé API incomplète", detail)
	}
}

func (r *apiKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// La valeur brute n'est pas relisible (hash en DB), mais le GET collection
	// confirme l'existence et réhydrate tous les champs non secrets.
	var state apiKeyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out any
	readPath := "/admin/orgs/" + apiPathSegment(state.OrgID.ValueString()) + "/api-keys"
	if _, err := r.data.apiDo(ctx, "GET", readPath, nil, &out); err != nil {
		resp.Diagnostics.AddError("Lecture clés API échouée", err.Error())
		return
	}
	item, found := collectionItem(out, "keys", state.ID.ValueString(), "id", "key_id")
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}
	if value := asString(item, "name"); value != "" {
		state.Name = types.StringValue(value)
	}
	if value := asString(item, "key_prefix", "prefix"); value != "" {
		state.KeyPrefix = types.StringValue(value)
	}
	if values, ok := item["scopes"].([]any); ok {
		scopes := make([]string, 0, len(values))
		for _, value := range values {
			scopes = append(scopes, fmt.Sprintf("%v", value))
		}
		list, diagnostics := types.ListValueFrom(ctx, types.StringType, scopes)
		resp.Diagnostics.Append(diagnostics...)
		if !diagnostics.HasError() {
			state.Scopes = list
		}
	}
	if value, ok := item["rate_limit_rpm"].(float64); ok {
		state.RateRPM = types.Int64Value(int64(value))
	}
	if value, ok := item["rate_limit_rpd"].(float64); ok {
		state.RateRPD = types.Int64Value(int64(value))
	}
	if value, ok := item["rate_limit_tpm"].(float64); ok {
		state.RateTPM = types.Int64Value(int64(value))
	} else if state.RateTPM.IsUnknown() {
		state.RateTPM = types.Int64Null()
	}
	if value, ok := item["max_budget_eur"].(float64); ok {
		state.MaxBudget = types.Float64Value(value)
	} else if state.MaxBudget.IsUnknown() {
		state.MaxBudget = types.Float64Null()
	}
	if values, ok := item["allowed_models"].([]any); ok {
		models := make([]string, 0, len(values))
		for _, value := range values {
			models = append(models, fmt.Sprintf("%v", value))
		}
		list, diagnostics := types.ListValueFrom(ctx, types.StringType, models)
		resp.Diagnostics.Append(diagnostics...)
		if !diagnostics.HasError() {
			state.AllowedModels = list
		}
	} else if state.AllowedModels.IsUnknown() {
		state.AllowedModels = types.ListNull(types.StringType)
	}
	if value := asString(item, "expires_at"); value != "" {
		state.ExpiresAt = types.StringValue(value)
	} else if state.ExpiresAt.IsUnknown() {
		state.ExpiresAt = types.StringNull()
	}
	if state.Key.IsUnknown() {
		state.Key = types.StringNull()
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

// ImportState : `org_id/key_id`; la valeur brute one-shot reste Null.
func (r *apiKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.SplitN(strings.TrimSpace(req.ID), "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Identifiant d'import invalide", "Utilisez exactement `org_id/key_id`.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(parts[1]))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), types.StringValue(parts[0]))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("key"), types.StringNull())...)
}

func (r *apiKeyResource) Update(ctx context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Tous les attributs mutables sont RequiresReplace → Update ne devrait pas être appelé.
	resp.Diagnostics.AddError("Mise à jour non supportée", "Les clés API AISIA sont immuables ; modifiez via recréation.")
}

func (r *apiKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state apiKeyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/orgs/"+apiPathSegment(state.OrgID.ValueString())+"/api-keys/"+apiPathSegment(state.ID.ValueString()), nil, nil); err != nil {
		resp.Diagnostics.AddError("Révocation clé API échouée", err.Error())
	}
}

// asString récupère la 1re clé présente (string) parmi des alias.
func asString(m map[string]any, keys ...string) string {
	for _, k := range keys {
		if v, ok := m[k]; ok {
			if s, ok2 := v.(string); ok2 {
				return s
			}
		}
	}
	return ""
}
