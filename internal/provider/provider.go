package provider

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// aisiaProvider implémente provider.Provider.
type aisiaProvider struct {
	version string
}

// providerData est injecté dans chaque resource/data source (Configure).
type providerData struct {
	endpoint string
	token    string
	http     *http.Client
}

type aisiaProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Token    types.String `tfsdk:"token"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &aisiaProvider{version: version}
	}
}

func (p *aisiaProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "aisia"
	resp.Version = p.version
}

func (p *aisiaProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provider AISIA — gère les ressources de la plateforme AISIA (organisations, clés providers par org) via l'API AISIA.",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "URL de base de l'API AISIA (défaut `https://api.aisia.fr`, ou env `AISIA_ENDPOINT`).",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "Jeton admin AISIA (Bearer). Préférez l'env `AISIA_TOKEN`. Sensible.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *aisiaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var cfg aisiaProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := os.Getenv("AISIA_ENDPOINT")
	if !cfg.Endpoint.IsNull() && cfg.Endpoint.ValueString() != "" {
		endpoint = cfg.Endpoint.ValueString()
	}
	if endpoint == "" {
		endpoint = "https://api.aisia.fr"
	}

	token := os.Getenv("AISIA_TOKEN")
	if !cfg.Token.IsNull() && cfg.Token.ValueString() != "" {
		token = cfg.Token.ValueString()
	}
	if token == "" {
		resp.Diagnostics.AddError(
			"Jeton AISIA manquant",
			"Renseignez l'attribut `token` ou la variable d'environnement `AISIA_TOKEN`.",
		)
		return
	}

	data := &providerData{
		endpoint: endpoint,
		token:    token,
		http:     &http.Client{Timeout: 30 * time.Second},
	}
	resp.ResourceData = data
	resp.DataSourceData = data
}

func (p *aisiaProvider) Resources(_ context.Context) []func() resource.Resource {
	hand := []func() resource.Resource{
		NewOrganizationResource,
		NewProviderKeyResource,
		NewUserResource,
		NewApiKeyResource,
		NewProviderResource,
		NewLocalModelResource,
		NewWebhookResource,
		NewResellerResource,
		NewGuardrailResource,
	}
	// generatedResources : resources CRUD générées depuis l'OpenAPI (resources_generated.go).
	return append(hand, generatedResources...)
}

func (p *aisiaProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	hand := []func() datasource.DataSource{
		NewOrganizationDataSource,
		NewProvidersDataSource,
		NewLocalModelsDataSource,
		NewPlansDataSource,
		NewAgentsDataSource,
	}
	// generatedDataSources : data sources GET générées depuis l'OpenAPI (catalog_generated.go).
	return append(hand, generatedDataSources...)
}
