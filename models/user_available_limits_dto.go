// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserAvailableLimitsDto user available limits dto
//
// swagger:model UserAvailableLimitsDto
type UserAvailableLimitsDto struct {

	// available CPU
	AvailableCPU float32 `json:"availableCPU,omitempty"`

	// available disk
	AvailableDisk float32 `json:"availableDisk,omitempty"`

	// available IP
	AvailableIP float32 `json:"availableIP,omitempty"`

	// available RAM
	AvailableRAM float32 `json:"availableRAM,omitempty"`
}

// Validate validates this user available limits dto
func (m *UserAvailableLimitsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserAvailableLimitsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAvailableLimitsDto) UnmarshalBinary(b []byte) error {
	var res UserAvailableLimitsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}