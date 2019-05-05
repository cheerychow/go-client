// Code generated by go-swagger; DO NOT EDIT.

package recipe_group

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

// NewSearchMyRecipeGroupForRecipeParams creates a new SearchMyRecipeGroupForRecipeParams object
// with the default values initialized.
func NewSearchMyRecipeGroupForRecipeParams() *SearchMyRecipeGroupForRecipeParams {
	var (
		includeMenuItemsDefault = bool(false)
		includeRecipeDefault    = bool(false)
	)
	return &SearchMyRecipeGroupForRecipeParams{
		IncludeMenuItems: &includeMenuItemsDefault,
		IncludeRecipe:    &includeRecipeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchMyRecipeGroupForRecipeParamsWithTimeout creates a new SearchMyRecipeGroupForRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchMyRecipeGroupForRecipeParamsWithTimeout(timeout time.Duration) *SearchMyRecipeGroupForRecipeParams {
	var (
		includeMenuItemsDefault = bool(false)
		includeRecipeDefault    = bool(false)
	)
	return &SearchMyRecipeGroupForRecipeParams{
		IncludeMenuItems: &includeMenuItemsDefault,
		IncludeRecipe:    &includeRecipeDefault,

		timeout: timeout,
	}
}

// NewSearchMyRecipeGroupForRecipeParamsWithContext creates a new SearchMyRecipeGroupForRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchMyRecipeGroupForRecipeParamsWithContext(ctx context.Context) *SearchMyRecipeGroupForRecipeParams {
	var (
		includeMenuItemsDefault = bool(false)
		includeRecipeDefault    = bool(false)
	)
	return &SearchMyRecipeGroupForRecipeParams{
		IncludeMenuItems: &includeMenuItemsDefault,
		IncludeRecipe:    &includeRecipeDefault,

		Context: ctx,
	}
}

// NewSearchMyRecipeGroupForRecipeParamsWithHTTPClient creates a new SearchMyRecipeGroupForRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchMyRecipeGroupForRecipeParamsWithHTTPClient(client *http.Client) *SearchMyRecipeGroupForRecipeParams {
	var (
		includeMenuItemsDefault = bool(false)
		includeRecipeDefault    = bool(false)
	)
	return &SearchMyRecipeGroupForRecipeParams{
		IncludeMenuItems: &includeMenuItemsDefault,
		IncludeRecipe:    &includeRecipeDefault,
		HTTPClient:       client,
	}
}

/*SearchMyRecipeGroupForRecipeParams contains all the parameters to send to the API endpoint
for the search my recipe group for recipe operation typically these are written to a http.Request
*/
type SearchMyRecipeGroupForRecipeParams struct {

	/*IncludeMenuItems
	  Should the recipe-group's menu items (1-to-many relationship) be included in the recipe-group object?

	*/
	IncludeMenuItems *bool
	/*IncludeRecipe
	  Should the menu-item's recipe object (1-to-1 relationship) be included in the menu-item object?

	*/
	IncludeRecipe *bool
	/*RecipeID
	  The recipe ID to fetch

	*/
	RecipeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) WithTimeout(timeout time.Duration) *SearchMyRecipeGroupForRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) WithContext(ctx context.Context) *SearchMyRecipeGroupForRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) WithHTTPClient(client *http.Client) *SearchMyRecipeGroupForRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeMenuItems adds the includeMenuItems to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) WithIncludeMenuItems(includeMenuItems *bool) *SearchMyRecipeGroupForRecipeParams {
	o.SetIncludeMenuItems(includeMenuItems)
	return o
}

// SetIncludeMenuItems adds the includeMenuItems to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) SetIncludeMenuItems(includeMenuItems *bool) {
	o.IncludeMenuItems = includeMenuItems
}

// WithIncludeRecipe adds the includeRecipe to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) WithIncludeRecipe(includeRecipe *bool) *SearchMyRecipeGroupForRecipeParams {
	o.SetIncludeRecipe(includeRecipe)
	return o
}

// SetIncludeRecipe adds the includeRecipe to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) SetIncludeRecipe(includeRecipe *bool) {
	o.IncludeRecipe = includeRecipe
}

// WithRecipeID adds the recipeID to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) WithRecipeID(recipeID string) *SearchMyRecipeGroupForRecipeParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the search my recipe group for recipe params
func (o *SearchMyRecipeGroupForRecipeParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchMyRecipeGroupForRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeMenuItems != nil {

		// query param include-menu-items
		var qrIncludeMenuItems bool
		if o.IncludeMenuItems != nil {
			qrIncludeMenuItems = *o.IncludeMenuItems
		}
		qIncludeMenuItems := swag.FormatBool(qrIncludeMenuItems)
		if qIncludeMenuItems != "" {
			if err := r.SetQueryParam("include-menu-items", qIncludeMenuItems); err != nil {
				return err
			}
		}

	}

	if o.IncludeRecipe != nil {

		// query param include-recipe
		var qrIncludeRecipe bool
		if o.IncludeRecipe != nil {
			qrIncludeRecipe = *o.IncludeRecipe
		}
		qIncludeRecipe := swag.FormatBool(qrIncludeRecipe)
		if qIncludeRecipe != "" {
			if err := r.SetQueryParam("include-recipe", qIncludeRecipe); err != nil {
				return err
			}
		}

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
