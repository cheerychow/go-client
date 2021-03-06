// Code generated by go-swagger; DO NOT EDIT.

package nutrition

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

// NewGetNutritionTipsForNutrientParams creates a new GetNutritionTipsForNutrientParams object
// with the default values initialized.
func NewGetNutritionTipsForNutrientParams() *GetNutritionTipsForNutrientParams {
	var ()
	return &GetNutritionTipsForNutrientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNutritionTipsForNutrientParamsWithTimeout creates a new GetNutritionTipsForNutrientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNutritionTipsForNutrientParamsWithTimeout(timeout time.Duration) *GetNutritionTipsForNutrientParams {
	var ()
	return &GetNutritionTipsForNutrientParams{

		timeout: timeout,
	}
}

// NewGetNutritionTipsForNutrientParamsWithContext creates a new GetNutritionTipsForNutrientParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNutritionTipsForNutrientParamsWithContext(ctx context.Context) *GetNutritionTipsForNutrientParams {
	var ()
	return &GetNutritionTipsForNutrientParams{

		Context: ctx,
	}
}

// NewGetNutritionTipsForNutrientParamsWithHTTPClient creates a new GetNutritionTipsForNutrientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNutritionTipsForNutrientParamsWithHTTPClient(client *http.Client) *GetNutritionTipsForNutrientParams {
	var ()
	return &GetNutritionTipsForNutrientParams{
		HTTPClient: client,
	}
}

/*GetNutritionTipsForNutrientParams contains all the parameters to send to the API endpoint
for the get nutrition tips for nutrient operation typically these are written to a http.Request
*/
type GetNutritionTipsForNutrientParams struct {

	/*NutrientID
	  The nutrient ID, either the nutrient ID number of the key name for the nutrient (string).

	*/
	NutrientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nutrition tips for nutrient params
func (o *GetNutritionTipsForNutrientParams) WithTimeout(timeout time.Duration) *GetNutritionTipsForNutrientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nutrition tips for nutrient params
func (o *GetNutritionTipsForNutrientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nutrition tips for nutrient params
func (o *GetNutritionTipsForNutrientParams) WithContext(ctx context.Context) *GetNutritionTipsForNutrientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nutrition tips for nutrient params
func (o *GetNutritionTipsForNutrientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nutrition tips for nutrient params
func (o *GetNutritionTipsForNutrientParams) WithHTTPClient(client *http.Client) *GetNutritionTipsForNutrientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nutrition tips for nutrient params
func (o *GetNutritionTipsForNutrientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNutrientID adds the nutrientID to the get nutrition tips for nutrient params
func (o *GetNutritionTipsForNutrientParams) WithNutrientID(nutrientID string) *GetNutritionTipsForNutrientParams {
	o.SetNutrientID(nutrientID)
	return o
}

// SetNutrientID adds the nutrientId to the get nutrition tips for nutrient params
func (o *GetNutritionTipsForNutrientParams) SetNutrientID(nutrientID string) {
	o.NutrientID = nutrientID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNutritionTipsForNutrientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nutrientId
	if err := r.SetPathParam("nutrientId", o.NutrientID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
