// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateResponseTransformerPluginJSONTypes string

const (
	CreateResponseTransformerPluginJSONTypesBoolean CreateResponseTransformerPluginJSONTypes = "boolean"
	CreateResponseTransformerPluginJSONTypesNumber  CreateResponseTransformerPluginJSONTypes = "number"
	CreateResponseTransformerPluginJSONTypesString  CreateResponseTransformerPluginJSONTypes = "string"
)

func (e CreateResponseTransformerPluginJSONTypes) ToPointer() *CreateResponseTransformerPluginJSONTypes {
	return &e
}
func (e *CreateResponseTransformerPluginJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = CreateResponseTransformerPluginJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseTransformerPluginJSONTypes: %v", v)
	}
}

type CreateResponseTransformerPluginAdd struct {
	Headers []string `json:"headers,omitempty"`
	JSON    []string `json:"json,omitempty"`
	// List of JSON type names. Specify the types of the JSON values returned when appending
	// JSON properties. Each string element can be one of: boolean, number, or string.
	JSONTypes []CreateResponseTransformerPluginJSONTypes `json:"json_types,omitempty"`
}

func (o *CreateResponseTransformerPluginAdd) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateResponseTransformerPluginAdd) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

func (o *CreateResponseTransformerPluginAdd) GetJSONTypes() []CreateResponseTransformerPluginJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

type CreateResponseTransformerPluginConfigJSONTypes string

const (
	CreateResponseTransformerPluginConfigJSONTypesBoolean CreateResponseTransformerPluginConfigJSONTypes = "boolean"
	CreateResponseTransformerPluginConfigJSONTypesNumber  CreateResponseTransformerPluginConfigJSONTypes = "number"
	CreateResponseTransformerPluginConfigJSONTypesString  CreateResponseTransformerPluginConfigJSONTypes = "string"
)

func (e CreateResponseTransformerPluginConfigJSONTypes) ToPointer() *CreateResponseTransformerPluginConfigJSONTypes {
	return &e
}
func (e *CreateResponseTransformerPluginConfigJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = CreateResponseTransformerPluginConfigJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseTransformerPluginConfigJSONTypes: %v", v)
	}
}

type CreateResponseTransformerPluginAppend struct {
	Headers []string `json:"headers,omitempty"`
	JSON    []string `json:"json,omitempty"`
	// List of JSON type names. Specify the types of the JSON values returned when appending
	// JSON properties. Each string element can be one of: boolean, number, or string.
	JSONTypes []CreateResponseTransformerPluginConfigJSONTypes `json:"json_types,omitempty"`
}

func (o *CreateResponseTransformerPluginAppend) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateResponseTransformerPluginAppend) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

func (o *CreateResponseTransformerPluginAppend) GetJSONTypes() []CreateResponseTransformerPluginConfigJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

type CreateResponseTransformerPluginRemove struct {
	Headers []string `json:"headers,omitempty"`
	JSON    []string `json:"json,omitempty"`
}

func (o *CreateResponseTransformerPluginRemove) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateResponseTransformerPluginRemove) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

type CreateResponseTransformerPluginRename struct {
	Headers []string `json:"headers,omitempty"`
}

func (o *CreateResponseTransformerPluginRename) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

type CreateResponseTransformerPluginConfigReplaceJSONTypes string

const (
	CreateResponseTransformerPluginConfigReplaceJSONTypesBoolean CreateResponseTransformerPluginConfigReplaceJSONTypes = "boolean"
	CreateResponseTransformerPluginConfigReplaceJSONTypesNumber  CreateResponseTransformerPluginConfigReplaceJSONTypes = "number"
	CreateResponseTransformerPluginConfigReplaceJSONTypesString  CreateResponseTransformerPluginConfigReplaceJSONTypes = "string"
)

func (e CreateResponseTransformerPluginConfigReplaceJSONTypes) ToPointer() *CreateResponseTransformerPluginConfigReplaceJSONTypes {
	return &e
}
func (e *CreateResponseTransformerPluginConfigReplaceJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = CreateResponseTransformerPluginConfigReplaceJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseTransformerPluginConfigReplaceJSONTypes: %v", v)
	}
}

type CreateResponseTransformerPluginReplace struct {
	Headers []string `json:"headers,omitempty"`
	JSON    []string `json:"json,omitempty"`
	// List of JSON type names. Specify the types of the JSON values returned when appending
	// JSON properties. Each string element can be one of: boolean, number, or string.
	JSONTypes []CreateResponseTransformerPluginConfigReplaceJSONTypes `json:"json_types,omitempty"`
}

func (o *CreateResponseTransformerPluginReplace) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateResponseTransformerPluginReplace) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

func (o *CreateResponseTransformerPluginReplace) GetJSONTypes() []CreateResponseTransformerPluginConfigReplaceJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

type CreateResponseTransformerPluginConfig struct {
	Add     *CreateResponseTransformerPluginAdd     `json:"add,omitempty"`
	Append  *CreateResponseTransformerPluginAppend  `json:"append,omitempty"`
	Remove  *CreateResponseTransformerPluginRemove  `json:"remove,omitempty"`
	Rename  *CreateResponseTransformerPluginRename  `json:"rename,omitempty"`
	Replace *CreateResponseTransformerPluginReplace `json:"replace,omitempty"`
}

func (o *CreateResponseTransformerPluginConfig) GetAdd() *CreateResponseTransformerPluginAdd {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *CreateResponseTransformerPluginConfig) GetAppend() *CreateResponseTransformerPluginAppend {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *CreateResponseTransformerPluginConfig) GetRemove() *CreateResponseTransformerPluginRemove {
	if o == nil {
		return nil
	}
	return o.Remove
}

func (o *CreateResponseTransformerPluginConfig) GetRename() *CreateResponseTransformerPluginRename {
	if o == nil {
		return nil
	}
	return o.Rename
}

func (o *CreateResponseTransformerPluginConfig) GetReplace() *CreateResponseTransformerPluginReplace {
	if o == nil {
		return nil
	}
	return o.Replace
}

type CreateResponseTransformerPluginProtocols string

const (
	CreateResponseTransformerPluginProtocolsGrpc           CreateResponseTransformerPluginProtocols = "grpc"
	CreateResponseTransformerPluginProtocolsGrpcs          CreateResponseTransformerPluginProtocols = "grpcs"
	CreateResponseTransformerPluginProtocolsHTTP           CreateResponseTransformerPluginProtocols = "http"
	CreateResponseTransformerPluginProtocolsHTTPS          CreateResponseTransformerPluginProtocols = "https"
	CreateResponseTransformerPluginProtocolsTCP            CreateResponseTransformerPluginProtocols = "tcp"
	CreateResponseTransformerPluginProtocolsTLS            CreateResponseTransformerPluginProtocols = "tls"
	CreateResponseTransformerPluginProtocolsTLSPassthrough CreateResponseTransformerPluginProtocols = "tls_passthrough"
	CreateResponseTransformerPluginProtocolsUDP            CreateResponseTransformerPluginProtocols = "udp"
	CreateResponseTransformerPluginProtocolsWs             CreateResponseTransformerPluginProtocols = "ws"
	CreateResponseTransformerPluginProtocolsWss            CreateResponseTransformerPluginProtocols = "wss"
)

func (e CreateResponseTransformerPluginProtocols) ToPointer() *CreateResponseTransformerPluginProtocols {
	return &e
}
func (e *CreateResponseTransformerPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = CreateResponseTransformerPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseTransformerPluginProtocols: %v", v)
	}
}

// CreateResponseTransformerPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateResponseTransformerPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseTransformerPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateResponseTransformerPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseTransformerPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateResponseTransformerPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateResponseTransformerPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseTransformerPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateResponseTransformerPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateResponseTransformerPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseTransformerPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateResponseTransformerPlugin struct {
	Config *CreateResponseTransformerPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"response-transformer" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateResponseTransformerPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateResponseTransformerPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateResponseTransformerPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateResponseTransformerPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateResponseTransformerPluginService `json:"service,omitempty"`
}

func (c CreateResponseTransformerPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateResponseTransformerPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateResponseTransformerPlugin) GetConfig() *CreateResponseTransformerPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateResponseTransformerPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateResponseTransformerPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateResponseTransformerPlugin) GetName() *string {
	return types.String("response-transformer")
}

func (o *CreateResponseTransformerPlugin) GetProtocols() []CreateResponseTransformerPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateResponseTransformerPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateResponseTransformerPlugin) GetConsumer() *CreateResponseTransformerPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateResponseTransformerPlugin) GetConsumerGroup() *CreateResponseTransformerPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateResponseTransformerPlugin) GetRoute() *CreateResponseTransformerPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateResponseTransformerPlugin) GetService() *CreateResponseTransformerPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
