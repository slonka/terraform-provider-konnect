// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GatewayPluginAiResponseTransformerDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginAiResponseTransformerDataSource{}

func NewGatewayPluginAiResponseTransformerDataSource() datasource.DataSource {
	return &GatewayPluginAiResponseTransformerDataSource{}
}

// GatewayPluginAiResponseTransformerDataSource is the data source implementation.
type GatewayPluginAiResponseTransformerDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginAiResponseTransformerDataSourceModel describes the data model.
type GatewayPluginAiResponseTransformerDataSourceModel struct {
	Config         tfTypes.AiResponseTransformerPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                      `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLConsumer                      `tfsdk:"consumer_group"`
	ControlPlaneID types.String                              `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                               `tfsdk:"created_at"`
	Enabled        types.Bool                                `tfsdk:"enabled"`
	ID             types.String                              `tfsdk:"id"`
	InstanceName   types.String                              `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering                `tfsdk:"ordering"`
	Protocols      []types.String                            `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                      `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                      `tfsdk:"service"`
	Tags           []types.String                            `tfsdk:"tags"`
	UpdatedAt      types.Int64                               `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginAiResponseTransformerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_ai_response_transformer"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginAiResponseTransformerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginAiResponseTransformer DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"http_proxy_host": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"http_proxy_port": schema.Int64Attribute{
						Computed:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
					},
					"http_timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Timeout in milliseconds for the AI upstream service.`,
					},
					"https_proxy_host": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"https_proxy_port": schema.Int64Attribute{
						Computed:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
					},
					"https_verify": schema.BoolAttribute{
						Computed:    true,
						Description: `Verify the TLS certificate of the AI upstream service.`,
					},
					"llm": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"allow_override": schema.BoolAttribute{
										Computed:    true,
										Description: `If enabled, the authorization header or parameter can be overridden in the request by the value configured in the plugin.`,
									},
									"aws_access_key_id": schema.StringAttribute{
										Computed:    true,
										Description: `Set this if you are using an AWS provider (Bedrock) and you are authenticating using static IAM User credentials. Setting this will override the AWS_ACCESS_KEY_ID environment variable for this plugin instance.`,
									},
									"aws_secret_access_key": schema.StringAttribute{
										Computed:    true,
										Description: `Set this if you are using an AWS provider (Bedrock) and you are authenticating using static IAM User credentials. Setting this will override the AWS_SECRET_ACCESS_KEY environment variable for this plugin instance.`,
									},
									"azure_client_id": schema.StringAttribute{
										Computed:    true,
										Description: `If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client ID.`,
									},
									"azure_client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client secret.`,
									},
									"azure_tenant_id": schema.StringAttribute{
										Computed:    true,
										Description: `If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the tenant ID.`,
									},
									"azure_use_managed_identity": schema.BoolAttribute{
										Computed:    true,
										Description: `Set true to use the Azure Cloud Managed Identity (or user-assigned identity) to authenticate with Azure-provider models.`,
									},
									"gcp_service_account_json": schema.StringAttribute{
										Computed:    true,
										Description: `Set this field to the full JSON of the GCP service account to authenticate, if required. If null (and gcp_use_service_account is true), Kong will attempt to read from environment variable ` + "`" + `GCP_SERVICE_ACCOUNT` + "`" + `.`,
									},
									"gcp_use_service_account": schema.BoolAttribute{
										Computed:    true,
										Description: `Use service account auth for GCP-based providers and models.`,
									},
									"header_name": schema.StringAttribute{
										Computed:    true,
										Description: `If AI model requires authentication via Authorization or API key header, specify its name here.`,
									},
									"header_value": schema.StringAttribute{
										Computed:    true,
										Description: `Specify the full auth header value for 'header_name', for example 'Bearer key' or just 'key'.`,
									},
									"param_location": schema.StringAttribute{
										Computed:    true,
										Description: `Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.`,
									},
									"param_name": schema.StringAttribute{
										Computed:    true,
										Description: `If AI model requires authentication via query parameter, specify its name here.`,
									},
									"param_value": schema.StringAttribute{
										Computed:    true,
										Description: `Specify the full parameter value for 'param_name'.`,
									},
								},
							},
							"logging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"log_payloads": schema.BoolAttribute{
										Computed:    true,
										Description: `If enabled, will log the request and response body into the Kong log plugin(s) output.`,
									},
									"log_statistics": schema.BoolAttribute{
										Computed:    true,
										Description: `If enabled and supported by the driver, will add model usage and token metrics into the Kong log plugin(s) output.`,
									},
								},
							},
							"model": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed:    true,
										Description: `Model name to execute.`,
									},
									"options": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"anthropic_version": schema.StringAttribute{
												Computed:    true,
												Description: `Defines the schema/API version, if using Anthropic provider.`,
											},
											"azure_api_version": schema.StringAttribute{
												Computed:    true,
												Description: `'api-version' for Azure OpenAI instances.`,
											},
											"azure_deployment_id": schema.StringAttribute{
												Computed:    true,
												Description: `Deployment ID for Azure OpenAI instances.`,
											},
											"azure_instance": schema.StringAttribute{
												Computed:    true,
												Description: `Instance name for Azure OpenAI hosted models.`,
											},
											"bedrock": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"aws_region": schema.StringAttribute{
														Computed:    true,
														Description: `If using AWS providers (Bedrock) you can override the ` + "`" + `AWS_REGION` + "`" + ` environment variable by setting this option.`,
													},
												},
											},
											"gemini": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"api_endpoint": schema.StringAttribute{
														Computed:    true,
														Description: `If running Gemini on Vertex, specify the regional API endpoint (hostname only).`,
													},
													"location_id": schema.StringAttribute{
														Computed:    true,
														Description: `If running Gemini on Vertex, specify the location ID.`,
													},
													"project_id": schema.StringAttribute{
														Computed:    true,
														Description: `If running Gemini on Vertex, specify the project ID.`,
													},
												},
											},
											"input_cost": schema.NumberAttribute{
												Computed:    true,
												Description: `Defines the cost per 1M tokens in your prompt.`,
											},
											"llama2_format": schema.StringAttribute{
												Computed:    true,
												Description: `If using llama2 provider, select the upstream message format.`,
											},
											"max_tokens": schema.Int64Attribute{
												Computed:    true,
												Description: `Defines the max_tokens, if using chat or completion models.`,
											},
											"mistral_format": schema.StringAttribute{
												Computed:    true,
												Description: `If using mistral provider, select the upstream message format.`,
											},
											"output_cost": schema.NumberAttribute{
												Computed:    true,
												Description: `Defines the cost per 1M tokens in the output of the AI.`,
											},
											"temperature": schema.NumberAttribute{
												Computed:    true,
												Description: `Defines the matching temperature, if using chat or completion models.`,
											},
											"top_k": schema.Int64Attribute{
												Computed:    true,
												Description: `Defines the top-k most likely tokens, if supported.`,
											},
											"top_p": schema.NumberAttribute{
												Computed:    true,
												Description: `Defines the top-p probability mass, if supported.`,
											},
											"upstream_path": schema.StringAttribute{
												Computed:    true,
												Description: `Manually specify or override the AI operation path, used when e.g. using the 'preserve' route_type.`,
											},
											"upstream_url": schema.StringAttribute{
												Computed:    true,
												Description: `Manually specify or override the full URL to the AI operation endpoints, when calling (self-)hosted models, or for running via a private endpoint.`,
											},
										},
										Description: `Key/value settings for the model`,
									},
									"provider": schema.StringAttribute{
										Computed:    true,
										Description: `AI provider request format - Kong translates requests to and from the specified backend compatible formats.`,
									},
								},
							},
							"route_type": schema.StringAttribute{
								Computed:    true,
								Description: `The model's operation implementation, for this provider. Set to ` + "`" + `preserve` + "`" + ` to pass through without transformation.`,
							},
						},
					},
					"max_request_body_size": schema.Int64Attribute{
						Computed:    true,
						Description: `max allowed body size allowed to be introspected`,
					},
					"parse_llm_response_json_instructions": schema.BoolAttribute{
						Computed:    true,
						Description: `Set true to read specific response format from the LLM, and accordingly set the status code / body / headers that proxy back to the client. You need to engineer your LLM prompt to return the correct format, see plugin docs 'Overview' page for usage instructions.`,
					},
					"prompt": schema.StringAttribute{
						Computed:    true,
						Description: `Use this prompt to tune the LLM system/assistant message for the returning proxy response (from the upstream), adn what response format you are expecting.`,
					},
					"transformation_extract_pattern": schema.StringAttribute{
						Computed:    true,
						Description: `Defines the regular expression that must match to indicate a successful AI transformation at the response phase. The first match will be set as the returning body. If the AI service's response doesn't match this pattern, a failure is returned to the client.`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
			},
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
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

func (r *GatewayPluginAiResponseTransformerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginAiResponseTransformerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginAiResponseTransformerDataSourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetAiresponsetransformerPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetAiresponsetransformerPlugin(ctx, request)
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
	if !(res.AiResponseTransformerPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAiResponseTransformerPlugin(res.AiResponseTransformerPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
