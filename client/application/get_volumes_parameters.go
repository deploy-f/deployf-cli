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

// NewGetVolumesParams creates a new GetVolumesParams object
// with the default values initialized.
func NewGetVolumesParams() *GetVolumesParams {
	var ()
	return &GetVolumesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVolumesParamsWithTimeout creates a new GetVolumesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVolumesParamsWithTimeout(timeout time.Duration) *GetVolumesParams {
	var ()
	return &GetVolumesParams{

		timeout: timeout,
	}
}

// NewGetVolumesParamsWithContext creates a new GetVolumesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVolumesParamsWithContext(ctx context.Context) *GetVolumesParams {
	var ()
	return &GetVolumesParams{

		Context: ctx,
	}
}

// NewGetVolumesParamsWithHTTPClient creates a new GetVolumesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVolumesParamsWithHTTPClient(client *http.Client) *GetVolumesParams {
	var ()
	return &GetVolumesParams{
		HTTPClient: client,
	}
}

/*GetVolumesParams contains all the parameters to send to the API endpoint
for the get volumes operation typically these are written to a http.Request
*/
type GetVolumesParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get volumes params
func (o *GetVolumesParams) WithTimeout(timeout time.Duration) *GetVolumesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get volumes params
func (o *GetVolumesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get volumes params
func (o *GetVolumesParams) WithContext(ctx context.Context) *GetVolumesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get volumes params
func (o *GetVolumesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get volumes params
func (o *GetVolumesParams) WithHTTPClient(client *http.Client) *GetVolumesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get volumes params
func (o *GetVolumesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get volumes params
func (o *GetVolumesParams) WithID(id int32) *GetVolumesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get volumes params
func (o *GetVolumesParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVolumesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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