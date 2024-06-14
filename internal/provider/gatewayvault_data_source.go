// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GatewayVaultDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayVaultDataSource{}

func NewGatewayVaultDataSource() datasource.DataSource {
	return &GatewayVaultDataSource{}
}

// GatewayVaultDataSource is the data source implementation.
type GatewayVaultDataSource struct {
	client *sdk.Konnect
}

// GatewayVaultDataSourceModel describes the data model.
type GatewayVaultDataSourceModel struct {
	Config         *tfTypes.VaultConfig `tfsdk:"config"`
	ControlPlaneID types.String         `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64          `tfsdk:"created_at"`
	Description    types.String         `tfsdk:"description"`
	ID             types.String         `tfsdk:"id"`
	Name           types.String         `tfsdk:"name"`
	Prefix         types.String         `tfsdk:"prefix"`
	Tags           []types.String       `tfsdk:"tags"`
	UpdatedAt      types.Int64          `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayVaultDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_vault"
}

// Schema defines the schema for the data source.
func (r *GatewayVaultDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayVault DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed:    true,
				Attributes:  map[string]schema.Attribute{},
				Description: `The configuration properties for the Vault which can be found on the vaults' documentation page.`,
			},
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `The description of the Vault entity.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the Vault to lookup`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.`,
			},
			"prefix": schema.StringAttribute{
				Computed:    true,
				Description: `The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Vault for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayVaultDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GatewayVaultDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayVaultDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	vaultID := data.ID.ValueString()
	request := operations.GetVaultRequest{
		ControlPlaneID: controlPlaneID,
		VaultID:        vaultID,
	}
	res, err := r.client.Vaults.GetVault(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Vault != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedVault(res.Vault)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
