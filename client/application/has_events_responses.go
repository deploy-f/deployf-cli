// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HasEventsReader is a Reader for the HasEvents structure.
type HasEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HasEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHasEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHasEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHasEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHasEventsOK creates a HasEventsOK with default headers values
func NewHasEventsOK() *HasEventsOK {
	return &HasEventsOK{}
}

/*HasEventsOK handles this case with default header values.

Success
*/
type HasEventsOK struct {
	Payload bool
}

func (o *HasEventsOK) Error() string {
	return fmt.Sprintf("[GET /api/application/{id}/has-events][%d] hasEventsOK  %+v", 200, o.Payload)
}

func (o *HasEventsOK) GetPayload() bool {
	return o.Payload
}

func (o *HasEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHasEventsUnauthorized creates a HasEventsUnauthorized with default headers values
func NewHasEventsUnauthorized() *HasEventsUnauthorized {
	return &HasEventsUnauthorized{}
}

/*HasEventsUnauthorized handles this case with default header values.

Unauthorized
*/
type HasEventsUnauthorized struct {
}

func (o *HasEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/application/{id}/has-events][%d] hasEventsUnauthorized ", 401)
}

func (o *HasEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHasEventsForbidden creates a HasEventsForbidden with default headers values
func NewHasEventsForbidden() *HasEventsForbidden {
	return &HasEventsForbidden{}
}

/*HasEventsForbidden handles this case with default header values.

Forbidden
*/
type HasEventsForbidden struct {
}

func (o *HasEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/application/{id}/has-events][%d] hasEventsForbidden ", 403)
}

func (o *HasEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}