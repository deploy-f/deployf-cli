// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cli/models"
)

// GetOccupiedResourcesReader is a Reader for the GetOccupiedResources structure.
type GetOccupiedResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOccupiedResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOccupiedResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOccupiedResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOccupiedResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOccupiedResourcesOK creates a GetOccupiedResourcesOK with default headers values
func NewGetOccupiedResourcesOK() *GetOccupiedResourcesOK {
	return &GetOccupiedResourcesOK{}
}

/*GetOccupiedResourcesOK handles this case with default header values.

Success
*/
type GetOccupiedResourcesOK struct {
	Payload *models.UserOccupiedResourcesDto
}

func (o *GetOccupiedResourcesOK) Error() string {
	return fmt.Sprintf("[GET /api/resources/occupied-resources/{user}][%d] getOccupiedResourcesOK  %+v", 200, o.Payload)
}

func (o *GetOccupiedResourcesOK) GetPayload() *models.UserOccupiedResourcesDto {
	return o.Payload
}

func (o *GetOccupiedResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserOccupiedResourcesDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOccupiedResourcesUnauthorized creates a GetOccupiedResourcesUnauthorized with default headers values
func NewGetOccupiedResourcesUnauthorized() *GetOccupiedResourcesUnauthorized {
	return &GetOccupiedResourcesUnauthorized{}
}

/*GetOccupiedResourcesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetOccupiedResourcesUnauthorized struct {
}

func (o *GetOccupiedResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/resources/occupied-resources/{user}][%d] getOccupiedResourcesUnauthorized ", 401)
}

func (o *GetOccupiedResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOccupiedResourcesForbidden creates a GetOccupiedResourcesForbidden with default headers values
func NewGetOccupiedResourcesForbidden() *GetOccupiedResourcesForbidden {
	return &GetOccupiedResourcesForbidden{}
}

/*GetOccupiedResourcesForbidden handles this case with default header values.

Forbidden
*/
type GetOccupiedResourcesForbidden struct {
}

func (o *GetOccupiedResourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/resources/occupied-resources/{user}][%d] getOccupiedResourcesForbidden ", 403)
}

func (o *GetOccupiedResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}