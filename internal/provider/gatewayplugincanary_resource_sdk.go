// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginCanaryResourceModel) ToSharedCanaryPluginInput() *shared.CanaryPluginInput {
	canaryByHeaderName := new(string)
	if !r.Config.CanaryByHeaderName.IsUnknown() && !r.Config.CanaryByHeaderName.IsNull() {
		*canaryByHeaderName = r.Config.CanaryByHeaderName.ValueString()
	} else {
		canaryByHeaderName = nil
	}
	duration := new(float64)
	if !r.Config.Duration.IsUnknown() && !r.Config.Duration.IsNull() {
		*duration, _ = r.Config.Duration.ValueBigFloat().Float64()
	} else {
		duration = nil
	}
	var groups []string = []string{}
	for _, groupsItem := range r.Config.Groups {
		groups = append(groups, groupsItem.ValueString())
	}
	hash := new(shared.Hash)
	if !r.Config.Hash.IsUnknown() && !r.Config.Hash.IsNull() {
		*hash = shared.Hash(r.Config.Hash.ValueString())
	} else {
		hash = nil
	}
	hashHeader := new(string)
	if !r.Config.HashHeader.IsUnknown() && !r.Config.HashHeader.IsNull() {
		*hashHeader = r.Config.HashHeader.ValueString()
	} else {
		hashHeader = nil
	}
	percentage := new(float64)
	if !r.Config.Percentage.IsUnknown() && !r.Config.Percentage.IsNull() {
		*percentage, _ = r.Config.Percentage.ValueBigFloat().Float64()
	} else {
		percentage = nil
	}
	start := new(float64)
	if !r.Config.Start.IsUnknown() && !r.Config.Start.IsNull() {
		*start, _ = r.Config.Start.ValueBigFloat().Float64()
	} else {
		start = nil
	}
	steps := new(float64)
	if !r.Config.Steps.IsUnknown() && !r.Config.Steps.IsNull() {
		*steps, _ = r.Config.Steps.ValueBigFloat().Float64()
	} else {
		steps = nil
	}
	upstreamFallback := new(bool)
	if !r.Config.UpstreamFallback.IsUnknown() && !r.Config.UpstreamFallback.IsNull() {
		*upstreamFallback = r.Config.UpstreamFallback.ValueBool()
	} else {
		upstreamFallback = nil
	}
	upstreamHost := new(string)
	if !r.Config.UpstreamHost.IsUnknown() && !r.Config.UpstreamHost.IsNull() {
		*upstreamHost = r.Config.UpstreamHost.ValueString()
	} else {
		upstreamHost = nil
	}
	upstreamPort := new(int64)
	if !r.Config.UpstreamPort.IsUnknown() && !r.Config.UpstreamPort.IsNull() {
		*upstreamPort = r.Config.UpstreamPort.ValueInt64()
	} else {
		upstreamPort = nil
	}
	upstreamURI := new(string)
	if !r.Config.UpstreamURI.IsUnknown() && !r.Config.UpstreamURI.IsNull() {
		*upstreamURI = r.Config.UpstreamURI.ValueString()
	} else {
		upstreamURI = nil
	}
	config := shared.CanaryPluginConfig{
		CanaryByHeaderName: canaryByHeaderName,
		Duration:           duration,
		Groups:             groups,
		Hash:               hash,
		HashHeader:         hashHeader,
		Percentage:         percentage,
		Start:              start,
		Steps:              steps,
		UpstreamFallback:   upstreamFallback,
		UpstreamHost:       upstreamHost,
		UpstreamPort:       upstreamPort,
		UpstreamURI:        upstreamURI,
	}
	var consumer *shared.CanaryPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CanaryPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CanaryPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CanaryPluginConsumerGroup{
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
	var ordering *shared.CanaryPluginOrdering
	if r.Ordering != nil {
		var after *shared.CanaryPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CanaryPluginAfter{
				Access: access,
			}
		}
		var before *shared.CanaryPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CanaryPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CanaryPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CanaryPluginProtocols = []shared.CanaryPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CanaryPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.CanaryPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.CanaryPluginRoute{
			ID: id3,
		}
	}
	var service *shared.CanaryPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.CanaryPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.CanaryPluginInput{
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

func (r *GatewayPluginCanaryResourceModel) RefreshFromSharedCanaryPlugin(resp *shared.CanaryPlugin) {
	if resp != nil {
		r.Config.CanaryByHeaderName = types.StringPointerValue(resp.Config.CanaryByHeaderName)
		if resp.Config.Duration != nil {
			r.Config.Duration = types.NumberValue(big.NewFloat(float64(*resp.Config.Duration)))
		} else {
			r.Config.Duration = types.NumberNull()
		}
		r.Config.Groups = []types.String{}
		for _, v := range resp.Config.Groups {
			r.Config.Groups = append(r.Config.Groups, types.StringValue(v))
		}
		if resp.Config.Hash != nil {
			r.Config.Hash = types.StringValue(string(*resp.Config.Hash))
		} else {
			r.Config.Hash = types.StringNull()
		}
		r.Config.HashHeader = types.StringPointerValue(resp.Config.HashHeader)
		if resp.Config.Percentage != nil {
			r.Config.Percentage = types.NumberValue(big.NewFloat(float64(*resp.Config.Percentage)))
		} else {
			r.Config.Percentage = types.NumberNull()
		}
		if resp.Config.Start != nil {
			r.Config.Start = types.NumberValue(big.NewFloat(float64(*resp.Config.Start)))
		} else {
			r.Config.Start = types.NumberNull()
		}
		if resp.Config.Steps != nil {
			r.Config.Steps = types.NumberValue(big.NewFloat(float64(*resp.Config.Steps)))
		} else {
			r.Config.Steps = types.NumberNull()
		}
		r.Config.UpstreamFallback = types.BoolPointerValue(resp.Config.UpstreamFallback)
		r.Config.UpstreamHost = types.StringPointerValue(resp.Config.UpstreamHost)
		r.Config.UpstreamPort = types.Int64PointerValue(resp.Config.UpstreamPort)
		r.Config.UpstreamURI = types.StringPointerValue(resp.Config.UpstreamURI)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLConsumer{}
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
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
