// Code generated by go-swagger; DO NOT EDIT.

package notification

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

// NewMarkAsReadedParams creates a new MarkAsReadedParams object
// with the default values initialized.
func NewMarkAsReadedParams() *MarkAsReadedParams {
	var ()
	return &MarkAsReadedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMarkAsReadedParamsWithTimeout creates a new MarkAsReadedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMarkAsReadedParamsWithTimeout(timeout time.Duration) *MarkAsReadedParams {
	var ()
	return &MarkAsReadedParams{

		timeout: timeout,
	}
}

// NewMarkAsReadedParamsWithContext creates a new MarkAsReadedParams object
// with the default values initialized, and the ability to set a context for a request
func NewMarkAsReadedParamsWithContext(ctx context.Context) *MarkAsReadedParams {
	var ()
	return &MarkAsReadedParams{

		Context: ctx,
	}
}

// NewMarkAsReadedParamsWithHTTPClient creates a new MarkAsReadedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMarkAsReadedParamsWithHTTPClient(client *http.Client) *MarkAsReadedParams {
	var ()
	return &MarkAsReadedParams{
		HTTPClient: client,
	}
}

/*MarkAsReadedParams contains all the parameters to send to the API endpoint
for the mark as readed operation typically these are written to a http.Request
*/
type MarkAsReadedParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the mark as readed params
func (o *MarkAsReadedParams) WithTimeout(timeout time.Duration) *MarkAsReadedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mark as readed params
func (o *MarkAsReadedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mark as readed params
func (o *MarkAsReadedParams) WithContext(ctx context.Context) *MarkAsReadedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mark as readed params
func (o *MarkAsReadedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mark as readed params
func (o *MarkAsReadedParams) WithHTTPClient(client *http.Client) *MarkAsReadedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mark as readed params
func (o *MarkAsReadedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the mark as readed params
func (o *MarkAsReadedParams) WithID(id int32) *MarkAsReadedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the mark as readed params
func (o *MarkAsReadedParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MarkAsReadedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
