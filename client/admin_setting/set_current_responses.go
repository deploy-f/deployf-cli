// Code generated by go-swagger; DO NOT EDIT.

package admin_setting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetCurrentReader is a Reader for the SetCurrent structure.
type SetCurrentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetCurrentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetCurrentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetCurrentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetCurrentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetCurrentOK creates a SetCurrentOK with default headers values
func NewSetCurrentOK() *SetCurrentOK {
	return &SetCurrentOK{}
}

/*SetCurrentOK handles this case with default header values.

Success
*/
type SetCurrentOK struct {
}

func (o *SetCurrentOK) Error() string {
	return fmt.Sprintf("[PUT /api/admin/settings][%d] setCurrentOK ", 200)
}

func (o *SetCurrentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCurrentUnauthorized creates a SetCurrentUnauthorized with default headers values
func NewSetCurrentUnauthorized() *SetCurrentUnauthorized {
	return &SetCurrentUnauthorized{}
}

/*SetCurrentUnauthorized handles this case with default header values.

Unauthorized
*/
type SetCurrentUnauthorized struct {
}

func (o *SetCurrentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/admin/settings][%d] setCurrentUnauthorized ", 401)
}

func (o *SetCurrentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCurrentForbidden creates a SetCurrentForbidden with default headers values
func NewSetCurrentForbidden() *SetCurrentForbidden {
	return &SetCurrentForbidden{}
}

/*SetCurrentForbidden handles this case with default header values.

Forbidden
*/
type SetCurrentForbidden struct {
}

func (o *SetCurrentForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/admin/settings][%d] setCurrentForbidden ", 403)
}

func (o *SetCurrentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
