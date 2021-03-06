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

// GetStatsReader is a Reader for the GetStats structure.
type GetStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetStatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStatsOK creates a GetStatsOK with default headers values
func NewGetStatsOK() *GetStatsOK {
	return &GetStatsOK{}
}

/*GetStatsOK handles this case with default header values.

Success
*/
type GetStatsOK struct {
	Payload *models.AdminStatsDto
}

func (o *GetStatsOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/stats][%d] getStatsOK  %+v", 200, o.Payload)
}

func (o *GetStatsOK) GetPayload() *models.AdminStatsDto {
	return o.Payload
}

func (o *GetStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminStatsDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatsUnauthorized creates a GetStatsUnauthorized with default headers values
func NewGetStatsUnauthorized() *GetStatsUnauthorized {
	return &GetStatsUnauthorized{}
}

/*GetStatsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetStatsUnauthorized struct {
}

func (o *GetStatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/stats][%d] getStatsUnauthorized ", 401)
}

func (o *GetStatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStatsForbidden creates a GetStatsForbidden with default headers values
func NewGetStatsForbidden() *GetStatsForbidden {
	return &GetStatsForbidden{}
}

/*GetStatsForbidden handles this case with default header values.

Forbidden
*/
type GetStatsForbidden struct {
}

func (o *GetStatsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/stats][%d] getStatsForbidden ", 403)
}

func (o *GetStatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
