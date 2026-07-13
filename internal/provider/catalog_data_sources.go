package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// catalogDataSource — data source générique LECTURE SEULE pour les catalogues AISIA
// servis par GET (providers LLM, modèles locaux, plans). Robuste au schéma : expose la
// réponse brute (`json`), le nombre d'items (`count`) et la liste d'identifiants (`ids`),
// quel que soit le format exact. À consommer via `jsondecode(data.aisia_X.json)`.
type catalogDataSource struct {
	data    *providerData
	name    string // suffixe TypeName, ex. "providers"
	path    string // endpoint GET, ex. "/v1/providers"
	listKey string // clé d'enveloppe de la liste, ex. "providers" (sinon liste nue)
	desc    string
}

func NewProvidersDataSource() datasource.DataSource {
	return &catalogDataSource{name: "providers", path: "/v1/providers", listKey: "providers",
		desc: "Catalogue des providers LLM enregistrés (providers.yaml)."}
}
func NewLocalModelsDataSource() datasource.DataSource {
	return &catalogDataSource{name: "local_models", path: "/admin/local-models", listKey: "local_models",
		desc: "Catalogue des modèles locaux (local_models.yaml)."}
}
func NewPlansDataSource() datasource.DataSource {
	return &catalogDataSource{name: "plans", path: "/admin/plans", listKey: "plans",
		desc: "Plans d'abonnement (Free/SaaS/BaaS/PaaS) et leurs quotas."}
}
func NewAgentsDataSource() datasource.DataSource {
	return &catalogDataSource{name: "agents", path: "/admin/agents", listKey: "agents",
		desc: "Catalogue des agents (identity.yaml). Lecture seule : les agents ne sont pas du config CRUD (définition versionnée + runtime orchestré)."}
}

var (
	_ datasource.DataSource              = &catalogDataSource{}
	_ datasource.DataSourceWithConfigure = &catalogDataSource{}
)

type catalogModel struct {
	JSON  types.String `tfsdk:"json"`
	Count types.Int64  `tfsdk:"item_count"`
	IDs   types.List   `tfsdk:"ids"`
}

func (d *catalogDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.name
}

func (d *catalogDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: d.desc + " Lecture seule (GET). Utilisez `jsondecode(...json)` pour exploiter les détails." + docLinksForEndpoint(d.path),
		Attributes: map[string]schema.Attribute{
			"json":       schema.StringAttribute{MarkdownDescription: "Réponse brute de l'API (JSON).", Computed: true},
			"item_count": schema.Int64Attribute{MarkdownDescription: "Nombre d'items.", Computed: true},
			"ids":        schema.ListAttribute{MarkdownDescription: "Identifiants des items (id/slug/name/model_id).", ElementType: types.StringType, Computed: true},
		},
	}
}

func (d *catalogDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	pd, ok := req.ProviderData.(*providerData)
	if !ok {
		resp.Diagnostics.AddError("Type de ProviderData inattendu", fmt.Sprintf("%T", req.ProviderData))
		return
	}
	d.data = pd
}

func (d *catalogDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	var out any
	if _, err := d.data.apiDo(ctx, "GET", d.path, nil, &out); err != nil {
		resp.Diagnostics.AddError("Lecture catalogue échouée", err.Error())
		return
	}
	raw, _ := json.Marshal(out)
	items := extractList(out, d.listKey)
	ids := make([]string, 0, len(items))
	for _, it := range items {
		if m, ok := it.(map[string]any); ok {
			ids = append(ids, asString(m, "id", "code", "slug", "name", "model_id", "provider"))
		}
	}
	idList, diags := types.ListValueFrom(ctx, types.StringType, ids)
	resp.Diagnostics.Append(diags...)
	state := catalogModel{
		JSON:  types.StringValue(string(raw)),
		Count: types.Int64Value(int64(len(items))),
		IDs:   idList,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

// extractList retourne la liste d'items, qu'elle soit nue ([...]) ou enveloppée
// ({"<listKey>": [...]} / {"data": [...]} / {"items": [...]}).
func extractList(out any, listKey string) []any {
	switch v := out.(type) {
	case []any:
		return v
	case map[string]any:
		for _, k := range []string{listKey, "data", "items", "results"} {
			if l, ok := v[k].([]any); ok {
				return l
			}
		}
	}
	return nil
}
