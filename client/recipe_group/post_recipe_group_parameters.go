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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// NewPostRecipeGroupParams creates a new PostRecipeGroupParams object
// with the default values initialized.
func NewPostRecipeGroupParams() *PostRecipeGroupParams {
	var ()
	return &PostRecipeGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRecipeGroupParamsWithTimeout creates a new PostRecipeGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRecipeGroupParamsWithTimeout(timeout time.Duration) *PostRecipeGroupParams {
	var ()
	return &PostRecipeGroupParams{

		timeout: timeout,
	}
}

// NewPostRecipeGroupParamsWithContext creates a new PostRecipeGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRecipeGroupParamsWithContext(ctx context.Context) *PostRecipeGroupParams {
	var ()
	return &PostRecipeGroupParams{

		Context: ctx,
	}
}

// NewPostRecipeGroupParamsWithHTTPClient creates a new PostRecipeGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRecipeGroupParamsWithHTTPClient(client *http.Client) *PostRecipeGroupParams {
	var ()
	return &PostRecipeGroupParams{
		HTTPClient: client,
	}
}

/*PostRecipeGroupParams contains all the parameters to send to the API endpoint
for the post recipe group operation typically these are written to a http.Request
*/
type PostRecipeGroupParams struct {

	/*JSONBody
	  The recipe group data. Some additional parameters are provide finer control over the recipe group insertion. Such as overwriting a previous slug, making sure only one ever exists, or is inserted

	*/
	JSONBody *models.RecipeGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post recipe group params
func (o *PostRecipeGroupParams) WithTimeout(timeout time.Duration) *PostRecipeGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post recipe group params
func (o *PostRecipeGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post recipe group params
func (o *PostRecipeGroupParams) WithContext(ctx context.Context) *PostRecipeGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post recipe group params
func (o *PostRecipeGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post recipe group params
func (o *PostRecipeGroupParams) WithHTTPClient(client *http.Client) *PostRecipeGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post recipe group params
func (o *PostRecipeGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the post recipe group params
func (o *PostRecipeGroupParams) WithJSONBody(jSONBody *models.RecipeGroup) *PostRecipeGroupParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the post recipe group params
func (o *PostRecipeGroupParams) SetJSONBody(jSONBody *models.RecipeGroup) {
	o.JSONBody = jSONBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostRecipeGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}