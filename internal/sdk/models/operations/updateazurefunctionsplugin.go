// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateAzurefunctionsPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID       string                            `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	AzureFunctionsPlugin *shared.AzureFunctionsPluginInput `request:"mediaType=application/json"`
}

func (o *UpdateAzurefunctionsPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateAzurefunctionsPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateAzurefunctionsPluginRequest) GetAzureFunctionsPlugin() *shared.AzureFunctionsPluginInput {
	if o == nil {
		return nil
	}
	return o.AzureFunctionsPlugin
}

type UpdateAzurefunctionsPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// AzureFunctions plugin
	AzureFunctionsPlugin *shared.AzureFunctionsPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateAzurefunctionsPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAzurefunctionsPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAzurefunctionsPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAzurefunctionsPluginResponse) GetAzureFunctionsPlugin() *shared.AzureFunctionsPlugin {
	if o == nil {
		return nil
	}
	return o.AzureFunctionsPlugin
}

func (o *UpdateAzurefunctionsPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
