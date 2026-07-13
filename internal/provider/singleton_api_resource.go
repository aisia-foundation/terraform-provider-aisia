package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// singletonApiResource — configuration singleton (GET + PUT/PATCH sur le même chemin, sans /{id}).
type singletonApiResource struct {
	data       *providerData
	name       string
	path       string
	updateVerb string
	desc       string
}

var (
	_ resource.Resource              = &singletonApiResource{}
	_ resource.ResourceWithConfigure = &singletonApiResource{}
)

type singletonApiResourceModel struct {
	ID   types.String `tfsdk:"id"`
	Body types.String `tfsdk:"body"`
	JSON types.String `tfsdk:"json"`
}

func (r *singletonApiResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.name
}

func (r *singletonApiResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: r.desc + docLinksForEndpoint(r.path),
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{MarkdownDescription: "Identifiant fixe (singleton).", Computed: true},
			"body": schema.StringAttribute{MarkdownDescription: "Corps JSON (payload PUT/PATCH).", Required: true},
			"json": schema.StringAttribute{MarkdownDescription: "Dernière réponse brute de l'API (JSON).", Computed: true},
		},
	}
}

func (r *singletonApiResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *singletonApiResource) singletonID() string { return "_singleton" }

func (r *singletonApiResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan singletonApiResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var body any
	if err := json.Unmarshal([]byte(plan.Body.ValueString()), &body); err != nil {
		resp.Diagnostics.AddError("body JSON invalide", err.Error())
		return
	}
	verb := r.updateVerb
	if verb == "" {
		verb = "PUT"
	}
	var out any
	if _, err := r.data.apiDo(ctx, verb, r.path, body, &out); err != nil {
		resp.Diagnostics.AddError("Création (upsert) échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	plan.ID = types.StringValue(r.singletonID())
	plan.JSON = types.StringValue(string(raw))
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *singletonApiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state singletonApiResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out any
	code, err := r.data.apiDo(ctx, "GET", r.path, nil, &out)
	if code == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	state.ID = types.StringValue(r.singletonID())
	state.JSON = types.StringValue(string(raw))
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *singletonApiResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state singletonApiResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
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
	if _, err := r.data.apiDo(ctx, verb, r.path, body, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	plan.ID = state.ID
	plan.JSON = types.StringValue(string(raw))
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *singletonApiResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
}