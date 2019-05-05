// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NutAbbrevName nut abbrev name
// swagger:model NutAbbrevName
type NutAbbrevName struct {

	// name
	Name string `json:"name,omitempty"`

	// nut food id
	NutFoodID int64 `json:"nut_food_id,omitempty"`

	// shrt desc
	ShrtDesc string `json:"shrt_desc,omitempty"`
}

// Validate validates this nut abbrev name
func (m *NutAbbrevName) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NutAbbrevName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NutAbbrevName) UnmarshalBinary(b []byte) error {
	var res NutAbbrevName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}