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

// NewSaveRecipeParams creates a new SaveRecipeParams object
// with the default values initialized.
func NewSaveRecipeParams() *SaveRecipeParams {
	var ()
	return &SaveRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSaveRecipeParamsWithTimeout creates a new SaveRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSaveRecipeParamsWithTimeout(timeout time.Duration) *SaveRecipeParams {
	var ()
	return &SaveRecipeParams{

		timeout: timeout,
	}
}

// NewSaveRecipeParamsWithContext creates a new SaveRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSaveRecipeParamsWithContext(ctx context.Context) *SaveRecipeParams {
	var ()
	return &SaveRecipeParams{

		Context: ctx,
	}
}

// NewSaveRecipeParamsWithHTTPClient creates a new SaveRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSaveRecipeParamsWithHTTPClient(client *http.Client) *SaveRecipeParams {
	var ()
	return &SaveRecipeParams{
		HTTPClient: client,
	}
}

/*SaveRecipeParams contains all the parameters to send to the API endpoint
for the save recipe operation typically these are written to a http.Request
*/
type SaveRecipeParams struct {

	/*JSONBody
	  The recipe to be saved.

	*/
	JSONBody *models.Recipe
	/*RecipeID
	  The recipe ID of the recipe you're updating, either the unique Recipe ID int64 or the oid string may be used here.

	*/
	RecipeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the save recipe params
func (o *SaveRecipeParams) WithTimeout(timeout time.Duration) *SaveRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save recipe params
func (o *SaveRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save recipe params
func (o *SaveRecipeParams) WithContext(ctx context.Context) *SaveRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save recipe params
func (o *SaveRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save recipe params
func (o *SaveRecipeParams) WithHTTPClient(client *http.Client) *SaveRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save recipe params
func (o *SaveRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the save recipe params
func (o *SaveRecipeParams) WithJSONBody(jSONBody *models.Recipe) *SaveRecipeParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the save recipe params
func (o *SaveRecipeParams) SetJSONBody(jSONBody *models.Recipe) {
	o.JSONBody = jSONBody
}

// WithRecipeID adds the recipeID to the save recipe params
func (o *SaveRecipeParams) WithRecipeID(recipeID string) *SaveRecipeParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the save recipe params
func (o *SaveRecipeParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WriteToRequest writes these params to a swagger request
func (o *SaveRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
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
