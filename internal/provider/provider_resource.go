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
	_ resource.Resource                = &providerResource{}
	_ resource.ResourceWithConfigure   = &providerResource{}
	_ resource.ResourceWithImportState = &providerResource{}
)

func NewProviderResource() resource.Resource { return &providerResource{} }

type providerResource struct{ data *providerData }

// aisia_provider — provider LLM géré en DB (overlay du providers.yaml, v6.9.64 Wave 2).
// `config_json` porte la définition complète (adapter_name, model, base_url, status…).
// Le `id` est immuable (RequiresReplace). Le YAML reste la source de défauts.
type providerModel struct {
	ID         types.String `tfsdk:"id"`
	ConfigJSON types.String `tfsdk:"config_json"`
	Name       types.String `tfsdk:"name"`
	Status     types.String `tfsdk:"status"`
}

func (r *providerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_provider"
}

func (r *providerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provider LLM AISIA géré en DB (overlay du `providers.yaml`). `config_json` = définition complète du provider (JSON). Le YAML reste les défauts versionnés.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifiant du provider (immuable).",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"config_json": schema.StringAttribute{
				MarkdownDescription: "Définition complète du provider en JSON (ex. `{\"adapter_name\":\"openai\",\"model\":\"gpt-4o\",\"status\":\"active\"}`).",
				Optional:            true,
				Computed:            true,
				Sensitive:           true,
			},
			"name":   schema.StringAttribute{MarkdownDescription: "Nom (dérivé de la config).", Computed: true},
			"status": schema.StringAttribute{MarkdownDescription: "Statut (active/inactive).", Computed: true},
		},
	}
}

func (r *providerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// write parse config_json, force l'id, et POST/PUT vers l'API ; met à jour name/status.
func (r *providerResource) write(ctx context.Context, plan *providerModel, method, path string) error {
	def, err := decodeConfigJSONObject(plan.ConfigJSON, "config_json")
	if err != nil {
		return err
	}
	def["id"] = plan.ID.ValueString()
	var out map[string]any
	if _, err := r.data.apiDo(ctx, method, path, def, &out); err != nil {
		return err
	}
	plan.Name = types.StringValue(asString(out, "name"))
	status := asString(out, "status")
	if status == "" {
		status = "active"
	}
	plan.Status = types.StringValue(status)
	return nil
}

func applyProviderComputed(state *providerModel, out map[string]any) {
	name := asString(out, "name", "display_name")
	if name == "" {
		name = state.ID.ValueString()
	}
	state.Name = types.StringValue(name)
	status := asString(out, "status")
	if status == "" {
		if enabled, ok := out["enabled"].(bool); ok && !enabled {
			status = "inactive"
		} else {
			status = "active"
		}
	}
	state.Status = types.StringValue(status)
}

func (r *providerResource) refresh(ctx context.Context, state *providerModel) (int, error) {
	var out map[string]any
	code, err := r.data.apiDo(ctx, "GET", "/admin/providers/"+apiPathSegment(state.ID.ValueString()), nil, &out)
	if err != nil {
		return code, err
	}
	reconciled, reconcileErr := reconcileConfigJSON(state.ConfigJSON, out, "id")
	if reconcileErr != nil {
		return code, reconcileErr
	}
	state.ConfigJSON = reconciled
	applyProviderComputed(state, out)
	return code, nil
}

func (r *providerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan providerModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.write(ctx, &plan, "POST", "/admin/providers"); err != nil {
		resp.Diagnostics.AddError("Création provider échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.refresh(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Relecture provider après création échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *providerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state providerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	code, err := r.refresh(ctx, &state)
	if code == 404 {
		resp.State.RemoveResource(ctx) // disparu côté serveur → drop de l'état
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture provider échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *providerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state providerModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = state.ID
	if plan.ConfigJSON.IsNull() || plan.ConfigJSON.IsUnknown() {
		plan.ConfigJSON = state.ConfigJSON
	}
	if err := r.write(ctx, &plan, "PUT", "/admin/providers/"+apiPathSegment(state.ID.ValueString())); err != nil {
		resp.Diagnostics.AddError("Mise à jour provider échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.refresh(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Relecture provider après mise à jour échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *providerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	identifier := strings.TrimSpace(req.ID)
	if identifier == "" {
		resp.Diagnostics.AddError("Identifiant d'import vide", "Fournissez l'identifiant exact du provider AISIA.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(identifier))...)
}

func (r *providerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state providerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/providers/"+apiPathSegment(state.ID.ValueString()), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression provider échouée", err.Error())
	}
}
