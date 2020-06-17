// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cli/models"
)

// GetPageReader is a Reader for the GetPage structure.
type GetPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPageOK creates a GetPageOK with default headers values
func NewGetPageOK() *GetPageOK {
	return &GetPageOK{}
}

/*GetPageOK handles this case with default header values.

Success
*/
type GetPageOK struct {
	Payload *models.PageDtoBillingAuditEventDto
}

func (o *GetPageOK) Error() string {
	return fmt.Sprintf("[GET /api/billing/{userId}/audit][%d] getPageOK  %+v", 200, o.Payload)
}

func (o *GetPageOK) GetPayload() *models.PageDtoBillingAuditEventDto {
	return o.Payload
}

func (o *GetPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageDtoBillingAuditEventDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPageUnauthorized creates a GetPageUnauthorized with default headers values
func NewGetPageUnauthorized() *GetPageUnauthorized {
	return &GetPageUnauthorized{}
}

/*GetPageUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPageUnauthorized struct {
}

func (o *GetPageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/billing/{userId}/audit][%d] getPageUnauthorized ", 401)
}

func (o *GetPageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPageForbidden creates a GetPageForbidden with default headers values
func NewGetPageForbidden() *GetPageForbidden {
	return &GetPageForbidden{}
}

/*GetPageForbidden handles this case with default header values.

Forbidden
*/
type GetPageForbidden struct {
}

func (o *GetPageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/billing/{userId}/audit][%d] getPageForbidden ", 403)
}

func (o *GetPageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}