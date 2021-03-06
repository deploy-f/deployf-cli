// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewStartParams creates a new StartParams object
// with the default values initialized.
func NewStartParams() *StartParams {
	var ()
	return &StartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartParamsWithTimeout creates a new StartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartParamsWithTimeout(timeout time.Duration) *StartParams {
	var ()
	return &StartParams{

		timeout: timeout,
	}
}

// NewStartParamsWithContext creates a new StartParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartParamsWithContext(ctx context.Context) *StartParams {
	var ()
	return &StartParams{

		Context: ctx,
	}
}

// NewStartParamsWithHTTPClient creates a new StartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartParamsWithHTTPClient(client *http.Client) *StartParams {
	var ()
	return &StartParams{
		HTTPClient: client,
	}
}

/*StartParams contains all the parameters to send to the API endpoint
for the start operation typically these are written to a http.Request
*/
type StartParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start params
func (o *StartParams) WithTimeout(timeout time.Duration) *StartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start params
func (o *StartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start params
func (o *StartParams) WithContext(ctx context.Context) *StartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start params
func (o *StartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start params
func (o *StartParams) WithHTTPClient(client *http.Client) *StartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start params
func (o *StartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the start params
func (o *StartParams) WithID(id int32) *StartParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the start params
func (o *StartParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
