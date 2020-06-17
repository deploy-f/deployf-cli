// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cli/models"
)

// UpdateReader is a Reader for the Update structure.
type UpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOK creates a UpdateOK with default headers values
func NewUpdateOK() *UpdateOK {
	return &UpdateOK{}
}

/*UpdateOK handles this case with default header values.

Success
*/
type UpdateOK struct {
	Payload *models.CustomDomainDto
}

func (o *UpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains][%d] updateOK  %+v", 200, o.Payload)
}

func (o *UpdateOK) GetPayload() *models.CustomDomainDto {
	return o.Payload
}

func (o *UpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomDomainDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUnauthorized creates a UpdateUnauthorized with default headers values
func NewUpdateUnauthorized() *UpdateUnauthorized {
	return &UpdateUnauthorized{}
}

/*UpdateUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateUnauthorized struct {
}

func (o *UpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains][%d] updateUnauthorized ", 401)
}

func (o *UpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateForbidden creates a UpdateForbidden with default headers values
func NewUpdateForbidden() *UpdateForbidden {
	return &UpdateForbidden{}
}

/*UpdateForbidden handles this case with default header values.

Forbidden
*/
type UpdateForbidden struct {
}

func (o *UpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains][%d] updateForbidden ", 403)
}

func (o *UpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}