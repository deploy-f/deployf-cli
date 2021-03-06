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

// NewUpdateEnvVarParams creates a new UpdateEnvVarParams object
// with the default values initialized.
func NewUpdateEnvVarParams() *UpdateEnvVarParams {
	var ()
	return &UpdateEnvVarParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEnvVarParamsWithTimeout creates a new UpdateEnvVarParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEnvVarParamsWithTimeout(timeout time.Duration) *UpdateEnvVarParams {
	var ()
	return &UpdateEnvVarParams{

		timeout: timeout,
	}
}

// NewUpdateEnvVarParamsWithContext creates a new UpdateEnvVarParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEnvVarParamsWithContext(ctx context.Context) *UpdateEnvVarParams {
	var ()
	return &UpdateEnvVarParams{

		Context: ctx,
	}
}

// NewUpdateEnvVarParamsWithHTTPClient creates a new UpdateEnvVarParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEnvVarParamsWithHTTPClient(client *http.Client) *UpdateEnvVarParams {
	var ()
	return &UpdateEnvVarParams{
		HTTPClient: client,
	}
}

/*UpdateEnvVarParams contains all the parameters to send to the API endpoint
for the update env var operation typically these are written to a http.Request
*/
type UpdateEnvVarParams struct {

	/*Data*/
	Data *models.ApplicationEnvironmentVariableDto
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update env var params
func (o *UpdateEnvVarParams) WithTimeout(timeout time.Duration) *UpdateEnvVarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update env var params
func (o *UpdateEnvVarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update env var params
func (o *UpdateEnvVarParams) WithContext(ctx context.Context) *UpdateEnvVarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update env var params
func (o *UpdateEnvVarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update env var params
func (o *UpdateEnvVarParams) WithHTTPClient(client *http.Client) *UpdateEnvVarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update env var params
func (o *UpdateEnvVarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the update env var params
func (o *UpdateEnvVarParams) WithData(data *models.ApplicationEnvironmentVariableDto) *UpdateEnvVarParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the update env var params
func (o *UpdateEnvVarParams) SetData(data *models.ApplicationEnvironmentVariableDto) {
	o.Data = data
}

// WithID adds the id to the update env var params
func (o *UpdateEnvVarParams) WithID(id int32) *UpdateEnvVarParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update env var params
func (o *UpdateEnvVarParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEnvVarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
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
