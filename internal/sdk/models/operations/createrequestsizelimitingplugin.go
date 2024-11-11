// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateRequestsizelimitingPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID            string                                 `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	RequestSizeLimitingPlugin *shared.RequestSizeLimitingPluginInput `request:"mediaType=application/json"`
}

func (o *CreateRequestsizelimitingPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateRequestsizelimitingPluginRequest) GetRequestSizeLimitingPlugin() *shared.RequestSizeLimitingPluginInput {
	if o == nil {
		return nil
	}
	return o.RequestSizeLimitingPlugin
}

type CreateRequestsizelimitingPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created RequestSizeLimiting plugin
	RequestSizeLimitingPlugin *shared.RequestSizeLimitingPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateRequestsizelimitingPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRequestsizelimitingPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRequestsizelimitingPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateRequestsizelimitingPluginResponse) GetRequestSizeLimitingPlugin() *shared.RequestSizeLimitingPlugin {
	if o == nil {
		return nil
	}
	return o.RequestSizeLimitingPlugin
}

func (o *CreateRequestsizelimitingPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
