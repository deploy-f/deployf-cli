// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReadyApplicationSummaryDto ready application summary dto
//
// swagger:model ReadyApplicationSummaryDto
type ReadyApplicationSummaryDto struct {

	// id
	ID int32 `json:"id,omitempty"`

	// image Url
	ImageURL string `json:"imageUrl,omitempty"`

	// labels
	Labels []string `json:"labels"`

	// name
	Name string `json:"name,omitempty"`

	// start limit CPU
	StartLimitCPU float32 `json:"startLimitCPU,omitempty"`

	// start limit RAM
	StartLimitRAM float32 `json:"startLimitRAM,omitempty"`
}

// Validate validates this ready application summary dto
func (m *ReadyApplicationSummaryDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReadyApplicationSummaryDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReadyApplicationSummaryDto) UnmarshalBinary(b []byte) error {
	var res ReadyApplicationSummaryDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
