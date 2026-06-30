package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource              = &organizationDataSource{}
	_ datasource.DataSourceWithConfigure = &organizationDataSource{}
)

func NewOrganizationDataSource() datasource.DataSource { return &organizationDataSource{} }

type organizationDataSource struct {
	data *providerData
}

func (d *organizationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization"
}

func (d *organizationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Lit une organisation AISIA existante par son `id`.",
		Attributes: map[string]schema.Attribute{
			"id":            schema.StringAttribute{MarkdownDescription: "ID de l'organisation.", Required: true},
			"name":          schema.StringAttribute{MarkdownDescription: "Nom.", Computed: true},
			"slug":          schema.StringAttribute{MarkdownDescription: "Slug.", Computed: true},
			"contract_type": schema.StringAttribute{MarkdownDescription: "Type de contrat.", Computed: true},
		},
	}
}

func (d *organizationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *organizationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var cfg organizationModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var out orgAPI
	if _, err := d.data.apiDo(ctx, "GET", "/admin/organizations/"+cfg.ID.ValueString(), nil, &out); err != nil {
		resp.Diagnostics.AddError("Lecture organisation échouée", err.Error())
		return
	}
	state := organizationModel{ID: cfg.ID}
	applyOrg(&state, out)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
