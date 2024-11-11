// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshControlPlaneDataSourceModel) RefreshFromSharedMeshControlPlane(resp *shared.MeshControlPlane) {
	if resp != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringValue(resp.ID)
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String)
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.Name = types.StringValue(resp.Name)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}
