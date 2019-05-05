// Code generated by go-swagger; DO NOT EDIT.

package recipe

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

	models "github.com/cheerychow/go-client/models"
)

// NewEditIngredientForRecipeParams creates a new EditIngredientForRecipeParams object
// with the default values initialized.
func NewEditIngredientForRecipeParams() *EditIngredientForRecipeParams {
	var ()
	return &EditIngredientForRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditIngredientForRecipeParamsWithTimeout creates a new EditIngredientForRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditIngredientForRecipeParamsWithTimeout(timeout time.Duration) *EditIngredientForRecipeParams {
	var ()
	return &EditIngredientForRecipeParams{

		timeout: timeout,
	}
}

// NewEditIngredientForRecipeParamsWithContext creates a new EditIngredientForRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditIngredientForRecipeParamsWithContext(ctx context.Context) *EditIngredientForRecipeParams {
	var ()
	return &EditIngredientForRecipeParams{

		Context: ctx,
	}
}

// NewEditIngredientForRecipeParamsWithHTTPClient creates a new EditIngredientForRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditIngredientForRecipeParamsWithHTTPClient(client *http.Client) *EditIngredientForRecipeParams {
	var ()
	return &EditIngredientForRecipeParams{
		HTTPClient: client,
	}
}

/*EditIngredientForRecipeParams contains all the parameters to send to the API endpoint
for the edit ingredient for recipe operation typically these are written to a http.Request
*/
type EditIngredientForRecipeParams struct {

	/*RecipeID
	  The recipe ID

	*/
	RecipeID string
	/*RecipeIngredients
	  The recipe ID

	*/
	RecipeIngredients []*models.Ingredient

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) WithTimeout(timeout time.Duration) *EditIngredientForRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) WithContext(ctx context.Context) *EditIngredientForRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) WithHTTPClient(client *http.Client) *EditIngredientForRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecipeID adds the recipeID to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) WithRecipeID(recipeID string) *EditIngredientForRecipeParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WithRecipeIngredients adds the recipeIngredients to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) WithRecipeIngredients(recipeIngredients []*models.Ingredient) *EditIngredientForRecipeParams {
	o.SetRecipeIngredients(recipeIngredients)
	return o
}

// SetRecipeIngredients adds the recipeIngredients to the edit ingredient for recipe params
func (o *EditIngredientForRecipeParams) SetRecipeIngredients(recipeIngredients []*models.Ingredient) {
	o.RecipeIngredients = recipeIngredients
}

// WriteToRequest writes these params to a swagger request
func (o *EditIngredientForRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param recipeId
	if err := r.SetPathParam("recipeId", o.RecipeID); err != nil {
		return err
	}

	if o.RecipeIngredients != nil {
		if err := r.SetBodyParam(o.RecipeIngredients); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
