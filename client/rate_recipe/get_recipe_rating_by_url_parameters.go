// Code generated by go-swagger; DO NOT EDIT.

package rate_recipe

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

// NewGetRecipeRatingByURLParams creates a new GetRecipeRatingByURLParams object
// with the default values initialized.
func NewGetRecipeRatingByURLParams() *GetRecipeRatingByURLParams {
	var ()
	return &GetRecipeRatingByURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecipeRatingByURLParamsWithTimeout creates a new GetRecipeRatingByURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecipeRatingByURLParamsWithTimeout(timeout time.Duration) *GetRecipeRatingByURLParams {
	var ()
	return &GetRecipeRatingByURLParams{

		timeout: timeout,
	}
}

// NewGetRecipeRatingByURLParamsWithContext creates a new GetRecipeRatingByURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecipeRatingByURLParamsWithContext(ctx context.Context) *GetRecipeRatingByURLParams {
	var ()
	return &GetRecipeRatingByURLParams{

		Context: ctx,
	}
}

// NewGetRecipeRatingByURLParamsWithHTTPClient creates a new GetRecipeRatingByURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecipeRatingByURLParamsWithHTTPClient(client *http.Client) *GetRecipeRatingByURLParams {
	var ()
	return &GetRecipeRatingByURLParams{
		HTTPClient: client,
	}
}

/*GetRecipeRatingByURLParams contains all the parameters to send to the API endpoint
for the get recipe rating by Url operation typically these are written to a http.Request
*/
type GetRecipeRatingByURLParams struct {

	/*URL
	  The url of the recipe you want to rate.

	*/
	URL string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recipe rating by Url params
func (o *GetRecipeRatingByURLParams) WithTimeout(timeout time.Duration) *GetRecipeRatingByURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recipe rating by Url params
func (o *GetRecipeRatingByURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recipe rating by Url params
func (o *GetRecipeRatingByURLParams) WithContext(ctx context.Context) *GetRecipeRatingByURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recipe rating by Url params
func (o *GetRecipeRatingByURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recipe rating by Url params
func (o *GetRecipeRatingByURLParams) WithHTTPClient(client *http.Client) *GetRecipeRatingByURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recipe rating by Url params
func (o *GetRecipeRatingByURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithURL adds the url to the get recipe rating by Url params
func (o *GetRecipeRatingByURLParams) WithURL(url string) *GetRecipeRatingByURLParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the get recipe rating by Url params
func (o *GetRecipeRatingByURLParams) SetURL(url string) {
	o.URL = url
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecipeRatingByURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param url
	if err := r.SetPathParam("url", o.URL); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}