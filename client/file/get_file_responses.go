// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetFileReader is a Reader for the GetFile structure.
type GetFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFileOK creates a GetFileOK with default headers values
func NewGetFileOK() *GetFileOK {
	return &GetFileOK{}
}

/*GetFileOK handles this case with default header values.

Success
*/
type GetFileOK struct {
}

func (o *GetFileOK) Error() string {
	return fmt.Sprintf("[GET /api/file/{fid}][%d] getFileOK ", 200)
}

func (o *GetFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFileUnauthorized creates a GetFileUnauthorized with default headers values
func NewGetFileUnauthorized() *GetFileUnauthorized {
	return &GetFileUnauthorized{}
}

/*GetFileUnauthorized handles this case with default header values.

Unauthorized
*/
type GetFileUnauthorized struct {
}

func (o *GetFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/file/{fid}][%d] getFileUnauthorized ", 401)
}

func (o *GetFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFileForbidden creates a GetFileForbidden with default headers values
func NewGetFileForbidden() *GetFileForbidden {
	return &GetFileForbidden{}
}

/*GetFileForbidden handles this case with default header values.

Forbidden
*/
type GetFileForbidden struct {
}

func (o *GetFileForbidden) Error() string {
	return fmt.Sprintf("[GET /api/file/{fid}][%d] getFileForbidden ", 403)
}

func (o *GetFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}