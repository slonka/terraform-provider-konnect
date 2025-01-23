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
	speakeasy_listvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/listvalidators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/objectvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayRouteResource{}
var _ resource.ResourceWithImportState = &GatewayRouteResource{}

func NewGatewayRouteResource() resource.Resource {
	return &GatewayRouteResource{}
}

// GatewayRouteResource defines the resource implementation.
type GatewayRouteResource struct {
	client *sdk.Konnect
}

// GatewayRouteResourceModel describes the resource data model.
type GatewayRouteResourceModel struct {
	ControlPlaneID          types.String                       `tfsdk:"control_plane_id"`
	CreatedAt               types.Int64                        `tfsdk:"created_at"`
	Destinations            []tfTypes.ClusterNodes             `tfsdk:"destinations"`
	Headers                 map[string][]types.String          `tfsdk:"headers"`
	Hosts                   []types.String                     `tfsdk:"hosts"`
	HTTPSRedirectStatusCode types.Int64                        `tfsdk:"https_redirect_status_code"`
	ID                      types.String                       `tfsdk:"id"`
	Methods                 []types.String                     `tfsdk:"methods"`
	Name                    types.String                       `tfsdk:"name"`
	PathHandling            types.String                       `tfsdk:"path_handling"`
	Paths                   []types.String                     `tfsdk:"paths"`
	PreserveHost            types.Bool                         `tfsdk:"preserve_host"`
	Protocols               []types.String                     `tfsdk:"protocols"`
	RegexPriority           types.Int64                        `tfsdk:"regex_priority"`
	RequestBuffering        types.Bool                         `tfsdk:"request_buffering"`
	ResponseBuffering       types.Bool                         `tfsdk:"response_buffering"`
	Service                 *tfTypes.ACLWithoutParentsConsumer `tfsdk:"service" tfPlanOnly:"true"`
	Snis                    []types.String                     `tfsdk:"snis"`
	Sources                 []tfTypes.ClusterNodes             `tfsdk:"sources"`
	StripPath               types.Bool                         `tfsdk:"strip_path"`
	Tags                    []types.String                     `tfsdk:"tags"`
	UpdatedAt               types.Int64                        `tfsdk:"updated_at"`
}

func (r *GatewayRouteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_route"
}

func (r *GatewayRouteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayRoute Resource",
		Attributes: map[string]schema.Attribute{
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
			"destinations": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.Object{
						speakeasy_objectvalidators.NotNull(),
					},
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"port": schema.Int64Attribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Description: `A list of IP destinations of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".`,
			},
			"headers": schema.MapNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.List{
						speakeasy_listvalidators.NotNull(),
					},
					Attributes: map[string]schema.Attribute{},
				},
				Description: `One or more lists of values indexed by header name that will cause this Route to match if present in the request. The ` + "`" + `Host` + "`" + ` header cannot be used with this attribute: hosts should be specified using the ` + "`" + `hosts` + "`" + ` attribute. When ` + "`" + `headers` + "`" + ` contains only one value and that value starts with the special prefix ` + "`" + `~*` + "`" + `, the value is interpreted as a regular expression.`,
			},
			"hosts": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of domain names that match this Route. Note that the hosts value is case sensitive.`,
			},
			"https_redirect_status_code": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Description: `The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is ` + "`" + `HTTP` + "`" + ` instead of ` + "`" + `HTTPS` + "`" + `. ` + "`" + `Location` + "`" + ` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the ` + "`" + `https` + "`" + ` protocol. must be one of ["426", "301", "302", "307", "308"]`,
				Validators: []validator.Int64{
					int64validator.OneOf(
						426,
						301,
						302,
						307,
						308,
					),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"methods": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of HTTP methods that match this Route.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The name of the Route. Route names must be unique, and they are case sensitive. For example, there can be two different Routes named "test" and "Test".`,
			},
			"path_handling": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior. must be one of ["v0", "v1"]`,
				Validators: []validator.String{
					stringvalidator.OneOf("v0", "v1"),
				},
			},
			"paths": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of paths that match this Route.`,
			},
			"preserve_host": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `When matching a Route via one of the ` + "`" + `hosts` + "`" + ` domain names, use the request ` + "`" + `Host` + "`" + ` header in the upstream request headers. If set to ` + "`" + `false` + "`" + `, the upstream ` + "`" + `Host` + "`" + ` header will be that of the Service's ` + "`" + `host` + "`" + `.`,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An array of the protocols this Route should allow. See the [Route Object](#route-object) section for a list of accepted protocols. When set to only ` + "`" + `"https"` + "`" + `, HTTP requests are answered with an upgrade error. When set to only ` + "`" + `"http"` + "`" + `, HTTPS requests are answered with an error.`,
			},
			"regex_priority": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Description: `A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same ` + "`" + `regex_priority` + "`" + `, the older one (lowest ` + "`" + `created_at` + "`" + `) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).`,
			},
			"request_buffering": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.`,
			},
			"response_buffering": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.`,
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
				Description: `The Service this Route is associated to. This is where the Route proxies traffic to.`,
			},
			"snis": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of SNIs that match this Route when using stream routing.`,
			},
			"sources": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.Object{
						speakeasy_objectvalidators.NotNull(),
					},
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"port": schema.Int64Attribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Description: `A list of IP sources of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".`,
			},
			"strip_path": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `When matching a Route via one of the ` + "`" + `paths` + "`" + `, strip the matching prefix from the upstream request URL.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Route for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayRouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayRouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayRouteResourceModel
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

	route := *data.ToSharedRouteInput()
	request := operations.CreateRouteRequest{
		ControlPlaneID: controlPlaneID,
		Route:          route,
	}
	res, err := r.client.Routes.CreateRoute(ctx, request)
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
	if !(res.Route != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRoute(res.Route)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayRouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayRouteResourceModel
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

	var routeID string
	routeID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetRouteRequest{
		RouteID:        routeID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Routes.GetRoute(ctx, request)
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
	if !(res.Route != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRoute(res.Route)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayRouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayRouteResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var routeID string
	routeID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	route := *data.ToSharedRouteInput()
	request := operations.UpsertRouteRequest{
		RouteID:        routeID,
		ControlPlaneID: controlPlaneID,
		Route:          route,
	}
	res, err := r.client.Routes.UpsertRoute(ctx, request)
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
	if !(res.Route != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRoute(res.Route)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayRouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayRouteResourceModel
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

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	var routeID string
	routeID = data.ID.ValueString()

	request := operations.DeleteRouteRequest{
		ControlPlaneID: controlPlaneID,
		RouteID:        routeID,
	}
	res, err := r.client.Routes.DeleteRoute(ctx, request)
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

func (r *GatewayRouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		ControlPlaneID string `json:"control_plane_id"`
		ID             string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "control_plane_id": "9524ec7d-36d9-465d-a8c5-83a3c9390458",  "route_id": "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7"}': `+err.Error())
		return
	}

	if len(data.ControlPlaneID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field control_plane_id is required but was not found in the json encoded ID. It's expected to be a value alike '"9524ec7d-36d9-465d-a8c5-83a3c9390458"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("control_plane_id"), data.ControlPlaneID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"a4326a41-aa12-44e3-93e4-6b6e58bfb9d7"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
