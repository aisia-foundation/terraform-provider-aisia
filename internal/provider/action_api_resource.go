package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// actionApiResource — mutation admin (POST/PUT/PATCH/DELETE) ou action idempotente.
// path = chemin littéral ; pathTemplate = chemin avec {placeholders} résolus depuis body.
type actionApiResource struct {
	data         *providerData
	name         string
	path         string
	pathTemplate string
	method       string
	readPath     string
	desc         string
}

var (
	_ resource.Resource              = &actionApiResource{}
	_ resource.ResourceWithConfigure = &actionApiResource{}
	_tplVar                        = regexp.MustCompile(`\{([a-zA-Z0-9_]+)\}`)
)

type actionApiResourceModel struct {
	ID   types.String `tfsdk:"id"`
	Body types.String `tfsdk:"body"`
	JSON types.String `tfsdk:"json"`
}

func (r *actionApiResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.name
}

func (r *actionApiResource) actionDocPath() string {
	if r.pathTemplate != "" {
		return r.pathTemplate
	}
	return r.path
}

func (r *actionApiResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: r.desc + docLinksForEndpoint(r.actionDocPath()),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifiant fixe de la mutation.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"body": schema.StringAttribute{
				MarkdownDescription: "Corps JSON (ex. `jsonencode({...})`). Clés requises si pathTemplate contient {placeholders}.",
				Required:            true,
			},
			"json": schema.StringAttribute{
				MarkdownDescription: "Dernière réponse brute de l'API (JSON).",
				Computed:            true,
			},
		},
	}
}

func (r *actionApiResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *actionApiResource) httpMethod() string {
	if r.method != "" {
		return strings.ToUpper(r.method)
	}
	return "POST"
}

func (r *actionApiResource) actionID() string { return "_action" }

func (r *actionApiResource) parseBody(plan actionApiResourceModel) (map[string]any, any, error) {
	raw := plan.Body.ValueString()
	if raw == "" || raw == "{}" {
		return map[string]any{}, map[string]any{}, nil
	}
	var body any
	if err := json.Unmarshal([]byte(raw), &body); err != nil {
		return nil, nil, err
	}
	m, ok := body.(map[string]any)
	if !ok {
		return map[string]any{}, body, nil
	}
	return m, body, nil
}

func (r *actionApiResource) resolvePath(fields map[string]any) (string, error) {
	tmpl := r.pathTemplate
	if tmpl == "" {
		tmpl = r.path
	}
	if tmpl == "" {
		return "", fmt.Errorf("chemin mutation non configuré")
	}
	var missing []string
	out := _tplVar.ReplaceAllStringFunc(tmpl, func(seg string) string {
		key := strings.Trim(seg, "{}")
		if v, ok := fields[key]; ok {
			return fmt.Sprintf("%v", v)
		}
		missing = append(missing, key)
		return seg
	})
	if len(missing) > 0 {
		return "", fmt.Errorf("body JSON : clés manquantes pour pathTemplate %s : %s", tmpl, strings.Join(missing, ", "))
	}
	return out, nil
}

func (r *actionApiResource) runMutation(ctx context.Context, plan actionApiResourceModel) (actionApiResourceModel, error) {
	fields, body, err := r.parseBody(plan)
	if err != nil {
		return plan, err
	}
	path, err := r.resolvePath(fields)
	if err != nil {
		return plan, err
	}
	var out any
	if _, err := r.data.apiDo(ctx, r.httpMethod(), path, body, &out); err != nil {
		return plan, err
	}
	raw, _ := json.Marshal(out)
	plan.ID = types.StringValue(r.actionID())
	plan.JSON = types.StringValue(string(raw))
	return plan, nil
}

func (r *actionApiResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan actionApiResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	out, err := r.runMutation(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("Mutation API échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, out)...)
}

func (r *actionApiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state actionApiResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.readPath == "" {
		resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
		return
	}
	var out any
	code, err := r.data.apiDo(ctx, "GET", r.readPath, nil, &out)
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

func (r *actionApiResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state actionApiResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	out, err := r.runMutation(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("Mutation API échouée", err.Error())
		return
	}
	out.ID = state.ID
	resp.Diagnostics.Append(resp.State.Set(ctx, out)...)
}

func (r *actionApiResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
}