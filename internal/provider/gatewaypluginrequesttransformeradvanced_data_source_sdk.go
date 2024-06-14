// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginRequestTransformerAdvancedDataSourceModel) RefreshFromSharedRequestTransformerAdvancedPlugin(resp *shared.RequestTransformerAdvancedPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateRequestTransformerAdvancedPluginConfig{}
			if resp.Config.Add == nil {
				r.Config.Add = nil
			} else {
				r.Config.Add = &tfTypes.CreateRequestTransformerAdvancedPluginAdd{}
				r.Config.Add.Body = []types.String{}
				for _, v := range resp.Config.Add.Body {
					r.Config.Add.Body = append(r.Config.Add.Body, types.StringValue(v))
				}
				r.Config.Add.Headers = []types.String{}
				for _, v := range resp.Config.Add.Headers {
					r.Config.Add.Headers = append(r.Config.Add.Headers, types.StringValue(v))
				}
				r.Config.Add.JSONTypes = []types.String{}
				for _, v := range resp.Config.Add.JSONTypes {
					r.Config.Add.JSONTypes = append(r.Config.Add.JSONTypes, types.StringValue(string(v)))
				}
				r.Config.Add.Querystring = []types.String{}
				for _, v := range resp.Config.Add.Querystring {
					r.Config.Add.Querystring = append(r.Config.Add.Querystring, types.StringValue(v))
				}
			}
			if resp.Config.Allow == nil {
				r.Config.Allow = nil
			} else {
				r.Config.Allow = &tfTypes.CreateRequestTransformerAdvancedPluginAllow{}
				r.Config.Allow.Body = []types.String{}
				for _, v := range resp.Config.Allow.Body {
					r.Config.Allow.Body = append(r.Config.Allow.Body, types.StringValue(v))
				}
			}
			if resp.Config.Append == nil {
				r.Config.Append = nil
			} else {
				r.Config.Append = &tfTypes.CreateRequestTransformerAdvancedPluginAdd{}
				r.Config.Append.Body = []types.String{}
				for _, v := range resp.Config.Append.Body {
					r.Config.Append.Body = append(r.Config.Append.Body, types.StringValue(v))
				}
				r.Config.Append.Headers = []types.String{}
				for _, v := range resp.Config.Append.Headers {
					r.Config.Append.Headers = append(r.Config.Append.Headers, types.StringValue(v))
				}
				r.Config.Append.JSONTypes = []types.String{}
				for _, v := range resp.Config.Append.JSONTypes {
					r.Config.Append.JSONTypes = append(r.Config.Append.JSONTypes, types.StringValue(string(v)))
				}
				r.Config.Append.Querystring = []types.String{}
				for _, v := range resp.Config.Append.Querystring {
					r.Config.Append.Querystring = append(r.Config.Append.Querystring, types.StringValue(v))
				}
			}
			r.Config.DotsInKeys = types.BoolPointerValue(resp.Config.DotsInKeys)
			r.Config.HTTPMethod = types.StringPointerValue(resp.Config.HTTPMethod)
			if resp.Config.Remove == nil {
				r.Config.Remove = nil
			} else {
				r.Config.Remove = &tfTypes.CreateRequestTransformerPluginAdd{}
				r.Config.Remove.Body = []types.String{}
				for _, v := range resp.Config.Remove.Body {
					r.Config.Remove.Body = append(r.Config.Remove.Body, types.StringValue(v))
				}
				r.Config.Remove.Headers = []types.String{}
				for _, v := range resp.Config.Remove.Headers {
					r.Config.Remove.Headers = append(r.Config.Remove.Headers, types.StringValue(v))
				}
				r.Config.Remove.Querystring = []types.String{}
				for _, v := range resp.Config.Remove.Querystring {
					r.Config.Remove.Querystring = append(r.Config.Remove.Querystring, types.StringValue(v))
				}
			}
			if resp.Config.Rename == nil {
				r.Config.Rename = nil
			} else {
				r.Config.Rename = &tfTypes.CreateRequestTransformerPluginAdd{}
				r.Config.Rename.Body = []types.String{}
				for _, v := range resp.Config.Rename.Body {
					r.Config.Rename.Body = append(r.Config.Rename.Body, types.StringValue(v))
				}
				r.Config.Rename.Headers = []types.String{}
				for _, v := range resp.Config.Rename.Headers {
					r.Config.Rename.Headers = append(r.Config.Rename.Headers, types.StringValue(v))
				}
				r.Config.Rename.Querystring = []types.String{}
				for _, v := range resp.Config.Rename.Querystring {
					r.Config.Rename.Querystring = append(r.Config.Rename.Querystring, types.StringValue(v))
				}
			}
			if resp.Config.Replace == nil {
				r.Config.Replace = nil
			} else {
				r.Config.Replace = &tfTypes.CreateRequestTransformerAdvancedPluginReplace{}
				r.Config.Replace.Body = []types.String{}
				for _, v := range resp.Config.Replace.Body {
					r.Config.Replace.Body = append(r.Config.Replace.Body, types.StringValue(v))
				}
				r.Config.Replace.Headers = []types.String{}
				for _, v := range resp.Config.Replace.Headers {
					r.Config.Replace.Headers = append(r.Config.Replace.Headers, types.StringValue(v))
				}
				r.Config.Replace.JSONTypes = []types.String{}
				for _, v := range resp.Config.Replace.JSONTypes {
					r.Config.Replace.JSONTypes = append(r.Config.Replace.JSONTypes, types.StringValue(string(v)))
				}
				r.Config.Replace.Querystring = []types.String{}
				for _, v := range resp.Config.Replace.Querystring {
					r.Config.Replace.Querystring = append(r.Config.Replace.Querystring, types.StringValue(v))
				}
				r.Config.Replace.URI = types.StringPointerValue(resp.Config.Replace.URI)
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
