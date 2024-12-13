// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateStatsdadvancedPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID       string                           `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	StatsdAdvancedPlugin shared.StatsdAdvancedPluginInput `request:"mediaType=application/json"`
}

func (o *CreateStatsdadvancedPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateStatsdadvancedPluginRequest) GetStatsdAdvancedPlugin() shared.StatsdAdvancedPluginInput {
	if o == nil {
		return shared.StatsdAdvancedPluginInput{}
	}
	return o.StatsdAdvancedPlugin
}

type CreateStatsdadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created StatsdAdvanced plugin
	StatsdAdvancedPlugin *shared.StatsdAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateStatsdadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateStatsdadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateStatsdadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateStatsdadvancedPluginResponse) GetStatsdAdvancedPlugin() *shared.StatsdAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.StatsdAdvancedPlugin
}

func (o *CreateStatsdadvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
