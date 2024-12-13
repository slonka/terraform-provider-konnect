// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateGraphqlratelimitingadvancedPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID                    string                                        `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	GraphqlRateLimitingAdvancedPlugin shared.GraphqlRateLimitingAdvancedPluginInput `request:"mediaType=application/json"`
}

func (o *CreateGraphqlratelimitingadvancedPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateGraphqlratelimitingadvancedPluginRequest) GetGraphqlRateLimitingAdvancedPlugin() shared.GraphqlRateLimitingAdvancedPluginInput {
	if o == nil {
		return shared.GraphqlRateLimitingAdvancedPluginInput{}
	}
	return o.GraphqlRateLimitingAdvancedPlugin
}

type CreateGraphqlratelimitingadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created GraphqlRateLimitingAdvanced plugin
	GraphqlRateLimitingAdvancedPlugin *shared.GraphqlRateLimitingAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateGraphqlratelimitingadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateGraphqlratelimitingadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateGraphqlratelimitingadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateGraphqlratelimitingadvancedPluginResponse) GetGraphqlRateLimitingAdvancedPlugin() *shared.GraphqlRateLimitingAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.GraphqlRateLimitingAdvancedPlugin
}

func (o *CreateGraphqlratelimitingadvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
