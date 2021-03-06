// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserOccupiedResourcesDto user occupied resources dto
//
// swagger:model UserOccupiedResourcesDto
type UserOccupiedResourcesDto struct {

	// occupied CPU
	OccupiedCPU float32 `json:"occupiedCPU,omitempty"`

	// occupied disk
	OccupiedDisk float32 `json:"occupiedDisk,omitempty"`

	// occupied i ps
	OccupiedIPs float32 `json:"occupiedIPs,omitempty"`

	// occupied RAM
	OccupiedRAM float32 `json:"occupiedRAM,omitempty"`
}

// Validate validates this user occupied resources dto
func (m *UserOccupiedResourcesDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserOccupiedResourcesDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserOccupiedResourcesDto) UnmarshalBinary(b []byte) error {
	var res UserOccupiedResourcesDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
