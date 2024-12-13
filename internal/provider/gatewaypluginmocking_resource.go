// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginMockingResource{}
var _ resource.ResourceWithImportState = &GatewayPluginMockingResource{}

func NewGatewayPluginMockingResource() resource.Resource {
	return &GatewayPluginMockingResource{}
}

// GatewayPluginMockingResource defines the resource implementation.
type GatewayPluginMockingResource struct {
	client *sdk.Konnect
}

// GatewayPluginMockingResourceModel describes the resource data model.
type GatewayPluginMockingResourceModel struct {
	Config         tfTypes.MockingPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer        `tfsdk:"consumer" tfPlanOnly:"true"`
	ConsumerGroup  *tfTypes.ACLConsumer        `tfsdk:"consumer_group" tfPlanOnly:"true"`
	ControlPlaneID types.String                `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                 `tfsdk:"created_at"`
	Enabled        types.Bool                  `tfsdk:"enabled"`
	ID             types.String                `tfsdk:"id"`
	InstanceName   types.String                `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering  `tfsdk:"ordering"`
	Protocols      []types.String              `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer        `tfsdk:"route" tfPlanOnly:"true"`
	Service        *tfTypes.ACLConsumer        `tfsdk:"service" tfPlanOnly:"true"`
	Tags           []types.String              `tfsdk:"tags"`
	UpdatedAt      types.Int64                 `tfsdk:"updated_at"`
}

func (r *GatewayPluginMockingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_mocking"
}

func (r *GatewayPluginMockingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginMocking Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"api_specification": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The contents of the specification file. You must use this option for hybrid or DB-less mode. You can include the full specification as part of the configuration. In Kong Manager, you can copy and paste the contents of the spec directly into the ` + "`" + `Config.Api Specification` + "`" + ` text field.`,
					},
					"api_specification_filename": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The path and name of the specification file loaded into Kong Gateway's database. You cannot use this option for DB-less or hybrid mode.`,
					},
					"custom_base_path": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The base path to be used for path match evaluation. This value is ignored if ` + "`" + `include_base_path` + "`" + ` is set to ` + "`" + `false` + "`" + `.`,
					},
					"include_base_path": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Indicates whether to include the base path when performing path match evaluation.`,
					},
					"included_status_codes": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.Int64Type,
						Description: `A global list of the HTTP status codes that can only be selected and returned.`,
					},
					"max_delay_time": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The maximum value in seconds of delay time. Set this value when ` + "`" + `random_delay` + "`" + ` is enabled and you want to adjust the default. The value must be greater than the ` + "`" + `min_delay_time` + "`" + `.`,
					},
					"min_delay_time": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The minimum value in seconds of delay time. Set this value when ` + "`" + `random_delay` + "`" + ` is enabled and you want to adjust the default. The value must be less than the ` + "`" + `max_delay_time` + "`" + `.`,
					},
					"random_delay": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Enables a random delay in the mocked response. Introduces delays to simulate real-time response times by APIs.`,
					},
					"random_examples": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Randomly selects one example and returns it. This parameter requires the spec to have multiple examples configured.`,
					},
					"random_status_code": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Determines whether to randomly select an HTTP status code from the responses of the corresponding API method. The default value is ` + "`" + `false` + "`" + `, which means the minimum HTTP status code is always selected and returned.`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{
					"id": types.StringType,
				})),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{
					"id": types.StringType,
				})),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"control_plane_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{
					"id": types.StringType,
				})),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{
					"id": types.StringType,
				})),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayPluginMockingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GatewayPluginMockingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginMockingResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	mockingPlugin := *data.ToSharedMockingPluginInput()
	request := operations.CreateMockingPluginRequest{
		ControlPlaneID: controlPlaneID,
		MockingPlugin:  mockingPlugin,
	}
	res, err := r.client.Plugins.CreateMockingPlugin(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.MockingPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMockingPlugin(res.MockingPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginMockingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginMockingResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetMockingPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetMockingPlugin(ctx, request)
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
	if !(res.MockingPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMockingPlugin(res.MockingPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginMockingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginMockingResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	mockingPlugin := *data.ToSharedMockingPluginInput()
	request := operations.UpdateMockingPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
		MockingPlugin:  mockingPlugin,
	}
	res, err := r.client.Plugins.UpdateMockingPlugin(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.MockingPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMockingPlugin(res.MockingPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginMockingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginMockingResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.DeleteMockingPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteMockingPlugin(ctx, request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *GatewayPluginMockingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		ControlPlaneID string `json:"control_plane_id"`
		ID             string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "control_plane_id": "9524ec7d-36d9-465d-a8c5-83a3c9390458",  "plugin_id": "3473c251-5b6c-4f45-b1ff-7ede735a366d"}': `+err.Error())
		return
	}

	if len(data.ControlPlaneID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field control_plane_id is required but was not found in the json encoded ID. It's expected to be a value alike '"9524ec7d-36d9-465d-a8c5-83a3c9390458"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("control_plane_id"), data.ControlPlaneID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"3473c251-5b6c-4f45-b1ff-7ede735a366d"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
