// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

// TLSClientCertificate - TLS Client Certificate
type TLSClientCertificate string

const (
	TLSClientCertificateRequest TLSClientCertificate = "REQUEST"
)

func (e TLSClientCertificate) ToPointer() *TLSClientCertificate {
	return &e
}
func (e *TLSClientCertificate) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "REQUEST":
		*e = TLSClientCertificate(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TLSClientCertificate: %v", v)
	}
}

type TLSHandshakeModifierPluginConfig struct {
	// TLS Client Certificate
	TLSClientCertificate *TLSClientCertificate `json:"tls_client_certificate,omitempty"`
}

func (o *TLSHandshakeModifierPluginConfig) GetTLSClientCertificate() *TLSClientCertificate {
	if o == nil {
		return nil
	}
	return o.TLSClientCertificate
}

// TLSHandshakeModifierPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type TLSHandshakeModifierPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *TLSHandshakeModifierPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type TLSHandshakeModifierPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *TLSHandshakeModifierPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type TLSHandshakeModifierPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *TLSHandshakeModifierPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type TLSHandshakeModifierPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *TLSHandshakeModifierPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type TLSHandshakeModifierPluginOrdering struct {
	After  *TLSHandshakeModifierPluginAfter  `json:"after,omitempty"`
	Before *TLSHandshakeModifierPluginBefore `json:"before,omitempty"`
}

func (o *TLSHandshakeModifierPluginOrdering) GetAfter() *TLSHandshakeModifierPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *TLSHandshakeModifierPluginOrdering) GetBefore() *TLSHandshakeModifierPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type TLSHandshakeModifierPluginProtocols string

const (
	TLSHandshakeModifierPluginProtocolsGrpc           TLSHandshakeModifierPluginProtocols = "grpc"
	TLSHandshakeModifierPluginProtocolsGrpcs          TLSHandshakeModifierPluginProtocols = "grpcs"
	TLSHandshakeModifierPluginProtocolsHTTP           TLSHandshakeModifierPluginProtocols = "http"
	TLSHandshakeModifierPluginProtocolsHTTPS          TLSHandshakeModifierPluginProtocols = "https"
	TLSHandshakeModifierPluginProtocolsTCP            TLSHandshakeModifierPluginProtocols = "tcp"
	TLSHandshakeModifierPluginProtocolsTLS            TLSHandshakeModifierPluginProtocols = "tls"
	TLSHandshakeModifierPluginProtocolsTLSPassthrough TLSHandshakeModifierPluginProtocols = "tls_passthrough"
	TLSHandshakeModifierPluginProtocolsUDP            TLSHandshakeModifierPluginProtocols = "udp"
	TLSHandshakeModifierPluginProtocolsWs             TLSHandshakeModifierPluginProtocols = "ws"
	TLSHandshakeModifierPluginProtocolsWss            TLSHandshakeModifierPluginProtocols = "wss"
)

func (e TLSHandshakeModifierPluginProtocols) ToPointer() *TLSHandshakeModifierPluginProtocols {
	return &e
}
func (e *TLSHandshakeModifierPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = TLSHandshakeModifierPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TLSHandshakeModifierPluginProtocols: %v", v)
	}
}

// TLSHandshakeModifierPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type TLSHandshakeModifierPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *TLSHandshakeModifierPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// TLSHandshakeModifierPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type TLSHandshakeModifierPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *TLSHandshakeModifierPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// TLSHandshakeModifierPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type TLSHandshakeModifierPlugin struct {
	Config TLSHandshakeModifierPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *TLSHandshakeModifierPluginConsumer      `json:"consumer"`
	ConsumerGroup *TLSHandshakeModifierPluginConsumerGroup `json:"consumer_group"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                               `json:"enabled,omitempty"`
	ID           *string                             `json:"id,omitempty"`
	InstanceName *string                             `json:"instance_name,omitempty"`
	name         string                              `const:"tls-handshake-modifier" json:"name"`
	Ordering     *TLSHandshakeModifierPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []TLSHandshakeModifierPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *TLSHandshakeModifierPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *TLSHandshakeModifierPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (t TLSHandshakeModifierPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSHandshakeModifierPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSHandshakeModifierPlugin) GetConfig() TLSHandshakeModifierPluginConfig {
	if o == nil {
		return TLSHandshakeModifierPluginConfig{}
	}
	return o.Config
}

func (o *TLSHandshakeModifierPlugin) GetConsumer() *TLSHandshakeModifierPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *TLSHandshakeModifierPlugin) GetConsumerGroup() *TLSHandshakeModifierPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *TLSHandshakeModifierPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TLSHandshakeModifierPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *TLSHandshakeModifierPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TLSHandshakeModifierPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *TLSHandshakeModifierPlugin) GetName() string {
	return "tls-handshake-modifier"
}

func (o *TLSHandshakeModifierPlugin) GetOrdering() *TLSHandshakeModifierPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *TLSHandshakeModifierPlugin) GetProtocols() []TLSHandshakeModifierPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *TLSHandshakeModifierPlugin) GetRoute() *TLSHandshakeModifierPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *TLSHandshakeModifierPlugin) GetService() *TLSHandshakeModifierPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *TLSHandshakeModifierPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *TLSHandshakeModifierPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// TLSHandshakeModifierPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type TLSHandshakeModifierPluginInput struct {
	Config TLSHandshakeModifierPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *TLSHandshakeModifierPluginConsumer      `json:"consumer"`
	ConsumerGroup *TLSHandshakeModifierPluginConsumerGroup `json:"consumer_group"`
	// Whether the plugin is applied.
	Enabled      *bool                               `json:"enabled,omitempty"`
	ID           *string                             `json:"id,omitempty"`
	InstanceName *string                             `json:"instance_name,omitempty"`
	name         string                              `const:"tls-handshake-modifier" json:"name"`
	Ordering     *TLSHandshakeModifierPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []TLSHandshakeModifierPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *TLSHandshakeModifierPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *TLSHandshakeModifierPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (t TLSHandshakeModifierPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSHandshakeModifierPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSHandshakeModifierPluginInput) GetConfig() TLSHandshakeModifierPluginConfig {
	if o == nil {
		return TLSHandshakeModifierPluginConfig{}
	}
	return o.Config
}

func (o *TLSHandshakeModifierPluginInput) GetConsumer() *TLSHandshakeModifierPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *TLSHandshakeModifierPluginInput) GetConsumerGroup() *TLSHandshakeModifierPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *TLSHandshakeModifierPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *TLSHandshakeModifierPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TLSHandshakeModifierPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *TLSHandshakeModifierPluginInput) GetName() string {
	return "tls-handshake-modifier"
}

func (o *TLSHandshakeModifierPluginInput) GetOrdering() *TLSHandshakeModifierPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *TLSHandshakeModifierPluginInput) GetProtocols() []TLSHandshakeModifierPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *TLSHandshakeModifierPluginInput) GetRoute() *TLSHandshakeModifierPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *TLSHandshakeModifierPluginInput) GetService() *TLSHandshakeModifierPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *TLSHandshakeModifierPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
