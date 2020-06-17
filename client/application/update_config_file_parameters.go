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

// NewUpdateConfigFileParams creates a new UpdateConfigFileParams object
// with the default values initialized.
func NewUpdateConfigFileParams() *UpdateConfigFileParams {
	var ()
	return &UpdateConfigFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigFileParamsWithTimeout creates a new UpdateConfigFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateConfigFileParamsWithTimeout(timeout time.Duration) *UpdateConfigFileParams {
	var ()
	return &UpdateConfigFileParams{

		timeout: timeout,
	}
}

// NewUpdateConfigFileParamsWithContext creates a new UpdateConfigFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateConfigFileParamsWithContext(ctx context.Context) *UpdateConfigFileParams {
	var ()
	return &UpdateConfigFileParams{

		Context: ctx,
	}
}

// NewUpdateConfigFileParamsWithHTTPClient creates a new UpdateConfigFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateConfigFileParamsWithHTTPClient(client *http.Client) *UpdateConfigFileParams {
	var ()
	return &UpdateConfigFileParams{
		HTTPClient: client,
	}
}

/*UpdateConfigFileParams contains all the parameters to send to the API endpoint
for the update config file operation typically these are written to a http.Request
*/
type UpdateConfigFileParams struct {

	/*Data*/
	Data *models.ApplicationConfigFileDto
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update config file params
func (o *UpdateConfigFileParams) WithTimeout(timeout time.Duration) *UpdateConfigFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update config file params
func (o *UpdateConfigFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update config file params
func (o *UpdateConfigFileParams) WithContext(ctx context.Context) *UpdateConfigFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update config file params
func (o *UpdateConfigFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update config file params
func (o *UpdateConfigFileParams) WithHTTPClient(client *http.Client) *UpdateConfigFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update config file params
func (o *UpdateConfigFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the update config file params
func (o *UpdateConfigFileParams) WithData(data *models.ApplicationConfigFileDto) *UpdateConfigFileParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the update config file params
func (o *UpdateConfigFileParams) SetData(data *models.ApplicationConfigFileDto) {
	o.Data = data
}

// WithID adds the id to the update config file params
func (o *UpdateConfigFileParams) WithID(id int32) *UpdateConfigFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update config file params
func (o *UpdateConfigFileParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
