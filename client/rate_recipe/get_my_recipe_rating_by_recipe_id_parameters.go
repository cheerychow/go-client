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

// NewGetMyRecipeRatingByRecipeIDParams creates a new GetMyRecipeRatingByRecipeIDParams object
// with the default values initialized.
func NewGetMyRecipeRatingByRecipeIDParams() *GetMyRecipeRatingByRecipeIDParams {
	var ()
	return &GetMyRecipeRatingByRecipeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMyRecipeRatingByRecipeIDParamsWithTimeout creates a new GetMyRecipeRatingByRecipeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMyRecipeRatingByRecipeIDParamsWithTimeout(timeout time.Duration) *GetMyRecipeRatingByRecipeIDParams {
	var ()
	return &GetMyRecipeRatingByRecipeIDParams{

		timeout: timeout,
	}
}

// NewGetMyRecipeRatingByRecipeIDParamsWithContext creates a new GetMyRecipeRatingByRecipeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMyRecipeRatingByRecipeIDParamsWithContext(ctx context.Context) *GetMyRecipeRatingByRecipeIDParams {
	var ()
	return &GetMyRecipeRatingByRecipeIDParams{

		Context: ctx,
	}
}

// NewGetMyRecipeRatingByRecipeIDParamsWithHTTPClient creates a new GetMyRecipeRatingByRecipeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMyRecipeRatingByRecipeIDParamsWithHTTPClient(client *http.Client) *GetMyRecipeRatingByRecipeIDParams {
	var ()
	return &GetMyRecipeRatingByRecipeIDParams{
		HTTPClient: client,
	}
}

/*GetMyRecipeRatingByRecipeIDParams contains all the parameters to send to the API endpoint
for the get my recipe rating by recipe Id operation typically these are written to a http.Request
*/
type GetMyRecipeRatingByRecipeIDParams struct {

	/*RecipeID
	  The url of the recipe, if it's already in the database.

	*/
	RecipeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get my recipe rating by recipe Id params
func (o *GetMyRecipeRatingByRecipeIDParams) WithTimeout(timeout time.Duration) *GetMyRecipeRatingByRecipeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get my recipe rating by recipe Id params
func (o *GetMyRecipeRatingByRecipeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get my recipe rating by recipe Id params
func (o *GetMyRecipeRatingByRecipeIDParams) WithContext(ctx context.Context) *GetMyRecipeRatingByRecipeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get my recipe rating by recipe Id params
func (o *GetMyRecipeRatingByRecipeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get my recipe rating by recipe Id params
func (o *GetMyRecipeRatingByRecipeIDParams) WithHTTPClient(client *http.Client) *GetMyRecipeRatingByRecipeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get my recipe rating by recipe Id params
func (o *GetMyRecipeRatingByRecipeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecipeID adds the recipeID to the get my recipe rating by recipe Id params
func (o *GetMyRecipeRatingByRecipeIDParams) WithRecipeID(recipeID string) *GetMyRecipeRatingByRecipeIDParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the get my recipe rating by recipe Id params
func (o *GetMyRecipeRatingByRecipeIDParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMyRecipeRatingByRecipeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param recipeId
	if err := r.SetPathParam("recipeId", o.RecipeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
