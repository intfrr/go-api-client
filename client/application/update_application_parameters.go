// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gotify/server/model"
)

// NewUpdateApplicationParams creates a new UpdateApplicationParams object
// with the default values initialized.
func NewUpdateApplicationParams() *UpdateApplicationParams {
	var ()
	return &UpdateApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateApplicationParamsWithTimeout creates a new UpdateApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateApplicationParamsWithTimeout(timeout time.Duration) *UpdateApplicationParams {
	var ()
	return &UpdateApplicationParams{

		timeout: timeout,
	}
}

// NewUpdateApplicationParamsWithContext creates a new UpdateApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateApplicationParamsWithContext(ctx context.Context) *UpdateApplicationParams {
	var ()
	return &UpdateApplicationParams{

		Context: ctx,
	}
}

// NewUpdateApplicationParamsWithHTTPClient creates a new UpdateApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateApplicationParamsWithHTTPClient(client *http.Client) *UpdateApplicationParams {
	var ()
	return &UpdateApplicationParams{
		HTTPClient: client,
	}
}

/*UpdateApplicationParams contains all the parameters to send to the API endpoint
for the update application operation typically these are written to a http.Request
*/
type UpdateApplicationParams struct {

	/*Body
	  the application to update

	*/
	Body *models.Application
	/*ID
	  the application id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update application params
func (o *UpdateApplicationParams) WithTimeout(timeout time.Duration) *UpdateApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update application params
func (o *UpdateApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update application params
func (o *UpdateApplicationParams) WithContext(ctx context.Context) *UpdateApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update application params
func (o *UpdateApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update application params
func (o *UpdateApplicationParams) WithHTTPClient(client *http.Client) *UpdateApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update application params
func (o *UpdateApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update application params
func (o *UpdateApplicationParams) WithBody(body *models.Application) *UpdateApplicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update application params
func (o *UpdateApplicationParams) SetBody(body *models.Application) {
	o.Body = body
}

// WithID adds the id to the update application params
func (o *UpdateApplicationParams) WithID(id int64) *UpdateApplicationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update application params
func (o *UpdateApplicationParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
