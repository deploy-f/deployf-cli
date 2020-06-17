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

// SetEnvVarsReader is a Reader for the SetEnvVars structure.
type SetEnvVarsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetEnvVarsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetEnvVarsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetEnvVarsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetEnvVarsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetEnvVarsOK creates a SetEnvVarsOK with default headers values
func NewSetEnvVarsOK() *SetEnvVarsOK {
	return &SetEnvVarsOK{}
}

/*SetEnvVarsOK handles this case with default header values.

Success
*/
type SetEnvVarsOK struct {
	Payload []*models.ApplicationEnvironmentVariableDto
}

func (o *SetEnvVarsOK) Error() string {
	return fmt.Sprintf("[PUT /api/application/{id}/env-vars][%d] setEnvVarsOK  %+v", 200, o.Payload)
}

func (o *SetEnvVarsOK) GetPayload() []*models.ApplicationEnvironmentVariableDto {
	return o.Payload
}

func (o *SetEnvVarsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetEnvVarsUnauthorized creates a SetEnvVarsUnauthorized with default headers values
func NewSetEnvVarsUnauthorized() *SetEnvVarsUnauthorized {
	return &SetEnvVarsUnauthorized{}
}

/*SetEnvVarsUnauthorized handles this case with default header values.

Unauthorized
*/
type SetEnvVarsUnauthorized struct {
}

func (o *SetEnvVarsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/application/{id}/env-vars][%d] setEnvVarsUnauthorized ", 401)
}

func (o *SetEnvVarsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetEnvVarsForbidden creates a SetEnvVarsForbidden with default headers values
func NewSetEnvVarsForbidden() *SetEnvVarsForbidden {
	return &SetEnvVarsForbidden{}
}

/*SetEnvVarsForbidden handles this case with default header values.

Forbidden
*/
type SetEnvVarsForbidden struct {
}

func (o *SetEnvVarsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/application/{id}/env-vars][%d] setEnvVarsForbidden ", 403)
}

func (o *SetEnvVarsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
