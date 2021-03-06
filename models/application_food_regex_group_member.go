// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ApplicationFoodRegexGroupMember application food regex group member
// swagger:model ApplicationFoodRegexGroupMember
type ApplicationFoodRegexGroupMember struct {

	// application food id
	ApplicationFoodID int64 `json:"application_food_id,omitempty"`

	// is regex
	IsRegex bool `json:"is_regex,omitempty"`

	// member group id
	MemberGroupID int64 `json:"member_group_id,omitempty"`

	// regex group id
	RegexGroupID int64 `json:"regex_group_id,omitempty"`
}

// Validate validates this application food regex group member
func (m *ApplicationFoodRegexGroupMember) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationFoodRegexGroupMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationFoodRegexGroupMember) UnmarshalBinary(b []byte) error {
	var res ApplicationFoodRegexGroupMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
