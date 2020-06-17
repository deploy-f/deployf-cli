// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cli/models"
)

// FindReader is a Reader for the Find structure.
type FindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFindOK creates a FindOK with default headers values
func NewFindOK() *FindOK {
	return &FindOK{}
}

/*FindOK handles this case with default header values.

Success
*/
type FindOK struct {
	Payload *models.PageDtoUserSummaryDto
}

func (o *FindOK) Error() string {
	return fmt.Sprintf("[GET /api/user/find][%d] findOK  %+v", 200, o.Payload)
}

func (o *FindOK) GetPayload() *models.PageDtoUserSummaryDto {
	return o.Payload
}

func (o *FindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageDtoUserSummaryDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUnauthorized creates a FindUnauthorized with default headers values
func NewFindUnauthorized() *FindUnauthorized {
	return &FindUnauthorized{}
}

/*FindUnauthorized handles this case with default header values.

Unauthorized
*/
type FindUnauthorized struct {
}

func (o *FindUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/user/find][%d] findUnauthorized ", 401)
}

func (o *FindUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindForbidden creates a FindForbidden with default headers values
func NewFindForbidden() *FindForbidden {
	return &FindForbidden{}
}

/*FindForbidden handles this case with default header values.

Forbidden
*/
type FindForbidden struct {
}

func (o *FindForbidden) Error() string {
	return fmt.Sprintf("[GET /api/user/find][%d] findForbidden ", 403)
}

func (o *FindForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}