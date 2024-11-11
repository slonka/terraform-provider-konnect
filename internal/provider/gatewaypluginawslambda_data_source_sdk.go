// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginAwsLambdaDataSourceModel) RefreshFromSharedAwsLambdaPlugin(resp *shared.AwsLambdaPlugin) {
	if resp != nil {
		r.Config.AwsAssumeRoleArn = types.StringPointerValue(resp.Config.AwsAssumeRoleArn)
		if resp.Config.AwsImdsProtocolVersion != nil {
			r.Config.AwsImdsProtocolVersion = types.StringValue(string(*resp.Config.AwsImdsProtocolVersion))
		} else {
			r.Config.AwsImdsProtocolVersion = types.StringNull()
		}
		r.Config.AwsKey = types.StringPointerValue(resp.Config.AwsKey)
		r.Config.AwsRegion = types.StringPointerValue(resp.Config.AwsRegion)
		r.Config.AwsRoleSessionName = types.StringPointerValue(resp.Config.AwsRoleSessionName)
		r.Config.AwsSecret = types.StringPointerValue(resp.Config.AwsSecret)
		r.Config.AwsStsEndpointURL = types.StringPointerValue(resp.Config.AwsStsEndpointURL)
		r.Config.AwsgatewayCompatible = types.BoolPointerValue(resp.Config.AwsgatewayCompatible)
		r.Config.Base64EncodeBody = types.BoolPointerValue(resp.Config.Base64EncodeBody)
		r.Config.DisableHTTPS = types.BoolPointerValue(resp.Config.DisableHTTPS)
		if resp.Config.EmptyArraysMode != nil {
			r.Config.EmptyArraysMode = types.StringValue(string(*resp.Config.EmptyArraysMode))
		} else {
			r.Config.EmptyArraysMode = types.StringNull()
		}
		r.Config.ForwardRequestBody = types.BoolPointerValue(resp.Config.ForwardRequestBody)
		r.Config.ForwardRequestHeaders = types.BoolPointerValue(resp.Config.ForwardRequestHeaders)
		r.Config.ForwardRequestMethod = types.BoolPointerValue(resp.Config.ForwardRequestMethod)
		r.Config.ForwardRequestURI = types.BoolPointerValue(resp.Config.ForwardRequestURI)
		r.Config.FunctionName = types.StringPointerValue(resp.Config.FunctionName)
		r.Config.Host = types.StringPointerValue(resp.Config.Host)
		if resp.Config.InvocationType != nil {
			r.Config.InvocationType = types.StringValue(string(*resp.Config.InvocationType))
		} else {
			r.Config.InvocationType = types.StringNull()
		}
		r.Config.IsProxyIntegration = types.BoolPointerValue(resp.Config.IsProxyIntegration)
		if resp.Config.Keepalive != nil {
			r.Config.Keepalive = types.NumberValue(big.NewFloat(float64(*resp.Config.Keepalive)))
		} else {
			r.Config.Keepalive = types.NumberNull()
		}
		if resp.Config.LogType != nil {
			r.Config.LogType = types.StringValue(string(*resp.Config.LogType))
		} else {
			r.Config.LogType = types.StringNull()
		}
		r.Config.Port = types.Int64PointerValue(resp.Config.Port)
		r.Config.ProxyURL = types.StringPointerValue(resp.Config.ProxyURL)
		r.Config.Qualifier = types.StringPointerValue(resp.Config.Qualifier)
		r.Config.SkipLargeBodies = types.BoolPointerValue(resp.Config.SkipLargeBodies)
		if resp.Config.Timeout != nil {
			r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
		} else {
			r.Config.Timeout = types.NumberNull()
		}
		r.Config.UnhandledStatus = types.Int64PointerValue(resp.Config.UnhandledStatus)
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
