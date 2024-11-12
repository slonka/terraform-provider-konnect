// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type ClaimsToVerify string

const (
	ClaimsToVerifyExp ClaimsToVerify = "exp"
	ClaimsToVerifyNbf ClaimsToVerify = "nbf"
)

func (e ClaimsToVerify) ToPointer() *ClaimsToVerify {
	return &e
}
func (e *ClaimsToVerify) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "exp":
		fallthrough
	case "nbf":
		*e = ClaimsToVerify(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClaimsToVerify: %v", v)
	}
}

type JwtPluginConfig struct {
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.
	Anonymous *string `json:"anonymous,omitempty"`
	// A list of registered claims (according to RFC 7519) that Kong can verify as well. Accepted values: one of exp or nbf.
	ClaimsToVerify []ClaimsToVerify `json:"claims_to_verify,omitempty"`
	// A list of cookie names that Kong will inspect to retrieve JWTs.
	CookieNames []string `json:"cookie_names,omitempty"`
	// A list of HTTP header names that Kong will inspect to retrieve JWTs.
	HeaderNames []string `json:"header_names,omitempty"`
	// The name of the claim in which the key identifying the secret must be passed. The plugin will attempt to read this claim from the JWT payload and the header, in that order.
	KeyClaimName *string `json:"key_claim_name,omitempty"`
	// A value between 0 and 31536000 (365 days) limiting the lifetime of the JWT to maximum_expiration seconds in the future.
	MaximumExpiration *float64 `json:"maximum_expiration,omitempty"`
	// When authentication fails the plugin sends `WWW-Authenticate` header with `realm` attribute value.
	Realm *string `json:"realm,omitempty"`
	// A boolean value that indicates whether the plugin should run (and try to authenticate) on OPTIONS preflight requests. If set to false, then OPTIONS requests will always be allowed.
	RunOnPreflight *bool `json:"run_on_preflight,omitempty"`
	// If true, the plugin assumes the credential’s secret to be base64 encoded. You will need to create a base64-encoded secret for your Consumer, and sign your JWT with the original secret.
	SecretIsBase64 *bool `json:"secret_is_base64,omitempty"`
	// A list of querystring parameters that Kong will inspect to retrieve JWTs.
	URIParamNames []string `json:"uri_param_names,omitempty"`
}

func (o *JwtPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *JwtPluginConfig) GetClaimsToVerify() []ClaimsToVerify {
	if o == nil {
		return nil
	}
	return o.ClaimsToVerify
}

func (o *JwtPluginConfig) GetCookieNames() []string {
	if o == nil {
		return nil
	}
	return o.CookieNames
}

func (o *JwtPluginConfig) GetHeaderNames() []string {
	if o == nil {
		return nil
	}
	return o.HeaderNames
}

func (o *JwtPluginConfig) GetKeyClaimName() *string {
	if o == nil {
		return nil
	}
	return o.KeyClaimName
}

func (o *JwtPluginConfig) GetMaximumExpiration() *float64 {
	if o == nil {
		return nil
	}
	return o.MaximumExpiration
}

func (o *JwtPluginConfig) GetRealm() *string {
	if o == nil {
		return nil
	}
	return o.Realm
}

func (o *JwtPluginConfig) GetRunOnPreflight() *bool {
	if o == nil {
		return nil
	}
	return o.RunOnPreflight
}

func (o *JwtPluginConfig) GetSecretIsBase64() *bool {
	if o == nil {
		return nil
	}
	return o.SecretIsBase64
}

func (o *JwtPluginConfig) GetURIParamNames() []string {
	if o == nil {
		return nil
	}
	return o.URIParamNames
}

// JwtPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type JwtPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *JwtPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type JwtPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *JwtPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type JwtPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *JwtPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type JwtPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *JwtPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type JwtPluginOrdering struct {
	After  *JwtPluginAfter  `json:"after,omitempty"`
	Before *JwtPluginBefore `json:"before,omitempty"`
}

func (o *JwtPluginOrdering) GetAfter() *JwtPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *JwtPluginOrdering) GetBefore() *JwtPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type JwtPluginProtocols string

const (
	JwtPluginProtocolsGrpc           JwtPluginProtocols = "grpc"
	JwtPluginProtocolsGrpcs          JwtPluginProtocols = "grpcs"
	JwtPluginProtocolsHTTP           JwtPluginProtocols = "http"
	JwtPluginProtocolsHTTPS          JwtPluginProtocols = "https"
	JwtPluginProtocolsTCP            JwtPluginProtocols = "tcp"
	JwtPluginProtocolsTLS            JwtPluginProtocols = "tls"
	JwtPluginProtocolsTLSPassthrough JwtPluginProtocols = "tls_passthrough"
	JwtPluginProtocolsUDP            JwtPluginProtocols = "udp"
	JwtPluginProtocolsWs             JwtPluginProtocols = "ws"
	JwtPluginProtocolsWss            JwtPluginProtocols = "wss"
)

func (e JwtPluginProtocols) ToPointer() *JwtPluginProtocols {
	return &e
}
func (e *JwtPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = JwtPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JwtPluginProtocols: %v", v)
	}
}

// JwtPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type JwtPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *JwtPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JwtPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type JwtPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *JwtPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JwtPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type JwtPlugin struct {
	Config JwtPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *JwtPluginConsumer      `json:"consumer"`
	ConsumerGroup *JwtPluginConsumerGroup `json:"consumer_group"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool              `json:"enabled,omitempty"`
	ID           *string            `json:"id,omitempty"`
	InstanceName *string            `json:"instance_name,omitempty"`
	name         string             `const:"jwt" json:"name"`
	Ordering     *JwtPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []JwtPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *JwtPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *JwtPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (j JwtPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JwtPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JwtPlugin) GetConfig() JwtPluginConfig {
	if o == nil {
		return JwtPluginConfig{}
	}
	return o.Config
}

func (o *JwtPlugin) GetConsumer() *JwtPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *JwtPlugin) GetConsumerGroup() *JwtPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *JwtPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *JwtPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *JwtPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JwtPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *JwtPlugin) GetName() string {
	return "jwt"
}

func (o *JwtPlugin) GetOrdering() *JwtPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *JwtPlugin) GetProtocols() []JwtPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *JwtPlugin) GetRoute() *JwtPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *JwtPlugin) GetService() *JwtPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *JwtPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *JwtPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// JwtPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type JwtPluginInput struct {
	Config JwtPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *JwtPluginConsumer      `json:"consumer"`
	ConsumerGroup *JwtPluginConsumerGroup `json:"consumer_group"`
	// Whether the plugin is applied.
	Enabled      *bool              `json:"enabled,omitempty"`
	ID           *string            `json:"id,omitempty"`
	InstanceName *string            `json:"instance_name,omitempty"`
	name         string             `const:"jwt" json:"name"`
	Ordering     *JwtPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []JwtPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *JwtPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *JwtPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (j JwtPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JwtPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JwtPluginInput) GetConfig() JwtPluginConfig {
	if o == nil {
		return JwtPluginConfig{}
	}
	return o.Config
}

func (o *JwtPluginInput) GetConsumer() *JwtPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *JwtPluginInput) GetConsumerGroup() *JwtPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *JwtPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *JwtPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JwtPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *JwtPluginInput) GetName() string {
	return "jwt"
}

func (o *JwtPluginInput) GetOrdering() *JwtPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *JwtPluginInput) GetProtocols() []JwtPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *JwtPluginInput) GetRoute() *JwtPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *JwtPluginInput) GetService() *JwtPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *JwtPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
