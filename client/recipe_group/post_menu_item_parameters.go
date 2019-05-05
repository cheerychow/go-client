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

	models "github.com/cheerychow/go-client/models"
)

// NewPostMenuItemParams creates a new PostMenuItemParams object
// with the default values initialized.
func NewPostMenuItemParams() *PostMenuItemParams {
	var ()
	return &PostMenuItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMenuItemParamsWithTimeout creates a new PostMenuItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMenuItemParamsWithTimeout(timeout time.Duration) *PostMenuItemParams {
	var ()
	return &PostMenuItemParams{

		timeout: timeout,
	}
}

// NewPostMenuItemParamsWithContext creates a new PostMenuItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMenuItemParamsWithContext(ctx context.Context) *PostMenuItemParams {
	var ()
	return &PostMenuItemParams{

		Context: ctx,
	}
}

// NewPostMenuItemParamsWithHTTPClient creates a new PostMenuItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMenuItemParamsWithHTTPClient(client *http.Client) *PostMenuItemParams {
	var ()
	return &PostMenuItemParams{
		HTTPClient: client,
	}
}

/*PostMenuItemParams contains all the parameters to send to the API endpoint
for the post menu item operation typically these are written to a http.Request
*/
type PostMenuItemParams struct {

	/*JSONBody
	  The recipe group data. Some additional parameters are provide finer control over the recipe group insertion. Such as overwriting a previous slug, making sure only one ever exists, or is inserted. The maximum number of menu-items per recipe group is 100.

	*/
	JSONBody *models.MenuItem
	/*RecipeGroupID
	  The group ID required

	*/
	RecipeGroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post menu item params
func (o *PostMenuItemParams) WithTimeout(timeout time.Duration) *PostMenuItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post menu item params
func (o *PostMenuItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post menu item params
func (o *PostMenuItemParams) WithContext(ctx context.Context) *PostMenuItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post menu item params
func (o *PostMenuItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post menu item params
func (o *PostMenuItemParams) WithHTTPClient(client *http.Client) *PostMenuItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post menu item params
func (o *PostMenuItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the post menu item params
func (o *PostMenuItemParams) WithJSONBody(jSONBody *models.MenuItem) *PostMenuItemParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the post menu item params
func (o *PostMenuItemParams) SetJSONBody(jSONBody *models.MenuItem) {
	o.JSONBody = jSONBody
}

// WithRecipeGroupID adds the recipeGroupID to the post menu item params
func (o *PostMenuItemParams) WithRecipeGroupID(recipeGroupID int64) *PostMenuItemParams {
	o.SetRecipeGroupID(recipeGroupID)
	return o
}

// SetRecipeGroupID adds the recipeGroupId to the post menu item params
func (o *PostMenuItemParams) SetRecipeGroupID(recipeGroupID int64) {
	o.RecipeGroupID = recipeGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PostMenuItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	// path param recipeGroupId
	if err := r.SetPathParam("recipeGroupId", swag.FormatInt64(o.RecipeGroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
