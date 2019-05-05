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

// PostRecipeGroupAndMenuItemReader is a Reader for the PostRecipeGroupAndMenuItem structure.
type PostRecipeGroupAndMenuItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecipeGroupAndMenuItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRecipeGroupAndMenuItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewPostRecipeGroupAndMenuItemCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostRecipeGroupAndMenuItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostRecipeGroupAndMenuItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRecipeGroupAndMenuItemOK creates a PostRecipeGroupAndMenuItemOK with default headers values
func NewPostRecipeGroupAndMenuItemOK() *PostRecipeGroupAndMenuItemOK {
	return &PostRecipeGroupAndMenuItemOK{}
}

/*PostRecipeGroupAndMenuItemOK handles this case with default header values.

Description was not specified
*/
type PostRecipeGroupAndMenuItemOK struct {
	Payload []*models.MenuItem
}

func (o *PostRecipeGroupAndMenuItemOK) Error() string {
	return fmt.Sprintf("[POST /recipe-group-and-menu-item][%d] postRecipeGroupAndMenuItemOK  %+v", 200, o.Payload)
}

func (o *PostRecipeGroupAndMenuItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecipeGroupAndMenuItemCreated creates a PostRecipeGroupAndMenuItemCreated with default headers values
func NewPostRecipeGroupAndMenuItemCreated() *PostRecipeGroupAndMenuItemCreated {
	return &PostRecipeGroupAndMenuItemCreated{}
}

/*PostRecipeGroupAndMenuItemCreated handles this case with default header values.

Description was not specified
*/
type PostRecipeGroupAndMenuItemCreated struct {
	Payload []*models.MenuItem
}

func (o *PostRecipeGroupAndMenuItemCreated) Error() string {
	return fmt.Sprintf("[POST /recipe-group-and-menu-item][%d] postRecipeGroupAndMenuItemCreated  %+v", 201, o.Payload)
}

func (o *PostRecipeGroupAndMenuItemCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecipeGroupAndMenuItemBadRequest creates a PostRecipeGroupAndMenuItemBadRequest with default headers values
func NewPostRecipeGroupAndMenuItemBadRequest() *PostRecipeGroupAndMenuItemBadRequest {
	return &PostRecipeGroupAndMenuItemBadRequest{}
}

/*PostRecipeGroupAndMenuItemBadRequest handles this case with default header values.

Description was not specified
*/
type PostRecipeGroupAndMenuItemBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *PostRecipeGroupAndMenuItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /recipe-group-and-menu-item][%d] postRecipeGroupAndMenuItemBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecipeGroupAndMenuItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecipeGroupAndMenuItemNotFound creates a PostRecipeGroupAndMenuItemNotFound with default headers values
func NewPostRecipeGroupAndMenuItemNotFound() *PostRecipeGroupAndMenuItemNotFound {
	return &PostRecipeGroupAndMenuItemNotFound{}
}

/*PostRecipeGroupAndMenuItemNotFound handles this case with default header values.

Description was not specified
*/
type PostRecipeGroupAndMenuItemNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *PostRecipeGroupAndMenuItemNotFound) Error() string {
	return fmt.Sprintf("[POST /recipe-group-and-menu-item][%d] postRecipeGroupAndMenuItemNotFound  %+v", 404, o.Payload)
}

func (o *PostRecipeGroupAndMenuItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}