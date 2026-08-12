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
	_ resource.Resource                = &guardrailResource{}
	_ resource.ResourceWithConfigure   = &guardrailResource{}
	_ resource.ResourceWithImportState = &guardrailResource{}
)

func NewGuardrailResource() resource.Resource { return &guardrailResource{} }

type guardrailResource struct {
	data *providerData
}

type guardrailModel struct {
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Pattern  types.String `tfsdk:"pattern"`
	Action   types.String `tfsdk:"action"`
	Enabled  types.Bool   `tfsdk:"enabled"`
	Category types.String `tfsdk:"category"`
}

// Forme JSON échangée avec l'API AISIA (/admin/guardrails).
type guardrailAPI struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Pattern  string `json:"pattern,omitempty"`
	Action   string `json:"action,omitempty"`
	Enabled  *bool  `json:"enabled,omitempty"`
	Category string `json:"category,omitempty"`
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
				Optional:            true,
				Computed:            true,
			},
			"pattern": schema.StringAttribute{
				MarkdownDescription: "Motif (regex ou chaîne) détecté par la règle.",
				Optional:            true,
				Computed:            true,
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
			"category": schema.StringAttribute{
				MarkdownDescription: "Catégorie métier optionnelle du guardrail.",
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
	if plan.Name.IsNull() || plan.Name.IsUnknown() || strings.TrimSpace(plan.Name.ValueString()) == "" || plan.Pattern.IsNull() || plan.Pattern.IsUnknown() || strings.TrimSpace(plan.Pattern.ValueString()) == "" {
		resp.Diagnostics.AddError("Guardrail incomplet", "`name` et `pattern` sont requis lors de la création.")
		return
	}
	payload := guardrailAPI{Name: plan.Name.ValueString(), Pattern: plan.Pattern.ValueString(), Action: knownStr(plan.Action), Category: knownStr(plan.Category)}
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
	if plan.ID.IsNull() || plan.ID.IsUnknown() || plan.ID.ValueString() == "" {
		resp.Diagnostics.AddError("Création guardrail sans identifiant", "POST /admin/guardrails n'a retourné aucun id exploitable.")
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *guardrailResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state guardrailModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out any
	if _, err := r.data.apiDo(ctx, "GET", "/admin/guardrails", nil, &out); err != nil {
		resp.Diagnostics.AddError("Lecture guardrails échouée", err.Error())
		return
	}
	item, found := collectionItem(out, "guardrails", state.ID.ValueString(), "id", "gid")
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}
	var current guardrailAPI
	if err := decodeAPIObject(item, &current); err != nil {
		resp.Diagnostics.AddError("Réponse guardrail invalide", err.Error())
		return
	}
	applyGuardrail(&state, current)
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
	if plan.Name.IsNull() || plan.Name.IsUnknown() {
		plan.Name = state.Name
	}
	if plan.Pattern.IsNull() || plan.Pattern.IsUnknown() {
		plan.Pattern = state.Pattern
	}
	if plan.Action.IsNull() || plan.Action.IsUnknown() {
		plan.Action = state.Action
	}
	if plan.Category.IsNull() || plan.Category.IsUnknown() {
		plan.Category = state.Category
	}
	if plan.Enabled.IsNull() || plan.Enabled.IsUnknown() {
		plan.Enabled = state.Enabled
	}
	payload := guardrailAPI{Name: plan.Name.ValueString(), Pattern: plan.Pattern.ValueString(), Action: knownStr(plan.Action), Category: knownStr(plan.Category)}
	if !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
		v := plan.Enabled.ValueBool()
		payload.Enabled = &v
	}
	var out guardrailAPI
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/guardrails/"+apiPathSegment(plan.ID.ValueString()), payload, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour guardrail échouée", err.Error())
		return
	}
	applyGuardrail(&plan, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *guardrailResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	identifier := strings.TrimSpace(req.ID)
	if identifier == "" {
		resp.Diagnostics.AddError("Identifiant d'import vide", "Fournissez l'identifiant exact du guardrail AISIA.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(identifier))...)
}

func (r *guardrailResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state guardrailModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/guardrails/"+apiPathSegment(state.ID.ValueString()), nil, nil); err != nil {
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
	m.Category = coalesceStr(api.Category, m.Category)
	if m.Action.ValueString() == "" {
		m.Action = types.StringValue("block")
	}
	if m.Category.ValueString() == "" {
		m.Category = types.StringValue("custom")
	}
	switch {
	case api.Enabled != nil:
		m.Enabled = types.BoolValue(*api.Enabled)
	case m.Enabled.IsNull() || m.Enabled.IsUnknown():
		m.Enabled = types.BoolValue(true)
	}
}
