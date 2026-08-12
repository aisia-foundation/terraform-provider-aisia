package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &resellerResource{}
	_ resource.ResourceWithConfigure   = &resellerResource{}
	_ resource.ResourceWithImportState = &resellerResource{}
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
				Optional:            true,
				Computed:            true,
			},
			"contact_email": schema.StringAttribute{
				MarkdownDescription: "Email de contact du revendeur.",
				Optional:            true,
				Computed:            true,
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
	if plan.Name.IsNull() || plan.Name.IsUnknown() || strings.TrimSpace(plan.Name.ValueString()) == "" || plan.ContactEmail.IsNull() || plan.ContactEmail.IsUnknown() || strings.TrimSpace(plan.ContactEmail.ValueString()) == "" {
		resp.Diagnostics.AddError("Revendeur incomplet", "`name` et `contact_email` sont requis lors de la création.")
		return
	}
	statusConfigured := !plan.Status.IsNull() && !plan.Status.IsUnknown()
	desiredStatus := plan.Status.ValueString()
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
	if plan.ID.IsNull() || plan.ID.IsUnknown() || plan.ID.ValueString() == "" {
		resp.Diagnostics.AddError("Création revendeur sans identifiant", "POST /admin/resellers n'a retourné aucun id exploitable.")
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if statusConfigured {
		var updateOut resellerAPI
		if _, err := r.data.apiDo(ctx, "PUT", "/admin/resellers/"+apiPathSegment(plan.ID.ValueString()), map[string]any{"status": desiredStatus}, &updateOut); err != nil {
			resp.Diagnostics.AddError("Statut revendeur après création échoué", err.Error())
			return
		}
		applyReseller(&plan, updateOut)
		if updateOut.Status == "" {
			plan.Status = types.StringValue(desiredStatus)
		}
		resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	}
}

func (r *resellerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resellerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out any
	if _, err := r.data.apiDo(ctx, "GET", "/admin/resellers", nil, &out); err != nil {
		resp.Diagnostics.AddError("Lecture revendeurs échouée", err.Error())
		return
	}
	item, found := collectionItem(out, "resellers", state.ID.ValueString(), "id", "reseller_id")
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}
	var current resellerAPI
	if err := decodeAPIObject(item, &current); err != nil {
		resp.Diagnostics.AddError("Réponse revendeur invalide", err.Error())
		return
	}
	applyReseller(&state, current)
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
	if plan.Name.IsNull() || plan.Name.IsUnknown() {
		plan.Name = state.Name
	}
	if plan.ContactEmail.IsNull() || plan.ContactEmail.IsUnknown() {
		plan.ContactEmail = state.ContactEmail
	}
	if plan.RevSharePct.IsNull() || plan.RevSharePct.IsUnknown() {
		plan.RevSharePct = state.RevSharePct
	}
	if plan.Domain.IsNull() || plan.Domain.IsUnknown() {
		plan.Domain = state.Domain
	}
	if plan.BrandingConfigID.IsNull() || plan.BrandingConfigID.IsUnknown() {
		plan.BrandingConfigID = state.BrandingConfigID
	}
	if plan.Status.IsNull() || plan.Status.IsUnknown() {
		plan.Status = state.Status
	}
	if plan.Notes.IsNull() || plan.Notes.IsUnknown() {
		plan.Notes = state.Notes
	}
	payload := map[string]any{
		"name": plan.Name.ValueString(), "contact_email": plan.ContactEmail.ValueString(),
		"rev_share_pct": plan.RevSharePct.ValueFloat64(), "domain": plan.Domain.ValueString(),
		"branding_config_id": plan.BrandingConfigID.ValueString(), "status": plan.Status.ValueString(),
		"notes": plan.Notes.ValueString(),
	}
	var out resellerAPI
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/resellers/"+apiPathSegment(plan.ID.ValueString()), payload, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour revendeur échouée", err.Error())
		return
	}
	applyReseller(&plan, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *resellerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	identifier := strings.TrimSpace(req.ID)
	if identifier == "" {
		resp.Diagnostics.AddError("Identifiant d'import vide", "Fournissez l'identifiant exact du revendeur AISIA.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(identifier))...)
}

func (r *resellerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resellerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/resellers/"+apiPathSegment(state.ID.ValueString()), nil, nil); err != nil {
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
	if m.Status.ValueString() == "" {
		m.Status = types.StringValue("active")
	}
}
