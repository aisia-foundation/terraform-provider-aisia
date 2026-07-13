package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// apiAttr — attribut typé d'une entité, dérivé du schéma OpenAPI (CONFIG-SCHEMA chantier B).
type apiAttr struct {
	Name     string // nom du champ (ex. "service_key")
	TFType   string // "string" | "bool" | "int" | "float" | "list"
	Required bool
	Desc     string
}

// apiResource — resource CRUD GÉNÉRIQUE pour toute entité AISIA administrable par API.
// Si `typedAttrs` est renseigné (généré depuis l'OpenAPI), le schéma expose des attributs
// TYPÉS champ par champ ; sinon fallback sur l'attribut `body` JSON opaque (rétro-compat).
type apiResource struct {
	data       *providerData
	name       string
	path       string
	updateVerb string
	canDelete  bool
	desc       string
	typedAttrs []apiAttr // vide = mode body-blob (rétro-compat)
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

func tfAttr(a apiAttr) schema.Attribute {
	switch a.TFType {
	case "bool":
		return schema.BoolAttribute{MarkdownDescription: a.Desc, Required: a.Required, Optional: !a.Required}
	case "int":
		return schema.Int64Attribute{MarkdownDescription: a.Desc, Required: a.Required, Optional: !a.Required}
	case "float":
		return schema.Float64Attribute{MarkdownDescription: a.Desc, Required: a.Required, Optional: !a.Required}
	case "list":
		return schema.ListAttribute{MarkdownDescription: a.Desc, ElementType: types.StringType, Required: a.Required, Optional: !a.Required}
	default:
		return schema.StringAttribute{MarkdownDescription: a.Desc, Required: a.Required, Optional: !a.Required}
	}
}

func (r *apiResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs := map[string]schema.Attribute{
		"id":   schema.StringAttribute{MarkdownDescription: "Identifiant de l'entité (retourné par l'API).", Computed: true},
		"json": schema.StringAttribute{MarkdownDescription: "Dernière réponse brute de l'API (JSON).", Computed: true},
	}
	if len(r.typedAttrs) > 0 {
		for _, a := range r.typedAttrs {
			if a.Name == "id" || a.Name == "json" {
				continue
			}
			attrs[a.Name] = tfAttr(a)
		}
	} else {
		bodyMods := []planmodifier.String{}
		if r.updateVerb == "" {
			bodyMods = append(bodyMods, stringplanmodifier.RequiresReplace())
		}
		attrs["body"] = schema.StringAttribute{MarkdownDescription: "Corps JSON de l'entité (payload create/update), ex. `jsonencode({...})`.", Required: true, PlanModifiers: bodyMods}
	}
	resp.Schema = schema.Schema{MarkdownDescription: r.desc + docLinksForEndpoint(r.path), Attributes: attrs}
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

// bodyFromTyped construit le payload JSON depuis les attributs typés (plan).
func (r *apiResource) bodyFromTyped(ctx context.Context, plan tfsdk.Plan) (map[string]any, diag.Diagnostics) {
	body := map[string]any{}
	var diags diag.Diagnostics
	for _, a := range r.typedAttrs {
		p := path.Root(a.Name)
		switch a.TFType {
		case "bool":
			var v types.Bool
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				body[a.Name] = v.ValueBool()
			}
		case "int":
			var v types.Int64
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				body[a.Name] = v.ValueInt64()
			}
		case "float":
			var v types.Float64
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				body[a.Name] = v.ValueFloat64()
			}
		case "list":
			var v types.List
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				out := []string{}
				v.ElementsAs(ctx, &out, false)
				body[a.Name] = out
			}
		default:
			var v types.String
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				body[a.Name] = v.ValueString()
			}
		}
	}
	return body, diags
}

func (r *apiResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var body any
	if len(r.typedAttrs) > 0 {
		m, d := r.bodyFromTyped(ctx, req.Plan)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		body = m
	} else {
		var plan apiResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if err := json.Unmarshal([]byte(plan.Body.ValueString()), &body); err != nil {
			resp.Diagnostics.AddError("body JSON invalide", err.Error())
			return
		}
	}
	var out any
	if _, err := r.data.apiDo(ctx, "POST", r.path, body, &out); err != nil {
		resp.Diagnostics.AddError("Création échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	// state = plan (copie tous les attributs typés/body) puis renseigner id/json calculés.
	resp.State.Raw = req.Plan.Raw
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(idFromResponse(out)))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("json"), types.StringValue(string(raw)))...)
}

func (r *apiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var id types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &id)...)
	if resp.Diagnostics.HasError() || id.ValueString() == "" {
		return
	}
	var out any
	code, err := r.data.apiDo(ctx, "GET", r.path+"/"+id.ValueString(), nil, &out)
	if code == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("json"), types.StringValue(string(raw)))...)
}

func (r *apiResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var id types.String
	req.State.GetAttribute(ctx, path.Root("id"), &id)
	verb := r.updateVerb
	if verb == "" {
		verb = "PUT"
	}
	var body any
	if len(r.typedAttrs) > 0 {
		m, d := r.bodyFromTyped(ctx, req.Plan)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		body = m
	} else {
		var plan apiResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if err := json.Unmarshal([]byte(plan.Body.ValueString()), &body); err != nil {
			resp.Diagnostics.AddError("body JSON invalide", err.Error())
			return
		}
	}
	var out any
	if _, err := r.data.apiDo(ctx, verb, r.path+"/"+id.ValueString(), body, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	resp.State.Raw = req.Plan.Raw
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("json"), types.StringValue(string(raw)))...)
}

func (r *apiResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var id types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &id)...)
	if resp.Diagnostics.HasError() || id.ValueString() == "" || !r.canDelete {
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", r.path+"/"+id.ValueString(), nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression échouée", err.Error())
	}
}
