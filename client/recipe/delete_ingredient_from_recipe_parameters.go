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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteIngredientFromRecipeParams creates a new DeleteIngredientFromRecipeParams object
// with the default values initialized.
func NewDeleteIngredientFromRecipeParams() *DeleteIngredientFromRecipeParams {
	var ()
	return &DeleteIngredientFromRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIngredientFromRecipeParamsWithTimeout creates a new DeleteIngredientFromRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIngredientFromRecipeParamsWithTimeout(timeout time.Duration) *DeleteIngredientFromRecipeParams {
	var ()
	return &DeleteIngredientFromRecipeParams{

		timeout: timeout,
	}
}

// NewDeleteIngredientFromRecipeParamsWithContext creates a new DeleteIngredientFromRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIngredientFromRecipeParamsWithContext(ctx context.Context) *DeleteIngredientFromRecipeParams {
	var ()
	return &DeleteIngredientFromRecipeParams{

		Context: ctx,
	}
}

// NewDeleteIngredientFromRecipeParamsWithHTTPClient creates a new DeleteIngredientFromRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIngredientFromRecipeParamsWithHTTPClient(client *http.Client) *DeleteIngredientFromRecipeParams {
	var ()
	return &DeleteIngredientFromRecipeParams{
		HTTPClient: client,
	}
}

/*DeleteIngredientFromRecipeParams contains all the parameters to send to the API endpoint
for the delete ingredient from recipe operation typically these are written to a http.Request
*/
type DeleteIngredientFromRecipeParams struct {

	/*IngredientID
	  The ingredient ID for the ingredient being edited.

	*/
	IngredientID int64
	/*RecipeID
	  The recipe ID

	*/
	RecipeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) WithTimeout(timeout time.Duration) *DeleteIngredientFromRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) WithContext(ctx context.Context) *DeleteIngredientFromRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) WithHTTPClient(client *http.Client) *DeleteIngredientFromRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIngredientID adds the ingredientID to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) WithIngredientID(ingredientID int64) *DeleteIngredientFromRecipeParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) SetIngredientID(ingredientID int64) {
	o.IngredientID = ingredientID
}

// WithRecipeID adds the recipeID to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) WithRecipeID(recipeID string) *DeleteIngredientFromRecipeParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the delete ingredient from recipe params
func (o *DeleteIngredientFromRecipeParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIngredientFromRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ingredient-id
	if err := r.SetPathParam("ingredient-id", swag.FormatInt64(o.IngredientID)); err != nil {
		return err
	}

	// path param recipeId
	if err := r.SetPathParam("recipeId", o.RecipeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
