// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetKonnectapplicationauthPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetKonnectapplicationauthPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *GetKonnectapplicationauthPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetKonnectapplicationauthPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// KonnectApplicationAuth plugin
	KonnectApplicationAuthPlugin *shared.KonnectApplicationAuthPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetKonnectapplicationauthPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKonnectapplicationauthPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKonnectapplicationauthPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetKonnectapplicationauthPluginResponse) GetKonnectApplicationAuthPlugin() *shared.KonnectApplicationAuthPlugin {
	if o == nil {
		return nil
	}
	return o.KonnectApplicationAuthPlugin
}

func (o *GetKonnectapplicationauthPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
