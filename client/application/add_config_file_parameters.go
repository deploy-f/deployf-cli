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

// NewAddConfigFileParams creates a new AddConfigFileParams object
// with the default values initialized.
func NewAddConfigFileParams() *AddConfigFileParams {
	var ()
	return &AddConfigFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddConfigFileParamsWithTimeout creates a new AddConfigFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddConfigFileParamsWithTimeout(timeout time.Duration) *AddConfigFileParams {
	var ()
	return &AddConfigFileParams{

		timeout: timeout,
	}
}

// NewAddConfigFileParamsWithContext creates a new AddConfigFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddConfigFileParamsWithContext(ctx context.Context) *AddConfigFileParams {
	var ()
	return &AddConfigFileParams{

		Context: ctx,
	}
}

// NewAddConfigFileParamsWithHTTPClient creates a new AddConfigFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddConfigFileParamsWithHTTPClient(client *http.Client) *AddConfigFileParams {
	var ()
	return &AddConfigFileParams{
		HTTPClient: client,
	}
}

/*AddConfigFileParams contains all the parameters to send to the API endpoint
for the add config file operation typically these are written to a http.Request
*/
type AddConfigFileParams struct {

	/*File*/
	File *models.ApplicationConfigFileDto
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add config file params
func (o *AddConfigFileParams) WithTimeout(timeout time.Duration) *AddConfigFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add config file params
func (o *AddConfigFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add config file params
func (o *AddConfigFileParams) WithContext(ctx context.Context) *AddConfigFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add config file params
func (o *AddConfigFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add config file params
func (o *AddConfigFileParams) WithHTTPClient(client *http.Client) *AddConfigFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add config file params
func (o *AddConfigFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the add config file params
func (o *AddConfigFileParams) WithFile(file *models.ApplicationConfigFileDto) *AddConfigFileParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the add config file params
func (o *AddConfigFileParams) SetFile(file *models.ApplicationConfigFileDto) {
	o.File = file
}

// WithID adds the id to the add config file params
func (o *AddConfigFileParams) WithID(id int32) *AddConfigFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add config file params
func (o *AddConfigFileParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddConfigFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.File != nil {
		if err := r.SetBodyParam(o.File); err != nil {
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
