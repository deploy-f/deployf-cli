// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KeyValuePairStringInt32 key value pair string int32
//
// swagger:model KeyValuePair[String,Int32]
type KeyValuePairStringInt32 struct {

	// key
	// Read Only: true
	Key string `json:"key,omitempty"`

	// value
	// Read Only: true
	Value int32 `json:"value,omitempty"`
}

// Validate validates this key value pair string int32
func (m *KeyValuePairStringInt32) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeyValuePairStringInt32) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyValuePairStringInt32) UnmarshalBinary(b []byte) error {
	var res KeyValuePairStringInt32
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}