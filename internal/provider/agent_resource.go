package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &agentResource{}
	_ resource.ResourceWithConfigure = &agentResource{}
)

func NewAgentResource() resource.Resource { return &agentResource{} }

type agentResource struct{ data *providerData }

// aisia_agent — état enabled d'un agent orchestré (GOV-01 Sprint S1).
// Lecture : GET /admin/agents/{code} · Écriture : POST /admin/agents/{code}/toggle.
type agentModel struct {
	Code    types.String `tfsdk:"code"`
	Enabled types.Bool   `tfsdk:"enabled"`
	JSON    types.String `tfsdk:"json"`
}

func (r *agentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_agent"
}

func (r *agentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Agent AISIA orchestré — gère l'état `enabled` via POST /admin/agents/{code}/toggle. Lecture via GET /admin/agents/{code}.",
		Attributes: map[string]schema.Attribute{
			"code": schema.StringAttribute{
				MarkdownDescription: "Code agent (ex. `agent_devops`, `devops`).",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Agent activé (persisté Redis + orchestrateur).",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"json": schema.StringAttribute{
				MarkdownDescription: "Dernière réponse GET /admin/agents/{code} (JSON).",
				Computed:            true,
			},
		},
	}
}

func (r *agentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func normalizeAgentCode(code string) string {
	code = strings.TrimSpace(code)
	if code == "" {
		return code
	}
	if strings.HasPrefix(code, "agent_") {
		return code
	}
	return "agent_" + code
}

func (r *agentResource) readAgent(ctx context.Context, code string) (map[string]any, int, error) {
	var out map[string]any
	path := "/admin/agents/" + code
	status, err := r.data.apiDo(ctx, "GET", path, nil, &out)
	return out, status, err
}

func (r *agentResource) toggle(ctx context.Context, code string, enabled bool) (map[string]any, error) {
	body := map[string]any{"enabled": enabled}
	var out map[string]any
	path := "/admin/agents/" + code + "/toggle"
	if _, err := r.data.apiDo(ctx, "POST", path, body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func agentEnabledFromStatus(info map[string]any) bool {
	if info == nil {
		return true
	}
	for _, key := range []string{"enabled", "active", "is_enabled"} {
		if v, ok := info[key]; ok {
			switch t := v.(type) {
			case bool:
				return t
			case float64:
				return t != 0
			case string:
				return t == "1" || strings.EqualFold(t, "true") || strings.EqualFold(t, "active")
			}
		}
	}
	return true
}

func (r *agentResource) apply(ctx context.Context, plan *agentModel) error {
	code := normalizeAgentCode(plan.Code.ValueString())
	info, status, err := r.readAgent(ctx, code)
	if status == 404 {
		return fmt.Errorf("agent %s introuvable", code)
	}
	if err != nil {
		return err
	}
	enabled := plan.Enabled.ValueBool()
	if !agentEnabledFromStatus(info) == enabled {
		if _, err := r.toggle(ctx, code, enabled); err != nil {
			return err
		}
		info, _, err = r.readAgent(ctx, code)
		if err != nil {
			return err
		}
	}
	raw, _ := json.Marshal(info)
	plan.Enabled = types.BoolValue(agentEnabledFromStatus(info))
	plan.JSON = types.StringValue(string(raw))
	return nil
}

func (r *agentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan agentModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.apply(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Création agent échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *agentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state agentModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	code := normalizeAgentCode(state.Code.ValueString())
	info, status, err := r.readAgent(ctx, code)
	if status == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture agent échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(info)
	state.Enabled = types.BoolValue(agentEnabledFromStatus(info))
	state.JSON = types.StringValue(string(raw))
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *agentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state agentModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.Code = state.Code
	if err := r.apply(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Mise à jour agent échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *agentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state agentModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	code := normalizeAgentCode(state.Code.ValueString())
	if _, err := r.toggle(ctx, code, false); err != nil {
		resp.Diagnostics.AddError("Désactivation agent échouée", err.Error())
	}
}