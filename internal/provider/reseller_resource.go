package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &resellerResource{}
	_ resource.ResourceWithConfigure = &resellerResource{}
)

func NewResellerResource() resource.Resource { return &resellerResource{} }

type resellerResource struct {
	data *providerData
}

type resellerModel struct {
	ID               types.String  `tfsdk:"id"`
	Name             types.String  `tfsdk:"name"`
	ContactEmail     types.String  `tfsdk:"contact_email"`
	RevSharePct      types.Float64 `tfsdk:"rev_share_pct"`
	Domain           types.String  `tfsdk:"domain"`
	BrandingConfigID types.String  `tfsdk:"branding_config_id"`
	Status           types.String  `tfsdk:"status"`
	Notes            types.String  `tfsdk:"notes"`
}

// Forme JSON échangée avec l'API AISIA (/admin/resellers).
type resellerAPI struct {
	ID               string   `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	ContactEmail     string   `json:"contact_email,omitempty"`
	RevSharePct      *float64 `json:"rev_share_pct,omitempty"`
	Domain           string   `json:"domain,omitempty"`
	BrandingConfigID string   `json:"branding_config_id,omitempty"`
	Status           string   `json:"status,omitempty"`
	Notes            string   `json:"notes,omitempty"`
}

func (r *resellerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reseller"
}

func (r *resellerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Un revendeur (reseller) AISIA (/admin/resellers), avec part de revenu et branding optionnel.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifiant du revendeur (généré par AISIA).",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Nom du revendeur.",
				Required:            true,
			},
			"contact_email": schema.StringAttribute{
				MarkdownDescription: "Email de contact du revendeur.",
				Required:            true,
			},
			"rev_share_pct": schema.Float64Attribute{
				MarkdownDescription: "Part de revenu reversée au revendeur, en pourcentage (défaut 25).",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Domaine personnalisé du revendeur (optionnel).",
				Optional:            true,
				Computed:            true,
			},
			"branding_config_id": schema.StringAttribute{
				MarkdownDescription: "Identifiant de la configuration de branding associée (optionnel).",
				Optional:            true,
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Statut du revendeur (ex. active). Piloté par l'API.",
				Optional:            true,
				Computed:            true,
			},
			"notes": schema.StringAttribute{
				MarkdownDescription: "Notes libres sur le revendeur (optionnel).",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (r *resellerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d, ok := req.ProviderData.(*providerData)
	if !ok {
		resp.Diagnostics.AddError("Type de ProviderData inattendu", fmt.Sprintf("%T", req.ProviderData))
		return
	}
	r.data = d
}

func (r *resellerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resellerModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := resellerAPI{Name: plan.Name.ValueString(), ContactEmail: plan.ContactEmail.ValueString()}
	if !plan.RevSharePct.IsNull() && !plan.RevSharePct.IsUnknown() {
		v := plan.RevSharePct.ValueFloat64()
		payload.RevSharePct = &v
	}
	payload.Domain = knownStr(plan.Domain)
	payload.BrandingConfigID = knownStr(plan.BrandingConfigID)
	payload.Notes = knownStr(plan.Notes)
	var out resellerAPI
	if _, err := r.data.apiDo(ctx, "POST", "/admin/resellers", payload, &out); err != nil {
		resp.Diagnostics.AddError("Création revendeur échouée", err.Error())
		return
	}
	applyReseller(&plan, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read est un no-op : l'API AISIA n'expose pas de GET /admin/resellers/{id}
// (seulement la collection). L'état Terraform est conservé tel quel.
func (r *resellerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resellerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *resellerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state resellerModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = state.ID
	payload := resellerAPI{Name: plan.Name.ValueString(), ContactEmail: plan.ContactEmail.ValueString()}
	if !plan.RevSharePct.IsNull() && !plan.RevSharePct.IsUnknown() {
		v := plan.RevSharePct.ValueFloat64()
		payload.RevSharePct = &v
	}
	payload.Domain = knownStr(plan.Domain)
	payload.BrandingConfigID = knownStr(plan.BrandingConfigID)
	payload.Status = knownStr(plan.Status)
	payload.Notes = knownStr(plan.Notes)
	var out resellerAPI
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/resellers/"+plan.ID.ValueString(), payload, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour revendeur échouée", err.Error())
		return
	}
	applyReseller(&plan, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *resellerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resellerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/resellers/"+state.ID.ValueString(), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression revendeur échouée", err.Error())
	}
}

// applyReseller copie la réponse API dans le modèle, en garantissant une valeur
// connue pour chaque attribut Computed (évite l'erreur "unknown after apply").
func applyReseller(m *resellerModel, api resellerAPI) {
	if api.ID != "" {
		m.ID = types.StringValue(api.ID)
	}
	if api.Name != "" {
		m.Name = types.StringValue(api.Name)
	}
	if api.ContactEmail != "" {
		m.ContactEmail = types.StringValue(api.ContactEmail)
	}
	switch {
	case api.RevSharePct != nil:
		m.RevSharePct = types.Float64Value(*api.RevSharePct)
	case m.RevSharePct.IsNull() || m.RevSharePct.IsUnknown():
		m.RevSharePct = types.Float64Value(25.0)
	}
	m.Domain = coalesceStr(api.Domain, m.Domain)
	m.BrandingConfigID = coalesceStr(api.BrandingConfigID, m.BrandingConfigID)
	m.Status = coalesceStr(api.Status, m.Status)
	m.Notes = coalesceStr(api.Notes, m.Notes)
}
