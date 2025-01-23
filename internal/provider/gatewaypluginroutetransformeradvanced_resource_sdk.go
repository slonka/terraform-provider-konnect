// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginRouteTransformerAdvancedResourceModel) ToSharedRouteTransformerAdvancedPluginInput() *shared.RouteTransformerAdvancedPluginInput {
	escapePath := new(bool)
	if !r.Config.EscapePath.IsUnknown() && !r.Config.EscapePath.IsNull() {
		*escapePath = r.Config.EscapePath.ValueBool()
	} else {
		escapePath = nil
	}
	host := new(string)
	if !r.Config.Host.IsUnknown() && !r.Config.Host.IsNull() {
		*host = r.Config.Host.ValueString()
	} else {
		host = nil
	}
	path := new(string)
	if !r.Config.Path.IsUnknown() && !r.Config.Path.IsNull() {
		*path = r.Config.Path.ValueString()
	} else {
		path = nil
	}
	port := new(string)
	if !r.Config.Port.IsUnknown() && !r.Config.Port.IsNull() {
		*port = r.Config.Port.ValueString()
	} else {
		port = nil
	}
	config := shared.RouteTransformerAdvancedPluginConfig{
		EscapePath: escapePath,
		Host:       host,
		Path:       path,
		Port:       port,
	}
	var consumer *shared.RouteTransformerAdvancedPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.RouteTransformerAdvancedPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.RouteTransformerAdvancedPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.RouteTransformerAdvancedPluginConsumerGroup{
			ID: id1,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id2 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id2 = r.ID.ValueString()
	} else {
		id2 = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.RouteTransformerAdvancedPluginOrdering
	if r.Ordering != nil {
		var after *shared.RouteTransformerAdvancedPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.RouteTransformerAdvancedPluginAfter{
				Access: access,
			}
		}
		var before *shared.RouteTransformerAdvancedPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.RouteTransformerAdvancedPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.RouteTransformerAdvancedPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.RouteTransformerAdvancedPluginProtocols = []shared.RouteTransformerAdvancedPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.RouteTransformerAdvancedPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.RouteTransformerAdvancedPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.RouteTransformerAdvancedPluginRoute{
			ID: id3,
		}
	}
	var service *shared.RouteTransformerAdvancedPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.RouteTransformerAdvancedPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.RouteTransformerAdvancedPluginInput{
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Enabled:       enabled,
		ID:            id2,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
		Tags:          tags,
	}
	return &out
}

func (r *GatewayPluginRouteTransformerAdvancedResourceModel) RefreshFromSharedRouteTransformerAdvancedPlugin(resp *shared.RouteTransformerAdvancedPlugin) {
	if resp != nil {
		r.Config.EscapePath = types.BoolPointerValue(resp.Config.EscapePath)
		r.Config.Host = types.StringPointerValue(resp.Config.Host)
		r.Config.Path = types.StringPointerValue(resp.Config.Path)
		r.Config.Port = types.StringPointerValue(resp.Config.Port)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLWithoutParentsConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.ACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.ACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.ACLPluginAfter{}
				r.Ordering.Before.Access = []types.String{}
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
