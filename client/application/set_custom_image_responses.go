// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetCustomImageReader is a Reader for the SetCustomImage structure.
type SetCustomImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetCustomImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetCustomImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetCustomImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetCustomImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetCustomImageOK creates a SetCustomImageOK with default headers values
func NewSetCustomImageOK() *SetCustomImageOK {
	return &SetCustomImageOK{}
}

/*SetCustomImageOK handles this case with default header values.

Success
*/
type SetCustomImageOK struct {
}

func (o *SetCustomImageOK) Error() string {
	return fmt.Sprintf("[POST /api/application/{id}/custom-image][%d] setCustomImageOK ", 200)
}

func (o *SetCustomImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCustomImageUnauthorized creates a SetCustomImageUnauthorized with default headers values
func NewSetCustomImageUnauthorized() *SetCustomImageUnauthorized {
	return &SetCustomImageUnauthorized{}
}

/*SetCustomImageUnauthorized handles this case with default header values.

Unauthorized
*/
type SetCustomImageUnauthorized struct {
}

func (o *SetCustomImageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/application/{id}/custom-image][%d] setCustomImageUnauthorized ", 401)
}

func (o *SetCustomImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCustomImageForbidden creates a SetCustomImageForbidden with default headers values
func NewSetCustomImageForbidden() *SetCustomImageForbidden {
	return &SetCustomImageForbidden{}
}

/*SetCustomImageForbidden handles this case with default header values.

Forbidden
*/
type SetCustomImageForbidden struct {
}

func (o *SetCustomImageForbidden) Error() string {
	return fmt.Sprintf("[POST /api/application/{id}/custom-image][%d] setCustomImageForbidden ", 403)
}

func (o *SetCustomImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}