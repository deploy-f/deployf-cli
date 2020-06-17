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

// DeleteVolumeReader is a Reader for the DeleteVolume structure.
type DeleteVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteVolumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteVolumeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteVolumeOK creates a DeleteVolumeOK with default headers values
func NewDeleteVolumeOK() *DeleteVolumeOK {
	return &DeleteVolumeOK{}
}

/*DeleteVolumeOK handles this case with default header values.

Success
*/
type DeleteVolumeOK struct {
	Payload bool
}

func (o *DeleteVolumeOK) Error() string {
	return fmt.Sprintf("[DELETE /api/application/{id}/volume/{volumeId}][%d] deleteVolumeOK  %+v", 200, o.Payload)
}

func (o *DeleteVolumeOK) GetPayload() bool {
	return o.Payload
}

func (o *DeleteVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeUnauthorized creates a DeleteVolumeUnauthorized with default headers values
func NewDeleteVolumeUnauthorized() *DeleteVolumeUnauthorized {
	return &DeleteVolumeUnauthorized{}
}

/*DeleteVolumeUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteVolumeUnauthorized struct {
}

func (o *DeleteVolumeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/application/{id}/volume/{volumeId}][%d] deleteVolumeUnauthorized ", 401)
}

func (o *DeleteVolumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVolumeForbidden creates a DeleteVolumeForbidden with default headers values
func NewDeleteVolumeForbidden() *DeleteVolumeForbidden {
	return &DeleteVolumeForbidden{}
}

/*DeleteVolumeForbidden handles this case with default header values.

Forbidden
*/
type DeleteVolumeForbidden struct {
}

func (o *DeleteVolumeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/application/{id}/volume/{volumeId}][%d] deleteVolumeForbidden ", 403)
}

func (o *DeleteVolumeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
