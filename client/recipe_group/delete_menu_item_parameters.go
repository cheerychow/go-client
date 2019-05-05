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

// NewDeleteMenuItemParams creates a new DeleteMenuItemParams object
// with the default values initialized.
func NewDeleteMenuItemParams() *DeleteMenuItemParams {
	var ()
	return &DeleteMenuItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMenuItemParamsWithTimeout creates a new DeleteMenuItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMenuItemParamsWithTimeout(timeout time.Duration) *DeleteMenuItemParams {
	var ()
	return &DeleteMenuItemParams{

		timeout: timeout,
	}
}

// NewDeleteMenuItemParamsWithContext creates a new DeleteMenuItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMenuItemParamsWithContext(ctx context.Context) *DeleteMenuItemParams {
	var ()
	return &DeleteMenuItemParams{

		Context: ctx,
	}
}

// NewDeleteMenuItemParamsWithHTTPClient creates a new DeleteMenuItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMenuItemParamsWithHTTPClient(client *http.Client) *DeleteMenuItemParams {
	var ()
	return &DeleteMenuItemParams{
		HTTPClient: client,
	}
}

/*DeleteMenuItemParams contains all the parameters to send to the API endpoint
for the delete menu item operation typically these are written to a http.Request
*/
type DeleteMenuItemParams struct {

	/*MenuItemID
	  The primary key of the menu item to be deleted.

	*/
	MenuItemID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete menu item params
func (o *DeleteMenuItemParams) WithTimeout(timeout time.Duration) *DeleteMenuItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete menu item params
func (o *DeleteMenuItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete menu item params
func (o *DeleteMenuItemParams) WithContext(ctx context.Context) *DeleteMenuItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete menu item params
func (o *DeleteMenuItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete menu item params
func (o *DeleteMenuItemParams) WithHTTPClient(client *http.Client) *DeleteMenuItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete menu item params
func (o *DeleteMenuItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMenuItemID adds the menuItemID to the delete menu item params
func (o *DeleteMenuItemParams) WithMenuItemID(menuItemID int64) *DeleteMenuItemParams {
	o.SetMenuItemID(menuItemID)
	return o
}

// SetMenuItemID adds the menuItemId to the delete menu item params
func (o *DeleteMenuItemParams) SetMenuItemID(menuItemID int64) {
	o.MenuItemID = menuItemID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMenuItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param menuItemId
	if err := r.SetPathParam("menuItemId", swag.FormatInt64(o.MenuItemID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
