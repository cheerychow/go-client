// Code generated by go-swagger; DO NOT EDIT.

package recipe_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/chow-go-client/models"
)

// ChangeAMenuItemsRecipeIDReader is a Reader for the ChangeAMenuItemsRecipeID structure.
type ChangeAMenuItemsRecipeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeAMenuItemsRecipeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeAMenuItemsRecipeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewChangeAMenuItemsRecipeIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewChangeAMenuItemsRecipeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeAMenuItemsRecipeIDOK creates a ChangeAMenuItemsRecipeIDOK with default headers values
func NewChangeAMenuItemsRecipeIDOK() *ChangeAMenuItemsRecipeIDOK {
	return &ChangeAMenuItemsRecipeIDOK{}
}

/*ChangeAMenuItemsRecipeIDOK handles this case with default header values.

The meu item to be fetched.
*/
type ChangeAMenuItemsRecipeIDOK struct {
	Payload *models.MenuItem
}

func (o *ChangeAMenuItemsRecipeIDOK) Error() string {
	return fmt.Sprintf("[PUT /menu-item/{menuItemId}/add-recipe-id/{recipeId}][%d] changeAMenuItemsRecipeIdOK  %+v", 200, o.Payload)
}

func (o *ChangeAMenuItemsRecipeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MenuItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeAMenuItemsRecipeIDBadRequest creates a ChangeAMenuItemsRecipeIDBadRequest with default headers values
func NewChangeAMenuItemsRecipeIDBadRequest() *ChangeAMenuItemsRecipeIDBadRequest {
	return &ChangeAMenuItemsRecipeIDBadRequest{}
}

/*ChangeAMenuItemsRecipeIDBadRequest handles this case with default header values.

Description was not specified.
*/
type ChangeAMenuItemsRecipeIDBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *ChangeAMenuItemsRecipeIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /menu-item/{menuItemId}/add-recipe-id/{recipeId}][%d] changeAMenuItemsRecipeIdBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeAMenuItemsRecipeIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeAMenuItemsRecipeIDNotFound creates a ChangeAMenuItemsRecipeIDNotFound with default headers values
func NewChangeAMenuItemsRecipeIDNotFound() *ChangeAMenuItemsRecipeIDNotFound {
	return &ChangeAMenuItemsRecipeIDNotFound{}
}

/*ChangeAMenuItemsRecipeIDNotFound handles this case with default header values.

Menu item or recipe id not found.
*/
type ChangeAMenuItemsRecipeIDNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *ChangeAMenuItemsRecipeIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /menu-item/{menuItemId}/add-recipe-id/{recipeId}][%d] changeAMenuItemsRecipeIdNotFound  %+v", 404, o.Payload)
}

func (o *ChangeAMenuItemsRecipeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
