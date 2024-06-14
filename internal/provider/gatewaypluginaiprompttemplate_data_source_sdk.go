// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginAIPromptTemplateDataSourceModel) RefreshFromSharedAIPromptTemplatePlugin(resp *shared.AIPromptTemplatePlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAIPromptTemplatePluginConfig{}
			r.Config.AllowUntemplatedRequests = types.BoolPointerValue(resp.Config.AllowUntemplatedRequests)
			r.Config.LogOriginalRequest = types.BoolPointerValue(resp.Config.LogOriginalRequest)
			r.Config.Templates = []tfTypes.Templates{}
			if len(r.Config.Templates) > len(resp.Config.Templates) {
				r.Config.Templates = r.Config.Templates[:len(resp.Config.Templates)]
			}
			for templatesCount, templatesItem := range resp.Config.Templates {
				var templates1 tfTypes.Templates
				templates1.Name = types.StringValue(templatesItem.Name)
				templates1.Template = types.StringValue(templatesItem.Template)
				if templatesCount+1 > len(r.Config.Templates) {
					r.Config.Templates = append(r.Config.Templates, templates1)
				} else {
					r.Config.Templates[templatesCount].Name = templates1.Name
					r.Config.Templates[templatesCount].Template = templates1.Template
				}
			}
		}
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
