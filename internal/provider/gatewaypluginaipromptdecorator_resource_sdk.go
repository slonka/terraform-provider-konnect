// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginAiPromptDecoratorResourceModel) ToSharedAiPromptDecoratorPluginInput() *shared.AiPromptDecoratorPluginInput {
	maxRequestBodySize := new(int64)
	if !r.Config.MaxRequestBodySize.IsUnknown() && !r.Config.MaxRequestBodySize.IsNull() {
		*maxRequestBodySize = r.Config.MaxRequestBodySize.ValueInt64()
	} else {
		maxRequestBodySize = nil
	}
	var prompts *shared.Prompts
	if r.Config.Prompts != nil {
		var append1 []shared.AiPromptDecoratorPluginAppend = []shared.AiPromptDecoratorPluginAppend{}
		for _, appendItem := range r.Config.Prompts.Append {
			var content string
			content = appendItem.Content.ValueString()

			role := new(shared.Role)
			if !appendItem.Role.IsUnknown() && !appendItem.Role.IsNull() {
				*role = shared.Role(appendItem.Role.ValueString())
			} else {
				role = nil
			}
			append1 = append(append1, shared.AiPromptDecoratorPluginAppend{
				Content: content,
				Role:    role,
			})
		}
		var prepend []shared.Prepend = []shared.Prepend{}
		for _, prependItem := range r.Config.Prompts.Prepend {
			var content1 string
			content1 = prependItem.Content.ValueString()

			role1 := new(shared.AiPromptDecoratorPluginRole)
			if !prependItem.Role.IsUnknown() && !prependItem.Role.IsNull() {
				*role1 = shared.AiPromptDecoratorPluginRole(prependItem.Role.ValueString())
			} else {
				role1 = nil
			}
			prepend = append(prepend, shared.Prepend{
				Content: content1,
				Role:    role1,
			})
		}
		prompts = &shared.Prompts{
			Append:  append1,
			Prepend: prepend,
		}
	}
	config := shared.AiPromptDecoratorPluginConfig{
		MaxRequestBodySize: maxRequestBodySize,
		Prompts:            prompts,
	}
	var consumer *shared.AiPromptDecoratorPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.AiPromptDecoratorPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.AiPromptDecoratorPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.AiPromptDecoratorPluginConsumerGroup{
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
	var ordering *shared.AiPromptDecoratorPluginOrdering
	if r.Ordering != nil {
		var after *shared.AiPromptDecoratorPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.AiPromptDecoratorPluginAfter{
				Access: access,
			}
		}
		var before *shared.AiPromptDecoratorPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.AiPromptDecoratorPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.AiPromptDecoratorPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.AiPromptDecoratorPluginProtocols = []shared.AiPromptDecoratorPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.AiPromptDecoratorPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.AiPromptDecoratorPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.AiPromptDecoratorPluginRoute{
			ID: id3,
		}
	}
	var service *shared.AiPromptDecoratorPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.AiPromptDecoratorPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.AiPromptDecoratorPluginInput{
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

func (r *GatewayPluginAiPromptDecoratorResourceModel) RefreshFromSharedAiPromptDecoratorPlugin(resp *shared.AiPromptDecoratorPlugin) {
	if resp != nil {
		r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
		if resp.Config.Prompts == nil {
			r.Config.Prompts = nil
		} else {
			r.Config.Prompts = &tfTypes.Prompts{}
			r.Config.Prompts.Append = []tfTypes.AiPromptDecoratorPluginAppend{}
			if len(r.Config.Prompts.Append) > len(resp.Config.Prompts.Append) {
				r.Config.Prompts.Append = r.Config.Prompts.Append[:len(resp.Config.Prompts.Append)]
			}
			for appendCount, appendItem := range resp.Config.Prompts.Append {
				var append2 tfTypes.AiPromptDecoratorPluginAppend
				append2.Content = types.StringValue(appendItem.Content)
				if appendItem.Role != nil {
					append2.Role = types.StringValue(string(*appendItem.Role))
				} else {
					append2.Role = types.StringNull()
				}
				if appendCount+1 > len(r.Config.Prompts.Append) {
					r.Config.Prompts.Append = append(r.Config.Prompts.Append, append2)
				} else {
					r.Config.Prompts.Append[appendCount].Content = append2.Content
					r.Config.Prompts.Append[appendCount].Role = append2.Role
				}
			}
			r.Config.Prompts.Prepend = []tfTypes.AiPromptDecoratorPluginAppend{}
			if len(r.Config.Prompts.Prepend) > len(resp.Config.Prompts.Prepend) {
				r.Config.Prompts.Prepend = r.Config.Prompts.Prepend[:len(resp.Config.Prompts.Prepend)]
			}
			for prependCount, prependItem := range resp.Config.Prompts.Prepend {
				var prepend1 tfTypes.AiPromptDecoratorPluginAppend
				prepend1.Content = types.StringValue(prependItem.Content)
				if prependItem.Role != nil {
					prepend1.Role = types.StringValue(string(*prependItem.Role))
				} else {
					prepend1.Role = types.StringNull()
				}
				if prependCount+1 > len(r.Config.Prompts.Prepend) {
					r.Config.Prompts.Prepend = append(r.Config.Prompts.Prepend, prepend1)
				} else {
					r.Config.Prompts.Prepend[prependCount].Content = prepend1.Content
					r.Config.Prompts.Prepend[prependCount].Role = prepend1.Role
				}
			}
		}
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
