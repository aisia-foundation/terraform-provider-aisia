package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// singletonApiResource — configuration singleton (GET + PUT/PATCH sur le même chemin, sans /{id}).
type singletonApiResource struct {
	data              *providerData
	name              string
	path              string
	updateVerb        string
	desc              string
	bodyFields        []string
	writeOnlyFields   []string
	deepSubsetFields  []string
	readFieldAliases  map[string]string
	responseObjectKey string
	sensitiveFields   []string
}

var (
	_ resource.Resource                = &singletonApiResource{}
	_ resource.ResourceWithConfigure   = &singletonApiResource{}
	_ resource.ResourceWithImportState = &singletonApiResource{}
)

type singletonApiResourceModel struct {
	ID   types.String `tfsdk:"id"`
	Body types.String `tfsdk:"body"`
	JSON types.String `tfsdk:"json"`
}

func (r *singletonApiResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.name
}

func (r *singletonApiResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: r.desc + docLinksForEndpoint(r.path),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{MarkdownDescription: "Identifiant fixe (singleton).", Computed: true},
			"body": schema.StringAttribute{
				MarkdownDescription: "Corps JSON (payload PUT/PATCH).",
				Required:            true,
				Sensitive:           len(r.sensitiveFields) > 0,
			},
			"json": schema.StringAttribute{
				MarkdownDescription: "Dernière réponse de l'API (JSON, champs sensibles expurgés).",
				Computed:            true,
				Sensitive:           len(r.sensitiveFields) > 0,
			},
		},
	}
}

func (r *singletonApiResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *singletonApiResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != r.singletonID() && req.ID != r.name && req.ID != r.path {
		resp.Diagnostics.AddError(
			"Identifiant singleton invalide",
			fmt.Sprintf("Utilisez %q, %q ou %q pour importer ce singleton.", r.singletonID(), r.name, r.path),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(r.singletonID()))...)
}

func (r *singletonApiResource) singletonID() string { return "_singleton" }

func (r *singletonApiResource) decodeBody(value types.String) (map[string]any, error) {
	if value.IsNull() || value.IsUnknown() {
		return nil, nil
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(value.ValueString()), &body); err != nil {
		return nil, fmt.Errorf("body doit être un objet JSON valide: %w", err)
	}
	allowed := stringSet(r.bodyFields)
	unknown := make([]string, 0)
	for field := range body {
		if _, ok := allowed[field]; !ok {
			unknown = append(unknown, field)
		}
	}
	if len(unknown) > 0 {
		sort.Strings(unknown)
		return nil, fmt.Errorf("champs hors contrat OpenAPI pour %s: %s", r.path, strings.Join(unknown, ", "))
	}
	return body, nil
}

func (r *singletonApiResource) remoteWritableFields(out any) map[string]any {
	object, ok := out.(map[string]any)
	if !ok {
		return map[string]any{}
	}
	if r.responseObjectKey != "" {
		nested, nestedOK := object[r.responseObjectKey].(map[string]any)
		if !nestedOK {
			return map[string]any{}
		}
		object = nested
	}
	sensitive := stringSet(r.sensitiveFields)
	writeOnly := stringSet(r.writeOnlyFields)
	remote := make(map[string]any)
	for _, field := range r.bodyFields {
		if _, secret := sensitive[field]; secret {
			continue
		}
		if _, opaque := writeOnly[field]; opaque {
			continue
		}
		readName := field
		if alias := r.readFieldAliases[field]; alias != "" {
			readName = alias
		}
		if value, found := object[readName]; found {
			remote[field] = value
		}
	}
	return remote
}

func (r *singletonApiResource) reconcileBody(current types.String, out any) (types.String, error) {
	remote := r.remoteWritableFields(out)
	if current.IsNull() || current.IsUnknown() {
		raw, err := json.Marshal(remote)
		if err != nil {
			return types.StringNull(), err
		}
		return types.StringValue(string(raw)), nil
	}
	desired, err := r.decodeBody(current)
	if err != nil {
		return types.StringNull(), err
	}
	changed := false
	deepSubset := stringSet(r.deepSubsetFields)
	for field, value := range remote {
		if configured, found := desired[field]; found {
			if _, deep := deepSubset[field]; deep {
				reconciled, nestedChanged := reconcileSubsetValue(configured, value)
				if nestedChanged {
					desired[field] = reconciled
					changed = true
				}
			} else if !reflect.DeepEqual(configured, value) {
				desired[field] = value
				changed = true
			}
		}
	}
	if !changed {
		// Conserver le texte exact configuré évite une différence artificielle
		// due uniquement à l'ordre/indentation du JSON.
		return current, nil
	}
	raw, err := json.Marshal(desired)
	if err != nil {
		return types.StringNull(), err
	}
	return types.StringValue(string(raw)), nil
}

func reconcileSubsetValue(configured any, observed any) (any, bool) {
	desiredObject, desiredOK := configured.(map[string]any)
	remoteObject, remoteOK := observed.(map[string]any)
	if !desiredOK || !remoteOK {
		return observed, !reflect.DeepEqual(configured, observed)
	}
	result := make(map[string]any, len(desiredObject))
	changed := false
	for key, desired := range desiredObject {
		result[key] = desired
		if remote, found := remoteObject[key]; found {
			reconciled, nestedChanged := reconcileSubsetValue(desired, remote)
			result[key] = reconciled
			changed = changed || nestedChanged
		}
	}
	return result, changed
}

func configuredSubsetMatches(configured any, observed any) bool {
	desiredObject, desiredOK := configured.(map[string]any)
	remoteObject, remoteOK := observed.(map[string]any)
	if !desiredOK || !remoteOK {
		return reflect.DeepEqual(configured, observed)
	}
	for key, desired := range desiredObject {
		remote, found := remoteObject[key]
		if !found || !configuredSubsetMatches(desired, remote) {
			return false
		}
	}
	return true
}

func (r *singletonApiResource) mutationMismatches(body map[string]any, out any) []string {
	remote := r.remoteWritableFields(out)
	deepSubset := stringSet(r.deepSubsetFields)
	mismatches := make([]string, 0)
	for field, desired := range body {
		observed, readable := remote[field]
		if readable {
			_, deep := deepSubset[field]
			if (deep && !configuredSubsetMatches(desired, observed)) || (!deep && !reflect.DeepEqual(desired, observed)) {
				mismatches = append(mismatches, field)
			}
		}
	}
	sort.Strings(mismatches)
	return mismatches
}

func (r *singletonApiResource) readRemote(ctx context.Context) (any, int, error) {
	var out any
	code, err := r.data.apiDo(ctx, "GET", r.path, nil, &out)
	return out, code, err
}

func (r *singletonApiResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan singletonApiResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	body, err := r.decodeBody(plan.Body)
	if err != nil {
		resp.Diagnostics.AddError("body JSON invalide", err.Error())
		return
	}
	verb := r.updateVerb
	if verb == "" {
		verb = "PUT"
	}
	var out any
	if _, err := r.data.apiDo(ctx, verb, r.path, body, &out); err != nil {
		resp.Diagnostics.AddError("Création (upsert) échouée", err.Error())
		return
	}
	plan.ID = types.StringValue(r.singletonID())
	plan.JSON = types.StringValue(marshalStateJSON(out, r.sensitiveFields))
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	observed, _, err := r.readRemote(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Relecture singleton après mutation échouée", err.Error())
		return
	}
	if mismatches := r.mutationMismatches(body, observed); len(mismatches) > 0 {
		resp.Diagnostics.AddError(
			"Mutation singleton non confirmée",
			"Les champs publics relus divergent de la valeur demandée : "+strings.Join(mismatches, ", "),
		)
		return
	}
	plan.JSON = types.StringValue(marshalStateJSON(observed, r.sensitiveFields))
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *singletonApiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state singletonApiResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	out, code, err := r.readRemote(ctx)
	if code == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Lecture échouée", err.Error())
		return
	}
	state.ID = types.StringValue(r.singletonID())
	state.JSON = types.StringValue(marshalStateJSON(out, r.sensitiveFields))
	reconciled, reconcileErr := r.reconcileBody(state.Body, out)
	if reconcileErr != nil {
		resp.Diagnostics.AddError("Réconciliation singleton échouée", reconcileErr.Error())
		return
	}
	state.Body = reconciled
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *singletonApiResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state singletonApiResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	verb := r.updateVerb
	if verb == "" {
		verb = "PUT"
	}
	body, err := r.decodeBody(plan.Body)
	if err != nil {
		resp.Diagnostics.AddError("body JSON invalide", err.Error())
		return
	}
	var out any
	if _, err := r.data.apiDo(ctx, verb, r.path, body, &out); err != nil {
		resp.Diagnostics.AddError("Mise à jour échouée", err.Error())
		return
	}
	plan.ID = state.ID
	plan.JSON = types.StringValue(marshalStateJSON(out, r.sensitiveFields))
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	observed, _, err := r.readRemote(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Relecture singleton après mutation échouée", err.Error())
		return
	}
	if mismatches := r.mutationMismatches(body, observed); len(mismatches) > 0 {
		resp.Diagnostics.AddError(
			"Mutation singleton non confirmée",
			"Les champs publics relus divergent de la valeur demandée : "+strings.Join(mismatches, ", "),
		)
		return
	}
	plan.JSON = types.StringValue(marshalStateJSON(observed, r.sensitiveFields))
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *singletonApiResource) Delete(_ context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError(
		"Suppression distante non représentable",
		fmt.Sprintf(
			"AISIA refuse de retirer aisia_%s du state en laissant silencieusement la configuration distante %s active : l'OpenAPI ne fournit aucun DELETE ni reset exact. Pour arrêter volontairement sa gestion sans mutation distante, utilisez un bloc removed avec destroy=false ou terraform state rm après revue explicite.",
			r.name,
			r.path,
		),
	)
}
