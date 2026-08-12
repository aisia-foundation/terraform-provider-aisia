package provider

import (
	"context"
	"crypto/rand"
	"encoding/hex"
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
	_ resource.Resource                = &webhookResource{}
	_ resource.ResourceWithConfigure   = &webhookResource{}
	_ resource.ResourceWithImportState = &webhookResource{}
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
	OrgID  types.String `tfsdk:"org_id"`
	Active types.Bool   `tfsdk:"active"`
}

// Forme JSON échangée avec l'API AISIA (/admin/webhooks).
type webhookAPI struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name"`
	URL    string   `json:"url"`
	Events []string `json:"events"`
	Secret string   `json:"secret,omitempty"`
	OrgID  string   `json:"org_id,omitempty"`
	Active *bool    `json:"active,omitempty"`
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
				MarkdownDescription: "Secret partagé pour signer les livraisons (HMAC). S'il est omis, le provider en génère un et le conserve comme sortie sensible.",
				Optional:            true,
				Computed:            true,
				Sensitive:           true,
			},
			"org_id": schema.StringAttribute{
				MarkdownDescription: "Organisation propriétaire (create-only).",
				Optional:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"active": schema.BoolAttribute{
				MarkdownDescription: "Webhook actif (champ UPDATE).",
				Optional:            true,
				Computed:            true,
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
	activeConfigured := !plan.Active.IsNull() && !plan.Active.IsUnknown()
	desiredActive := plan.Active.ValueBool()
	if !plan.Secret.IsNull() && !plan.Secret.IsUnknown() && strings.TrimSpace(plan.Secret.ValueString()) == "" {
		resp.Diagnostics.AddError("Secret webhook invalide", "Un secret explicitement configuré ne peut pas être vide ou blanc ; omettez-le pour laisser le provider en générer un.")
		return
	}
	if plan.Secret.IsNull() || plan.Secret.IsUnknown() {
		secret, err := generateWebhookSecret()
		if err != nil {
			resp.Diagnostics.AddError("Génération secret webhook échouée", err.Error())
			return
		}
		plan.Secret = types.StringValue(secret)
	}
	payload := webhookAPI{Name: plan.Name.ValueString(), URL: plan.URL.ValueString(), Events: events, Secret: plan.Secret.ValueString(), OrgID: knownStr(plan.OrgID)}
	var out map[string]any
	if _, err := r.data.apiDo(ctx, "POST", "/admin/webhooks", payload, &out); err != nil {
		resp.Diagnostics.AddError("Création webhook échouée", err.Error())
		return
	}
	identifier := idFromResponse(out)
	if identifier == "" {
		resp.Diagnostics.AddError("Création webhook sans identifiant", "POST /admin/webhooks n'a retourné aucun id exploitable.")
		return
	}
	plan.ID = types.StringValue(identifier)
	if active, ok := out["active"].(bool); ok {
		plan.Active = types.BoolValue(active)
	} else if plan.Active.IsNull() || plan.Active.IsUnknown() {
		// create_webhook persiste active=true par défaut.
		plan.Active = types.BoolValue(true)
	}
	// Le POST a déjà créé une entité distante. Conserver immédiatement son ID
	// et le secret one-shot permet à Terraform de la reprendre/détruire même si
	// le PUT complémentaire d'activation échoue ensuite (state alors tainted).
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if activeConfigured {
		active := desiredActive
		payload.Active = &active
		if _, err := r.data.apiDo(ctx, "PUT", "/admin/webhooks/"+apiPathSegment(identifier), payload, nil); err != nil {
			resp.Diagnostics.AddError("Activation webhook après création échouée", err.Error())
			return
		}
		plan.Active = types.BoolValue(desiredActive)
		resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	}
}

func generateWebhookSecret() (string, error) {
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return "", err
	}
	return hex.EncodeToString(raw), nil
}

func (r *webhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state webhookModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out any
	if _, err := r.data.apiDo(ctx, "GET", "/admin/webhooks", nil, &out); err != nil {
		resp.Diagnostics.AddError("Lecture webhooks échouée", err.Error())
		return
	}
	item, found := collectionItem(out, "webhooks", state.ID.ValueString(), "id", "webhook_id")
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}
	var current webhookAPI
	if err := decodeAPIObject(item, &current); err != nil {
		resp.Diagnostics.AddError("Réponse webhook invalide", err.Error())
		return
	}
	state.Name = types.StringValue(current.Name)
	state.URL = types.StringValue(current.URL)
	if current.OrgID != "" {
		state.OrgID = types.StringValue(current.OrgID)
	}
	if current.Active != nil {
		state.Active = types.BoolValue(*current.Active)
	}
	if state.Secret.IsUnknown() {
		// Les imports ne peuvent pas relire le secret HMAC stocké côté serveur.
		// Null exprime cette absence sans inventer ni exposer une valeur.
		state.Secret = types.StringNull()
	}
	events, diagnostics := types.ListValueFrom(ctx, types.StringType, current.Events)
	resp.Diagnostics.Append(diagnostics...)
	if !diagnostics.HasError() {
		state.Events = events
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
	if !plan.Secret.IsNull() && !plan.Secret.IsUnknown() && strings.TrimSpace(plan.Secret.ValueString()) == "" {
		resp.Diagnostics.AddError("Secret webhook invalide", "Un secret explicitement configuré ne peut pas être vide ou blanc.")
		return
	}
	if plan.Secret.IsNull() || plan.Secret.IsUnknown() {
		plan.Secret = state.Secret
		if plan.Secret.IsUnknown() {
			plan.Secret = types.StringNull()
		}
	}
	if plan.Active.IsNull() || plan.Active.IsUnknown() {
		plan.Active = state.Active
		if plan.Active.IsUnknown() {
			plan.Active = types.BoolNull()
		}
	}
	plan.ID = state.ID
	var events []string
	resp.Diagnostics.Append(plan.Events.ElementsAs(ctx, &events, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := webhookAPI{Name: plan.Name.ValueString(), URL: plan.URL.ValueString(), Events: events, Secret: plan.Secret.ValueString()}
	if !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		active := plan.Active.ValueBool()
		payload.Active = &active
	}
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/webhooks/"+apiPathSegment(plan.ID.ValueString()), payload, nil); err != nil {
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
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/webhooks/"+apiPathSegment(state.ID.ValueString()), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression webhook échouée", err.Error())
	}
}

func (r *webhookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	identifier := strings.TrimSpace(req.ID)
	if identifier == "" {
		resp.Diagnostics.AddError("Identifiant d'import vide", "Fournissez l'identifiant exact du webhook AISIA.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(identifier))...)
}
