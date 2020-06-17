// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DockerRegistryAuthDto docker registry auth dto
//
// swagger:model DockerRegistryAuthDto
type DockerRegistryAuthDto struct {

	// address
	Address string `json:"address,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// image prefix
	ImagePrefix string `json:"imagePrefix,omitempty"`

	// pass
	Pass string `json:"pass,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this docker registry auth dto
func (m *DockerRegistryAuthDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DockerRegistryAuthDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DockerRegistryAuthDto) UnmarshalBinary(b []byte) error {
	var res DockerRegistryAuthDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}