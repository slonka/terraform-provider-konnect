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
var _ resource.Resource = &GatewayKeyResource{}
var _ resource.ResourceWithImportState = &GatewayKeyResource{}

func NewGatewayKeyResource() resource.Resource {
	return &GatewayKeyResource{}
}

// GatewayKeyResource defines the resource implementation.
type GatewayKeyResource struct {
	client *sdk.Konnect
}

// GatewayKeyResourceModel describes the resource data model.
type GatewayKeyResourceModel struct {
	ControlPlaneID types.String         `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64          `tfsdk:"created_at"`
	ID             types.String         `tfsdk:"id"`
	Jwk            types.String         `tfsdk:"jwk"`
	Kid            types.String         `tfsdk:"kid"`
	Name           types.String         `tfsdk:"name"`
	Pem            *tfTypes.Pem         `tfsdk:"pem"`
	Set            *tfTypes.ACLConsumer `tfsdk:"set" tfPlanOnly:"true"`
	Tags           []types.String       `tfsdk:"tags"`
	UpdatedAt      types.Int64          `tfsdk:"updated_at"`
}

func (r *GatewayKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_key"
}

func (r *GatewayKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayKey Resource",
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
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"jwk": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `A JSON Web Key represented as a string.`,
			},
			"kid": schema.StringAttribute{
				Required:    true,
				Description: `A unique identifier for a key.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The name to associate with the given keys.`,
			},
			"pem": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"private_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"public_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `A keypair in PEM format.`,
			},
			"set": schema.SingleNestedAttribute{
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
				Description: `The id (an UUID) of the key-set with which to associate the key.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
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

func (r *GatewayKeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayKeyResourceModel
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

	key := *data.ToSharedKeyInput()
	request := operations.CreateKeyRequest{
		ControlPlaneID: controlPlaneID,
		Key:            key,
	}
	res, err := r.client.Keys.CreateKey(ctx, request)
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
	if !(res.Key != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedKey(res.Key)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayKeyResourceModel
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

	var keyID string
	keyID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetKeyRequest{
		KeyID:          keyID,
		ControlPlaneID: controlPlaneID,
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

func (r *GatewayKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayKeyResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var keyID string
	keyID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	key := *data.ToSharedKeyInput()
	request := operations.UpsertKeyRequest{
		KeyID:          keyID,
		ControlPlaneID: controlPlaneID,
		Key:            key,
	}
	res, err := r.client.Keys.UpsertKey(ctx, request)
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
	if !(res.Key != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedKey(res.Key)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayKeyResourceModel
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

	var keyID string
	keyID = data.ID.ValueString()

	request := operations.DeleteKeyRequest{
		ControlPlaneID: controlPlaneID,
		KeyID:          keyID,
	}
	res, err := r.client.Keys.DeleteKey(ctx, request)
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

func (r *GatewayKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		ControlPlaneID string `json:"control_plane_id"`
		ID             string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "control_plane_id": "9524ec7d-36d9-465d-a8c5-83a3c9390458",  "key_id": "bba22c06-a632-42be-a018-1b9ff357b5b9"}': `+err.Error())
		return
	}

	if len(data.ControlPlaneID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field control_plane_id is required but was not found in the json encoded ID. It's expected to be a value alike '"9524ec7d-36d9-465d-a8c5-83a3c9390458"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("control_plane_id"), data.ControlPlaneID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"bba22c06-a632-42be-a018-1b9ff357b5b9"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
