package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &providerKeyResource{}
	_ resource.ResourceWithConfigure = &providerKeyResource{}
)

func NewProviderKeyResource() resource.Resource { return &providerKeyResource{} }

type providerKeyResource struct {
	data *providerData
}

// aisia_provider_key — clé API provider PAR organisation (KEY-2 / isolation tenant).
type providerKeyModel struct {
	ID           types.String `tfsdk:"id"`
	OrgID        types.String `tfsdk:"org_id"`
	ProviderID   types.String `tfsdk:"provider_id"`
	KeyValue     types.String `tfsdk:"key_value"`
	AccountEmail types.String `tfsdk:"account_email"`
}

type providerKeyAPI struct {
	KeyValue     string `json:"key_value,omitempty"`
	AccountEmail string `json:"account_email,omitempty"`
}

func (r *providerKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_provider_key"
}

func (r *providerKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Clé API d'un provider LLM **spécifique à une organisation** (isolation multi-tenant KEY-2). Validée par AISIA via un appel réel au provider avant stockage.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifiant composite `org_id/provider_id`.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				MarkdownDescription: "ID de l'organisation propriétaire de la clé.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"provider_id": schema.StringAttribute{
				MarkdownDescription: "ID du provider (ex. `openai`, `cohere`, `groq`).",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"key_value": schema.StringAttribute{
				MarkdownDescription: "Valeur de la clé API (sensible, write-only — non relue depuis l'API qui la masque).",
				Required:            true,
				Sensitive:           true,
			},
			"account_email": schema.StringAttribute{
				MarkdownDescription: "Email du compte SaaS propriétaire de la clé (traçabilité).",
				Optional:            true,
			},
		},
	}
}

func (r *providerKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *providerKeyResource) keyPath(m providerKeyModel) string {
	return "/admin/orgs/" + m.OrgID.ValueString() + "/provider-keys/" + m.ProviderID.ValueString()
}

func (r *providerKeyResource) upsert(ctx context.Context, m *providerKeyModel) error {
	payload := providerKeyAPI{KeyValue: m.KeyValue.ValueString(), AccountEmail: m.AccountEmail.ValueString()}
	_, err := r.data.apiDo(ctx, "PUT", r.keyPath(*m), payload, nil)
	if err != nil {
		return err
	}
	m.ID = types.StringValue(m.OrgID.ValueString() + "/" + m.ProviderID.ValueString())
	return nil
}

func (r *providerKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan providerKeyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.upsert(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Création clé provider échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *providerKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// La valeur de clé est masquée côté API → on conserve l'état (write-only secret).
	// On pourrait confirmer l'existence via GET liste, omis ici (clé sensible).
	var state providerKeyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *providerKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan providerKeyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.upsert(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Mise à jour clé provider échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *providerKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state providerKeyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", r.keyPath(state), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression clé provider échouée", err.Error())
	}
}

// ImportState : "org_id/provider_id".
func (r *providerKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
