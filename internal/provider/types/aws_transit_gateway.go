// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AWSTransitGateway struct {
	CidrBlocks                     []types.String                    `tfsdk:"cidr_blocks"`
	DNSConfig                      []TransitGatewayDNSConfig         `tfsdk:"dns_config"`
	Name                           types.String                      `tfsdk:"name"`
	TransitGatewayAttachmentConfig AwsTransitGatewayAttachmentConfig `tfsdk:"transit_gateway_attachment_config"`
}
