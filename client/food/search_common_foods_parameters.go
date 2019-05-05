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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSearchCommonFoodsParams creates a new SearchCommonFoodsParams object
// with the default values initialized.
func NewSearchCommonFoodsParams() *SearchCommonFoodsParams {
	var ()
	return &SearchCommonFoodsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchCommonFoodsParamsWithTimeout creates a new SearchCommonFoodsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchCommonFoodsParamsWithTimeout(timeout time.Duration) *SearchCommonFoodsParams {
	var ()
	return &SearchCommonFoodsParams{

		timeout: timeout,
	}
}

// NewSearchCommonFoodsParamsWithContext creates a new SearchCommonFoodsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchCommonFoodsParamsWithContext(ctx context.Context) *SearchCommonFoodsParams {
	var ()
	return &SearchCommonFoodsParams{

		Context: ctx,
	}
}

// NewSearchCommonFoodsParamsWithHTTPClient creates a new SearchCommonFoodsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchCommonFoodsParamsWithHTTPClient(client *http.Client) *SearchCommonFoodsParams {
	var ()
	return &SearchCommonFoodsParams{
		HTTPClient: client,
	}
}

/*SearchCommonFoodsParams contains all the parameters to send to the API endpoint
for the search common foods operation typically these are written to a http.Request
*/
type SearchCommonFoodsParams struct {

	/*Limit
	  Limit the number of results returned.

	*/
	Limit *int32
	/*Offset
	  The offset into search results.

	*/
	Offset *int64
	/*Search
	  Case insensitive search term, for ingredients or foods. Ingredients will be parsed as ingredient lines and returned as IngredientWithNutritionAbbrev objects.

	*/
	Search string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search common foods params
func (o *SearchCommonFoodsParams) WithTimeout(timeout time.Duration) *SearchCommonFoodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search common foods params
func (o *SearchCommonFoodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search common foods params
func (o *SearchCommonFoodsParams) WithContext(ctx context.Context) *SearchCommonFoodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search common foods params
func (o *SearchCommonFoodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search common foods params
func (o *SearchCommonFoodsParams) WithHTTPClient(client *http.Client) *SearchCommonFoodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search common foods params
func (o *SearchCommonFoodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the search common foods params
func (o *SearchCommonFoodsParams) WithLimit(limit *int32) *SearchCommonFoodsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search common foods params
func (o *SearchCommonFoodsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the search common foods params
func (o *SearchCommonFoodsParams) WithOffset(offset *int64) *SearchCommonFoodsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search common foods params
func (o *SearchCommonFoodsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSearch adds the search to the search common foods params
func (o *SearchCommonFoodsParams) WithSearch(search string) *SearchCommonFoodsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the search common foods params
func (o *SearchCommonFoodsParams) SetSearch(search string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SearchCommonFoodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// path param search
	if err := r.SetPathParam("search", o.Search); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
