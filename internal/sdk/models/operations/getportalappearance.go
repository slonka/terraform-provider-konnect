// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetPortalAppearanceRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

func (o *GetPortalAppearanceRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

type GetPortalAppearanceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Details about a portal's current appearance settings.
	GetPortalAppearanceResponse *shared.GetPortalAppearanceResponse
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetPortalAppearanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPortalAppearanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPortalAppearanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPortalAppearanceResponse) GetGetPortalAppearanceResponse() *shared.GetPortalAppearanceResponse {
	if o == nil {
		return nil
	}
	return o.GetPortalAppearanceResponse
}

func (o *GetPortalAppearanceResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *GetPortalAppearanceResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *GetPortalAppearanceResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
