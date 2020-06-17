// Code generated by go-swagger; DO NOT EDIT.

package admin_setting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"cli/models"
)

// NewCreateKubeClusterParams creates a new CreateKubeClusterParams object
// with the default values initialized.
func NewCreateKubeClusterParams() *CreateKubeClusterParams {
	var ()
	return &CreateKubeClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateKubeClusterParamsWithTimeout creates a new CreateKubeClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateKubeClusterParamsWithTimeout(timeout time.Duration) *CreateKubeClusterParams {
	var ()
	return &CreateKubeClusterParams{

		timeout: timeout,
	}
}

// NewCreateKubeClusterParamsWithContext creates a new CreateKubeClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateKubeClusterParamsWithContext(ctx context.Context) *CreateKubeClusterParams {
	var ()
	return &CreateKubeClusterParams{

		Context: ctx,
	}
}

// NewCreateKubeClusterParamsWithHTTPClient creates a new CreateKubeClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateKubeClusterParamsWithHTTPClient(client *http.Client) *CreateKubeClusterParams {
	var ()
	return &CreateKubeClusterParams{
		HTTPClient: client,
	}
}

/*CreateKubeClusterParams contains all the parameters to send to the API endpoint
for the create kube cluster operation typically these are written to a http.Request
*/
type CreateKubeClusterParams struct {

	/*Data*/
	Data *models.KubernetesClusterSummaryDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create kube cluster params
func (o *CreateKubeClusterParams) WithTimeout(timeout time.Duration) *CreateKubeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create kube cluster params
func (o *CreateKubeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create kube cluster params
func (o *CreateKubeClusterParams) WithContext(ctx context.Context) *CreateKubeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create kube cluster params
func (o *CreateKubeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create kube cluster params
func (o *CreateKubeClusterParams) WithHTTPClient(client *http.Client) *CreateKubeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create kube cluster params
func (o *CreateKubeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create kube cluster params
func (o *CreateKubeClusterParams) WithData(data *models.KubernetesClusterSummaryDto) *CreateKubeClusterParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create kube cluster params
func (o *CreateKubeClusterParams) SetData(data *models.KubernetesClusterSummaryDto) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CreateKubeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}