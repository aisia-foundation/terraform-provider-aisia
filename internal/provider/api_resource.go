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

// apiResource — resource CRUD GÉNÉRIQUE pour toute entité AISIA administrable par API
// (POST collection, GET/{id}, PUT|PATCH/{id}, DELETE/{id}). Robuste au schéma : l'entité
// est décrite par son corps JSON (`body`, ex. jsonencode({...})) ; le provider gère le
// cycle create/read/update/delete. Registrations générées dans resources_generated.go.
type apiResource struct {
	data       *providerData
	name       string // suffixe TypeName, ex. "admin_webhooks"
	path       string // collection, ex. "/admin/webhooks"
	updateVerb string // "PUT" | "PATCH" | "" (aucun update → ForceNew)
	canDelete  bool
	desc       string
}

var (
	_ resource.Resource              = &apiResource{}
	_ resource.ResourceWithConfigure = &apiResource{}
)

type apiResourceModel struct {
	ID   types.String `tfsdk:"id"`
	Body types.String `tfsdk:"body"`
	JSON types.String `tfsdk:"json"`
}

func (r *apiResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.name
}

func (r *apiResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	bodyMods := []planmodifier.String{}
	if r.updateVerb == "" { // pas d'endpoint update → toute modif = recréation
		bodyMods = append(bodyMods, stringplanmodifier.RequiresReplace())
	}
	resp.Schema = schema.Schema{
		MarkdownDescription: r.desc,
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{MarkdownDescription: "Identifiant de l'entité (retourné par l'API).", Computed: true},
			"body": schema.StringAttribute{MarkdownDescription: "Corps JSON de l'entité (payload create/update), ex. `jsonencode({...})`.", Required: true, PlanModifiers: bodyMods},
			"json": schema.StringAttribute{MarkdownDescription: "Dernière réponse brute de l'API (JSON).", Computed: true},
		},
	}
}

func (r *apiResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	pd, ok := req.ProviderData.(*providerData)
	if !ok {
		resp.Diagnostics.AddError("ProviderData inattendu", fmt.Sprintf("%T", req.ProviderData))
		return
	}
	r.data = pd
}

func idFromResponse(out any) string {
	if m, ok := out.(map[string]any); ok {
		for _, k := range []string{"id", "uuid", "code", "slug", "name"} {
			if v, ok := m[k]; ok {
				return fmt.Sprintf("%v", v)
			}
		}
	}
	return ""
}

func (r *apiResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan apiResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var body any
	if err := json.Unmarshal([]byte(plan.Body.ValueString()), &body); err != nil {
		resp.Diagnostics.AddError("body JSON invalide", err.Error())
		return
	}
	var out any
	if _, err := r.data.apiDo(ctx, "POST", r.path, body, &out); err != nil {
		resp.Diagnostics.AddError("Création échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	plan.ID = types.StringValue(idFromResponse(out))
	plan.JSON = types.StringValue(string(raw))
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *apiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state apiResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || state.ID.ValueString() == "" {
		return
	}
	var out any
	code, err := r.data.apiDo(ctx, "GET", r.path+"/"+state.ID.ValueString(), nil, &out)
	if code == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	state.JSON = types.StringValue(string(raw))
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *apiResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state apiResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = state.ID
	verb := r.updateVerb
	if verb == "" {
		verb = "PUT"
	}
	var body any
	if err := json.Unmarshal([]byte(plan.Body.ValueString()), &body); err != nil {
		resp.Diagnostics.AddError("body JSON invalide", err.Error())
		return
	}
	var out any
	if _, err := r.data.apiDo(ctx, verb, r.path+"/"+plan.ID.ValueString(), body, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	plan.JSON = types.StringValue(string(raw))
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *apiResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state apiResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || state.ID.ValueString() == "" || !r.canDelete {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", r.path+"/"+state.ID.ValueString(), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression échouée", err.Error())
	}
}
