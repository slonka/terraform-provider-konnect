// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *APIProductSpecificationDataSourceModel) RefreshFromSharedAPIProductVersionSpec(resp *shared.APIProductVersionSpec) {
	if resp != nil {
		r.Content = types.StringValue(resp.Content)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}
