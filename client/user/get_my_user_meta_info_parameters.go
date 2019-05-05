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
)

// NewGetMyUserMetaInfoParams creates a new GetMyUserMetaInfoParams object
// with the default values initialized.
func NewGetMyUserMetaInfoParams() *GetMyUserMetaInfoParams {

	return &GetMyUserMetaInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMyUserMetaInfoParamsWithTimeout creates a new GetMyUserMetaInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMyUserMetaInfoParamsWithTimeout(timeout time.Duration) *GetMyUserMetaInfoParams {

	return &GetMyUserMetaInfoParams{

		timeout: timeout,
	}
}

// NewGetMyUserMetaInfoParamsWithContext creates a new GetMyUserMetaInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMyUserMetaInfoParamsWithContext(ctx context.Context) *GetMyUserMetaInfoParams {

	return &GetMyUserMetaInfoParams{

		Context: ctx,
	}
}

// NewGetMyUserMetaInfoParamsWithHTTPClient creates a new GetMyUserMetaInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMyUserMetaInfoParamsWithHTTPClient(client *http.Client) *GetMyUserMetaInfoParams {

	return &GetMyUserMetaInfoParams{
		HTTPClient: client,
	}
}

/*GetMyUserMetaInfoParams contains all the parameters to send to the API endpoint
for the get my user meta info operation typically these are written to a http.Request
*/
type GetMyUserMetaInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get my user meta info params
func (o *GetMyUserMetaInfoParams) WithTimeout(timeout time.Duration) *GetMyUserMetaInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get my user meta info params
func (o *GetMyUserMetaInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get my user meta info params
func (o *GetMyUserMetaInfoParams) WithContext(ctx context.Context) *GetMyUserMetaInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get my user meta info params
func (o *GetMyUserMetaInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get my user meta info params
func (o *GetMyUserMetaInfoParams) WithHTTPClient(client *http.Client) *GetMyUserMetaInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get my user meta info params
func (o *GetMyUserMetaInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMyUserMetaInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
