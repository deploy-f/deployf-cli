// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewFindParams creates a new FindParams object
// with the default values initialized.
func NewFindParams() *FindParams {
	var ()
	return &FindParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindParamsWithTimeout creates a new FindParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindParamsWithTimeout(timeout time.Duration) *FindParams {
	var ()
	return &FindParams{

		timeout: timeout,
	}
}

// NewFindParamsWithContext creates a new FindParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindParamsWithContext(ctx context.Context) *FindParams {
	var ()
	return &FindParams{

		Context: ctx,
	}
}

// NewFindParamsWithHTTPClient creates a new FindParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindParamsWithHTTPClient(client *http.Client) *FindParams {
	var ()
	return &FindParams{
		HTTPClient: client,
	}
}

/*FindParams contains all the parameters to send to the API endpoint
for the find operation typically these are written to a http.Request
*/
type FindParams struct {

	/*Count*/
	Count *int32
	/*Name*/
	Name *string
	/*Page*/
	Page *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find params
func (o *FindParams) WithTimeout(timeout time.Duration) *FindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find params
func (o *FindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find params
func (o *FindParams) WithContext(ctx context.Context) *FindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find params
func (o *FindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find params
func (o *FindParams) WithHTTPClient(client *http.Client) *FindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find params
func (o *FindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the find params
func (o *FindParams) WithCount(count *int32) *FindParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the find params
func (o *FindParams) SetCount(count *int32) {
	o.Count = count
}

// WithName adds the name to the find params
func (o *FindParams) WithName(name *string) *FindParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find params
func (o *FindParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the find params
func (o *FindParams) WithPage(page *int32) *FindParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the find params
func (o *FindParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *FindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param Count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("Count", qCount); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param Name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("Name", qName); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param Page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("Page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}