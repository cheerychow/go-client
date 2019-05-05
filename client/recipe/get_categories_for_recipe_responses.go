// Code generated by go-swagger; DO NOT EDIT.

package recipe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// GetCategoriesForRecipeReader is a Reader for the GetCategoriesForRecipe structure.
type GetCategoriesForRecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCategoriesForRecipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCategoriesForRecipeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetCategoriesForRecipeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCategoriesForRecipeOK creates a GetCategoriesForRecipeOK with default headers values
func NewGetCategoriesForRecipeOK() *GetCategoriesForRecipeOK {
	return &GetCategoriesForRecipeOK{}
}

/*GetCategoriesForRecipeOK handles this case with default header values.

Description was not specified
*/
type GetCategoriesForRecipeOK struct {
	Payload []string
}

func (o *GetCategoriesForRecipeOK) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/categories][%d] getCategoriesForRecipeOK  %+v", 200, o.Payload)
}

func (o *GetCategoriesForRecipeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCategoriesForRecipeBadRequest creates a GetCategoriesForRecipeBadRequest with default headers values
func NewGetCategoriesForRecipeBadRequest() *GetCategoriesForRecipeBadRequest {
	return &GetCategoriesForRecipeBadRequest{}
}

/*GetCategoriesForRecipeBadRequest handles this case with default header values.

Description was not specified
*/
type GetCategoriesForRecipeBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *GetCategoriesForRecipeBadRequest) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/categories][%d] getCategoriesForRecipeBadRequest  %+v", 400, o.Payload)
}

func (o *GetCategoriesForRecipeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}