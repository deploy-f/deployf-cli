// Code generated by go-swagger; DO NOT EDIT.

package admin_setting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cli/models"
)

// GetCurrentReader is a Reader for the GetCurrent structure.
type GetCurrentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCurrentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCurrentOK creates a GetCurrentOK with default headers values
func NewGetCurrentOK() *GetCurrentOK {
	return &GetCurrentOK{}
}

/*GetCurrentOK handles this case with default header values.

Success
*/
type GetCurrentOK struct {
	Payload *models.ServiceParametersModelDto
}

func (o *GetCurrentOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/settings][%d] getCurrentOK  %+v", 200, o.Payload)
}

func (o *GetCurrentOK) GetPayload() *models.ServiceParametersModelDto {
	return o.Payload
}

func (o *GetCurrentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceParametersModelDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUnauthorized creates a GetCurrentUnauthorized with default headers values
func NewGetCurrentUnauthorized() *GetCurrentUnauthorized {
	return &GetCurrentUnauthorized{}
}

/*GetCurrentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCurrentUnauthorized struct {
}

func (o *GetCurrentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/settings][%d] getCurrentUnauthorized ", 401)
}

func (o *GetCurrentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentForbidden creates a GetCurrentForbidden with default headers values
func NewGetCurrentForbidden() *GetCurrentForbidden {
	return &GetCurrentForbidden{}
}

/*GetCurrentForbidden handles this case with default header values.

Forbidden
*/
type GetCurrentForbidden struct {
}

func (o *GetCurrentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/settings][%d] getCurrentForbidden ", 403)
}

func (o *GetCurrentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}