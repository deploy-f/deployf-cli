// Code generated by go-swagger; DO NOT EDIT.

package home

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
)

// NewGetHealthParams creates a new GetHealthParams object
// with the default values initialized.
func NewGetHealthParams() *GetHealthParams {

	return &GetHealthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHealthParamsWithTimeout creates a new GetHealthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHealthParamsWithTimeout(timeout time.Duration) *GetHealthParams {

	return &GetHealthParams{

		timeout: timeout,
	}
}

// NewGetHealthParamsWithContext creates a new GetHealthParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHealthParamsWithContext(ctx context.Context) *GetHealthParams {

	return &GetHealthParams{

		Context: ctx,
	}
}

// NewGetHealthParamsWithHTTPClient creates a new GetHealthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHealthParamsWithHTTPClient(client *http.Client) *GetHealthParams {

	return &GetHealthParams{
		HTTPClient: client,
	}
}

/*GetHealthParams contains all the parameters to send to the API endpoint
for the get health operation typically these are written to a http.Request
*/
type GetHealthParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get health params
func (o *GetHealthParams) WithTimeout(timeout time.Duration) *GetHealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get health params
func (o *GetHealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get health params
func (o *GetHealthParams) WithContext(ctx context.Context) *GetHealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get health params
func (o *GetHealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get health params
func (o *GetHealthParams) WithHTTPClient(client *http.Client) *GetHealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get health params
func (o *GetHealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
