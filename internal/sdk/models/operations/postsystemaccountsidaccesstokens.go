// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

var PostSystemAccountsIDAccessTokensServerList = []string{
	"https://global.api.konghq.com/",
}

type PostSystemAccountsIDAccessTokensRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// The request body to create a system account access token.
	CreateSystemAccountAccessToken *shared.CreateSystemAccountAccessToken `request:"mediaType=application/json"`
}

func (o *PostSystemAccountsIDAccessTokensRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *PostSystemAccountsIDAccessTokensRequest) GetCreateSystemAccountAccessToken() *shared.CreateSystemAccountAccessToken {
	if o == nil {
		return nil
	}
	return o.CreateSystemAccountAccessToken
}

type PostSystemAccountsIDAccessTokensResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response including a single system account access token with the token.
	SystemAccountAccessTokenCreated *shared.SystemAccountAccessTokenCreated
	// Unauthenticated
	UnauthorizedError *shared.UnauthorizedError
	// Not Found
	NotFoundError *shared.NotFoundError
	// Conflict
	ConflictError *shared.ConflictError
}

func (o *PostSystemAccountsIDAccessTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostSystemAccountsIDAccessTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostSystemAccountsIDAccessTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostSystemAccountsIDAccessTokensResponse) GetSystemAccountAccessTokenCreated() *shared.SystemAccountAccessTokenCreated {
	if o == nil {
		return nil
	}
	return o.SystemAccountAccessTokenCreated
}

func (o *PostSystemAccountsIDAccessTokensResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *PostSystemAccountsIDAccessTokensResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *PostSystemAccountsIDAccessTokensResponse) GetConflictError() *shared.ConflictError {
	if o == nil {
		return nil
	}
	return o.ConflictError
}
