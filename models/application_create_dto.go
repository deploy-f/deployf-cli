// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApplicationCreateDto application create dto
//
// swagger:model ApplicationCreateDto
type ApplicationCreateDto struct {

	// image
	Image string `json:"image,omitempty"`

	// limit CPU
	// Maximum: 3.40282346638529e+38
	// Minimum: 0.00499999988824129
	LimitCPU float32 `json:"limitCPU,omitempty"`

	// limit RAM
	// Maximum: 3.40282346638529e+38
	// Minimum: 0.00499999988824129
	LimitRAM float32 `json:"limitRAM,omitempty"`

	// name
	// Required: true
	// Max Length: 150
	Name *string `json:"name"`
}

// Validate validates this application create dto
func (m *ApplicationCreateDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimitCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitRAM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationCreateDto) validateLimitCPU(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitCPU) { // not required
		return nil
	}

	if err := validate.Minimum("limitCPU", "body", float64(m.LimitCPU), 0.00499999988824129, false); err != nil {
		return err
	}

	if err := validate.Maximum("limitCPU", "body", float64(m.LimitCPU), 3.40282346638529e+38, false); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationCreateDto) validateLimitRAM(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitRAM) { // not required
		return nil
	}

	if err := validate.Minimum("limitRAM", "body", float64(m.LimitRAM), 0.00499999988824129, false); err != nil {
		return err
	}

	if err := validate.Maximum("limitRAM", "body", float64(m.LimitRAM), 3.40282346638529e+38, false); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationCreateDto) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 150); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationCreateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationCreateDto) UnmarshalBinary(b []byte) error {
	var res ApplicationCreateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
