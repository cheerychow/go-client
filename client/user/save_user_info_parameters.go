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

	models "github.com/cheerychow/chow-go-client/models"
)

// NewSaveUserInfoParams creates a new SaveUserInfoParams object
// with the default values initialized.
func NewSaveUserInfoParams() *SaveUserInfoParams {
	var ()
	return &SaveUserInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSaveUserInfoParamsWithTimeout creates a new SaveUserInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSaveUserInfoParamsWithTimeout(timeout time.Duration) *SaveUserInfoParams {
	var ()
	return &SaveUserInfoParams{

		timeout: timeout,
	}
}

// NewSaveUserInfoParamsWithContext creates a new SaveUserInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewSaveUserInfoParamsWithContext(ctx context.Context) *SaveUserInfoParams {
	var ()
	return &SaveUserInfoParams{

		Context: ctx,
	}
}

// NewSaveUserInfoParamsWithHTTPClient creates a new SaveUserInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSaveUserInfoParamsWithHTTPClient(client *http.Client) *SaveUserInfoParams {
	var ()
	return &SaveUserInfoParams{
		HTTPClient: client,
	}
}

/*SaveUserInfoParams contains all the parameters to send to the API endpoint
for the save user info operation typically these are written to a http.Request
*/
type SaveUserInfoParams struct {

	/*JSONBody
	  A user object containing the changes to be made.

	*/
	JSONBody *models.SuggestUser
	/*UserID
	  The ID of the user to which is to be changed.

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the save user info params
func (o *SaveUserInfoParams) WithTimeout(timeout time.Duration) *SaveUserInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save user info params
func (o *SaveUserInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save user info params
func (o *SaveUserInfoParams) WithContext(ctx context.Context) *SaveUserInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save user info params
func (o *SaveUserInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save user info params
func (o *SaveUserInfoParams) WithHTTPClient(client *http.Client) *SaveUserInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save user info params
func (o *SaveUserInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the save user info params
func (o *SaveUserInfoParams) WithJSONBody(jSONBody *models.SuggestUser) *SaveUserInfoParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the save user info params
func (o *SaveUserInfoParams) SetJSONBody(jSONBody *models.SuggestUser) {
	o.JSONBody = jSONBody
}

// WithUserID adds the userID to the save user info params
func (o *SaveUserInfoParams) WithUserID(userID int64) *SaveUserInfoParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the save user info params
func (o *SaveUserInfoParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SaveUserInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	// path param userId
	if err := r.SetPathParam("userId", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
