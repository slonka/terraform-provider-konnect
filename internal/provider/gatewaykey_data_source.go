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
var _ datasource.DataSource = &GatewayKeyDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayKeyDataSource{}

func NewGatewayKeyDataSource() datasource.DataSource {
	return &GatewayKeyDataSource{}
}

// GatewayKeyDataSource is the data source implementation.
type GatewayKeyDataSource struct {
	client *sdk.Konnect
}

// GatewayKeyDataSourceModel describes the data model.
type GatewayKeyDataSourceModel struct {
	ControlPlaneID types.String         `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64          `tfsdk:"created_at"`
	ID             types.String         `tfsdk:"id"`
	Jwk            types.String         `tfsdk:"jwk"`
	Kid            types.String         `tfsdk:"kid"`
	Name           types.String         `tfsdk:"name"`
	Pem            *tfTypes.Pem         `tfsdk:"pem"`
	Set            *tfTypes.ACLConsumer `tfsdk:"set"`
	Tags           []types.String       `tfsdk:"tags"`
	UpdatedAt      types.Int64          `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayKeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_key"
}

// Schema defines the schema for the data source.
func (r *GatewayKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayKey DataSource",

		Attributes: map[string]schema.Attribute{
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the Key to lookup`,
			},
			"jwk": schema.StringAttribute{
				Computed:    true,
				Description: `A JSON Web Key represented as a string.`,
			},
			"kid": schema.StringAttribute{
				Computed:    true,
				Description: `A unique identifier for a key.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name to associate with the given keys.`,
			},
			"pem": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"private_key": schema.StringAttribute{
						Computed: true,
					},
					"public_key": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `A keypair in PEM format.`,
			},
			"set": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `The id (an UUID) of the key-set with which to associate the key.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Key for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayKeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayKeyDataSourceModel
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
	keyID := data.ID.ValueString()
	request := operations.GetKeyRequest{
		ControlPlaneID: controlPlaneID,
		KeyID:          keyID,
	}
	res, err := r.client.Keys.GetKey(ctx, request)
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
	if !(res.Key != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedKey(res.Key)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
