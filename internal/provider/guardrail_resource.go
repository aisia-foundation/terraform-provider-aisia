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
	_ resource.Resource              = &guardrailResource{}
	_ resource.ResourceWithConfigure = &guardrailResource{}
)

func NewGuardrailResource() resource.Resource { return &guardrailResource{} }

type guardrailResource struct {
	data *providerData
}

type guardrailModel struct {
	ID      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Pattern types.String `tfsdk:"pattern"`
	Action  types.String `tfsdk:"action"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

// Forme JSON échangée avec l'API AISIA (/admin/guardrails).
type guardrailAPI struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	Action  string `json:"action,omitempty"`
	Enabled *bool  `json:"enabled,omitempty"`
}

func (r *guardrailResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_guardrail"
}

func (r *guardrailResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Une règle de garde-fou (guardrail) AISIA (/admin/guardrails) : motif filtré et action associée.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifiant du guardrail (généré par AISIA).",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Nom du guardrail.",
				Required:            true,
			},
			"pattern": schema.StringAttribute{
				MarkdownDescription: "Motif (regex ou chaîne) détecté par la règle.",
				Required:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Action appliquée en cas de correspondance : block | flag | redact (défaut block).",
				Optional:            true,
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Règle active (défaut true).",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (r *guardrailResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *guardrailResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan guardrailModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := guardrailAPI{Name: plan.Name.ValueString(), Pattern: plan.Pattern.ValueString(), Action: knownStr(plan.Action)}
	if !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
		v := plan.Enabled.ValueBool()
		payload.Enabled = &v
	}
	var out guardrailAPI
	if _, err := r.data.apiDo(ctx, "POST", "/admin/guardrails", payload, &out); err != nil {
		resp.Diagnostics.AddError("Création guardrail échouée", err.Error())
		return
	}
	applyGuardrail(&plan, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read est un no-op : l'API AISIA n'expose pas de GET /admin/guardrails/{id}
// (seulement la collection). L'état Terraform est conservé tel quel.
func (r *guardrailResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state guardrailModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *guardrailResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state guardrailModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = state.ID
	payload := guardrailAPI{Name: plan.Name.ValueString(), Pattern: plan.Pattern.ValueString(), Action: knownStr(plan.Action)}
	if !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
		v := plan.Enabled.ValueBool()
		payload.Enabled = &v
	}
	var out guardrailAPI
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/guardrails/"+plan.ID.ValueString(), payload, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour guardrail échouée", err.Error())
		return
	}
	applyGuardrail(&plan, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *guardrailResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state guardrailModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/guardrails/"+state.ID.ValueString(), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression guardrail échouée", err.Error())
	}
}

// applyGuardrail copie la réponse API dans le modèle, en garantissant une valeur
// connue pour chaque attribut Computed (évite l'erreur "unknown after apply").
func applyGuardrail(m *guardrailModel, api guardrailAPI) {
	if api.ID != "" {
		m.ID = types.StringValue(api.ID)
	}
	if api.Name != "" {
		m.Name = types.StringValue(api.Name)
	}
	if api.Pattern != "" {
		m.Pattern = types.StringValue(api.Pattern)
	}
	m.Action = coalesceStr(api.Action, m.Action)
	if m.Action.ValueString() == "" {
		m.Action = types.StringValue("block")
	}
	switch {
	case api.Enabled != nil:
		m.Enabled = types.BoolValue(*api.Enabled)
	case m.Enabled.IsNull() || m.Enabled.IsUnknown():
		m.Enabled = types.BoolValue(true)
	}
}
