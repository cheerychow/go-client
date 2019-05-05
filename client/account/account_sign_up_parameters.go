// Code generated by go-swagger; DO NOT EDIT.

package account

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

	models "github.com/cheerychow/chow-go-client/models"
)

// NewAccountSignUpParams creates a new AccountSignUpParams object
// with the default values initialized.
func NewAccountSignUpParams() *AccountSignUpParams {
	var ()
	return &AccountSignUpParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountSignUpParamsWithTimeout creates a new AccountSignUpParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountSignUpParamsWithTimeout(timeout time.Duration) *AccountSignUpParams {
	var ()
	return &AccountSignUpParams{

		timeout: timeout,
	}
}

// NewAccountSignUpParamsWithContext creates a new AccountSignUpParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountSignUpParamsWithContext(ctx context.Context) *AccountSignUpParams {
	var ()
	return &AccountSignUpParams{

		Context: ctx,
	}
}

// NewAccountSignUpParamsWithHTTPClient creates a new AccountSignUpParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountSignUpParamsWithHTTPClient(client *http.Client) *AccountSignUpParams {
	var ()
	return &AccountSignUpParams{
		HTTPClient: client,
	}
}

/*AccountSignUpParams contains all the parameters to send to the API endpoint
for the account sign up operation typically these are written to a http.Request
*/
type AccountSignUpParams struct {

	/*JSONBody
	  The username, must be unique.

	*/
	JSONBody *models.User

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the account sign up params
func (o *AccountSignUpParams) WithTimeout(timeout time.Duration) *AccountSignUpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account sign up params
func (o *AccountSignUpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account sign up params
func (o *AccountSignUpParams) WithContext(ctx context.Context) *AccountSignUpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account sign up params
func (o *AccountSignUpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account sign up params
func (o *AccountSignUpParams) WithHTTPClient(client *http.Client) *AccountSignUpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account sign up params
func (o *AccountSignUpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the account sign up params
func (o *AccountSignUpParams) WithJSONBody(jSONBody *models.User) *AccountSignUpParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the account sign up params
func (o *AccountSignUpParams) SetJSONBody(jSONBody *models.User) {
	o.JSONBody = jSONBody
}

// WriteToRequest writes these params to a swagger request
func (o *AccountSignUpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
