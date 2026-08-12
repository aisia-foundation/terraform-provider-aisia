package provider

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_              resource.Resource                = &organizationResource{}
	_              resource.ResourceWithConfigure   = &organizationResource{}
	_              resource.ResourceWithImportState = &organizationResource{}
	orgSlugPattern                                  = regexp.MustCompile(`^[a-z0-9][a-z0-9_-]{1,62}[a-z0-9]$`)
)

func NewOrganizationResource() resource.Resource { return &organizationResource{} }

type organizationResource struct {
	data *providerData
}

type organizationModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Slug           types.String `tfsdk:"slug"`
	Plan           types.String `tfsdk:"plan"`
	ContractType   types.String `tfsdk:"contract_type"`
	Description    types.String `tfsdk:"description"`
	MaxUsers       types.Int64  `tfsdk:"max_users"`
	MaxRequestsDay types.Int64  `tfsdk:"max_requests_day"`
	MaxTokensDay   types.Int64  `tfsdk:"max_tokens_day"`
	SupportTier    types.String `tfsdk:"support_tier"`
	DeployChannel  types.String `tfsdk:"deploy_channel"`
	Status         types.String `tfsdk:"status"`
}

// orgAPI conserve des pointeurs pour distinguer une valeur vide/0 réellement
// renvoyée d'un champ absent de la réponse JSON.
type orgAPI struct {
	ID             string  `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	Slug           *string `json:"slug,omitempty"`
	Plan           *string `json:"plan,omitempty"`
	ContractType   *string `json:"contract_type,omitempty"`
	Description    *string `json:"description,omitempty"`
	MaxUsers       *int64  `json:"max_users,omitempty"`
	MaxRequestsDay *int64  `json:"max_requests_day,omitempty"`
	MaxTokensDay   *int64  `json:"max_tokens_day,omitempty"`
	SupportTier    *string `json:"support_tier,omitempty"`
	DeployChannel  *string `json:"deploy_channel,omitempty"`
	Status         *string `json:"status,omitempty"`
}

func stringPointer(value string) *string { return &value }
func int64Pointer(value int64) *int64    { return &value }

func (r *organizationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization"
}

func (r *organizationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	useStringState := []planmodifier.String{stringplanmodifier.UseStateForUnknown()}
	useIntState := []planmodifier.Int64{int64planmodifier.UseStateForUnknown()}
	resp.Schema = schema.Schema{
		MarkdownDescription: "Une organisation AISIA (tenant), avec son plan, ses quotas et ses canaux d'exploitation. Gérée via le CRUD admin exact.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifiant de l'organisation (généré par AISIA).",
				Computed:            true,
				PlanModifiers:       useStringState,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Nom de l'organisation.",
				Required:            true,
			},
			"slug": schema.StringAttribute{
				MarkdownDescription: "Slug 3–64 caractères. S'il est omis, le provider en dérive un du nom selon le contrat backend.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       useStringState,
			},
			"plan": schema.StringAttribute{
				MarkdownDescription: "Plan AISIA (défaut API `b2c_free`).",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       useStringState,
			},
			"contract_type": schema.StringAttribute{
				MarkdownDescription: "Type de contrat (défaut API `saas`).",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       useStringState,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description métier de l'organisation.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       useStringState,
			},
			"max_users": schema.Int64Attribute{
				MarkdownDescription: "Quota maximal d'utilisateurs (défaut API 5).",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       useIntState,
			},
			"max_requests_day": schema.Int64Attribute{
				MarkdownDescription: "Quota maximal de requêtes par jour (défaut API 1000).",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       useIntState,
			},
			"max_tokens_day": schema.Int64Attribute{
				MarkdownDescription: "Quota maximal de tokens par jour (défaut API 1000000).",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       useIntState,
			},
			"support_tier": schema.StringAttribute{
				MarkdownDescription: "Niveau de support (champ de mise à jour OrgUpdate).",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       useStringState,
			},
			"deploy_channel": schema.StringAttribute{
				MarkdownDescription: "Canal de déploiement (champ de mise à jour OrgUpdate).",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       useStringState,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Statut distant observé de l'organisation.",
				Computed:            true,
				PlanModifiers:       useStringState,
			},
		},
	}
}

func (r *organizationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func normalizeOrgSlug(value string) (string, error) {
	normalized := strings.ToLower(strings.TrimSpace(value))
	normalized = strings.ReplaceAll(normalized, " ", "-")
	if !orgSlugPattern.MatchString(normalized) {
		return "", fmt.Errorf("slug invalide %q : 3–64 caractères ASCII minuscules, chiffres, tirets ou underscores, avec début/fin alphanumériques", value)
	}
	return normalized, nil
}

func resolveOrgCreateDefaults(plan *organizationModel) error {
	slug := strings.TrimSpace(plan.Slug.ValueString())
	if plan.Slug.IsNull() || plan.Slug.IsUnknown() || slug == "" {
		slug = deriveOrgSlug(plan.Name.ValueString())
	}
	normalizedSlug, err := normalizeOrgSlug(slug)
	if err != nil {
		return err
	}
	plan.Slug = types.StringValue(normalizedSlug)
	if plan.Plan.IsNull() || plan.Plan.IsUnknown() {
		plan.Plan = types.StringValue("b2c_free")
	}
	if plan.ContractType.IsNull() || plan.ContractType.IsUnknown() {
		plan.ContractType = types.StringValue("saas")
	}
	if plan.Description.IsNull() || plan.Description.IsUnknown() {
		plan.Description = types.StringValue("")
	}
	if plan.MaxUsers.IsNull() || plan.MaxUsers.IsUnknown() {
		plan.MaxUsers = types.Int64Value(5)
	}
	if plan.MaxRequestsDay.IsNull() || plan.MaxRequestsDay.IsUnknown() {
		plan.MaxRequestsDay = types.Int64Value(1000)
	}
	if plan.MaxTokensDay.IsNull() || plan.MaxTokensDay.IsUnknown() {
		plan.MaxTokensDay = types.Int64Value(1000000)
	}
	if plan.SupportTier.IsUnknown() {
		plan.SupportTier = types.StringNull()
	}
	if plan.DeployChannel.IsUnknown() {
		plan.DeployChannel = types.StringNull()
	}
	if plan.Status.IsUnknown() {
		plan.Status = types.StringNull()
	}
	return nil
}

func orgCreatePayload(plan organizationModel) orgAPI {
	return orgAPI{
		Name:           stringPointer(plan.Name.ValueString()),
		Slug:           stringPointer(plan.Slug.ValueString()),
		Plan:           stringPointer(plan.Plan.ValueString()),
		ContractType:   stringPointer(plan.ContractType.ValueString()),
		Description:    stringPointer(plan.Description.ValueString()),
		MaxUsers:       int64Pointer(plan.MaxUsers.ValueInt64()),
		MaxRequestsDay: int64Pointer(plan.MaxRequestsDay.ValueInt64()),
		MaxTokensDay:   int64Pointer(plan.MaxTokensDay.ValueInt64()),
	}
}

func orgUpdatePayload(plan organizationModel) map[string]any {
	payload := map[string]any{}
	stringsByName := map[string]types.String{
		"name": plan.Name, "slug": plan.Slug, "plan": plan.Plan,
		"contract_type": plan.ContractType, "description": plan.Description,
		"support_tier": plan.SupportTier, "deploy_channel": plan.DeployChannel,
	}
	for name, value := range stringsByName {
		if !value.IsNull() && !value.IsUnknown() {
			payload[name] = value.ValueString()
		}
	}
	intsByName := map[string]types.Int64{
		"max_users": plan.MaxUsers, "max_requests_day": plan.MaxRequestsDay,
		"max_tokens_day": plan.MaxTokensDay,
	}
	for name, value := range intsByName {
		if !value.IsNull() && !value.IsUnknown() {
			payload[name] = value.ValueInt64()
		}
	}
	return payload
}

func mergeOrgUnknowns(plan *organizationModel, state organizationModel) {
	mergeString := func(target *types.String, prior types.String) {
		if target.IsNull() || target.IsUnknown() {
			*target = prior
		}
	}
	mergeInt := func(target *types.Int64, prior types.Int64) {
		if target.IsNull() || target.IsUnknown() {
			*target = prior
		}
	}
	plan.ID = state.ID
	mergeString(&plan.Name, state.Name)
	mergeString(&plan.Slug, state.Slug)
	mergeString(&plan.Plan, state.Plan)
	mergeString(&plan.ContractType, state.ContractType)
	mergeString(&plan.Description, state.Description)
	mergeInt(&plan.MaxUsers, state.MaxUsers)
	mergeInt(&plan.MaxRequestsDay, state.MaxRequestsDay)
	mergeInt(&plan.MaxTokensDay, state.MaxTokensDay)
	mergeString(&plan.SupportTier, state.SupportTier)
	mergeString(&plan.DeployChannel, state.DeployChannel)
	mergeString(&plan.Status, state.Status)
}

func (r *organizationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan organizationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	supportConfigured := !plan.SupportTier.IsNull() && !plan.SupportTier.IsUnknown()
	deployConfigured := !plan.DeployChannel.IsNull() && !plan.DeployChannel.IsUnknown()
	if err := resolveOrgCreateDefaults(&plan); err != nil {
		resp.Diagnostics.AddError("Configuration organisation invalide", err.Error())
		return
	}
	var out orgAPI
	if _, err := r.data.apiDo(ctx, "POST", "/admin/organizations/", orgCreatePayload(plan), &out); err != nil {
		resp.Diagnostics.AddError("Création organisation échouée", err.Error())
		return
	}
	applyOrg(&plan, out)
	if plan.ID.IsNull() || plan.ID.IsUnknown() || strings.TrimSpace(plan.ID.ValueString()) == "" {
		resp.Diagnostics.AddError("Création organisation sans identifiant", "POST /admin/organizations/ n'a retourné aucun id exploitable ; une revue distante est requise avant relance.")
		return
	}
	// Le POST est déjà effectif : conserver immédiatement l'ID et les valeurs
	// avant l'éventuel PUT complémentaire support/deploy.
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if supportConfigured || deployConfigured {
		supplement := map[string]any{}
		if supportConfigured {
			supplement["support_tier"] = plan.SupportTier.ValueString()
		}
		if deployConfigured {
			supplement["deploy_channel"] = plan.DeployChannel.ValueString()
		}
		var updateOut orgAPI
		if _, err := r.data.apiDo(ctx, "PUT", "/admin/organizations/"+apiPathSegment(plan.ID.ValueString()), supplement, &updateOut); err != nil {
			resp.Diagnostics.AddError("Mise à jour complémentaire organisation échouée", err.Error())
			return
		}
		applyOrg(&plan, updateOut)
		resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.refresh(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Relecture organisation après création échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func deriveOrgSlug(name string) string {
	var slug strings.Builder
	lastDash := false
	for _, char := range strings.ToLower(strings.TrimSpace(name)) {
		isASCIIAlphaNumeric := char <= unicode.MaxASCII && (unicode.IsLetter(char) || unicode.IsDigit(char))
		if isASCIIAlphaNumeric {
			slug.WriteRune(char)
			lastDash = false
			continue
		}
		if slug.Len() > 0 && !lastDash {
			slug.WriteByte('-')
			lastDash = true
		}
	}
	value := strings.Trim(slug.String(), "-")
	if len(value) > 64 {
		value = strings.Trim(value[:64], "-")
	}
	if value != "" && len(value) < 3 {
		value = strings.Trim(value+"-org", "-")
	}
	return value
}

func (r *organizationResource) refresh(ctx context.Context, state *organizationModel) error {
	if state.ID.IsNull() || state.ID.IsUnknown() || state.ID.ValueString() == "" {
		return fmt.Errorf("identifiant organisation vide")
	}
	var out orgAPI
	_, err := r.data.apiDo(ctx, "GET", "/admin/organizations/"+apiPathSegment(state.ID.ValueString()), nil, &out)
	if err != nil {
		return err
	}
	applyOrg(state, out)
	if state.SupportTier.IsUnknown() {
		state.SupportTier = types.StringNull()
	}
	if state.DeployChannel.IsUnknown() {
		state.DeployChannel = types.StringNull()
	}
	if state.Status.IsUnknown() {
		state.Status = types.StringNull()
	}
	return nil
}

func (r *organizationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state organizationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out orgAPI
	code, err := r.data.apiDo(ctx, "GET", "/admin/organizations/"+apiPathSegment(state.ID.ValueString()), nil, &out)
	if code == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture organisation échouée", err.Error())
		return
	}
	applyOrg(&state, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *organizationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state organizationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	mergeOrgUnknowns(&plan, state)
	if normalized, err := normalizeOrgSlug(plan.Slug.ValueString()); err != nil {
		resp.Diagnostics.AddError("Configuration organisation invalide", err.Error())
		return
	} else {
		plan.Slug = types.StringValue(normalized)
	}
	var out orgAPI
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/organizations/"+apiPathSegment(state.ID.ValueString()), orgUpdatePayload(plan), &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour organisation échouée", err.Error())
		return
	}
	applyOrg(&plan, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.refresh(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Relecture organisation après mise à jour échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *organizationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state organizationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/organizations/"+apiPathSegment(state.ID.ValueString()), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression organisation échouée", err.Error())
	}
}

func (r *organizationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	identifier := strings.TrimSpace(req.ID)
	if identifier == "" {
		resp.Diagnostics.AddError("Identifiant d'import vide", "Fournissez l'identifiant exact de l'organisation AISIA.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(identifier))...)
}

// applyOrg copie la réponse API dans le modèle d'état, y compris les valeurs
// vides et les quotas égaux à zéro lorsqu'ils sont explicitement présents.
func applyOrg(model *organizationModel, api orgAPI) {
	if api.ID != "" {
		model.ID = types.StringValue(api.ID)
	}
	if api.Name != nil {
		model.Name = types.StringValue(*api.Name)
	}
	if api.Slug != nil {
		model.Slug = types.StringValue(*api.Slug)
	}
	if api.Plan != nil {
		model.Plan = types.StringValue(*api.Plan)
	}
	if api.ContractType != nil {
		model.ContractType = types.StringValue(*api.ContractType)
	}
	if api.Description != nil {
		model.Description = types.StringValue(*api.Description)
	}
	if api.MaxUsers != nil {
		model.MaxUsers = types.Int64Value(*api.MaxUsers)
	}
	if api.MaxRequestsDay != nil {
		model.MaxRequestsDay = types.Int64Value(*api.MaxRequestsDay)
	}
	if api.MaxTokensDay != nil {
		model.MaxTokensDay = types.Int64Value(*api.MaxTokensDay)
	}
	if api.SupportTier != nil {
		model.SupportTier = types.StringValue(*api.SupportTier)
	}
	if api.DeployChannel != nil {
		model.DeployChannel = types.StringValue(*api.DeployChannel)
	}
	if api.Status != nil {
		model.Status = types.StringValue(*api.Status)
	}
}
