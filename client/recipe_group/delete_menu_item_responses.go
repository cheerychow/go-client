// Code generated by go-swagger; DO NOT EDIT.

package recipe_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// DeleteMenuItemReader is a Reader for the DeleteMenuItem structure.
type DeleteMenuItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMenuItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMenuItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteMenuItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteMenuItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMenuItemOK creates a DeleteMenuItemOK with default headers values
func NewDeleteMenuItemOK() *DeleteMenuItemOK {
	return &DeleteMenuItemOK{}
}

/*DeleteMenuItemOK handles this case with default header values.

Description was not specified
*/
type DeleteMenuItemOK struct {
	Payload *models.MenuItem
}

func (o *DeleteMenuItemOK) Error() string {
	return fmt.Sprintf("[DELETE /menu-item/{menuItemId}][%d] deleteMenuItemOK  %+v", 200, o.Payload)
}

func (o *DeleteMenuItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MenuItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMenuItemForbidden creates a DeleteMenuItemForbidden with default headers values
func NewDeleteMenuItemForbidden() *DeleteMenuItemForbidden {
	return &DeleteMenuItemForbidden{}
}

/*DeleteMenuItemForbidden handles this case with default header values.

You're not the group owner.
*/
type DeleteMenuItemForbidden struct {
	Payload *models.HTTPAPIError
}

func (o *DeleteMenuItemForbidden) Error() string {
	return fmt.Sprintf("[DELETE /menu-item/{menuItemId}][%d] deleteMenuItemForbidden  %+v", 403, o.Payload)
}

func (o *DeleteMenuItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMenuItemNotFound creates a DeleteMenuItemNotFound with default headers values
func NewDeleteMenuItemNotFound() *DeleteMenuItemNotFound {
	return &DeleteMenuItemNotFound{}
}

/*DeleteMenuItemNotFound handles this case with default header values.

Menu item doesn't exist
*/
type DeleteMenuItemNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *DeleteMenuItemNotFound) Error() string {
	return fmt.Sprintf("[DELETE /menu-item/{menuItemId}][%d] deleteMenuItemNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMenuItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
