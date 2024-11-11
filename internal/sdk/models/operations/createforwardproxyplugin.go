// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateForwardproxyPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID     string                          `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	ForwardProxyPlugin *shared.ForwardProxyPluginInput `request:"mediaType=application/json"`
}

func (o *CreateForwardproxyPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateForwardproxyPluginRequest) GetForwardProxyPlugin() *shared.ForwardProxyPluginInput {
	if o == nil {
		return nil
	}
	return o.ForwardProxyPlugin
}

type CreateForwardproxyPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created ForwardProxy plugin
	ForwardProxyPlugin *shared.ForwardProxyPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateForwardproxyPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateForwardproxyPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateForwardproxyPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateForwardproxyPluginResponse) GetForwardProxyPlugin() *shared.ForwardProxyPlugin {
	if o == nil {
		return nil
	}
	return o.ForwardProxyPlugin
}

func (o *CreateForwardproxyPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
