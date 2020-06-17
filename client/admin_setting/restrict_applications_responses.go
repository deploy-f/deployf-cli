// Code generated by go-swagger; DO NOT EDIT.

package admin_setting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RestrictApplicationsReader is a Reader for the RestrictApplications structure.
type RestrictApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestrictApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestrictApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRestrictApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRestrictApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRestrictApplicationsOK creates a RestrictApplicationsOK with default headers values
func NewRestrictApplicationsOK() *RestrictApplicationsOK {
	return &RestrictApplicationsOK{}
}

/*RestrictApplicationsOK handles this case with default header values.

Success
*/
type RestrictApplicationsOK struct {
	Payload []string
}

func (o *RestrictApplicationsOK) Error() string {
	return fmt.Sprintf("[PUT /api/admin/applications/restrict][%d] restrictApplicationsOK  %+v", 200, o.Payload)
}

func (o *RestrictApplicationsOK) GetPayload() []string {
	return o.Payload
}

func (o *RestrictApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestrictApplicationsUnauthorized creates a RestrictApplicationsUnauthorized with default headers values
func NewRestrictApplicationsUnauthorized() *RestrictApplicationsUnauthorized {
	return &RestrictApplicationsUnauthorized{}
}

/*RestrictApplicationsUnauthorized handles this case with default header values.

Unauthorized
*/
type RestrictApplicationsUnauthorized struct {
}

func (o *RestrictApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/admin/applications/restrict][%d] restrictApplicationsUnauthorized ", 401)
}

func (o *RestrictApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestrictApplicationsForbidden creates a RestrictApplicationsForbidden with default headers values
func NewRestrictApplicationsForbidden() *RestrictApplicationsForbidden {
	return &RestrictApplicationsForbidden{}
}

/*RestrictApplicationsForbidden handles this case with default header values.

Forbidden
*/
type RestrictApplicationsForbidden struct {
}

func (o *RestrictApplicationsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/admin/applications/restrict][%d] restrictApplicationsForbidden ", 403)
}

func (o *RestrictApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}