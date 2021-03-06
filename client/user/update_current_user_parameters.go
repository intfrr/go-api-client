// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gotify/server/model"
)

// NewUpdateCurrentUserParams creates a new UpdateCurrentUserParams object
// with the default values initialized.
func NewUpdateCurrentUserParams() *UpdateCurrentUserParams {
	var ()
	return &UpdateCurrentUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCurrentUserParamsWithTimeout creates a new UpdateCurrentUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCurrentUserParamsWithTimeout(timeout time.Duration) *UpdateCurrentUserParams {
	var ()
	return &UpdateCurrentUserParams{

		timeout: timeout,
	}
}

// NewUpdateCurrentUserParamsWithContext creates a new UpdateCurrentUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCurrentUserParamsWithContext(ctx context.Context) *UpdateCurrentUserParams {
	var ()
	return &UpdateCurrentUserParams{

		Context: ctx,
	}
}

// NewUpdateCurrentUserParamsWithHTTPClient creates a new UpdateCurrentUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCurrentUserParamsWithHTTPClient(client *http.Client) *UpdateCurrentUserParams {
	var ()
	return &UpdateCurrentUserParams{
		HTTPClient: client,
	}
}

/*UpdateCurrentUserParams contains all the parameters to send to the API endpoint
for the update current user operation typically these are written to a http.Request
*/
type UpdateCurrentUserParams struct {

	/*Body
	  the user

	*/
	Body *models.UserExternalPass

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update current user params
func (o *UpdateCurrentUserParams) WithTimeout(timeout time.Duration) *UpdateCurrentUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update current user params
func (o *UpdateCurrentUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update current user params
func (o *UpdateCurrentUserParams) WithContext(ctx context.Context) *UpdateCurrentUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update current user params
func (o *UpdateCurrentUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update current user params
func (o *UpdateCurrentUserParams) WithHTTPClient(client *http.Client) *UpdateCurrentUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update current user params
func (o *UpdateCurrentUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update current user params
func (o *UpdateCurrentUserParams) WithBody(body *models.UserExternalPass) *UpdateCurrentUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update current user params
func (o *UpdateCurrentUserParams) SetBody(body *models.UserExternalPass) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCurrentUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
