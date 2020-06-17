// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ConsistencyCheckAndResolveReader is a Reader for the ConsistencyCheckAndResolve structure.
type ConsistencyCheckAndResolveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyCheckAndResolveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsistencyCheckAndResolveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConsistencyCheckAndResolveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewConsistencyCheckAndResolveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConsistencyCheckAndResolveOK creates a ConsistencyCheckAndResolveOK with default headers values
func NewConsistencyCheckAndResolveOK() *ConsistencyCheckAndResolveOK {
	return &ConsistencyCheckAndResolveOK{}
}

/*ConsistencyCheckAndResolveOK handles this case with default header values.

Success
*/
type ConsistencyCheckAndResolveOK struct {
}

func (o *ConsistencyCheckAndResolveOK) Error() string {
	return fmt.Sprintf("[GET /api/user/consistency][%d] consistencyCheckAndResolveOK ", 200)
}

func (o *ConsistencyCheckAndResolveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsistencyCheckAndResolveUnauthorized creates a ConsistencyCheckAndResolveUnauthorized with default headers values
func NewConsistencyCheckAndResolveUnauthorized() *ConsistencyCheckAndResolveUnauthorized {
	return &ConsistencyCheckAndResolveUnauthorized{}
}

/*ConsistencyCheckAndResolveUnauthorized handles this case with default header values.

Unauthorized
*/
type ConsistencyCheckAndResolveUnauthorized struct {
}

func (o *ConsistencyCheckAndResolveUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/user/consistency][%d] consistencyCheckAndResolveUnauthorized ", 401)
}

func (o *ConsistencyCheckAndResolveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsistencyCheckAndResolveForbidden creates a ConsistencyCheckAndResolveForbidden with default headers values
func NewConsistencyCheckAndResolveForbidden() *ConsistencyCheckAndResolveForbidden {
	return &ConsistencyCheckAndResolveForbidden{}
}

/*ConsistencyCheckAndResolveForbidden handles this case with default header values.

Forbidden
*/
type ConsistencyCheckAndResolveForbidden struct {
}

func (o *ConsistencyCheckAndResolveForbidden) Error() string {
	return fmt.Sprintf("[GET /api/user/consistency][%d] consistencyCheckAndResolveForbidden ", 403)
}

func (o *ConsistencyCheckAndResolveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}