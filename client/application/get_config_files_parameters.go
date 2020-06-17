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

// NewGetConfigFilesParams creates a new GetConfigFilesParams object
// with the default values initialized.
func NewGetConfigFilesParams() *GetConfigFilesParams {
	var ()
	return &GetConfigFilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigFilesParamsWithTimeout creates a new GetConfigFilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConfigFilesParamsWithTimeout(timeout time.Duration) *GetConfigFilesParams {
	var ()
	return &GetConfigFilesParams{

		timeout: timeout,
	}
}

// NewGetConfigFilesParamsWithContext creates a new GetConfigFilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConfigFilesParamsWithContext(ctx context.Context) *GetConfigFilesParams {
	var ()
	return &GetConfigFilesParams{

		Context: ctx,
	}
}

// NewGetConfigFilesParamsWithHTTPClient creates a new GetConfigFilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConfigFilesParamsWithHTTPClient(client *http.Client) *GetConfigFilesParams {
	var ()
	return &GetConfigFilesParams{
		HTTPClient: client,
	}
}

/*GetConfigFilesParams contains all the parameters to send to the API endpoint
for the get config files operation typically these are written to a http.Request
*/
type GetConfigFilesParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get config files params
func (o *GetConfigFilesParams) WithTimeout(timeout time.Duration) *GetConfigFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get config files params
func (o *GetConfigFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get config files params
func (o *GetConfigFilesParams) WithContext(ctx context.Context) *GetConfigFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get config files params
func (o *GetConfigFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get config files params
func (o *GetConfigFilesParams) WithHTTPClient(client *http.Client) *GetConfigFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get config files params
func (o *GetConfigFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get config files params
func (o *GetConfigFilesParams) WithID(id int32) *GetConfigFilesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get config files params
func (o *GetConfigFilesParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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