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

// GetRecipeGroupsAndRecipesReader is a Reader for the GetRecipeGroupsAndRecipes structure.
type GetRecipeGroupsAndRecipesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecipeGroupsAndRecipesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRecipeGroupsAndRecipesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRecipeGroupsAndRecipesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecipeGroupsAndRecipesOK creates a GetRecipeGroupsAndRecipesOK with default headers values
func NewGetRecipeGroupsAndRecipesOK() *GetRecipeGroupsAndRecipesOK {
	return &GetRecipeGroupsAndRecipesOK{}
}

/*GetRecipeGroupsAndRecipesOK handles this case with default header values.

Description was not specified
*/
type GetRecipeGroupsAndRecipesOK struct {
	Payload []*models.RecipeGroup
}

func (o *GetRecipeGroupsAndRecipesOK) Error() string {
	return fmt.Sprintf("[GET /recipe-group-and-menu-items][%d] getRecipeGroupsAndRecipesOK  %+v", 200, o.Payload)
}

func (o *GetRecipeGroupsAndRecipesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipeGroupsAndRecipesBadRequest creates a GetRecipeGroupsAndRecipesBadRequest with default headers values
func NewGetRecipeGroupsAndRecipesBadRequest() *GetRecipeGroupsAndRecipesBadRequest {
	return &GetRecipeGroupsAndRecipesBadRequest{}
}

/*GetRecipeGroupsAndRecipesBadRequest handles this case with default header values.

Description was not specified
*/
type GetRecipeGroupsAndRecipesBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *GetRecipeGroupsAndRecipesBadRequest) Error() string {
	return fmt.Sprintf("[GET /recipe-group-and-menu-items][%d] getRecipeGroupsAndRecipesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecipeGroupsAndRecipesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}