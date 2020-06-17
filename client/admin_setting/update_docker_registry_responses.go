// Code generated by go-swagger; DO NOT EDIT.

package admin_setting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateDockerRegistryReader is a Reader for the UpdateDockerRegistry structure.
type UpdateDockerRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDockerRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDockerRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDockerRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDockerRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDockerRegistryOK creates a UpdateDockerRegistryOK with default headers values
func NewUpdateDockerRegistryOK() *UpdateDockerRegistryOK {
	return &UpdateDockerRegistryOK{}
}

/*UpdateDockerRegistryOK handles this case with default header values.

Success
*/
type UpdateDockerRegistryOK struct {
}

func (o *UpdateDockerRegistryOK) Error() string {
	return fmt.Sprintf("[PATCH /api/admin/docker/registry][%d] updateDockerRegistryOK ", 200)
}

func (o *UpdateDockerRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDockerRegistryUnauthorized creates a UpdateDockerRegistryUnauthorized with default headers values
func NewUpdateDockerRegistryUnauthorized() *UpdateDockerRegistryUnauthorized {
	return &UpdateDockerRegistryUnauthorized{}
}

/*UpdateDockerRegistryUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateDockerRegistryUnauthorized struct {
}

func (o *UpdateDockerRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/admin/docker/registry][%d] updateDockerRegistryUnauthorized ", 401)
}

func (o *UpdateDockerRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDockerRegistryForbidden creates a UpdateDockerRegistryForbidden with default headers values
func NewUpdateDockerRegistryForbidden() *UpdateDockerRegistryForbidden {
	return &UpdateDockerRegistryForbidden{}
}

/*UpdateDockerRegistryForbidden handles this case with default header values.

Forbidden
*/
type UpdateDockerRegistryForbidden struct {
}

func (o *UpdateDockerRegistryForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/admin/docker/registry][%d] updateDockerRegistryForbidden ", 403)
}

func (o *UpdateDockerRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
