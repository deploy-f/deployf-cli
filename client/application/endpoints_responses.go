// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cli/models"
)

// EndpointsReader is a Reader for the Endpoints structure.
type EndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEndpointsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEndpointsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEndpointsOK creates a EndpointsOK with default headers values
func NewEndpointsOK() *EndpointsOK {
	return &EndpointsOK{}
}

/*EndpointsOK handles this case with default header values.

Success
*/
type EndpointsOK struct {
	Payload *models.ApplicationEndpointDto
}

func (o *EndpointsOK) Error() string {
	return fmt.Sprintf("[GET /api/application/{id}/endpoints][%d] endpointsOK  %+v", 200, o.Payload)
}

func (o *EndpointsOK) GetPayload() *models.ApplicationEndpointDto {
	return o.Payload
}

func (o *EndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationEndpointDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointsUnauthorized creates a EndpointsUnauthorized with default headers values
func NewEndpointsUnauthorized() *EndpointsUnauthorized {
	return &EndpointsUnauthorized{}
}

/*EndpointsUnauthorized handles this case with default header values.

Unauthorized
*/
type EndpointsUnauthorized struct {
}

func (o *EndpointsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/application/{id}/endpoints][%d] endpointsUnauthorized ", 401)
}

func (o *EndpointsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointsForbidden creates a EndpointsForbidden with default headers values
func NewEndpointsForbidden() *EndpointsForbidden {
	return &EndpointsForbidden{}
}

/*EndpointsForbidden handles this case with default header values.

Forbidden
*/
type EndpointsForbidden struct {
}

func (o *EndpointsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/application/{id}/endpoints][%d] endpointsForbidden ", 403)
}

func (o *EndpointsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
