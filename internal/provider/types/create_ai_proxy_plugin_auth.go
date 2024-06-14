// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateAIProxyPluginAuth struct {
	AzureClientID           types.String `tfsdk:"azure_client_id"`
	AzureClientSecret       types.String `tfsdk:"azure_client_secret"`
	AzureTenantID           types.String `tfsdk:"azure_tenant_id"`
	AzureUseManagedIdentity types.Bool   `tfsdk:"azure_use_managed_identity"`
	HeaderName              types.String `tfsdk:"header_name"`
	HeaderValue             types.String `tfsdk:"header_value"`
	ParamLocation           types.String `tfsdk:"param_location"`
	ParamName               types.String `tfsdk:"param_name"`
	ParamValue              types.String `tfsdk:"param_value"`
}
