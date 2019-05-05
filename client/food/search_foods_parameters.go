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

// NewSearchFoodsParams creates a new SearchFoodsParams object
// with the default values initialized.
func NewSearchFoodsParams() *SearchFoodsParams {
	var ()
	return &SearchFoodsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchFoodsParamsWithTimeout creates a new SearchFoodsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchFoodsParamsWithTimeout(timeout time.Duration) *SearchFoodsParams {
	var ()
	return &SearchFoodsParams{

		timeout: timeout,
	}
}

// NewSearchFoodsParamsWithContext creates a new SearchFoodsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchFoodsParamsWithContext(ctx context.Context) *SearchFoodsParams {
	var ()
	return &SearchFoodsParams{

		Context: ctx,
	}
}

// NewSearchFoodsParamsWithHTTPClient creates a new SearchFoodsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchFoodsParamsWithHTTPClient(client *http.Client) *SearchFoodsParams {
	var ()
	return &SearchFoodsParams{
		HTTPClient: client,
	}
}

/*SearchFoodsParams contains all the parameters to send to the API endpoint
for the search foods operation typically these are written to a http.Request
*/
type SearchFoodsParams struct {

	/*Limit
	  Limit the number of results returned.

	*/
	Limit *int32
	/*Offset
	  The offset into search results.

	*/
	Offset *int64
	/*SearchTerm
	  A search term for searching for foods.

	*/
	SearchTerm string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search foods params
func (o *SearchFoodsParams) WithTimeout(timeout time.Duration) *SearchFoodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search foods params
func (o *SearchFoodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search foods params
func (o *SearchFoodsParams) WithContext(ctx context.Context) *SearchFoodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search foods params
func (o *SearchFoodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search foods params
func (o *SearchFoodsParams) WithHTTPClient(client *http.Client) *SearchFoodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search foods params
func (o *SearchFoodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the search foods params
func (o *SearchFoodsParams) WithLimit(limit *int32) *SearchFoodsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search foods params
func (o *SearchFoodsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the search foods params
func (o *SearchFoodsParams) WithOffset(offset *int64) *SearchFoodsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search foods params
func (o *SearchFoodsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSearchTerm adds the searchTerm to the search foods params
func (o *SearchFoodsParams) WithSearchTerm(searchTerm string) *SearchFoodsParams {
	o.SetSearchTerm(searchTerm)
	return o
}

// SetSearchTerm adds the searchTerm to the search foods params
func (o *SearchFoodsParams) SetSearchTerm(searchTerm string) {
	o.SearchTerm = searchTerm
}

// WriteToRequest writes these params to a swagger request
func (o *SearchFoodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param searchTerm
	if err := r.SetPathParam("searchTerm", o.SearchTerm); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}