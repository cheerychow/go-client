// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SuggestUser suggest user
// swagger:model SuggestUser
type SuggestUser struct {

	// email
	// Pattern: ^.+@.+$
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// handle
	Handle string `json:"handle,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// source
	Source string `json:"source,omitempty"`
}

// Validate validates this suggest user
func (m *SuggestUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SuggestUser) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.Pattern("email", "body", string(m.Email), `^.+@.+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SuggestUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuggestUser) UnmarshalBinary(b []byte) error {
	var res SuggestUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}