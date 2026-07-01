package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &providerResource{}
	_ resource.ResourceWithConfigure = &providerResource{}
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
				Required:            true,
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
	def := map[string]any{}
	if err := json.Unmarshal([]byte(plan.ConfigJSON.ValueString()), &def); err != nil {
		return fmt.Errorf("config_json invalide (JSON): %w", err)
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
}

func (r *providerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state providerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out map[string]any
	code, err := r.data.apiDo(ctx, "GET", "/admin/providers/"+state.ID.ValueString(), nil, &out)
	if code == 404 {
		resp.State.RemoveResource(ctx) // disparu côté serveur → drop de l'état
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture provider échouée", err.Error())
		return
	}
	// config_json conservé tel quel (désiré) ; on rafraîchit seulement les computed.
	state.Name = types.StringValue(asString(out, "name"))
	if s := asString(out, "status"); s != "" {
		state.Status = types.StringValue(s)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *providerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan providerModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.write(ctx, &plan, "PUT", "/admin/providers/"+plan.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Mise à jour provider échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *providerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state providerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/providers/"+state.ID.ValueString(), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression provider échouée", err.Error())
	}
}
