// Code generated by go-swagger; DO NOT EDIT.

package promocode

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

// NewActivateParams creates a new ActivateParams object
// with the default values initialized.
func NewActivateParams() *ActivateParams {
	var ()
	return &ActivateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActivateParamsWithTimeout creates a new ActivateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActivateParamsWithTimeout(timeout time.Duration) *ActivateParams {
	var ()
	return &ActivateParams{

		timeout: timeout,
	}
}

// NewActivateParamsWithContext creates a new ActivateParams object
// with the default values initialized, and the ability to set a context for a request
func NewActivateParamsWithContext(ctx context.Context) *ActivateParams {
	var ()
	return &ActivateParams{

		Context: ctx,
	}
}

// NewActivateParamsWithHTTPClient creates a new ActivateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActivateParamsWithHTTPClient(client *http.Client) *ActivateParams {
	var ()
	return &ActivateParams{
		HTTPClient: client,
	}
}

/*ActivateParams contains all the parameters to send to the API endpoint
for the activate operation typically these are written to a http.Request
*/
type ActivateParams struct {

	/*Code*/
	Code string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the activate params
func (o *ActivateParams) WithTimeout(timeout time.Duration) *ActivateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the activate params
func (o *ActivateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the activate params
func (o *ActivateParams) WithContext(ctx context.Context) *ActivateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the activate params
func (o *ActivateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the activate params
func (o *ActivateParams) WithHTTPClient(client *http.Client) *ActivateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the activate params
func (o *ActivateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCode adds the code to the activate params
func (o *ActivateParams) WithCode(code string) *ActivateParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the activate params
func (o *ActivateParams) SetCode(code string) {
	o.Code = code
}

// WriteToRequest writes these params to a swagger request
func (o *ActivateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param code
	if err := r.SetPathParam("code", o.Code); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}