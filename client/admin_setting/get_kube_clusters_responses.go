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

// GetKubeClustersReader is a Reader for the GetKubeClusters structure.
type GetKubeClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubeClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubeClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetKubeClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubeClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKubeClustersOK creates a GetKubeClustersOK with default headers values
func NewGetKubeClustersOK() *GetKubeClustersOK {
	return &GetKubeClustersOK{}
}

/*GetKubeClustersOK handles this case with default header values.

Success
*/
type GetKubeClustersOK struct {
	Payload []*models.KubernetesClusterSummaryDto
}

func (o *GetKubeClustersOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/clusters/k8s][%d] getKubeClustersOK  %+v", 200, o.Payload)
}

func (o *GetKubeClustersOK) GetPayload() []*models.KubernetesClusterSummaryDto {
	return o.Payload
}

func (o *GetKubeClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubeClustersUnauthorized creates a GetKubeClustersUnauthorized with default headers values
func NewGetKubeClustersUnauthorized() *GetKubeClustersUnauthorized {
	return &GetKubeClustersUnauthorized{}
}

/*GetKubeClustersUnauthorized handles this case with default header values.

Unauthorized
*/
type GetKubeClustersUnauthorized struct {
}

func (o *GetKubeClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/clusters/k8s][%d] getKubeClustersUnauthorized ", 401)
}

func (o *GetKubeClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubeClustersForbidden creates a GetKubeClustersForbidden with default headers values
func NewGetKubeClustersForbidden() *GetKubeClustersForbidden {
	return &GetKubeClustersForbidden{}
}

/*GetKubeClustersForbidden handles this case with default header values.

Forbidden
*/
type GetKubeClustersForbidden struct {
}

func (o *GetKubeClustersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/clusters/k8s][%d] getKubeClustersForbidden ", 403)
}

func (o *GetKubeClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
