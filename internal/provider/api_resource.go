package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// apiAttr — attribut typé d'une entité, dérivé du schéma OpenAPI (CONFIG-SCHEMA chantier B).
type apiAttr struct {
	Name        string // nom de l'attribut Terraform (ex. "service_key")
	WireName    string // nom JSON OpenAPI si différent (ex. request_id -> id)
	TFType      string // "string" | "bool" | "int" | "float" | "list"
	Required    bool
	Sensitive   bool
	JSONEncoded bool // string Terraform contenant jsonencode(object)
	Desc        string
}

func (a apiAttr) wireName() string {
	if a.WireName != "" {
		return a.WireName
	}
	return a.Name
}

// apiResource — resource CRUD GÉNÉRIQUE pour toute entité AISIA administrable par API.
// Si `typedAttrs` est renseigné (généré depuis l'OpenAPI), le schéma expose des attributs
// TYPÉS champ par champ ; sinon fallback sur l'attribut `body` JSON opaque (rétro-compat).
type apiResource struct {
	data             *providerData
	name             string
	path             string
	itemPathTemplate string
	idParam          string
	updateVerb       string
	canDelete        bool
	desc             string
	typedAttrs       []apiAttr // vide = mode body-blob (rétro-compat)
	// Les contrats de corps POST et UPDATE sont distincts : un champ create-only
	// ne doit jamais être envoyé au PUT, et inversement.
	createFields         []string
	updateFields         []string
	updateRequiredFields []string
	// bodyFields est réservé aux rares contrats runtime explicites que FastAPI
	// ne publie pas dans OpenAPI (handlers request.json directs, notamment SCIM).
	// Le body historique reste présent pour préserver configurations et state.
	bodyFields               []string
	createBodyRequiredFields []string
	updateBodyRequiredFields []string
	// sensitiveFields protège les clés secrètes des payloads body-blob.
	sensitiveFields []string
}

var (
	_ resource.Resource                = &apiResource{}
	_ resource.ResourceWithConfigure   = &apiResource{}
	_ resource.ResourceWithImportState = &apiResource{}
)

type apiResourceModel struct {
	ID   types.String `tfsdk:"id"`
	Body types.String `tfsdk:"body"`
	JSON types.String `tfsdk:"json"`
}

func (r *apiResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.name
}

func tfAttrWithLifecycle(a apiAttr, requiresReplace bool) schema.Attribute {
	switch a.TFType {
	case "bool":
		modifiers := []planmodifier.Bool{}
		if requiresReplace {
			modifiers = append(modifiers, boolplanmodifier.RequiresReplace())
		}
		return schema.BoolAttribute{MarkdownDescription: a.Desc, Required: a.Required, Optional: !a.Required, Computed: !a.Required, Sensitive: a.Sensitive, PlanModifiers: modifiers}
	case "int":
		modifiers := []planmodifier.Int64{}
		if requiresReplace {
			modifiers = append(modifiers, int64planmodifier.RequiresReplace())
		}
		return schema.Int64Attribute{MarkdownDescription: a.Desc, Required: a.Required, Optional: !a.Required, Computed: !a.Required, Sensitive: a.Sensitive, PlanModifiers: modifiers}
	case "float":
		modifiers := []planmodifier.Float64{}
		if requiresReplace {
			modifiers = append(modifiers, float64planmodifier.RequiresReplace())
		}
		return schema.Float64Attribute{MarkdownDescription: a.Desc, Required: a.Required, Optional: !a.Required, Computed: !a.Required, Sensitive: a.Sensitive, PlanModifiers: modifiers}
	case "list":
		modifiers := []planmodifier.List{}
		if requiresReplace {
			modifiers = append(modifiers, listplanmodifier.RequiresReplace())
		}
		return schema.ListAttribute{MarkdownDescription: a.Desc, ElementType: types.StringType, Required: a.Required, Optional: !a.Required, Computed: !a.Required, Sensitive: a.Sensitive, PlanModifiers: modifiers}
	default:
		modifiers := []planmodifier.String{}
		if requiresReplace {
			modifiers = append(modifiers, stringplanmodifier.RequiresReplace())
		}
		return schema.StringAttribute{MarkdownDescription: a.Desc, Required: a.Required, Optional: !a.Required, Computed: !a.Required, Sensitive: a.Sensitive, PlanModifiers: modifiers}
	}
}

func tfAttr(a apiAttr) schema.Attribute { return tfAttrWithLifecycle(a, false) }

func containsString(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func (r *apiResource) hasSensitiveAttrs() bool {
	if len(r.sensitiveFields) > 0 {
		return true
	}
	for _, a := range r.typedAttrs {
		if a.Sensitive {
			return true
		}
	}
	return false
}

func (r *apiResource) sensitiveAttrNames() []string {
	names := append([]string(nil), r.sensitiveFields...)
	for _, a := range r.typedAttrs {
		if a.Sensitive {
			names = append(names, a.wireName())
		}
	}
	return names
}

var defaultSensitiveResponseFields = []string{
	"access_token",
	"api_key",
	"api_token",
	"client_secret",
	"config",
	"credential",
	"credentials",
	"headers",
	"key",
	"key_value",
	"password",
	"private_key",
	"recovery_codes",
	"refresh_token",
	"secret",
	"token",
}

func sensitiveFieldSet(fields []string) map[string]struct{} {
	names := make(map[string]struct{}, len(fields)+len(defaultSensitiveResponseFields))
	for _, field := range defaultSensitiveResponseFields {
		names[field] = struct{}{}
	}
	for _, field := range fields {
		names[field] = struct{}{}
	}
	return names
}

// redactSensitiveResponse évite de dupliquer dans l'attribut JSON calculé une
// valeur secrète que Terraform conserve déjà comme entrée Sensitive. Le state
// contient nécessairement les arguments configurés avec Terraform < 1.11 ; il
// ne doit pas en plus conserver une réponse API brute susceptible de les échoer.
func redactSensitiveResponse(value any, sensitive map[string]struct{}) any {
	switch typed := value.(type) {
	case map[string]any:
		out := make(map[string]any, len(typed))
		for key, item := range typed {
			normalized := strings.ToLower(strings.TrimSpace(key))
			_, exact := sensitive[normalized]
			secretSuffix := strings.HasSuffix(normalized, "_password") ||
				strings.HasSuffix(normalized, "_secret") ||
				strings.HasSuffix(normalized, "_token") ||
				strings.HasSuffix(normalized, "_api_key") ||
				strings.HasSuffix(normalized, "_private_key")
			if exact || secretSuffix {
				out[key] = "[REDACTED]"
				continue
			}
			out[key] = redactSensitiveResponse(item, sensitive)
		}
		return out
	case []any:
		out := make([]any, len(typed))
		for i, item := range typed {
			out[i] = redactSensitiveResponse(item, sensitive)
		}
		return out
	default:
		return value
	}
}

func marshalStateJSON(out any, sensitiveFields []string) string {
	out = redactSensitiveResponse(out, sensitiveFieldSet(sensitiveFields))
	raw, _ := json.Marshal(out)
	return string(raw)
}

func (r *apiResource) responseJSON(out any) string {
	return marshalStateJSON(out, r.sensitiveAttrNames())
}

func (r *apiResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs := map[string]schema.Attribute{
		"id":   schema.StringAttribute{MarkdownDescription: "Identifiant de l'entité (retourné par l'API).", Computed: true},
		"json": schema.StringAttribute{MarkdownDescription: "Dernière réponse de l'API (JSON, champs sensibles expurgés).", Computed: true, Sensitive: r.hasSensitiveAttrs()},
	}
	if len(r.typedAttrs) > 0 {
		for _, a := range r.typedAttrs {
			if a.Name == "id" || a.Name == "json" || a.Name == "body" {
				resp.Diagnostics.AddError(
					"Attribut Terraform réservé",
					fmt.Sprintf("Le champ OpenAPI %q de aisia_%s doit être exposé sous un alias Terraform non réservé.", a.wireName(), r.name),
				)
				continue
			}
			createOnly := containsString(r.createFields, a.wireName()) && !containsString(r.updateFields, a.wireName())
			attrs[a.Name] = tfAttrWithLifecycle(a, createOnly)
		}
	} else {
		bodyMods := []planmodifier.String{}
		if r.updateVerb == "" {
			bodyMods = append(bodyMods, stringplanmodifier.RequiresReplace())
		}
		attrs["body"] = schema.StringAttribute{
			MarkdownDescription: "Corps JSON de l'entité (payload create/update), ex. `jsonencode({...})`.",
			Required:            len(r.bodyFields) == 0,
			Optional:            len(r.bodyFields) > 0,
			Computed:            len(r.bodyFields) > 0,
			Sensitive:           len(r.sensitiveFields) > 0,
			PlanModifiers:       bodyMods,
		}
	}
	resp.Schema = schema.Schema{MarkdownDescription: r.desc + docLinksForEndpoint(r.path), Attributes: attrs}
}

func requiredBodyFieldsMissing(body map[string]any, required []string) []string {
	missing := make([]string, 0)
	for _, field := range required {
		value, found := body[field]
		if !found || value == nil {
			missing = append(missing, field)
			continue
		}
		if text, ok := value.(string); ok && strings.TrimSpace(text) == "" {
			missing = append(missing, field)
		}
	}
	sort.Strings(missing)
	return missing
}

func (r *apiResource) decodeBody(value types.String, required []string) (map[string]any, error) {
	if value.IsNull() || value.IsUnknown() || strings.TrimSpace(value.ValueString()) == "" {
		return nil, fmt.Errorf("body JSON requis pour %s", r.path)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(value.ValueString()), &body); err != nil {
		return nil, fmt.Errorf("body doit être un objet JSON valide: %w", err)
	}
	if len(r.bodyFields) > 0 {
		allowed := stringSet(r.bodyFields)
		unknown := make([]string, 0)
		for field := range body {
			if _, ok := allowed[field]; !ok {
				unknown = append(unknown, field)
			}
		}
		if len(unknown) > 0 {
			sort.Strings(unknown)
			return nil, fmt.Errorf("champs hors contrat runtime exact pour %s: %s", r.path, strings.Join(unknown, ", "))
		}
	}
	if missing := requiredBodyFieldsMissing(body, required); len(missing) > 0 {
		return nil, fmt.Errorf("champs JSON requis absents pour %s: %s", r.path, strings.Join(missing, ", "))
	}
	return body, nil
}

func (r *apiResource) remoteBodyFields(out any) map[string]any {
	object, ok := out.(map[string]any)
	if !ok {
		return map[string]any{}
	}
	allowed := stringSet(r.bodyFields)
	remote := make(map[string]any)
	for field, value := range object {
		if _, ok := allowed[field]; ok {
			remote[field] = value
		}
	}
	return remote
}

func (r *apiResource) reconcileBody(current types.String, out any) (types.String, error) {
	remote := r.remoteBodyFields(out)
	if current.IsNull() || current.IsUnknown() {
		raw, err := json.Marshal(remote)
		if err != nil {
			return types.StringNull(), err
		}
		return types.StringValue(string(raw)), nil
	}
	desired, err := r.decodeBody(current, nil)
	if err != nil {
		return types.StringNull(), err
	}
	changed := false
	for field, value := range remote {
		if configured, found := desired[field]; found && !reflect.DeepEqual(configured, value) {
			desired[field] = value
			changed = true
		}
	}
	if !changed {
		return current, nil
	}
	raw, err := json.Marshal(desired)
	if err != nil {
		return types.StringNull(), err
	}
	return types.StringValue(string(raw)), nil
}

func (r *apiResource) reconcileBodyState(ctx context.Context, state *tfsdk.State, out any) diag.Diagnostics {
	var diags diag.Diagnostics
	if len(r.bodyFields) == 0 {
		return diags
	}
	var current types.String
	diags.Append(state.GetAttribute(ctx, path.Root("body"), &current)...)
	if diags.HasError() {
		return diags
	}
	reconciled, err := r.reconcileBody(current, out)
	if err != nil {
		diags.AddError("Réconciliation body échouée", err.Error())
		return diags
	}
	diags.Append(state.SetAttribute(ctx, path.Root("body"), reconciled)...)
	return diags
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

func (r *apiResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	identifier := strings.TrimSpace(req.ID)
	if identifier == "" {
		resp.Diagnostics.AddError("Identifiant d'import vide", "Fournissez l'identifiant exact de l'entité AISIA.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(identifier))...)
}

func idFromResponse(out any) string {
	if m, ok := out.(map[string]any); ok {
		for _, k := range []string{"id", "uuid", "code", "slug", "name"} {
			if v, ok := m[k]; ok && v != nil {
				if value := strings.TrimSpace(fmt.Sprintf("%v", v)); value != "" {
					return value
				}
			}
		}
	}
	return ""
}

// idFromResponseForParam privilégie le nom d'identifiant déclaré par le
// chemin OpenAPI (ex. group_id), puis les alias historiques communs. Une
// création sans identifiant reste une erreur : Terraform ne peut pas inventer
// un chemin item fiable pour Read/Update/Delete.
func idFromResponseForParam(out any, parameter string) string {
	if m, ok := out.(map[string]any); ok && parameter != "" {
		if value, ok := m[parameter]; ok && value != nil {
			if identifier := strings.TrimSpace(fmt.Sprintf("%v", value)); identifier != "" {
				return identifier
			}
		}
	}
	return idFromResponse(out)
}

func (r *apiResource) itemPath(identifier string) (string, error) {
	identifier = strings.TrimSpace(identifier)
	if identifier == "" {
		return "", fmt.Errorf("identifiant item vide")
	}
	if r.itemPathTemplate == "" || r.idParam == "" {
		return "", fmt.Errorf("contrat item OpenAPI incomplet")
	}
	marker := "{" + r.idParam + "}"
	if strings.Count(r.itemPathTemplate, marker) != 1 {
		return "", fmt.Errorf(
			"pathTemplate %q ne contient pas exactement une fois %s",
			r.itemPathTemplate,
			marker,
		)
	}
	resolved := strings.Replace(r.itemPathTemplate, marker, url.PathEscape(identifier), 1)
	if strings.ContainsAny(resolved, "{}") {
		return "", fmt.Errorf("pathTemplate %q contient un placeholder non résolu", r.itemPathTemplate)
	}
	return resolved, nil
}

// bodyFromTypedFields construit un payload limité au contrat OpenAPI du verbe.
func (r *apiResource) bodyFromTypedFields(ctx context.Context, plan tfsdk.Plan, allowedFields []string) (map[string]any, diag.Diagnostics) {
	body := map[string]any{}
	var diags diag.Diagnostics
	for _, a := range r.typedAttrs {
		wireName := a.wireName()
		if !containsString(allowedFields, wireName) {
			continue
		}
		p := path.Root(a.Name)
		switch a.TFType {
		case "bool":
			var v types.Bool
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				body[wireName] = v.ValueBool()
			}
		case "int":
			var v types.Int64
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				body[wireName] = v.ValueInt64()
			}
		case "float":
			var v types.Float64
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				body[wireName] = v.ValueFloat64()
			}
		case "list":
			var v types.List
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				out := []string{}
				v.ElementsAs(ctx, &out, false)
				body[wireName] = out
			}
		default:
			var v types.String
			diags.Append(plan.GetAttribute(ctx, p, &v)...)
			if !v.IsNull() && !v.IsUnknown() {
				if a.JSONEncoded {
					var decoded any
					if err := json.Unmarshal([]byte(v.ValueString()), &decoded); err != nil {
						diags.AddAttributeError(
							p,
							"JSON invalide",
							fmt.Sprintf("Le champ %s doit contenir un jsonencode(...) valide: %v", a.Name, err),
						)
					} else {
						body[wireName] = decoded
					}
				} else {
					body[wireName] = v.ValueString()
				}
			}
		}
	}
	return body, diags
}

// bodyFromTyped conserve le comportement historique pour les actions typées,
// qui utilisent l'intégralité de leur unique schéma de mutation.
func (r *apiResource) bodyFromTyped(ctx context.Context, plan tfsdk.Plan) (map[string]any, diag.Diagnostics) {
	fields := make([]string, 0, len(r.typedAttrs))
	for _, attr := range r.typedAttrs {
		fields = append(fields, attr.wireName())
	}
	return r.bodyFromTypedFields(ctx, plan, fields)
}

func missingRequiredFields(body map[string]any, required []string) []string {
	missing := make([]string, 0)
	for _, field := range required {
		if _, found := body[field]; !found {
			missing = append(missing, field)
		}
	}
	return missing
}

func hasConfiguredUpdateOnly(body map[string]any, createFields, updateFields []string) bool {
	for _, field := range updateFields {
		if containsString(createFields, field) {
			continue
		}
		if _, configured := body[field]; configured {
			return true
		}
	}
	return false
}

func typedStateValue(ctx context.Context, attribute apiAttr, value any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	if value == nil {
		switch attribute.TFType {
		case "bool":
			return types.BoolNull(), diags
		case "int":
			return types.Int64Null(), diags
		case "float":
			return types.Float64Null(), diags
		case "list":
			return types.ListNull(types.StringType), diags
		default:
			return types.StringNull(), diags
		}
	}

	switch attribute.TFType {
	case "bool":
		if typed, ok := value.(bool); ok {
			return types.BoolValue(typed), diags
		}
	case "int":
		switch typed := value.(type) {
		case float64:
			if math.Trunc(typed) == typed {
				return types.Int64Value(int64(typed)), diags
			}
		case int:
			return types.Int64Value(int64(typed)), diags
		case int64:
			return types.Int64Value(typed), diags
		case json.Number:
			if parsed, err := typed.Int64(); err == nil {
				return types.Int64Value(parsed), diags
			}
		}
	case "float":
		switch typed := value.(type) {
		case float64:
			return types.Float64Value(typed), diags
		case int:
			return types.Float64Value(float64(typed)), diags
		case int64:
			return types.Float64Value(float64(typed)), diags
		case json.Number:
			if parsed, err := typed.Float64(); err == nil {
				return types.Float64Value(parsed), diags
			}
		}
	case "list":
		var items []string
		switch typed := value.(type) {
		case []any:
			items = make([]string, 0, len(typed))
			for _, item := range typed {
				items = append(items, fmt.Sprintf("%v", item))
			}
		case []string:
			items = typed
		default:
			break
		}
		if items != nil {
			result, listDiags := types.ListValueFrom(ctx, types.StringType, items)
			diags.Append(listDiags...)
			return result, diags
		}
	default:
		if attribute.JSONEncoded {
			encoded, err := json.Marshal(value)
			if err == nil {
				return types.StringValue(string(encoded)), diags
			}
		} else if typed, ok := value.(string); ok {
			return types.StringValue(typed), diags
		}
	}

	diags.AddError(
		"Réponse API incompatible",
		fmt.Sprintf("Le champ %s ne correspond pas au type Terraform %s.", attribute.Name, attribute.TFType),
	)
	return types.StringNull(), diags
}

func typedStateIsUnknown(ctx context.Context, state *tfsdk.State, attribute apiAttr) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	attributePath := path.Root(attribute.Name)
	switch attribute.TFType {
	case "bool":
		var value types.Bool
		diags.Append(state.GetAttribute(ctx, attributePath, &value)...)
		return value.IsUnknown(), diags
	case "int":
		var value types.Int64
		diags.Append(state.GetAttribute(ctx, attributePath, &value)...)
		return value.IsUnknown(), diags
	case "float":
		var value types.Float64
		diags.Append(state.GetAttribute(ctx, attributePath, &value)...)
		return value.IsUnknown(), diags
	case "list":
		var value types.List
		diags.Append(state.GetAttribute(ctx, attributePath, &value)...)
		return value.IsUnknown(), diags
	default:
		var value types.String
		diags.Append(state.GetAttribute(ctx, attributePath, &value)...)
		return value.IsUnknown(), diags
	}
}

func (r *apiResource) hydrateTypedState(ctx context.Context, state *tfsdk.State, out any) diag.Diagnostics {
	var diags diag.Diagnostics
	object, ok := out.(map[string]any)
	if !ok {
		return diags
	}
	for _, attribute := range r.typedAttrs {
		value, found := object[attribute.wireName()]
		if !found {
			unknown, unknownDiags := typedStateIsUnknown(ctx, state, attribute)
			diags.Append(unknownDiags...)
			if unknown && !unknownDiags.HasError() {
				nullValue, nullDiags := typedStateValue(ctx, attribute, nil)
				diags.Append(nullDiags...)
				if !nullDiags.HasError() {
					diags.Append(state.SetAttribute(ctx, path.Root(attribute.Name), nullValue)...)
				}
			}
			continue
		}
		stateValue, valueDiags := typedStateValue(ctx, attribute, value)
		diags.Append(valueDiags...)
		if valueDiags.HasError() {
			continue
		}
		diags.Append(state.SetAttribute(ctx, path.Root(attribute.Name), stateValue)...)
	}
	return diags
}

func (r *apiResource) refreshTypedState(ctx context.Context, state *tfsdk.State, identifier string) diag.Diagnostics {
	var diags diag.Diagnostics
	itemPath, err := r.itemPath(identifier)
	if err != nil {
		diags.AddError("Contrat de relecture invalide", err.Error())
		return diags
	}
	var out any
	code, err := r.data.apiDo(ctx, "GET", itemPath, nil, &out)
	if err != nil {
		if code == 404 {
			diags.AddError("Entité absente après mutation", fmt.Sprintf("GET %s a retourné 404 immédiatement après la mutation confirmée.", itemPath))
		} else {
			diags.AddError("Relecture après mutation échouée", err.Error())
		}
		return diags
	}
	diags.Append(state.SetAttribute(ctx, path.Root("json"), types.StringValue(r.responseJSON(out)))...)
	diags.Append(r.hydrateTypedState(ctx, state, out)...)
	diags.Append(r.reconcileBodyState(ctx, state, out)...)
	return diags
}

func (r *apiResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var body any
	var updateBody map[string]any
	if len(r.typedAttrs) > 0 {
		m, d := r.bodyFromTypedFields(ctx, req.Plan, r.createFields)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		body = m
		updateBody, d = r.bodyFromTypedFields(ctx, req.Plan, r.updateFields)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		var plan apiResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		decoded, err := r.decodeBody(plan.Body, r.createBodyRequiredFields)
		if err != nil {
			resp.Diagnostics.AddError("body JSON invalide", err.Error())
			return
		}
		body = decoded
	}
	var out any
	if _, err := r.data.apiDo(ctx, "POST", r.path, body, &out); err != nil {
		resp.Diagnostics.AddError("Création échouée", err.Error())
		return
	}
	identifier := idFromResponseForParam(out, r.idParam)
	if identifier == "" {
		resp.Diagnostics.AddError(
			"Création sans identifiant exploitable",
			fmt.Sprintf(
				"POST %s n'a retourné ni %q ni un alias d'identifiant reconnu; le state Terraform n'a pas été créé.",
				r.path,
				r.idParam,
			),
		)
		return
	}
	// state = plan (copie tous les attributs typés/body) puis renseigner id/json calculés.
	resp.State.Raw = req.Plan.Raw
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(identifier))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("json"), types.StringValue(r.responseJSON(out)))...)
	if resp.Diagnostics.HasError() {
		return
	}
	if hasConfiguredUpdateOnly(updateBody, r.createFields, r.updateFields) {
		if missing := missingRequiredFields(updateBody, r.updateRequiredFields); len(missing) > 0 {
			resp.Diagnostics.AddError(
				"Mise à jour complémentaire impossible",
				"Champs UPDATE requis absents après création : "+strings.Join(missing, ", "),
			)
			return
		}
		itemPath, err := r.itemPath(identifier)
		if err != nil {
			resp.Diagnostics.AddError("Contrat de mise à jour complémentaire invalide", err.Error())
			return
		}
		var updateOut any
		if _, err := r.data.apiDo(ctx, r.updateVerb, itemPath, updateBody, &updateOut); err != nil {
			resp.Diagnostics.AddError("Mise à jour complémentaire échouée", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("json"), types.StringValue(r.responseJSON(updateOut)))...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	// Le GET item exact est l'autorité de state : il hydrate les defaults serveur
	// et confirme que le POST/PUT a réellement appliqué les attributs attendus.
	resp.Diagnostics.Append(r.refreshTypedState(ctx, &resp.State, identifier)...)
}

func (r *apiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var id types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &id)...)
	if resp.Diagnostics.HasError() || id.ValueString() == "" {
		return
	}
	itemPath, err := r.itemPath(id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Contrat de lecture invalide", err.Error())
		return
	}
	var out any
	code, err := r.data.apiDo(ctx, "GET", itemPath, nil, &out)
	if code == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture échouée", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("json"), types.StringValue(r.responseJSON(out)))...)
	resp.Diagnostics.Append(r.hydrateTypedState(ctx, &resp.State, out)...)
	resp.Diagnostics.Append(r.reconcileBodyState(ctx, &resp.State, out)...)
}

func (r *apiResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var id types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &id)...)
	if resp.Diagnostics.HasError() {
		return
	}
	itemPath, err := r.itemPath(id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Contrat de mise à jour invalide", err.Error())
		return
	}
	verb := r.updateVerb
	if verb == "" {
		verb = "PUT"
	}
	var body any
	if len(r.typedAttrs) > 0 {
		m, d := r.bodyFromTypedFields(ctx, req.Plan, r.updateFields)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		if missing := missingRequiredFields(m, r.updateRequiredFields); len(missing) > 0 {
			resp.Diagnostics.AddError(
				"Corps UPDATE incomplet",
				"Champs OpenAPI requis absents : "+strings.Join(missing, ", "),
			)
			return
		}
		body = m
	} else {
		var plan, state apiResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if plan.Body.IsNull() || plan.Body.IsUnknown() {
			plan.Body = state.Body
		}
		decoded, err := r.decodeBody(plan.Body, r.updateBodyRequiredFields)
		if err != nil {
			resp.Diagnostics.AddError("body JSON invalide", err.Error())
			return
		}
		body = decoded
	}
	var out any
	if _, err := r.data.apiDo(ctx, verb, itemPath, body, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour échouée", err.Error())
		return
	}
	resp.State.Raw = req.Plan.Raw
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
	if len(r.bodyFields) > 0 {
		var bodyValue types.String
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("body"), &bodyValue)...)
		var plannedBody types.String
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("body"), &plannedBody)...)
		if !plannedBody.IsNull() && !plannedBody.IsUnknown() {
			bodyValue = plannedBody
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("body"), bodyValue)...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("json"), types.StringValue(r.responseJSON(out)))...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(r.refreshTypedState(ctx, &resp.State, id.ValueString())...)
}

func (r *apiResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var id types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &id)...)
	if resp.Diagnostics.HasError() || id.ValueString() == "" || !r.canDelete {
		return
	}
	itemPath, err := r.itemPath(id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Contrat de suppression invalide", err.Error())
		return
	}
	if _, err := r.data.apiDo(ctx, "DELETE", itemPath, nil, nil); err != nil {
		resp.Diagnostics.AddError("Suppression échouée", err.Error())
	}
}
