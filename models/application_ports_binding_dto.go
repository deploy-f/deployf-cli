// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationPortsBindingDto application ports binding dto
//
// swagger:model ApplicationPortsBindingDto
type ApplicationPortsBindingDto struct {

	// application Id
	ApplicationID int32 `json:"applicationId,omitempty"`

	// bindings
	Bindings []*ApplicationPortBindingDto `json:"bindings"`
}

// Validate validates this application ports binding dto
func (m *ApplicationPortsBindingDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationPortsBindingDto) validateBindings(formats strfmt.Registry) error {

	if swag.IsZero(m.Bindings) { // not required
		return nil
	}

	for i := 0; i < len(m.Bindings); i++ {
		if swag.IsZero(m.Bindings[i]) { // not required
			continue
		}

		if m.Bindings[i] != nil {
			if err := m.Bindings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bindings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationPortsBindingDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationPortsBindingDto) UnmarshalBinary(b []byte) error {
	var res ApplicationPortsBindingDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}