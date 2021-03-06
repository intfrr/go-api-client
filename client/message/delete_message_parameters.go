// Code generated by go-swagger; DO NOT EDIT.

package message

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
)

// NewDeleteMessageParams creates a new DeleteMessageParams object
// with the default values initialized.
func NewDeleteMessageParams() *DeleteMessageParams {
	var ()
	return &DeleteMessageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMessageParamsWithTimeout creates a new DeleteMessageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMessageParamsWithTimeout(timeout time.Duration) *DeleteMessageParams {
	var ()
	return &DeleteMessageParams{

		timeout: timeout,
	}
}

// NewDeleteMessageParamsWithContext creates a new DeleteMessageParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMessageParamsWithContext(ctx context.Context) *DeleteMessageParams {
	var ()
	return &DeleteMessageParams{

		Context: ctx,
	}
}

// NewDeleteMessageParamsWithHTTPClient creates a new DeleteMessageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMessageParamsWithHTTPClient(client *http.Client) *DeleteMessageParams {
	var ()
	return &DeleteMessageParams{
		HTTPClient: client,
	}
}

/*DeleteMessageParams contains all the parameters to send to the API endpoint
for the delete message operation typically these are written to a http.Request
*/
type DeleteMessageParams struct {

	/*ID
	  the message id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete message params
func (o *DeleteMessageParams) WithTimeout(timeout time.Duration) *DeleteMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete message params
func (o *DeleteMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete message params
func (o *DeleteMessageParams) WithContext(ctx context.Context) *DeleteMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete message params
func (o *DeleteMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete message params
func (o *DeleteMessageParams) WithHTTPClient(client *http.Client) *DeleteMessageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete message params
func (o *DeleteMessageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete message params
func (o *DeleteMessageParams) WithID(id int64) *DeleteMessageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete message params
func (o *DeleteMessageParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
