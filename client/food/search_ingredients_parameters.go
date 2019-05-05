// Code generated by go-swagger; DO NOT EDIT.

package food

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

// NewSearchIngredientsParams creates a new SearchIngredientsParams object
// with the default values initialized.
func NewSearchIngredientsParams() *SearchIngredientsParams {
	var ()
	return &SearchIngredientsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchIngredientsParamsWithTimeout creates a new SearchIngredientsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchIngredientsParamsWithTimeout(timeout time.Duration) *SearchIngredientsParams {
	var ()
	return &SearchIngredientsParams{

		timeout: timeout,
	}
}

// NewSearchIngredientsParamsWithContext creates a new SearchIngredientsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchIngredientsParamsWithContext(ctx context.Context) *SearchIngredientsParams {
	var ()
	return &SearchIngredientsParams{

		Context: ctx,
	}
}

// NewSearchIngredientsParamsWithHTTPClient creates a new SearchIngredientsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchIngredientsParamsWithHTTPClient(client *http.Client) *SearchIngredientsParams {
	var ()
	return &SearchIngredientsParams{
		HTTPClient: client,
	}
}

/*SearchIngredientsParams contains all the parameters to send to the API endpoint
for the search ingredients operation typically these are written to a http.Request
*/
type SearchIngredientsParams struct {

	/*Search
	  Case insensitive search term, for ingredients or foods. Ingredients will be parsed as ingredient lines and returned as IngredientWithNutritionAbbrev objects.

	*/
	Search string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search ingredients params
func (o *SearchIngredientsParams) WithTimeout(timeout time.Duration) *SearchIngredientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search ingredients params
func (o *SearchIngredientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search ingredients params
func (o *SearchIngredientsParams) WithContext(ctx context.Context) *SearchIngredientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search ingredients params
func (o *SearchIngredientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search ingredients params
func (o *SearchIngredientsParams) WithHTTPClient(client *http.Client) *SearchIngredientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search ingredients params
func (o *SearchIngredientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearch adds the search to the search ingredients params
func (o *SearchIngredientsParams) WithSearch(search string) *SearchIngredientsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the search ingredients params
func (o *SearchIngredientsParams) SetSearch(search string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SearchIngredientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param search
	if err := r.SetPathParam("search", o.Search); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
