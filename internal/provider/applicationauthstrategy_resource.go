// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_boolplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/boolplanmodifier"
	speakeasy_listplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/listplanmodifier"
	speakeasy_objectplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
	speakeasy_listvalidators "github.com/kong/terraform-provider-konnect/internal/validators/listvalidators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ApplicationAuthStrategyResource{}
var _ resource.ResourceWithImportState = &ApplicationAuthStrategyResource{}

func NewApplicationAuthStrategyResource() resource.Resource {
	return &ApplicationAuthStrategyResource{}
}

// ApplicationAuthStrategyResource defines the resource implementation.
type ApplicationAuthStrategyResource struct {
	client *sdk.Konnect
}

// ApplicationAuthStrategyResourceModel describes the resource data model.
type ApplicationAuthStrategyResourceModel struct {
	Active        types.Bool                                   `tfsdk:"active"`
	DisplayName   types.String                                 `tfsdk:"display_name"`
	ID            types.String                                 `tfsdk:"id"`
	KeyAuth       *tfTypes.AppAuthStrategyKeyAuthRequest       `tfsdk:"key_auth" tfPlanOnly:"true"`
	Name          types.String                                 `tfsdk:"name"`
	OpenidConnect *tfTypes.AppAuthStrategyOpenIDConnectRequest `tfsdk:"openid_connect" tfPlanOnly:"true"`
}

func (r *ApplicationAuthStrategyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_auth_strategy"
}

func (r *ApplicationAuthStrategyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ApplicationAuthStrategy Resource",
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: `At least one published product version is using this auth strategy.`,
			},
			"display_name": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.` + "\n" +
					``,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Application auth strategy identifier`,
			},
			"key_auth": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
						},
						Description: `At least one published product version is using this auth strategy.`,
					},
					"configs": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"key_auth": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"key_names": schema.ListAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										ElementType: types.StringType,
										Description: `The names of the headers containing the API key. You can specify multiple header names. Requires replacement if changed. `,
										Validators: []validator.List{
											listvalidator.SizeAtLeast(1),
											listvalidator.SizeAtMost(10),
										},
									},
								},
								MarkdownDescription: `The most basic mode to configure an Application Auth Strategy for an API Product Version. ` + "\n" +
									`Using this mode will allow developers to generate API keys that will authenticate their application requests. ` + "\n" +
									`Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for Key Auth.` + "\n" +
									`` + "\n" +
									`Requires replacement if changed. ; Not Null`,
								Validators: []validator.Object{
									speakeasy_objectvalidators.NotNull(),
								},
							},
						},
						Description: `JSON-B object containing the configuration for the Key Auth strategy. Requires replacement if changed. ; Not Null`,
						Validators: []validator.Object{
							speakeasy_objectvalidators.NotNull(),
						},
					},
					"created_at": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `An ISO-8601 timestamp representation of entity creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dcr_provider": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Attributes: map[string]schema.Attribute{
							"display_name": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								MarkdownDescription: `The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.` + "\n" +
									``,
							},
							"id": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Description: `Contains a unique identifier used for this resource.`,
							},
							"name": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
							},
							"provider_type": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Description: `The type of DCR provider. Can be one of the following - auth0, azureAd, curity, okta, http. must be one of ["auth0", "azureAd", "curity", "okta", "http"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"auth0",
										"azureAd",
										"curity",
										"okta",
										"http",
									),
								},
							},
						},
					},
					"display_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						MarkdownDescription: `The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.` + "\n" +
							`` + "\n" +
							`Requires replacement if changed. ; Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
						},
					},
					"id": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `Contains a unique identifier used for this resource.`,
					},
					"name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						MarkdownDescription: `The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.` + "\n" +
							`` + "\n" +
							`Requires replacement if changed. ; Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
						},
					},
					"strategy_type": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Optional:    true,
						Description: `Requires replacement if changed. ; Not Null; must be one of ["key_auth"]`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.OneOf(
								"key_auth",
							),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `An ISO-8601 timestamp representation of entity update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Description: `Request for creating a Key Auth Application Auth Strategy. Requires replacement if changed. `,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("openid_connect"),
					}...),
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.` + "\n" +
					``,
			},
			"openid_connect": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
						},
						Description: `At least one published product version is using this auth strategy.`,
					},
					"configs": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"openid_connect": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Description: `Parsed as JSON.`,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
									},
									"auth_methods": schema.ListAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										ElementType: types.StringType,
										Description: `Requires replacement if changed. ; Not Null`,
										Validators: []validator.List{
											speakeasy_listvalidators.NotNull(),
											listvalidator.SizeAtMost(10),
										},
									},
									"credential_claim": schema.ListAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										ElementType: types.StringType,
										Description: `Requires replacement if changed. ; Not Null`,
										Validators: []validator.List{
											speakeasy_listvalidators.NotNull(),
											listvalidator.SizeAtMost(10),
										},
									},
									"issuer": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Description: `Requires replacement if changed. ; Not Null`,
										Validators: []validator.String{
											speakeasy_stringvalidators.NotNull(),
										},
									},
									"scopes": schema.ListAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										ElementType: types.StringType,
										Description: `Requires replacement if changed. ; Not Null`,
										Validators: []validator.List{
											speakeasy_listvalidators.NotNull(),
											listvalidator.SizeAtMost(50),
										},
									},
								},
								MarkdownDescription: `A more advanced mode to configure an API Product Version’s Application Auth Strategy. ` + "\n" +
									`Using this mode will allow developers to use API credentials issued from an external IdP that will authenticate their application requests. ` + "\n" +
									`Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for the same Auth Strategy. ` + "\n" +
									`An OIDC strategy may be used in conjunction with a DCR provider to automatically create the IdP application.` + "\n" +
									`` + "\n" +
									`Requires replacement if changed. ; Not Null`,
								Validators: []validator.Object{
									speakeasy_objectvalidators.NotNull(),
								},
							},
						},
						Description: `JSON-B object containing the configuration for the OIDC strategy. Requires replacement if changed. ; Not Null`,
						Validators: []validator.Object{
							speakeasy_objectvalidators.NotNull(),
						},
					},
					"created_at": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `An ISO-8601 timestamp representation of entity creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dcr_provider": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Attributes: map[string]schema.Attribute{
							"display_name": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								MarkdownDescription: `The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.` + "\n" +
									``,
							},
							"id": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Description: `Contains a unique identifier used for this resource.`,
							},
							"name": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
							},
							"provider_type": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Description: `The type of DCR provider. Can be one of the following - auth0, azureAd, curity, okta, http. must be one of ["auth0", "azureAd", "curity", "okta", "http"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"auth0",
										"azureAd",
										"curity",
										"okta",
										"http",
									),
								},
							},
						},
					},
					"dcr_provider_id": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Optional:    true,
						Description: `Requires replacement if changed. `,
					},
					"display_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						MarkdownDescription: `The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.` + "\n" +
							`` + "\n" +
							`Requires replacement if changed. ; Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
						},
					},
					"id": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `Contains a unique identifier used for this resource.`,
					},
					"name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						MarkdownDescription: `The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.` + "\n" +
							`` + "\n" +
							`Requires replacement if changed. ; Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
						},
					},
					"strategy_type": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Optional:    true,
						Description: `Requires replacement if changed. ; Not Null; must be one of ["openid_connect"]`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.OneOf(
								"openid_connect",
							),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `An ISO-8601 timestamp representation of entity update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Description: `Payload for creating an OIDC Application Auth Strategy. Requires replacement if changed. `,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("key_auth"),
					}...),
				},
			},
		},
	}
}

func (r *ApplicationAuthStrategyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ApplicationAuthStrategyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ApplicationAuthStrategyResourceModel
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

	request := *data.ToSharedCreateAppAuthStrategyRequest()
	res, err := r.client.AppAuthStrategies.CreateAppAuthStrategy(ctx, request)
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
	if !(res.CreateAppAuthStrategyResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCreateAppAuthStrategyResponse(res.CreateAppAuthStrategyResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	authStrategyID := data.ID.ValueString()
	request1 := operations.GetAppAuthStrategyRequest{
		AuthStrategyID: authStrategyID,
	}
	res1, err := r.client.AppAuthStrategies.GetAppAuthStrategy(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.GetAppAuthStrategyResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedGetAppAuthStrategyResponse(res1.GetAppAuthStrategyResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationAuthStrategyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ApplicationAuthStrategyResourceModel
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

	authStrategyID := data.ID.ValueString()
	request := operations.GetAppAuthStrategyRequest{
		AuthStrategyID: authStrategyID,
	}
	res, err := r.client.AppAuthStrategies.GetAppAuthStrategy(ctx, request)
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
	if !(res.GetAppAuthStrategyResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedGetAppAuthStrategyResponse(res.GetAppAuthStrategyResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationAuthStrategyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ApplicationAuthStrategyResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	authStrategyID := data.ID.ValueString()
	updateAppAuthStrategyRequest := *data.ToSharedUpdateAppAuthStrategyRequest()
	request := operations.UpdateAppAuthStrategyRequest{
		AuthStrategyID:               authStrategyID,
		UpdateAppAuthStrategyRequest: updateAppAuthStrategyRequest,
	}
	res, err := r.client.AppAuthStrategies.UpdateAppAuthStrategy(ctx, request)
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
	if !(res.UpdateAppAuthStrategyResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUpdateAppAuthStrategyResponse(res.UpdateAppAuthStrategyResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	authStrategyId1 := data.ID.ValueString()
	request1 := operations.GetAppAuthStrategyRequest{
		AuthStrategyID: authStrategyId1,
	}
	res1, err := r.client.AppAuthStrategies.GetAppAuthStrategy(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.GetAppAuthStrategyResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedGetAppAuthStrategyResponse(res1.GetAppAuthStrategyResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationAuthStrategyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ApplicationAuthStrategyResourceModel
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

	authStrategyID := data.ID.ValueString()
	request := operations.DeleteAppAuthStrategyRequest{
		AuthStrategyID: authStrategyID,
	}
	res, err := r.client.AppAuthStrategies.DeleteAppAuthStrategy(ctx, request)
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

func (r *ApplicationAuthStrategyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
