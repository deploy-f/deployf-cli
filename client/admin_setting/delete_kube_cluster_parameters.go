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
	"github.com/go-openapi/swag"
)

// NewDeleteKubeClusterParams creates a new DeleteKubeClusterParams object
// with the default values initialized.
func NewDeleteKubeClusterParams() *DeleteKubeClusterParams {
	var ()
	return &DeleteKubeClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKubeClusterParamsWithTimeout creates a new DeleteKubeClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteKubeClusterParamsWithTimeout(timeout time.Duration) *DeleteKubeClusterParams {
	var ()
	return &DeleteKubeClusterParams{

		timeout: timeout,
	}
}

// NewDeleteKubeClusterParamsWithContext creates a new DeleteKubeClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteKubeClusterParamsWithContext(ctx context.Context) *DeleteKubeClusterParams {
	var ()
	return &DeleteKubeClusterParams{

		Context: ctx,
	}
}

// NewDeleteKubeClusterParamsWithHTTPClient creates a new DeleteKubeClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteKubeClusterParamsWithHTTPClient(client *http.Client) *DeleteKubeClusterParams {
	var ()
	return &DeleteKubeClusterParams{
		HTTPClient: client,
	}
}

/*DeleteKubeClusterParams contains all the parameters to send to the API endpoint
for the delete kube cluster operation typically these are written to a http.Request
*/
type DeleteKubeClusterParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete kube cluster params
func (o *DeleteKubeClusterParams) WithTimeout(timeout time.Duration) *DeleteKubeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete kube cluster params
func (o *DeleteKubeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete kube cluster params
func (o *DeleteKubeClusterParams) WithContext(ctx context.Context) *DeleteKubeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete kube cluster params
func (o *DeleteKubeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete kube cluster params
func (o *DeleteKubeClusterParams) WithHTTPClient(client *http.Client) *DeleteKubeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete kube cluster params
func (o *DeleteKubeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete kube cluster params
func (o *DeleteKubeClusterParams) WithID(id int32) *DeleteKubeClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete kube cluster params
func (o *DeleteKubeClusterParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKubeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
