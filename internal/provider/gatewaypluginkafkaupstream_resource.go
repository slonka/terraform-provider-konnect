// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	speakeasy_int64validators "github.com/kong/terraform-provider-konnect/v2/internal/validators/int64validators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginKafkaUpstreamResource{}
var _ resource.ResourceWithImportState = &GatewayPluginKafkaUpstreamResource{}

func NewGatewayPluginKafkaUpstreamResource() resource.Resource {
	return &GatewayPluginKafkaUpstreamResource{}
}

// GatewayPluginKafkaUpstreamResource defines the resource implementation.
type GatewayPluginKafkaUpstreamResource struct {
	client *sdk.Konnect
}

// GatewayPluginKafkaUpstreamResourceModel describes the resource data model.
type GatewayPluginKafkaUpstreamResourceModel struct {
	Config         tfTypes.KafkaUpstreamPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer              `tfsdk:"consumer" tfPlanOnly:"true"`
	ConsumerGroup  *tfTypes.ACLConsumer              `tfsdk:"consumer_group" tfPlanOnly:"true"`
	ControlPlaneID types.String                      `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                       `tfsdk:"created_at"`
	Enabled        types.Bool                        `tfsdk:"enabled"`
	ID             types.String                      `tfsdk:"id"`
	InstanceName   types.String                      `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering        `tfsdk:"ordering"`
	Protocols      []types.String                    `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer              `tfsdk:"route" tfPlanOnly:"true"`
	Service        *tfTypes.ACLConsumer              `tfsdk:"service" tfPlanOnly:"true"`
	Tags           []types.String                    `tfsdk:"tags"`
	UpdatedAt      types.Int64                       `tfsdk:"updated_at"`
}

func (r *GatewayPluginKafkaUpstreamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_kafka_upstream"
}

func (r *GatewayPluginKafkaUpstreamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginKafkaUpstream Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"authentication": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"mechanism": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The SASL authentication mechanism.  Supported options: ` + "`" + `PLAIN` + "`" + `, ` + "`" + `SCRAM-SHA-256` + "`" + `, or ` + "`" + `SCRAM-SHA-512` + "`" + `. must be one of ["PLAIN", "SCRAM-SHA-256", "SCRAM-SHA-512"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"PLAIN",
										"SCRAM-SHA-256",
										"SCRAM-SHA-512",
									),
								},
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Password for SASL authentication.`,
							},
							"strategy": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The authentication strategy for the plugin, the only option for the value is ` + "`" + `sasl` + "`" + `. must be "sasl"`,
								Validators: []validator.String{
									stringvalidator.OneOf("sasl"),
								},
							},
							"tokenauth": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Enable this to indicate ` + "`" + `DelegationToken` + "`" + ` authentication.`,
							},
							"user": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Username for SASL authentication.`,
							},
						},
					},
					"bootstrap_servers": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `A string representing a host name, such as example.com. Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
									},
								},
								"port": schema.Int64Attribute{
									Computed:    true,
									Optional:    true,
									Description: `An integer representing a port number between 0 and 65535, inclusive. Not Null`,
									Validators: []validator.Int64{
										speakeasy_int64validators.NotNull(),
										int64validator.AtMost(65535),
									},
								},
							},
						},
						Description: `Set of bootstrap brokers in a ` + "`" + `{host: host, port: port}` + "`" + ` list format.`,
					},
					"cluster_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `An identifier for the Kafka cluster. By default, this field generates a random string. You can also set your own custom cluster identifier.  If more than one Kafka plugin is configured without a ` + "`" + `cluster_name` + "`" + ` (that is, if the default autogenerated value is removed), these plugins will use the same producer, and by extension, the same cluster. Logs will be sent to the leader of the cluster.`,
					},
					"forward_body": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Include the request body in the message. At least one of these must be true: ` + "`" + `forward_method` + "`" + `, ` + "`" + `forward_uri` + "`" + `, ` + "`" + `forward_headers` + "`" + `, ` + "`" + `forward_body` + "`" + `.`,
					},
					"forward_headers": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Include the request headers in the message. At least one of these must be true: ` + "`" + `forward_method` + "`" + `, ` + "`" + `forward_uri` + "`" + `, ` + "`" + `forward_headers` + "`" + `, ` + "`" + `forward_body` + "`" + `.`,
					},
					"forward_method": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Include the request method in the message. At least one of these must be true: ` + "`" + `forward_method` + "`" + `, ` + "`" + `forward_uri` + "`" + `, ` + "`" + `forward_headers` + "`" + `, ` + "`" + `forward_body` + "`" + `.`,
					},
					"forward_uri": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Include the request URI and URI arguments (as in, query arguments) in the message. At least one of these must be true: ` + "`" + `forward_method` + "`" + `, ` + "`" + `forward_uri` + "`" + `, ` + "`" + `forward_headers` + "`" + `, ` + "`" + `forward_body` + "`" + `.`,
					},
					"keepalive": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Keepalive timeout in milliseconds.`,
					},
					"keepalive_enabled": schema.BoolAttribute{
						Computed: true,
						Optional: true,
					},
					"producer_async": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Flag to enable asynchronous mode.`,
					},
					"producer_async_buffering_limits_messages_in_memory": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum number of messages that can be buffered in memory in asynchronous mode.`,
					},
					"producer_async_flush_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum time interval in milliseconds between buffer flushes in asynchronous mode.`,
					},
					"producer_request_acks": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments; 1 for only the leader; and -1 for the full ISR (In-Sync Replica set). must be one of ["-1", "0", "1"]`,
						Validators: []validator.Int64{
							int64validator.OneOf(-1, 0, 1),
						},
					},
					"producer_request_limits_bytes_per_request": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of a Produce request in bytes.`,
					},
					"producer_request_limits_messages_per_request": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum number of messages to include into a single producer request.`,
					},
					"producer_request_retries_backoff_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Backoff interval between retry attempts in milliseconds.`,
					},
					"producer_request_retries_max_attempts": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum number of retry attempts per single Produce request.`,
					},
					"producer_request_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Time to wait for a Produce response in milliseconds.`,
					},
					"security": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"certificate_id": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `UUID of certificate entity for mTLS authentication.`,
							},
							"ssl": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Enables TLS.`,
							},
						},
					},
					"timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Socket timeout in milliseconds.`,
					},
					"topic": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The Kafka topic to publish to.`,
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

func (r *GatewayPluginKafkaUpstreamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginKafkaUpstreamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginKafkaUpstreamResourceModel
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

	kafkaUpstreamPlugin := *data.ToSharedKafkaUpstreamPluginInput()
	request := operations.CreateKafkaupstreamPluginRequest{
		ControlPlaneID:      controlPlaneID,
		KafkaUpstreamPlugin: kafkaUpstreamPlugin,
	}
	res, err := r.client.Plugins.CreateKafkaupstreamPlugin(ctx, request)
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
	if !(res.KafkaUpstreamPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedKafkaUpstreamPlugin(res.KafkaUpstreamPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginKafkaUpstreamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginKafkaUpstreamResourceModel
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

	request := operations.GetKafkaupstreamPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetKafkaupstreamPlugin(ctx, request)
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
	if !(res.KafkaUpstreamPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedKafkaUpstreamPlugin(res.KafkaUpstreamPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginKafkaUpstreamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginKafkaUpstreamResourceModel
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

	kafkaUpstreamPlugin := *data.ToSharedKafkaUpstreamPluginInput()
	request := operations.UpdateKafkaupstreamPluginRequest{
		PluginID:            pluginID,
		ControlPlaneID:      controlPlaneID,
		KafkaUpstreamPlugin: kafkaUpstreamPlugin,
	}
	res, err := r.client.Plugins.UpdateKafkaupstreamPlugin(ctx, request)
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
	if !(res.KafkaUpstreamPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedKafkaUpstreamPlugin(res.KafkaUpstreamPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginKafkaUpstreamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginKafkaUpstreamResourceModel
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

	request := operations.DeleteKafkaupstreamPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteKafkaupstreamPlugin(ctx, request)
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

func (r *GatewayPluginKafkaUpstreamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
