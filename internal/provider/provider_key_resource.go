package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &providerKeyResource{}
	_ resource.ResourceWithConfigure   = &providerKeyResource{}
	_ resource.ResourceWithImportState = &providerKeyResource{}
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
	Owner        types.String `tfsdk:"owner"`
	RotationDays types.Int64  `tfsdk:"rotation_days"`
	KeyMasked    types.String `tfsdk:"key_masked"`
	ExpiresAt    types.String `tfsdk:"expires_at"`
}

type providerKeyAPI struct {
	KeyValue     string `json:"key_value,omitempty"`
	AccountEmail string `json:"account_email,omitempty"`
	Owner        string `json:"owner,omitempty"`
	RotationDays int64  `json:"rotation_days,omitempty"`
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
				Optional:            true,
				Computed:            true,
				Sensitive:           true,
			},
			"account_email": schema.StringAttribute{
				MarkdownDescription: "Email du compte SaaS propriétaire de la clé (traçabilité).",
				Optional:            true,
				Computed:            true,
			},
			"owner": schema.StringAttribute{
				MarkdownDescription: "Propriétaire métier de la clé.",
				Optional:            true,
				Computed:            true,
			},
			"rotation_days": schema.Int64Attribute{
				MarkdownDescription: "Période de rotation (défaut API 90 jours).",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"key_masked": schema.StringAttribute{MarkdownDescription: "Représentation masquée renvoyée par l'API.", Computed: true},
			"expires_at": schema.StringAttribute{MarkdownDescription: "Échéance calculée de la clé.", Computed: true},
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
	return "/admin/orgs/" + apiPathSegment(m.OrgID.ValueString()) + "/provider-keys/" + apiPathSegment(m.ProviderID.ValueString())
}

func (r *providerKeyResource) upsert(ctx context.Context, m *providerKeyModel) error {
	if m.KeyValue.IsNull() || m.KeyValue.IsUnknown() || strings.TrimSpace(m.KeyValue.ValueString()) == "" {
		return fmt.Errorf("key_value est requis pour créer ou faire tourner une clé provider")
	}
	if m.RotationDays.IsNull() || m.RotationDays.IsUnknown() {
		m.RotationDays = types.Int64Value(90)
	}
	payload := providerKeyAPI{
		KeyValue: m.KeyValue.ValueString(), AccountEmail: knownStr(m.AccountEmail),
		Owner: knownStr(m.Owner), RotationDays: m.RotationDays.ValueInt64(),
	}
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
	if plan.AccountEmail.IsUnknown() {
		plan.AccountEmail = types.StringNull()
	}
	if plan.Owner.IsUnknown() {
		plan.Owner = types.StringNull()
	}
	if plan.KeyMasked.IsUnknown() {
		plan.KeyMasked = types.StringNull()
	}
	if plan.ExpiresAt.IsUnknown() {
		plan.ExpiresAt = types.StringNull()
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *providerKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// La valeur de clé est masquée côté API ; le GET collection confirme
	// néanmoins l'existence et réhydrate les métadonnées non secrètes.
	var state providerKeyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out any
	readPath := "/admin/orgs/" + apiPathSegment(state.OrgID.ValueString()) + "/provider-keys"
	if _, err := r.data.apiDo(ctx, "GET", readPath, nil, &out); err != nil {
		resp.Diagnostics.AddError("Lecture clés provider échouée", err.Error())
		return
	}
	item, found := collectionItem(out, "keys", state.ProviderID.ValueString(), "provider_id", "id")
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}
	if value := asString(item, "account_email"); value != "" {
		state.AccountEmail = types.StringValue(value)
	} else if state.AccountEmail.IsUnknown() {
		state.AccountEmail = types.StringNull()
	}
	if value := asString(item, "owner"); value != "" {
		state.Owner = types.StringValue(value)
	} else if state.Owner.IsUnknown() {
		state.Owner = types.StringNull()
	}
	if value := asString(item, "key_masked"); value != "" {
		state.KeyMasked = types.StringValue(value)
	} else if state.KeyMasked.IsUnknown() {
		state.KeyMasked = types.StringNull()
	}
	if value := asString(item, "expires_at"); value != "" {
		state.ExpiresAt = types.StringValue(value)
	} else if state.ExpiresAt.IsUnknown() {
		state.ExpiresAt = types.StringNull()
	}
	if value, ok := item["rotation_days"].(float64); ok {
		state.RotationDays = types.Int64Value(int64(value))
	} else if state.RotationDays.IsUnknown() {
		state.RotationDays = types.Int64Value(90)
	}
	if state.KeyValue.IsUnknown() {
		state.KeyValue = types.StringNull()
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *providerKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state providerKeyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = state.ID
	plan.OrgID = state.OrgID
	plan.ProviderID = state.ProviderID
	if plan.KeyValue.IsNull() || plan.KeyValue.IsUnknown() {
		plan.KeyValue = state.KeyValue
	}
	if plan.AccountEmail.IsNull() || plan.AccountEmail.IsUnknown() {
		plan.AccountEmail = state.AccountEmail
	}
	if plan.Owner.IsNull() || plan.Owner.IsUnknown() {
		plan.Owner = state.Owner
	}
	if plan.RotationDays.IsNull() || plan.RotationDays.IsUnknown() {
		plan.RotationDays = state.RotationDays
	}
	plan.KeyMasked = state.KeyMasked
	plan.ExpiresAt = state.ExpiresAt
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
	parts := strings.SplitN(strings.TrimSpace(req.ID), "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Identifiant d'import invalide", "Utilisez exactement `org_id/provider_id`.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(req.ID))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), types.StringValue(parts[0]))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("provider_id"), types.StringValue(parts[1]))...)
}
