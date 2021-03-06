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

	"cli/models"
)

// NewSetLimitsParams creates a new SetLimitsParams object
// with the default values initialized.
func NewSetLimitsParams() *SetLimitsParams {
	var ()
	return &SetLimitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetLimitsParamsWithTimeout creates a new SetLimitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetLimitsParamsWithTimeout(timeout time.Duration) *SetLimitsParams {
	var ()
	return &SetLimitsParams{

		timeout: timeout,
	}
}

// NewSetLimitsParamsWithContext creates a new SetLimitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetLimitsParamsWithContext(ctx context.Context) *SetLimitsParams {
	var ()
	return &SetLimitsParams{

		Context: ctx,
	}
}

// NewSetLimitsParamsWithHTTPClient creates a new SetLimitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetLimitsParamsWithHTTPClient(client *http.Client) *SetLimitsParams {
	var ()
	return &SetLimitsParams{
		HTTPClient: client,
	}
}

/*SetLimitsParams contains all the parameters to send to the API endpoint
for the set limits operation typically these are written to a http.Request
*/
type SetLimitsParams struct {

	/*ID*/
	ID int32
	/*LimitsDto*/
	LimitsDto *models.ApplicationLimitsDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set limits params
func (o *SetLimitsParams) WithTimeout(timeout time.Duration) *SetLimitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set limits params
func (o *SetLimitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set limits params
func (o *SetLimitsParams) WithContext(ctx context.Context) *SetLimitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set limits params
func (o *SetLimitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set limits params
func (o *SetLimitsParams) WithHTTPClient(client *http.Client) *SetLimitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set limits params
func (o *SetLimitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the set limits params
func (o *SetLimitsParams) WithID(id int32) *SetLimitsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the set limits params
func (o *SetLimitsParams) SetID(id int32) {
	o.ID = id
}

// WithLimitsDto adds the limitsDto to the set limits params
func (o *SetLimitsParams) WithLimitsDto(limitsDto *models.ApplicationLimitsDto) *SetLimitsParams {
	o.SetLimitsDto(limitsDto)
	return o
}

// SetLimitsDto adds the limitsDto to the set limits params
func (o *SetLimitsParams) SetLimitsDto(limitsDto *models.ApplicationLimitsDto) {
	o.LimitsDto = limitsDto
}

// WriteToRequest writes these params to a swagger request
func (o *SetLimitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.LimitsDto != nil {
		if err := r.SetBodyParam(o.LimitsDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
