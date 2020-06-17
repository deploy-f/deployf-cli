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

// CreateDockerRegistryReader is a Reader for the CreateDockerRegistry structure.
type CreateDockerRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDockerRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDockerRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateDockerRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDockerRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDockerRegistryOK creates a CreateDockerRegistryOK with default headers values
func NewCreateDockerRegistryOK() *CreateDockerRegistryOK {
	return &CreateDockerRegistryOK{}
}

/*CreateDockerRegistryOK handles this case with default header values.

Success
*/
type CreateDockerRegistryOK struct {
	Payload *models.DockerRegistryAuthDto
}

func (o *CreateDockerRegistryOK) Error() string {
	return fmt.Sprintf("[POST /api/admin/docker/registry][%d] createDockerRegistryOK  %+v", 200, o.Payload)
}

func (o *CreateDockerRegistryOK) GetPayload() *models.DockerRegistryAuthDto {
	return o.Payload
}

func (o *CreateDockerRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DockerRegistryAuthDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDockerRegistryUnauthorized creates a CreateDockerRegistryUnauthorized with default headers values
func NewCreateDockerRegistryUnauthorized() *CreateDockerRegistryUnauthorized {
	return &CreateDockerRegistryUnauthorized{}
}

/*CreateDockerRegistryUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateDockerRegistryUnauthorized struct {
}

func (o *CreateDockerRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/admin/docker/registry][%d] createDockerRegistryUnauthorized ", 401)
}

func (o *CreateDockerRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDockerRegistryForbidden creates a CreateDockerRegistryForbidden with default headers values
func NewCreateDockerRegistryForbidden() *CreateDockerRegistryForbidden {
	return &CreateDockerRegistryForbidden{}
}

/*CreateDockerRegistryForbidden handles this case with default header values.

Forbidden
*/
type CreateDockerRegistryForbidden struct {
}

func (o *CreateDockerRegistryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/admin/docker/registry][%d] createDockerRegistryForbidden ", 403)
}

func (o *CreateDockerRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}