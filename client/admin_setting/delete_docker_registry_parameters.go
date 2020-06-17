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

// NewDeleteDockerRegistryParams creates a new DeleteDockerRegistryParams object
// with the default values initialized.
func NewDeleteDockerRegistryParams() *DeleteDockerRegistryParams {
	var ()
	return &DeleteDockerRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDockerRegistryParamsWithTimeout creates a new DeleteDockerRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDockerRegistryParamsWithTimeout(timeout time.Duration) *DeleteDockerRegistryParams {
	var ()
	return &DeleteDockerRegistryParams{

		timeout: timeout,
	}
}

// NewDeleteDockerRegistryParamsWithContext creates a new DeleteDockerRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDockerRegistryParamsWithContext(ctx context.Context) *DeleteDockerRegistryParams {
	var ()
	return &DeleteDockerRegistryParams{

		Context: ctx,
	}
}

// NewDeleteDockerRegistryParamsWithHTTPClient creates a new DeleteDockerRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDockerRegistryParamsWithHTTPClient(client *http.Client) *DeleteDockerRegistryParams {
	var ()
	return &DeleteDockerRegistryParams{
		HTTPClient: client,
	}
}

/*DeleteDockerRegistryParams contains all the parameters to send to the API endpoint
for the delete docker registry operation typically these are written to a http.Request
*/
type DeleteDockerRegistryParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete docker registry params
func (o *DeleteDockerRegistryParams) WithTimeout(timeout time.Duration) *DeleteDockerRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete docker registry params
func (o *DeleteDockerRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete docker registry params
func (o *DeleteDockerRegistryParams) WithContext(ctx context.Context) *DeleteDockerRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete docker registry params
func (o *DeleteDockerRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete docker registry params
func (o *DeleteDockerRegistryParams) WithHTTPClient(client *http.Client) *DeleteDockerRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete docker registry params
func (o *DeleteDockerRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete docker registry params
func (o *DeleteDockerRegistryParams) WithID(id int32) *DeleteDockerRegistryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete docker registry params
func (o *DeleteDockerRegistryParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDockerRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
