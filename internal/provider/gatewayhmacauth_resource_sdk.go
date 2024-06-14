// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayHMACAuthResourceModel) ToSharedHMACAuthWithoutParents() *shared.HMACAuthWithoutParents {
	secret := new(string)
	if !r.Secret.IsUnknown() && !r.Secret.IsNull() {
		*secret = r.Secret.ValueString()
	} else {
		secret = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	username := new(string)
	if !r.Username.IsUnknown() && !r.Username.IsNull() {
		*username = r.Username.ValueString()
	} else {
		username = nil
	}
	out := shared.HMACAuthWithoutParents{
		Secret:   secret,
		Tags:     tags,
		Username: username,
	}
	return &out
}

func (r *GatewayHMACAuthResourceModel) RefreshFromSharedHMACAuth(resp *shared.HMACAuth) {
	if resp != nil {
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Secret = types.StringPointerValue(resp.Secret)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.Username = types.StringPointerValue(resp.Username)
	}
}
