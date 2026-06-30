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
	_ resource.Resource              = &organizationResource{}
	_ resource.ResourceWithConfigure = &organizationResource{}
)

func NewOrganizationResource() resource.Resource { return &organizationResource{} }

type organizationResource struct {
	data *providerData
}

type organizationModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Slug         types.String `tfsdk:"slug"`
	ContractType types.String `tfsdk:"contract_type"`
}

// Forme JSON échangée avec l'API AISIA (/admin/organizations).
type orgAPI struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Slug         string `json:"slug,omitempty"`
	ContractType string `json:"contract_type,omitempty"`
}

func (r *organizationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization"
}

func (r *organizationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Une organisation AISIA (tenant). Gérée via l'API admin AISIA.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifiant de l'organisation (généré par AISIA).",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Nom de l'organisation.",
				Required:            true,
			},
			"slug": schema.StringAttribute{
				MarkdownDescription: "Slug (sous-domaine) de l'organisation.",
				Optional:            true,
				Computed:            true,
			},
			"contract_type": schema.StringAttribute{
				MarkdownDescription: "Type de contrat : saas | baas | paas (défaut saas).",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (r *organizationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *organizationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan organizationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := orgAPI{Name: plan.Name.ValueString(), Slug: plan.Slug.ValueString(), ContractType: plan.ContractType.ValueString()}
	var out orgAPI
	if _, err := r.data.apiDo(ctx, "POST", "/admin/organizations", payload, &out); err != nil {
		resp.Diagnostics.AddError("Création organisation échouée", err.Error())
		return
	}
	applyOrg(&plan, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *organizationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state organizationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out orgAPI
	code, err := r.data.apiDo(ctx, "GET", "/admin/organizations/"+state.ID.ValueString(), nil, &out)
	if code == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture organisation échouée", err.Error())
		return
	}
	applyOrg(&state, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *organizationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan organizationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := orgAPI{Name: plan.Name.ValueString(), Slug: plan.Slug.ValueString(), ContractType: plan.ContractType.ValueString()}
	var out orgAPI
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/organizations/"+plan.ID.ValueString(), payload, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour organisation échouée", err.Error())
		return
	}
	applyOrg(&plan, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *organizationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state organizationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/organizations/"+state.ID.ValueString(), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression organisation échouée", err.Error())
	}
}

// applyOrg copie la réponse API dans le modèle d'état (en conservant les champs connus).
func applyOrg(m *organizationModel, api orgAPI) {
	if api.ID != "" {
		m.ID = types.StringValue(api.ID)
	}
	if api.Name != "" {
		m.Name = types.StringValue(api.Name)
	}
	if api.Slug != "" {
		m.Slug = types.StringValue(api.Slug)
	}
	if api.ContractType != "" {
		m.ContractType = types.StringValue(api.ContractType)
	}
}
