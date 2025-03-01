// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginResponseTransformerResourceModel) ToSharedResponseTransformerPluginInput() *shared.ResponseTransformerPluginInput {
	var add *shared.ResponseTransformerPluginAdd
	if r.Config.Add != nil {
		var headers []string = []string{}
		for _, headersItem := range r.Config.Add.Headers {
			headers = append(headers, headersItem.ValueString())
		}
		var json []string = []string{}
		for _, jsonItem := range r.Config.Add.JSON {
			json = append(json, jsonItem.ValueString())
		}
		var jsonTypes []shared.ResponseTransformerPluginJSONTypes = []shared.ResponseTransformerPluginJSONTypes{}
		for _, jsonTypesItem := range r.Config.Add.JSONTypes {
			jsonTypes = append(jsonTypes, shared.ResponseTransformerPluginJSONTypes(jsonTypesItem.ValueString()))
		}
		add = &shared.ResponseTransformerPluginAdd{
			Headers:   headers,
			JSON:      json,
			JSONTypes: jsonTypes,
		}
	}
	var append1 *shared.ResponseTransformerPluginAppend
	if r.Config.Append != nil {
		var headers1 []string = []string{}
		for _, headersItem1 := range r.Config.Append.Headers {
			headers1 = append(headers1, headersItem1.ValueString())
		}
		var json1 []string = []string{}
		for _, jsonItem1 := range r.Config.Append.JSON {
			json1 = append(json1, jsonItem1.ValueString())
		}
		var jsonTypes1 []shared.ResponseTransformerPluginConfigJSONTypes = []shared.ResponseTransformerPluginConfigJSONTypes{}
		for _, jsonTypesItem1 := range r.Config.Append.JSONTypes {
			jsonTypes1 = append(jsonTypes1, shared.ResponseTransformerPluginConfigJSONTypes(jsonTypesItem1.ValueString()))
		}
		append1 = &shared.ResponseTransformerPluginAppend{
			Headers:   headers1,
			JSON:      json1,
			JSONTypes: jsonTypes1,
		}
	}
	var remove *shared.ResponseTransformerPluginRemove
	if r.Config.Remove != nil {
		var headers2 []string = []string{}
		for _, headersItem2 := range r.Config.Remove.Headers {
			headers2 = append(headers2, headersItem2.ValueString())
		}
		var json2 []string = []string{}
		for _, jsonItem2 := range r.Config.Remove.JSON {
			json2 = append(json2, jsonItem2.ValueString())
		}
		remove = &shared.ResponseTransformerPluginRemove{
			Headers: headers2,
			JSON:    json2,
		}
	}
	var rename *shared.ResponseTransformerPluginRename
	if r.Config.Rename != nil {
		var headers3 []string = []string{}
		for _, headersItem3 := range r.Config.Rename.Headers {
			headers3 = append(headers3, headersItem3.ValueString())
		}
		var json3 []string = []string{}
		for _, jsonItem3 := range r.Config.Rename.JSON {
			json3 = append(json3, jsonItem3.ValueString())
		}
		rename = &shared.ResponseTransformerPluginRename{
			Headers: headers3,
			JSON:    json3,
		}
	}
	var replace *shared.ResponseTransformerPluginReplace
	if r.Config.Replace != nil {
		var headers4 []string = []string{}
		for _, headersItem4 := range r.Config.Replace.Headers {
			headers4 = append(headers4, headersItem4.ValueString())
		}
		var json4 []string = []string{}
		for _, jsonItem4 := range r.Config.Replace.JSON {
			json4 = append(json4, jsonItem4.ValueString())
		}
		var jsonTypes2 []shared.ResponseTransformerPluginConfigReplaceJSONTypes = []shared.ResponseTransformerPluginConfigReplaceJSONTypes{}
		for _, jsonTypesItem2 := range r.Config.Replace.JSONTypes {
			jsonTypes2 = append(jsonTypes2, shared.ResponseTransformerPluginConfigReplaceJSONTypes(jsonTypesItem2.ValueString()))
		}
		replace = &shared.ResponseTransformerPluginReplace{
			Headers:   headers4,
			JSON:      json4,
			JSONTypes: jsonTypes2,
		}
	}
	config := shared.ResponseTransformerPluginConfig{
		Add:     add,
		Append:  append1,
		Remove:  remove,
		Rename:  rename,
		Replace: replace,
	}
	var consumer *shared.ResponseTransformerPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.ResponseTransformerPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.ResponseTransformerPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.ResponseTransformerPluginConsumerGroup{
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
	var ordering *shared.ResponseTransformerPluginOrdering
	if r.Ordering != nil {
		var after *shared.ResponseTransformerPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.ResponseTransformerPluginAfter{
				Access: access,
			}
		}
		var before *shared.ResponseTransformerPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.ResponseTransformerPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.ResponseTransformerPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.ResponseTransformerPluginProtocols = []shared.ResponseTransformerPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.ResponseTransformerPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.ResponseTransformerPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.ResponseTransformerPluginRoute{
			ID: id3,
		}
	}
	var service *shared.ResponseTransformerPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.ResponseTransformerPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.ResponseTransformerPluginInput{
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

func (r *GatewayPluginResponseTransformerResourceModel) RefreshFromSharedResponseTransformerPlugin(resp *shared.ResponseTransformerPlugin) {
	if resp != nil {
		if resp.Config.Add == nil {
			r.Config.Add = nil
		} else {
			r.Config.Add = &tfTypes.ResponseTransformerPluginAdd{}
			r.Config.Add.Headers = []types.String{}
			for _, v := range resp.Config.Add.Headers {
				r.Config.Add.Headers = append(r.Config.Add.Headers, types.StringValue(v))
			}
			r.Config.Add.JSON = []types.String{}
			for _, v := range resp.Config.Add.JSON {
				r.Config.Add.JSON = append(r.Config.Add.JSON, types.StringValue(v))
			}
			r.Config.Add.JSONTypes = []types.String{}
			for _, v := range resp.Config.Add.JSONTypes {
				r.Config.Add.JSONTypes = append(r.Config.Add.JSONTypes, types.StringValue(string(v)))
			}
		}
		if resp.Config.Append == nil {
			r.Config.Append = nil
		} else {
			r.Config.Append = &tfTypes.ResponseTransformerPluginAdd{}
			r.Config.Append.Headers = []types.String{}
			for _, v := range resp.Config.Append.Headers {
				r.Config.Append.Headers = append(r.Config.Append.Headers, types.StringValue(v))
			}
			r.Config.Append.JSON = []types.String{}
			for _, v := range resp.Config.Append.JSON {
				r.Config.Append.JSON = append(r.Config.Append.JSON, types.StringValue(v))
			}
			r.Config.Append.JSONTypes = []types.String{}
			for _, v := range resp.Config.Append.JSONTypes {
				r.Config.Append.JSONTypes = append(r.Config.Append.JSONTypes, types.StringValue(string(v)))
			}
		}
		if resp.Config.Remove == nil {
			r.Config.Remove = nil
		} else {
			r.Config.Remove = &tfTypes.ResponseTransformerPluginRemove{}
			r.Config.Remove.Headers = []types.String{}
			for _, v := range resp.Config.Remove.Headers {
				r.Config.Remove.Headers = append(r.Config.Remove.Headers, types.StringValue(v))
			}
			r.Config.Remove.JSON = []types.String{}
			for _, v := range resp.Config.Remove.JSON {
				r.Config.Remove.JSON = append(r.Config.Remove.JSON, types.StringValue(v))
			}
		}
		if resp.Config.Rename == nil {
			r.Config.Rename = nil
		} else {
			r.Config.Rename = &tfTypes.ResponseTransformerPluginRemove{}
			r.Config.Rename.Headers = []types.String{}
			for _, v := range resp.Config.Rename.Headers {
				r.Config.Rename.Headers = append(r.Config.Rename.Headers, types.StringValue(v))
			}
			r.Config.Rename.JSON = []types.String{}
			for _, v := range resp.Config.Rename.JSON {
				r.Config.Rename.JSON = append(r.Config.Rename.JSON, types.StringValue(v))
			}
		}
		if resp.Config.Replace == nil {
			r.Config.Replace = nil
		} else {
			r.Config.Replace = &tfTypes.ResponseTransformerPluginAdd{}
			r.Config.Replace.Headers = []types.String{}
			for _, v := range resp.Config.Replace.Headers {
				r.Config.Replace.Headers = append(r.Config.Replace.Headers, types.StringValue(v))
			}
			r.Config.Replace.JSON = []types.String{}
			for _, v := range resp.Config.Replace.JSON {
				r.Config.Replace.JSON = append(r.Config.Replace.JSON, types.StringValue(v))
			}
			r.Config.Replace.JSONTypes = []types.String{}
			for _, v := range resp.Config.Replace.JSONTypes {
				r.Config.Replace.JSONTypes = append(r.Config.Replace.JSONTypes, types.StringValue(string(v)))
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
