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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPublicAccountInfoParams creates a new GetPublicAccountInfoParams object
// with the default values initialized.
func NewGetPublicAccountInfoParams() *GetPublicAccountInfoParams {
	var ()
	return &GetPublicAccountInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicAccountInfoParamsWithTimeout creates a new GetPublicAccountInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicAccountInfoParamsWithTimeout(timeout time.Duration) *GetPublicAccountInfoParams {
	var ()
	return &GetPublicAccountInfoParams{

		timeout: timeout,
	}
}

// NewGetPublicAccountInfoParamsWithContext creates a new GetPublicAccountInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicAccountInfoParamsWithContext(ctx context.Context) *GetPublicAccountInfoParams {
	var ()
	return &GetPublicAccountInfoParams{

		Context: ctx,
	}
}

// NewGetPublicAccountInfoParamsWithHTTPClient creates a new GetPublicAccountInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicAccountInfoParamsWithHTTPClient(client *http.Client) *GetPublicAccountInfoParams {
	var ()
	return &GetPublicAccountInfoParams{
		HTTPClient: client,
	}
}

/*GetPublicAccountInfoParams contains all the parameters to send to the API endpoint
for the get public account info operation typically these are written to a http.Request
*/
type GetPublicAccountInfoParams struct {

	/*Handle
	  The account handle of the user to fetch.

	*/
	Handle *string
	/*UserID
	  The account's primary key of the user to fetch.

	*/
	UserID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public account info params
func (o *GetPublicAccountInfoParams) WithTimeout(timeout time.Duration) *GetPublicAccountInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public account info params
func (o *GetPublicAccountInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public account info params
func (o *GetPublicAccountInfoParams) WithContext(ctx context.Context) *GetPublicAccountInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public account info params
func (o *GetPublicAccountInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public account info params
func (o *GetPublicAccountInfoParams) WithHTTPClient(client *http.Client) *GetPublicAccountInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public account info params
func (o *GetPublicAccountInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHandle adds the handle to the get public account info params
func (o *GetPublicAccountInfoParams) WithHandle(handle *string) *GetPublicAccountInfoParams {
	o.SetHandle(handle)
	return o
}

// SetHandle adds the handle to the get public account info params
func (o *GetPublicAccountInfoParams) SetHandle(handle *string) {
	o.Handle = handle
}

// WithUserID adds the userID to the get public account info params
func (o *GetPublicAccountInfoParams) WithUserID(userID *int64) *GetPublicAccountInfoParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get public account info params
func (o *GetPublicAccountInfoParams) SetUserID(userID *int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicAccountInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Handle != nil {

		// query param handle
		var qrHandle string
		if o.Handle != nil {
			qrHandle = *o.Handle
		}
		qHandle := qrHandle
		if qHandle != "" {
			if err := r.SetQueryParam("handle", qHandle); err != nil {
				return err
			}
		}

	}

	if o.UserID != nil {

		// query param userId
		var qrUserID int64
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := swag.FormatInt64(qrUserID)
		if qUserID != "" {
			if err := r.SetQueryParam("userId", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
