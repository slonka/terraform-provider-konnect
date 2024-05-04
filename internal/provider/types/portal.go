// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Portal struct {
	ApplicationCount                 types.Number `tfsdk:"application_count"`
	AutoApproveApplications          types.Bool   `tfsdk:"auto_approve_applications"`
	AutoApproveDevelopers            types.Bool   `tfsdk:"auto_approve_developers"`
	CreatedAt                        types.String `tfsdk:"created_at"`
	CustomClientDomain               types.String `tfsdk:"custom_client_domain"`
	CustomDomain                     types.String `tfsdk:"custom_domain"`
	DefaultApplicationAuthStrategyID types.String `tfsdk:"default_application_auth_strategy_id"`
	DefaultDomain                    types.String `tfsdk:"default_domain"`
	Description                      types.String `tfsdk:"description"`
	DeveloperCount                   types.Number `tfsdk:"developer_count"`
	DisplayName                      types.String `tfsdk:"display_name"`
	ID                               types.String `tfsdk:"id"`
	IsPublic                         types.Bool   `tfsdk:"is_public"`
	Name                             types.String `tfsdk:"name"`
	PublishedProductCount            types.Number `tfsdk:"published_product_count"`
	RbacEnabled                      types.Bool   `tfsdk:"rbac_enabled"`
	UpdatedAt                        types.String `tfsdk:"updated_at"`
}
