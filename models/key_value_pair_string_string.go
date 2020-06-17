// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KeyValuePairStringString key value pair string string
//
// swagger:model KeyValuePair[String,String]
type KeyValuePairStringString struct {

	// key
	// Read Only: true
	Key string `json:"key,omitempty"`

	// value
	// Read Only: true
	Value string `json:"value,omitempty"`
}

// Validate validates this key value pair string string
func (m *KeyValuePairStringString) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeyValuePairStringString) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyValuePairStringString) UnmarshalBinary(b []byte) error {
	var res KeyValuePairStringString
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}