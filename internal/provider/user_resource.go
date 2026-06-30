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
	_ resource.Resource              = &userResource{}
	_ resource.ResourceWithConfigure = &userResource{}
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
}

type userAPI struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Role        string `json:"role,omitempty"`
	OrgID       string `json:"org_id,omitempty"`
	Password    string `json:"password,omitempty"`           // (entrée)
	GeneratedPw string `json:"generated_password,omitempty"` // (sortie one-shot)
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
				MarkdownDescription: "Email (immuable).",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"display_name": schema.StringAttribute{MarkdownDescription: "Nom affiché.", Optional: true, Computed: true},
			"role": schema.StringAttribute{
				MarkdownDescription: "Rôle : super_admin|org_owner|org_admin|org_user|b2c_free|investor (défaut org_user).",
				Optional:            true, Computed: true,
			},
			"org_id":   schema.StringAttribute{MarkdownDescription: "Org de rattachement.", Optional: true},
			"password": schema.StringAttribute{MarkdownDescription: "Mot de passe (sensible). Vide = auto-généré.", Optional: true, Sensitive: true},
			"generated_password": schema.StringAttribute{
				MarkdownDescription: "Mot de passe auto-généré (one-shot, sensible).",
				Computed:            true, Sensitive: true,
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
	payload := userAPI{
		Email: plan.Email.ValueString(), DisplayName: plan.DisplayName.ValueString(),
		Role: plan.Role.ValueString(), OrgID: plan.OrgID.ValueString(), Password: plan.Password.ValueString(),
	}
	var out userAPI
	if _, err := r.data.apiDo(ctx, "POST", "/admin/users", payload, &out); err != nil {
		resp.Diagnostics.AddError("Création utilisateur échouée", err.Error())
		return
	}
	if out.ID != "" {
		plan.ID = types.StringValue(out.ID)
	}
	if out.DisplayName != "" {
		plan.DisplayName = types.StringValue(out.DisplayName)
	}
	if out.Role != "" {
		plan.Role = types.StringValue(out.Role)
	}
	// mot de passe auto-généré renvoyé une seule fois.
	plan.GeneratedPassword = types.StringValue(out.GeneratedPw)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *userResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state userModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out userAPI
	code, err := r.data.apiDo(ctx, "GET", "/admin/users/"+state.ID.ValueString(), nil, &out)
	if code == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture utilisateur échouée", err.Error())
		return
	}
	if out.Email != "" {
		state.Email = types.StringValue(out.Email)
	}
	if out.DisplayName != "" {
		state.DisplayName = types.StringValue(out.DisplayName)
	}
	if out.Role != "" {
		state.Role = types.StringValue(out.Role)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *userResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan userModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Le rôle est l'attribut mutable principal (endpoint dédié).
	if _, err := r.data.apiDo(ctx, "PUT", "/admin/users/"+plan.ID.ValueString()+"/role",
		map[string]string{"role": plan.Role.ValueString()}, nil); err != nil {
		resp.Diagnostics.AddError("Mise à jour rôle échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *userResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state userModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", "/admin/users/"+state.ID.ValueString(), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression utilisateur échouée", err.Error())
	}
}
