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
	_ resource.Resource                = &userResource{}
	_ resource.ResourceWithConfigure   = &userResource{}
	_ resource.ResourceWithImportState = &userResource{}
)

func NewUserResource() resource.Resource { return &userResource{} }

type userResource struct {
	data *providerData
}

type userModel struct {
	ID                types.String `tfsdk:"id"`
	Email             types.String `tfsdk:"email"`
	DisplayName       types.String `tfsdk:"display_name"`
	Role              types.String `tfsdk:"role"`
	OrgID             types.String `tfsdk:"org_id"`
	Password          types.String `tfsdk:"password"`
	GeneratedPassword types.String `tfsdk:"generated_password"`
	Active            types.Bool   `tfsdk:"active"`
	UserType          types.String `tfsdk:"user_type"`
}

type userAPI struct {
	ID          string `json:"user_id,omitempty"`
	Email       string `json:"email,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Role        string `json:"role,omitempty"`
	OrgID       string `json:"org_id,omitempty"`
	Password    string `json:"password,omitempty"`      // (entrée)
	GeneratedPw string `json:"temp_password,omitempty"` // (sortie one-shot)
	Active      *bool  `json:"active,omitempty"`
	UserType    string `json:"user_type,omitempty"`
}

func (r *userResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *userResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Un utilisateur AISIA. Si `password` n'est pas fourni, AISIA en génère un (exposé une seule fois dans `generated_password`).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "ID utilisateur (généré).",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "Email utilisateur (créable et modifiable par l'API admin).",
				Optional:            true,
				Computed:            true,
			},
			"display_name": schema.StringAttribute{MarkdownDescription: "Nom affiché.", Optional: true, Computed: true},
			"role": schema.StringAttribute{
				MarkdownDescription: "Rôle : super_admin|org_owner|org_admin|org_user|b2c_free|investor (défaut org_user).",
				Optional:            true, Computed: true,
			},
			"org_id": schema.StringAttribute{
				MarkdownDescription: "Org de rattachement (create-only).",
				Optional:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Attribut historique conservé pour décoder les anciens states. Le contrat OpenAPI courant ne permet pas de fixer le mot de passe ; laissez-le vide pour obtenir `generated_password`.",
				Optional:            true,
				Sensitive:           true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"generated_password": schema.StringAttribute{
				MarkdownDescription: "Mot de passe auto-généré (one-shot, sensible).",
				Computed:            true, Sensitive: true,
			},
			"active": schema.BoolAttribute{
				MarkdownDescription: "Compte actif (champ AdminUserUpdate).",
				Optional:            true,
				Computed:            true,
			},
			"user_type": schema.StringAttribute{
				MarkdownDescription: "Type utilisateur persistant (champ AdminUserUpdate).",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (r *userResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *userResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan userModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !plan.Password.IsNull() && !plan.Password.IsUnknown() && plan.Password.ValueString() != "" {
		resp.Diagnostics.AddError(
			"Mot de passe fourni non supporté",
			"Le requestBody OpenAPI de POST /admin/users n'accepte pas `password`. AISIA refuse de l'ignorer silencieusement ; laissez ce champ vide et utilisez le `generated_password` one-shot.",
		)
		return
	}
	if plan.Email.IsNull() || plan.Email.IsUnknown() || strings.TrimSpace(plan.Email.ValueString()) == "" {
		resp.Diagnostics.AddError("Email utilisateur manquant", "`email` est requis lors de la création d'un utilisateur AISIA.")
		return
	}
	activeConfigured := !plan.Active.IsNull() && !plan.Active.IsUnknown()
	userTypeConfigured := !plan.UserType.IsNull() && !plan.UserType.IsUnknown()
	if plan.DisplayName.IsNull() || plan.DisplayName.IsUnknown() || plan.DisplayName.ValueString() == "" {
		defaultName := strings.SplitN(plan.Email.ValueString(), "@", 2)[0]
		plan.DisplayName = types.StringValue(defaultName)
	}
	if plan.Role.IsNull() || plan.Role.IsUnknown() || plan.Role.ValueString() == "" {
		plan.Role = types.StringValue("org_user")
	}
	payload := map[string]any{
		"email": plan.Email.ValueString(), "display_name": plan.DisplayName.ValueString(),
		"role": plan.Role.ValueString(),
	}
	if !plan.OrgID.IsNull() && !plan.OrgID.IsUnknown() && strings.TrimSpace(plan.OrgID.ValueString()) != "" {
		payload["org_id"] = plan.OrgID.ValueString()
	}
	var out userAPI
	if _, err := r.data.apiDo(ctx, "POST", "/admin/users", payload, &out); err != nil {
		resp.Diagnostics.AddError("Création utilisateur échouée", err.Error())
		return
	}
	if out.ID == "" {
		resp.Diagnostics.AddError(
			"Création utilisateur sans identifiant",
			"POST /admin/users n'a pas retourné le champ user_id requis ; aucun state incomplet n'a été enregistré.",
		)
		return
	}
	plan.ID = types.StringValue(out.ID)
	if out.DisplayName != "" {
		plan.DisplayName = types.StringValue(out.DisplayName)
	}
	if out.Role != "" {
		plan.Role = types.StringValue(out.Role)
	}
	// Mot de passe auto-généré renvoyé une seule fois. Résoudre toutes les
	// valeurs Computed avant le premier State.Set, même lors d'une réponse partielle.
	if out.GeneratedPw != "" {
		plan.GeneratedPassword = types.StringValue(out.GeneratedPw)
	} else {
		plan.GeneratedPassword = types.StringNull()
	}
	if plan.OrgID.IsUnknown() {
		plan.OrgID = types.StringNull()
	}
	if plan.Active.IsUnknown() {
		plan.Active = types.BoolNull()
	}
	if plan.UserType.IsUnknown() {
		plan.UserType = types.StringNull()
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if out.GeneratedPw == "" {
		resp.Diagnostics.AddError("Création utilisateur sans secret one-shot", "La réponse 2xx ne contient pas `temp_password`. L'ID a été conservé dans le state pour empêcher un second POST ; auditez ou réinitialisez explicitement le mot de passe.")
		return
	}
	if activeConfigured || userTypeConfigured {
		supplement := map[string]any{}
		if activeConfigured {
			supplement["active"] = plan.Active.ValueBool()
		}
		if userTypeConfigured {
			supplement["user_type"] = plan.UserType.ValueString()
		}
		if _, err := r.data.apiDo(ctx, "PUT", "/admin/users/"+apiPathSegment(plan.ID.ValueString()), supplement, nil); err != nil {
			resp.Diagnostics.AddError("Mise à jour complémentaire utilisateur échouée", err.Error())
			return
		}
	}
	if err := r.refresh(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Relecture utilisateur après création échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func boolFromAPI(object map[string]any, key string) (bool, bool) {
	value, found := object[key]
	if !found || value == nil {
		return false, false
	}
	switch typed := value.(type) {
	case bool:
		return typed, true
	case float64:
		return typed != 0, true
	case string:
		return typed == "1" || strings.EqualFold(typed, "true") || strings.EqualFold(typed, "active"), true
	default:
		return false, false
	}
}

func applyUserPublic(state *userModel, item map[string]any) {
	if value, found := item["email"].(string); found {
		state.Email = types.StringValue(value)
	}
	if value, found := item["display_name"].(string); found {
		state.DisplayName = types.StringValue(value)
	}
	if value, found := item["role"].(string); found {
		state.Role = types.StringValue(value)
	}
	if value, found := item["org_id"].(string); found {
		state.OrgID = types.StringValue(value)
	} else if state.OrgID.IsUnknown() {
		state.OrgID = types.StringNull()
	}
	if value, found := item["user_type"].(string); found {
		state.UserType = types.StringValue(value)
	} else if state.UserType.IsUnknown() {
		state.UserType = types.StringNull()
	}
	if active, found := boolFromAPI(item, "active"); found {
		state.Active = types.BoolValue(active)
	} else if state.Active.IsUnknown() {
		state.Active = types.BoolNull()
	}
	if state.GeneratedPassword.IsUnknown() {
		state.GeneratedPassword = types.StringNull()
	}
}

func (r *userResource) refresh(ctx context.Context, state *userModel) error {
	var out any
	if _, err := r.data.apiDo(ctx, "GET", "/admin/users", nil, &out); err != nil {
		return err
	}
	item, found := collectionItem(out, "users", state.ID.ValueString(), "id", "user_id")
	if !found {
		return fmt.Errorf("utilisateur %s absent de GET /admin/users", state.ID.ValueString())
	}
	applyUserPublic(state, item)
	return nil
}

func (r *userResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state userModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out any
	if _, err := r.data.apiDo(ctx, "GET", "/admin/users", nil, &out); err != nil {
		resp.Diagnostics.AddError("Lecture utilisateur échouée", err.Error())
		return
	}
	item, found := collectionItem(out, "users", state.ID.ValueString(), "id", "user_id")
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}
	applyUserPublic(&state, item)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *userResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state userModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = state.ID
	if plan.Email.IsNull() || plan.Email.IsUnknown() {
		plan.Email = state.Email
	}
	if plan.DisplayName.IsNull() || plan.DisplayName.IsUnknown() {
		plan.DisplayName = state.DisplayName
	}
	if plan.Role.IsNull() || plan.Role.IsUnknown() {
		plan.Role = state.Role
	}
	if plan.OrgID.IsNull() || plan.OrgID.IsUnknown() {
		plan.OrgID = state.OrgID
	}
	if plan.Password.IsNull() || plan.Password.IsUnknown() {
		plan.Password = state.Password
	}
	if plan.GeneratedPassword.IsNull() || plan.GeneratedPassword.IsUnknown() {
		plan.GeneratedPassword = state.GeneratedPassword
	}
	if plan.Active.IsNull() || plan.Active.IsUnknown() {
		plan.Active = state.Active
	}
	if plan.UserType.IsNull() || plan.UserType.IsUnknown() {
		plan.UserType = state.UserType
	}
	payload := map[string]any{}
	for name, value := range map[string]types.String{
		"email": plan.Email, "display_name": plan.DisplayName,
		"role": plan.Role, "user_type": plan.UserType,
	} {
		if !value.IsNull() && !value.IsUnknown() {
			payload[name] = value.ValueString()
		}
	}
	if !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		payload["active"] = plan.Active.ValueBool()
	}
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/users/"+apiPathSegment(plan.ID.ValueString()), payload, nil); err != nil {
		resp.Diagnostics.AddError("Mise à jour utilisateur échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.refresh(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Relecture utilisateur après mise à jour échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *userResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	identifier := strings.TrimSpace(req.ID)
	if identifier == "" {
		resp.Diagnostics.AddError("Identifiant d'import vide", "Fournissez l'identifiant exact de l'utilisateur AISIA.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(identifier))...)
}

func (r *userResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state userModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/users/"+apiPathSegment(state.ID.ValueString()), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression utilisateur échouée", err.Error())
	}
}
