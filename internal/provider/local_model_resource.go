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
	_ resource.Resource              = &localModelResource{}
	_ resource.ResourceWithConfigure = &localModelResource{}
)

func NewLocalModelResource() resource.Resource { return &localModelResource{} }

type localModelResource struct{ data *providerData }

// aisia_local_model — modèle local géré en DB (overlay du local_models.yaml, v6.9.64 Wave 2).
// `config_json` porte la définition complète (runtime, endpoint, model_name, priority…).
type localModelModel struct {
	ModelID     types.String `tfsdk:"model_id"`
	ConfigJSON  types.String `tfsdk:"config_json"`
	DisplayName types.String `tfsdk:"display_name"`
	Status      types.String `tfsdk:"status"`
}

func (r *localModelResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_local_model"
}

func (r *localModelResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Modèle local AISIA géré en DB (overlay du `local_models.yaml`). `config_json` = définition complète (runtime, endpoint, model_name, priority…).",
		Attributes: map[string]schema.Attribute{
			"model_id": schema.StringAttribute{
				MarkdownDescription: "Identifiant du modèle local (immuable).",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"config_json": schema.StringAttribute{
				MarkdownDescription: "Définition complète du modèle en JSON (ex. `{\"runtime\":\"ollama\",\"model_name\":\"llama3.3:70b\",\"priority\":120,\"status\":\"active\"}`).",
				Required:            true,
			},
			"display_name": schema.StringAttribute{MarkdownDescription: "Nom affiché (dérivé).", Computed: true},
			"status":       schema.StringAttribute{MarkdownDescription: "Statut (active/inactive).", Computed: true},
		},
	}
}

func (r *localModelResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// write parse config_json, force le model_id et POST (upsert) vers l'API.
func (r *localModelResource) write(ctx context.Context, plan *localModelModel) error {
	def := map[string]any{}
	if err := json.Unmarshal([]byte(plan.ConfigJSON.ValueString()), &def); err != nil {
		return fmt.Errorf("config_json invalide (JSON): %w", err)
	}
	def["model_id"] = plan.ModelID.ValueString()
	var out map[string]any
	if _, err := r.data.apiDo(ctx, "POST", "/admin/local-models", def, &out); err != nil {
		return err
	}
	plan.DisplayName = types.StringValue(asString(out, "display_name", "name"))
	status := asString(out, "status")
	if status == "" {
		status = "active"
	}
	plan.Status = types.StringValue(status)
	return nil
}

func (r *localModelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan localModelModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.write(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Création modèle local échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *localModelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state localModelModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out map[string]any
	code, err := r.data.apiDo(ctx, "GET", "/admin/local-models/"+state.ModelID.ValueString(), nil, &out)
	if code == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture modèle local échouée", err.Error())
		return
	}
	state.DisplayName = types.StringValue(asString(out, "display_name", "name"))
	if s := asString(out, "status"); s != "" {
		state.Status = types.StringValue(s)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *localModelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan localModelModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.write(ctx, &plan); err != nil { // POST = upsert
		resp.Diagnostics.AddError("Mise à jour modèle local échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *localModelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state localModelModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/local-models/"+state.ModelID.ValueString(), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression modèle local échouée", err.Error())
	}
}
