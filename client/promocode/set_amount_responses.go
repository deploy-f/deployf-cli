// Code generated by go-swagger; DO NOT EDIT.

package promocode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetAmountReader is a Reader for the SetAmount structure.
type SetAmountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAmountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetAmountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetAmountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetAmountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAmountOK creates a SetAmountOK with default headers values
func NewSetAmountOK() *SetAmountOK {
	return &SetAmountOK{}
}

/*SetAmountOK handles this case with default header values.

Success
*/
type SetAmountOK struct {
}

func (o *SetAmountOK) Error() string {
	return fmt.Sprintf("[PUT /api/promocode/set-amount/{code}][%d] setAmountOK ", 200)
}

func (o *SetAmountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetAmountUnauthorized creates a SetAmountUnauthorized with default headers values
func NewSetAmountUnauthorized() *SetAmountUnauthorized {
	return &SetAmountUnauthorized{}
}

/*SetAmountUnauthorized handles this case with default header values.

Unauthorized
*/
type SetAmountUnauthorized struct {
}

func (o *SetAmountUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/promocode/set-amount/{code}][%d] setAmountUnauthorized ", 401)
}

func (o *SetAmountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetAmountForbidden creates a SetAmountForbidden with default headers values
func NewSetAmountForbidden() *SetAmountForbidden {
	return &SetAmountForbidden{}
}

/*SetAmountForbidden handles this case with default header values.

Forbidden
*/
type SetAmountForbidden struct {
}

func (o *SetAmountForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/promocode/set-amount/{code}][%d] setAmountForbidden ", 403)
}

func (o *SetAmountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}