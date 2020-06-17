// Code generated by go-swagger; DO NOT EDIT.

package ready_application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cli/models"
)

// CreateReader is a Reader for the Create structure.
type CreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOK creates a CreateOK with default headers values
func NewCreateOK() *CreateOK {
	return &CreateOK{}
}

/*CreateOK handles this case with default header values.

Success
*/
type CreateOK struct {
	Payload *models.ReadyApplicationSummaryDto
}

func (o *CreateOK) Error() string {
	return fmt.Sprintf("[POST /api/ready-applications][%d] createOK  %+v", 200, o.Payload)
}

func (o *CreateOK) GetPayload() *models.ReadyApplicationSummaryDto {
	return o.Payload
}

func (o *CreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReadyApplicationSummaryDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnauthorized creates a CreateUnauthorized with default headers values
func NewCreateUnauthorized() *CreateUnauthorized {
	return &CreateUnauthorized{}
}

/*CreateUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateUnauthorized struct {
}

func (o *CreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/ready-applications][%d] createUnauthorized ", 401)
}

func (o *CreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateForbidden creates a CreateForbidden with default headers values
func NewCreateForbidden() *CreateForbidden {
	return &CreateForbidden{}
}

/*CreateForbidden handles this case with default header values.

Forbidden
*/
type CreateForbidden struct {
}

func (o *CreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/ready-applications][%d] createForbidden ", 403)
}

func (o *CreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
