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
	_ resource.Resource              = &webhookResource{}
	_ resource.ResourceWithConfigure = &webhookResource{}
)

func NewWebhookResource() resource.Resource { return &webhookResource{} }

type webhookResource struct {
	data *providerData
}

type webhookModel struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	URL    types.String `tfsdk:"url"`
	Events types.List   `tfsdk:"events"`
	Secret types.String `tfsdk:"secret"`
}

// Forme JSON échangée avec l'API AISIA (/admin/webhooks).
type webhookAPI struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name"`
	URL    string   `json:"url"`
	Events []string `json:"events"`
	Secret string   `json:"secret,omitempty"`
}

func (r *webhookResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

func (r *webhookResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Un webhook sortant AISIA (/admin/webhooks). Notifie une URL externe sur les événements souscrits.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifiant du webhook (généré par AISIA).",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Nom du webhook.",
				Required:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL de destination des notifications (endpoint HTTPS).",
				Required:            true,
			},
			"events": schema.ListAttribute{
				MarkdownDescription: "Liste des événements souscrits (ex. `[\"invoke.completed\", \"org.created\"]`).",
				Required:            true,
				ElementType:         types.StringType,
			},
			"secret": schema.StringAttribute{
				MarkdownDescription: "Secret partagé pour signer les livraisons (HMAC). Sensible.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (r *webhookResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *webhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan webhookModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var events []string
	resp.Diagnostics.Append(plan.Events.ElementsAs(ctx, &events, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := webhookAPI{Name: plan.Name.ValueString(), URL: plan.URL.ValueString(), Events: events, Secret: plan.Secret.ValueString()}
	var out map[string]any
	if _, err := r.data.apiDo(ctx, "POST", "/admin/webhooks", payload, &out); err != nil {
		resp.Diagnostics.AddError("Création webhook échouée", err.Error())
		return
	}
	plan.ID = types.StringValue(idFromResponse(out))
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read est un no-op : l'API AISIA n'expose pas de GET /admin/webhooks/{id}
// (seulement la collection). L'état Terraform est conservé tel quel.
func (r *webhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state webhookModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *webhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state webhookModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = state.ID
	var events []string
	resp.Diagnostics.Append(plan.Events.ElementsAs(ctx, &events, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := webhookAPI{Name: plan.Name.ValueString(), URL: plan.URL.ValueString(), Events: events, Secret: plan.Secret.ValueString()}
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/webhooks/"+plan.ID.ValueString(), payload, nil); err != nil {
		resp.Diagnostics.AddError("Mise à jour webhook échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *webhookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state webhookModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/webhooks/"+state.ID.ValueString(), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression webhook échouée", err.Error())
	}
}
