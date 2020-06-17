// Code generated by go-swagger; DO NOT EDIT.

package domains

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

// NewAddTLSCetrificateParams creates a new AddTLSCetrificateParams object
// with the default values initialized.
func NewAddTLSCetrificateParams() *AddTLSCetrificateParams {
	var ()
	return &AddTLSCetrificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddTLSCetrificateParamsWithTimeout creates a new AddTLSCetrificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddTLSCetrificateParamsWithTimeout(timeout time.Duration) *AddTLSCetrificateParams {
	var ()
	return &AddTLSCetrificateParams{

		timeout: timeout,
	}
}

// NewAddTLSCetrificateParamsWithContext creates a new AddTLSCetrificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddTLSCetrificateParamsWithContext(ctx context.Context) *AddTLSCetrificateParams {
	var ()
	return &AddTLSCetrificateParams{

		Context: ctx,
	}
}

// NewAddTLSCetrificateParamsWithHTTPClient creates a new AddTLSCetrificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddTLSCetrificateParamsWithHTTPClient(client *http.Client) *AddTLSCetrificateParams {
	var ()
	return &AddTLSCetrificateParams{
		HTTPClient: client,
	}
}

/*AddTLSCetrificateParams contains all the parameters to send to the API endpoint
for the add Tls cetrificate operation typically these are written to a http.Request
*/
type AddTLSCetrificateParams struct {

	/*File*/
	File runtime.NamedReadCloser
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) WithTimeout(timeout time.Duration) *AddTLSCetrificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) WithContext(ctx context.Context) *AddTLSCetrificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) WithHTTPClient(client *http.Client) *AddTLSCetrificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) WithFile(file runtime.NamedReadCloser) *AddTLSCetrificateParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithID adds the id to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) WithID(id int32) *AddTLSCetrificateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add Tls cetrificate params
func (o *AddTLSCetrificateParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddTLSCetrificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.File != nil {

		if o.File != nil {

			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}

		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}