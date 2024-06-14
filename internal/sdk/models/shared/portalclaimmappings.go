// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PortalClaimMappings - Mappings from a portal developer atribute to an Identity Provider claim.
type PortalClaimMappings struct {
	Name   *string `json:"name,omitempty"`
	Email  *string `json:"email,omitempty"`
	Groups *string `json:"groups,omitempty"`
}

func (o *PortalClaimMappings) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PortalClaimMappings) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *PortalClaimMappings) GetGroups() *string {
	if o == nil {
		return nil
	}
	return o.Groups
}
