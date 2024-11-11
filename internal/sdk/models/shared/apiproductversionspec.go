// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// APIProductVersionSpec - API product version specification
type APIProductVersionSpec struct {
	// The API product version specification identifier.
	ID string `json:"id"`
	// The name of the API product version specification
	Name string `json:"name"`
	// The contents of the API product version specification
	Content string `json:"content"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (a APIProductVersionSpec) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIProductVersionSpec) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIProductVersionSpec) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *APIProductVersionSpec) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *APIProductVersionSpec) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *APIProductVersionSpec) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *APIProductVersionSpec) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
