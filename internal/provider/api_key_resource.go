package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &apiKeyResource{}
	_ resource.ResourceWithConfigure = &apiKeyResource{}
)

func NewApiKeyResource() resource.Resource { return &apiKeyResource{} }

type apiKeyResource struct {
	data *providerData
}

// aisia_api_key — clé d'accès programmatique (aisia_sk_...) scoped à une org.
// Immuable : tout changement => recréation. La valeur brute n'est renvoyée QU'UNE
// fois à la création (stockée en sortie sensible `key`).
type apiKeyModel struct {
	ID         types.String   `tfsdk:"id"`
	OrgID      types.String   `tfsdk:"org_id"`
	Name       types.String   `tfsdk:"name"`
	Scopes     types.List     `tfsdk:"scopes"`
	RateRPM    types.Int64    `tfsdk:"rate_limit_rpm"`
	RateRPD    types.Int64    `tfsdk:"rate_limit_rpd"`
	Key        types.String   `tfsdk:"key"`
	KeyPrefix  types.String   `tfsdk:"key_prefix"`
}

func (r *apiKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_key"
}

func (r *apiKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	rr := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{
		MarkdownDescription: "Clé d'accès programmatique AISIA (`aisia_sk_…`) scoped à une organisation. Immuable. La valeur brute n'est exposée qu'à la création (`key`, sensible).",
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{MarkdownDescription: "ID de la clé.", Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"org_id":      schema.StringAttribute{MarkdownDescription: "Org propriétaire.", Required: true, PlanModifiers: rr},
			"name":        schema.StringAttribute{MarkdownDescription: "Nom de la clé.", Required: true, PlanModifiers: rr},
			"scopes":      schema.ListAttribute{MarkdownDescription: "Scopes (défaut [\"invoke\"]).", ElementType: types.StringType, Optional: true, Computed: true, PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()}},
			"rate_limit_rpm": schema.Int64Attribute{MarkdownDescription: "Limite req/min.", Optional: true, Computed: true},
			"rate_limit_rpd": schema.Int64Attribute{MarkdownDescription: "Limite req/jour.", Optional: true, Computed: true},
			"key":         schema.StringAttribute{MarkdownDescription: "Valeur brute (one-shot, sensible).", Computed: true, Sensitive: true},
			"key_prefix":  schema.StringAttribute{MarkdownDescription: "Préfixe public de la clé.", Computed: true},
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
	var scopes []string
	resp.Diagnostics.Append(plan.Scopes.ElementsAs(ctx, &scopes, false)...)
	if len(scopes) == 0 {
		scopes = []string{"invoke"}
	}
	payload := map[string]any{"name": plan.Name.ValueString(), "scopes": scopes}
	if !plan.RateRPM.IsNull() {
		payload["rate_limit_rpm"] = plan.RateRPM.ValueInt64()
	}
	if !plan.RateRPD.IsNull() {
		payload["rate_limit_rpd"] = plan.RateRPD.ValueInt64()
	}
	var out map[string]any
	if _, err := r.data.apiDo(ctx, "POST", "/admin/orgs/"+plan.OrgID.ValueString()+"/api-keys", payload, &out); err != nil {
		resp.Diagnostics.AddError("Création clé API échouée", err.Error())
		return
	}
	plan.ID = types.StringValue(asString(out, "id", "key_id"))
	plan.Key = types.StringValue(asString(out, "key", "full_key", "api_key"))
	plan.KeyPrefix = types.StringValue(asString(out, "key_prefix", "prefix"))
	if sc, ok := types.ListValueFrom(ctx, types.StringType, scopes); ok.HasError() {
		resp.Diagnostics.Append(ok...)
	} else {
		plan.Scopes = sc
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *apiKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// La valeur brute n'est pas relisible (hash en DB) → on conserve l'état.
	var state apiKeyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
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
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/orgs/"+state.OrgID.ValueString()+"/api-keys/"+state.ID.ValueString(), nil, nil); err != nil {
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
