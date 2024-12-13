// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreatePostfunctionPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID     string                         `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	PostFunctionPlugin shared.PostFunctionPluginInput `request:"mediaType=application/json"`
}

func (o *CreatePostfunctionPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreatePostfunctionPluginRequest) GetPostFunctionPlugin() shared.PostFunctionPluginInput {
	if o == nil {
		return shared.PostFunctionPluginInput{}
	}
	return o.PostFunctionPlugin
}

type CreatePostfunctionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created PostFunction plugin
	PostFunctionPlugin *shared.PostFunctionPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreatePostfunctionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePostfunctionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePostfunctionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePostfunctionPluginResponse) GetPostFunctionPlugin() *shared.PostFunctionPlugin {
	if o == nil {
		return nil
	}
	return o.PostFunctionPlugin
}

func (o *CreatePostfunctionPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
