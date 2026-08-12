package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"

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

// actionApiResource — mutation admin (POST/PUT/PATCH/DELETE) ou action idempotente.
// path = chemin littéral ; pathTemplate = chemin avec {placeholders} résolus depuis body.
type actionApiResource struct {
	data                      *providerData
	name                      string
	path                      string
	pathTemplate              string
	method                    string
	readPath                  string
	desc                      string
	hasJSONBody               bool
	bodyRequired              bool
	allowAdditionalBodyFields bool
	queryParams               []actionQueryParam
	bodyFields                []string
	bodyRequiredFields        []string
	typedAttrs                []apiAttr
	fieldAliases              map[string]string
	sensitiveFields           []string
	secretOutputs             []string
}

type actionQueryParam struct {
	Name     string
	Required bool
}

var (
	_       resource.Resource               = &actionApiResource{}
	_       resource.ResourceWithConfigure  = &actionApiResource{}
	_       resource.ResourceWithModifyPlan = &actionApiResource{}
	_tplVar                                 = regexp.MustCompile(`\{([a-zA-Z0-9_]+)\}`)
)

func (r *actionApiResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.name
}

func (r *actionApiResource) actionDocPath() string {
	if r.pathTemplate != "" {
		return r.pathTemplate
	}
	return r.path
}

func (r *actionApiResource) bodyDescription() string {
	description := "Corps JSON complémentaire (ex. `jsonencode({...})`). Les clés de chemin/query sont utilisées comme paramètres de transport et retirées du corps HTTP sauf si l'OpenAPI les déclare aussi dans le body."
	required := make([]string, 0, len(r.queryParams))
	for _, parameter := range r.queryParams {
		if parameter.Required {
			required = append(required, "`"+parameter.Name+"`")
		}
	}
	if len(required) > 0 {
		description += " Paramètres query OpenAPI requis à fournir dans ce JSON : " + strings.Join(required, ", ") + "."
	}
	return description
}

func (r *actionApiResource) hasSensitiveAttrs() bool {
	if len(r.sensitiveFields) > 0 || len(r.secretOutputs) > 0 {
		return true
	}
	for _, attr := range r.typedAttrs {
		if attr.Sensitive {
			return true
		}
	}
	return false
}

func (r *actionApiResource) sensitiveAttrNames() []string {
	names := append([]string(nil), r.sensitiveFields...)
	for _, attr := range r.typedAttrs {
		if attr.Sensitive {
			names = append(names, attr.wireName())
		}
	}
	return names
}

func actionTFAttr(attr apiAttr) schema.Attribute {
	switch attr.TFType {
	case "bool":
		return schema.BoolAttribute{MarkdownDescription: attr.Desc, Required: attr.Required, Optional: !attr.Required, Sensitive: attr.Sensitive, PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()}}
	case "int":
		return schema.Int64Attribute{MarkdownDescription: attr.Desc, Required: attr.Required, Optional: !attr.Required, Sensitive: attr.Sensitive, PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()}}
	case "float":
		return schema.Float64Attribute{MarkdownDescription: attr.Desc, Required: attr.Required, Optional: !attr.Required, Sensitive: attr.Sensitive, PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()}}
	case "list":
		return schema.ListAttribute{MarkdownDescription: attr.Desc, ElementType: types.StringType, Required: attr.Required, Optional: !attr.Required, Sensitive: attr.Sensitive, PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()}}
	default:
		return schema.StringAttribute{MarkdownDescription: attr.Desc, Required: attr.Required, Optional: !attr.Required, Sensitive: attr.Sensitive, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}}
	}
}

func (r *actionApiResource) bodyInputRequired() bool {
	available := make(map[string]struct{}, len(r.typedAttrs))
	for _, attr := range r.typedAttrs {
		available[attr.wireName()] = struct{}{}
	}
	for source, target := range r.fieldAliases {
		if _, ok := available[source]; ok {
			available[target] = struct{}{}
		}
	}
	for _, match := range _tplVar.FindAllStringSubmatch(r.pathTemplate, -1) {
		if _, ok := available[match[1]]; !ok {
			return true
		}
	}
	for _, parameter := range r.queryParams {
		if parameter.Required {
			if _, ok := available[parameter.Name]; !ok {
				return true
			}
		}
	}
	return r.bodyRequired && len(r.typedAttrs) == 0
}

func (r *actionApiResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Identifiant fixe de la mutation.",
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"json": schema.StringAttribute{
			MarkdownDescription: "Dernière réponse de l'API (JSON, champs sensibles expurgés).",
			Computed:            true,
			Sensitive:           r.hasSensitiveAttrs(),
		},
	}
	bodyRequired := r.bodyInputRequired()
	attrs["body"] = schema.StringAttribute{
		MarkdownDescription: r.bodyDescription(),
		Required:            bodyRequired,
		Optional:            !bodyRequired,
		Sensitive:           r.hasSensitiveAttrs(),
		PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
	}
	for _, attr := range r.typedAttrs {
		if attr.Name == "id" || attr.Name == "json" || attr.Name == "body" {
			resp.Diagnostics.AddError(
				"Attribut Terraform réservé",
				fmt.Sprintf("Le champ OpenAPI %q de aisia_%s doit être exposé sous un alias Terraform non réservé.", attr.wireName(), r.name),
			)
			continue
		}
		attrs[attr.Name] = actionTFAttr(attr)
	}
	for _, output := range r.secretOutputs {
		if _, reserved := attrs[output]; reserved {
			resp.Diagnostics.AddError(
				"Sortie secrète en collision",
				fmt.Sprintf("La sortie calculée %s entre en collision avec un attribut de aisia_%s.", output, r.name),
			)
			continue
		}
		attrs[output] = schema.StringAttribute{
			MarkdownDescription: "Sortie one-shot de l'API, conservée séparément de `json` expurgé.",
			Computed:            true,
			Sensitive:           true,
		}
	}
	resp.Schema = schema.Schema{
		MarkdownDescription: r.desc + docLinksForEndpoint(r.actionDocPath()),
		Attributes:          attrs,
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

func (r *actionApiResource) isLegacyStateID(id types.String) bool {
	return !id.IsNull() && !id.IsUnknown() && id.ValueString() != "" && id.ValueString() != r.actionID()
}

func (r *actionApiResource) legacyStateMessage(identifier string) string {
	return fmt.Sprintf(
		"Le state existant de aisia_%s porte l'identifiant d'entité historique %q. Ce type est désormais une action exacte et non importable ; AISIA refuse de rejouer automatiquement sa mutation POST, car cela dupliquerait l'entité. Retirez d'abord cette ressource de la configuration sans destruction distante (bloc removed avec destroy=false), puis retirez/migrez explicitement son state après revue.",
		r.name,
		identifier,
	)
}

func (r *actionApiResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || !req.State.Raw.IsKnown() {
		return
	}
	var identifier types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &identifier)...)
	if resp.Diagnostics.HasError() || !r.isLegacyStateID(identifier) {
		return
	}
	resp.Diagnostics.AddError(
		"Migration explicite du state requise",
		r.legacyStateMessage(identifier.ValueString()),
	)
}

func (r *actionApiResource) httpMethod() string {
	if r.method != "" {
		return strings.ToUpper(r.method)
	}
	return "POST"
}

func (r *actionApiResource) actionID() string { return "_action" }

func (r *actionApiResource) parseBody(rawValue types.String) (map[string]any, any, error) {
	raw := rawValue.ValueString()
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

func stringSet(values []string) map[string]struct{} {
	result := make(map[string]struct{}, len(values))
	for _, value := range values {
		result[value] = struct{}{}
	}
	return result
}

func (r *actionApiResource) requestBody(body any) any {
	if !r.hasJSONBody {
		return nil
	}
	object, ok := body.(map[string]any)
	if !ok {
		return body
	}
	clean := make(map[string]any, len(object))
	for key, value := range object {
		clean[key] = value
	}
	bodyFields := stringSet(r.bodyFields)
	for _, match := range _tplVar.FindAllStringSubmatch(r.pathTemplate, -1) {
		if _, belongsToBody := bodyFields[match[1]]; !belongsToBody {
			delete(clean, match[1])
		}
	}
	for _, parameter := range r.queryParams {
		if _, belongsToBody := bodyFields[parameter.Name]; !belongsToBody {
			delete(clean, parameter.Name)
		}
	}
	return clean
}

func (r *actionApiResource) validateRequestBody(body any) error {
	if !r.hasJSONBody || len(r.bodyFields) == 0 || r.allowAdditionalBodyFields {
		return nil
	}
	object, ok := body.(map[string]any)
	if !ok {
		return fmt.Errorf("body JSON doit être un objet pour %s %s", r.httpMethod(), r.actionDocPath())
	}
	allowed := stringSet(r.bodyFields)
	unknown := make([]string, 0)
	for field := range object {
		if _, ok := allowed[field]; !ok {
			unknown = append(unknown, field)
		}
	}
	if len(unknown) > 0 {
		sort.Strings(unknown)
		return fmt.Errorf("body JSON contient des champs hors contrat pour %s %s: %s", r.httpMethod(), r.actionDocPath(), strings.Join(unknown, ", "))
	}
	return nil
}

func (r *actionApiResource) applyFieldAliases(fields map[string]any, body any) any {
	object, objectBody := body.(map[string]any)
	for source, target := range r.fieldAliases {
		value, exists := fields[source]
		if !exists {
			continue
		}
		fields[target] = value
		delete(fields, source)
		if objectBody {
			object[target] = value
			delete(object, source)
		}
	}
	return body
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
			return url.PathEscape(fmt.Sprintf("%v", v))
		}
		missing = append(missing, key)
		return seg
	})
	if len(missing) > 0 {
		return "", fmt.Errorf("body JSON : clés manquantes pour pathTemplate %s : %s", tmpl, strings.Join(missing, ", "))
	}
	return out, nil
}

func actionQueryValues(value any) []string {
	switch typed := value.(type) {
	case nil:
		return nil
	case []any:
		values := make([]string, 0, len(typed))
		for _, item := range typed {
			values = append(values, fmt.Sprintf("%v", item))
		}
		return values
	case []string:
		return typed
	default:
		return []string{fmt.Sprintf("%v", typed)}
	}
}

func (r *actionApiResource) resolveQuery(rawPath string, fields map[string]any) (string, error) {
	if len(r.queryParams) == 0 {
		return rawPath, nil
	}
	parsed, err := url.Parse(rawPath)
	if err != nil {
		return "", fmt.Errorf("chemin mutation invalide %q: %w", rawPath, err)
	}
	query := parsed.Query()
	for _, parameter := range r.queryParams {
		values := actionQueryValues(fields[parameter.Name])
		if len(values) == 0 {
			if parameter.Required {
				return "", fmt.Errorf("body JSON : query param requis manquant : %s", parameter.Name)
			}
			continue
		}
		for _, value := range values {
			query.Add(parameter.Name, value)
		}
	}
	parsed.RawQuery = query.Encode()
	return parsed.String(), nil
}

func (r *actionApiResource) runMutation(ctx context.Context, plan tfsdk.Plan) (string, map[string]string, diag.Diagnostics, bool, error) {
	var diagnostics diag.Diagnostics
	var rawBody types.String
	diagnostics.Append(plan.GetAttribute(ctx, path.Root("body"), &rawBody)...)
	if diagnostics.HasError() {
		return "", nil, diagnostics, false, nil
	}
	fields, body, err := r.parseBody(rawBody)
	if err != nil {
		return "", nil, diagnostics, false, err
	}
	if len(r.typedAttrs) > 0 {
		typed, typedDiagnostics := (&apiResource{typedAttrs: r.typedAttrs}).bodyFromTyped(ctx, plan)
		diagnostics.Append(typedDiagnostics...)
		if diagnostics.HasError() {
			return "", nil, diagnostics, false, nil
		}
		for key, value := range typed {
			fields[key] = value
		}
		if object, ok := body.(map[string]any); ok {
			for key, value := range typed {
				object[key] = value
			}
			body = object
		} else if len(typed) > 0 {
			body = typed
		}
	}
	body = r.applyFieldAliases(fields, body)
	requestPath, err := r.resolvePath(fields)
	if err != nil {
		return "", nil, diagnostics, false, err
	}
	requestPath, err = r.resolveQuery(requestPath, fields)
	if err != nil {
		return "", nil, diagnostics, false, err
	}
	requestBody := r.requestBody(body)
	if err := r.validateRequestBody(requestBody); err != nil {
		return "", nil, diagnostics, false, err
	}
	if len(r.bodyRequiredFields) > 0 {
		object, ok := requestBody.(map[string]any)
		if !ok {
			return "", nil, diagnostics, false, fmt.Errorf("body JSON doit être un objet contenant: %s", strings.Join(r.bodyRequiredFields, ", "))
		}
		if missing := requiredBodyFieldsMissing(object, r.bodyRequiredFields); len(missing) > 0 {
			return "", nil, diagnostics, false, fmt.Errorf("body JSON : champs requis absents: %s", strings.Join(missing, ", "))
		}
	}
	var out any
	if _, err := r.data.apiDo(ctx, r.httpMethod(), requestPath, requestBody, &out); err != nil {
		return "", nil, diagnostics, false, err
	}
	jsonState := marshalStateJSON(out, r.sensitiveAttrNames())
	secretValues := make(map[string]string, len(r.secretOutputs))
	object, _ := out.(map[string]any)
	for _, output := range r.secretOutputs {
		value, found := object[output]
		if !found || value == nil || strings.TrimSpace(fmt.Sprintf("%v", value)) == "" {
			return jsonState, secretValues, diagnostics, true, fmt.Errorf(
				"réponse %s %s sans sortie one-shot %q requise",
				r.httpMethod(),
				r.actionDocPath(),
				output,
			)
		}
		secretValues[output] = fmt.Sprintf("%v", value)
	}
	return jsonState, secretValues, diagnostics, true, nil
}

func (r *actionApiResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	jsonState, secretValues, diagnostics, mutationSucceeded, err := r.runMutation(ctx, req.Plan)
	resp.Diagnostics.Append(diagnostics...)
	if mutationSucceeded {
		// Une réponse HTTP 2xx prouve l'effet distant. Persister immédiatement un
		// state récupérable avant de valider les sorties one-shot empêche Terraform
		// de rejouer aveuglément le POST si la réponse est incomplète.
		resp.State.Raw = req.Plan.Raw
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(r.actionID()))...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("json"), types.StringValue(jsonState))...)
		for _, output := range r.secretOutputs {
			if value, found := secretValues[output]; found {
				resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(output), types.StringValue(value))...)
			} else {
				resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(output), types.StringNull())...)
			}
		}
	}
	if err != nil {
		resp.Diagnostics.AddError("Mutation API échouée", err.Error())
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *actionApiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	if r.readPath == "" {
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
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("json"), types.StringValue(marshalStateJSON(out, r.sensitiveAttrNames())))...)
}

func (r *actionApiResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Action Terraform immuable",
		"Une action ne peut pas être mise à jour en rejouant aveuglément sa mutation. Tous ses attributs configurables imposent un remplacement explicite.",
	)
}

func (r *actionApiResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var identifier types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &identifier)...)
	if resp.Diagnostics.HasError() || !r.isLegacyStateID(identifier) {
		return
	}
	resp.Diagnostics.AddError(
		"Migration explicite du state requise",
		r.legacyStateMessage(identifier.ValueString()),
	)
}
