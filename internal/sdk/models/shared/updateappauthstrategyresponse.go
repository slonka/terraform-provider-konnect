// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"time"
)

type AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType string

const (
	AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyTypeOpenidConnect AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType = "openid_connect"
)

func (e AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType) ToPointer() *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType {
	return &e
}
func (e *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openid_connect":
		*e = AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType: %v", v)
	}
}

// AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseConfigs - JSON-B object containing the configuration for the OIDC strategy
type AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseConfigs struct {
	// A more advanced mode to configure an API Product Version’s Application Auth Strategy.
	// Using this mode will allow developers to use API credentials issued from an external IdP that will authenticate their application requests.
	// Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for the same Auth Strategy.
	// An OIDC strategy may be used in conjunction with a DCR provider to automatically create the IdP application.
	//
	OpenidConnect AppAuthStrategyConfigOpenIDConnect `json:"openid-connect"`
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseConfigs) GetOpenidConnect() AppAuthStrategyConfigOpenIDConnect {
	if o == nil {
		return AppAuthStrategyConfigOpenIDConnect{}
	}
	return o.OpenidConnect
}

type AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseDcrProvider struct {
	// Contains a unique identifier used for this resource.
	ID   string `json:"id"`
	Name string `json:"name"`
	// The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.
	//
	DisplayName *string `json:"display_name,omitempty"`
	// The type of DCR provider. Can be one of the following - auth0, azureAd, curity, okta, http
	ProviderType DcrProviderType `json:"provider_type"`
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseDcrProvider) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseDcrProvider) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseDcrProvider) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseDcrProvider) GetProviderType() DcrProviderType {
	if o == nil {
		return DcrProviderType("")
	}
	return o.ProviderType
}

// AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse - Response payload from creating an OIDC Application Auth Strategy
type AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse struct {
	// Contains a unique identifier used for this resource.
	ID string `json:"id"`
	// The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.
	//
	Name string `json:"name"`
	// The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.
	//
	DisplayName  string                                                                        `json:"display_name"`
	StrategyType AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType `json:"strategy_type"`
	// JSON-B object containing the configuration for the OIDC strategy
	Configs AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseConfigs `json:"configs"`
	// At least one published product version is using this auth strategy.
	Active      bool                                                                          `json:"active"`
	DcrProvider *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseDcrProvider `json:"dcr_provider"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (a AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) GetStrategyType() AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType {
	if o == nil {
		return AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType("")
	}
	return o.StrategyType
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) GetConfigs() AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseConfigs {
	if o == nil {
		return AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseConfigs{}
	}
	return o.Configs
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) GetDcrProvider() *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseDcrProvider {
	if o == nil {
		return nil
	}
	return o.DcrProvider
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType string

const (
	AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyTypeKeyAuth AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType = "key_auth"
)

func (e AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType) ToPointer() *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType {
	return &e
}
func (e *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "key_auth":
		*e = AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType: %v", v)
	}
}

// AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseConfigs - JSON-B object containing the configuration for the Key Auth strategy
type AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseConfigs struct {
	// The most basic mode to configure an Application Auth Strategy for an API Product Version.
	// Using this mode will allow developers to generate API keys that will authenticate their application requests.
	// Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for Key Auth.
	//
	KeyAuth AppAuthStrategyConfigKeyAuth `json:"key-auth"`
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseConfigs) GetKeyAuth() AppAuthStrategyConfigKeyAuth {
	if o == nil {
		return AppAuthStrategyConfigKeyAuth{}
	}
	return o.KeyAuth
}

type AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseDcrProvider struct {
	// Contains a unique identifier used for this resource.
	ID   string `json:"id"`
	Name string `json:"name"`
	// The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.
	//
	DisplayName *string `json:"display_name,omitempty"`
	// The type of DCR provider. Can be one of the following - auth0, azureAd, curity, okta, http
	ProviderType DcrProviderType `json:"provider_type"`
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseDcrProvider) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseDcrProvider) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseDcrProvider) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseDcrProvider) GetProviderType() DcrProviderType {
	if o == nil {
		return DcrProviderType("")
	}
	return o.ProviderType
}

// AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse - Response payload from creating or updating a Key Auth Application Auth Strategy
type AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse struct {
	// Contains a unique identifier used for this resource.
	ID string `json:"id"`
	// The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.
	//
	Name string `json:"name"`
	// The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.
	//
	DisplayName  string                                                                  `json:"display_name"`
	StrategyType AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType `json:"strategy_type"`
	// JSON-B object containing the configuration for the Key Auth strategy
	Configs AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseConfigs `json:"configs"`
	// At least one published product version is using this auth strategy.
	Active      bool                                                                    `json:"active"`
	DcrProvider *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseDcrProvider `json:"dcr_provider"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (a AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) GetStrategyType() AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType {
	if o == nil {
		return AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType("")
	}
	return o.StrategyType
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) GetConfigs() AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseConfigs {
	if o == nil {
		return AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseConfigs{}
	}
	return o.Configs
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) GetDcrProvider() *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseDcrProvider {
	if o == nil {
		return nil
	}
	return o.DcrProvider
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type UpdateAppAuthStrategyResponseType string

const (
	UpdateAppAuthStrategyResponseTypeKeyAuth       UpdateAppAuthStrategyResponseType = "key_auth"
	UpdateAppAuthStrategyResponseTypeOpenidConnect UpdateAppAuthStrategyResponseType = "openid_connect"
)

// UpdateAppAuthStrategyResponse - A set of plugin configurations that represent how the gateway will perform authentication and authorization for a Product Version. Called “Auth Strategy” for short in the context of portals/applications. The plugins are synced to any Gateway Service that is currently linked or becomes linked to the Product Version.
type UpdateAppAuthStrategyResponse struct {
	AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse             *AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse
	AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse *AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse

	Type UpdateAppAuthStrategyResponseType
}

func CreateUpdateAppAuthStrategyResponseKeyAuth(keyAuth AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse) UpdateAppAuthStrategyResponse {
	typ := UpdateAppAuthStrategyResponseTypeKeyAuth

	typStr := AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseStrategyType(typ)
	keyAuth.StrategyType = typStr

	return UpdateAppAuthStrategyResponse{
		AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse: &keyAuth,
		Type: typ,
	}
}

func CreateUpdateAppAuthStrategyResponseOpenidConnect(openidConnect AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse) UpdateAppAuthStrategyResponse {
	typ := UpdateAppAuthStrategyResponseTypeOpenidConnect

	typStr := AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseStrategyType(typ)
	openidConnect.StrategyType = typStr

	return UpdateAppAuthStrategyResponse{
		AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse: &openidConnect,
		Type: typ,
	}
}

func (u *UpdateAppAuthStrategyResponse) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		StrategyType string `json:"strategy_type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.StrategyType {
	case "key_auth":
		appAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse := new(AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse)
		if err := utils.UnmarshalJSON(data, &appAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (StrategyType == key_auth) type AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse within UpdateAppAuthStrategyResponse: %w", string(data), err)
		}

		u.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse = appAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse
		u.Type = UpdateAppAuthStrategyResponseTypeKeyAuth
		return nil
	case "openid_connect":
		appAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse := new(AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse)
		if err := utils.UnmarshalJSON(data, &appAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (StrategyType == openid_connect) type AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse within UpdateAppAuthStrategyResponse: %w", string(data), err)
		}

		u.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse = appAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse
		u.Type = UpdateAppAuthStrategyResponseTypeOpenidConnect
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateAppAuthStrategyResponse", string(data))
}

func (u UpdateAppAuthStrategyResponse) MarshalJSON() ([]byte, error) {
	if u.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse != nil {
		return utils.MarshalJSON(u.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse, "", true)
	}

	if u.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse != nil {
		return utils.MarshalJSON(u.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateAppAuthStrategyResponse: all fields are null")
}
