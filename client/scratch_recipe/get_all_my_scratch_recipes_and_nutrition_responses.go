// Code generated by go-swagger; DO NOT EDIT.

package scratch_recipe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// GetAllMyScratchRecipesAndNutritionReader is a Reader for the GetAllMyScratchRecipesAndNutrition structure.
type GetAllMyScratchRecipesAndNutritionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllMyScratchRecipesAndNutritionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllMyScratchRecipesAndNutritionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAllMyScratchRecipesAndNutritionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAllMyScratchRecipesAndNutritionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllMyScratchRecipesAndNutritionOK creates a GetAllMyScratchRecipesAndNutritionOK with default headers values
func NewGetAllMyScratchRecipesAndNutritionOK() *GetAllMyScratchRecipesAndNutritionOK {
	return &GetAllMyScratchRecipesAndNutritionOK{}
}

/*GetAllMyScratchRecipesAndNutritionOK handles this case with default header values.

Retreives a list of the currenty logged in user's scratch pad
*/
type GetAllMyScratchRecipesAndNutritionOK struct {
	Payload []*models.Recipe
}

func (o *GetAllMyScratchRecipesAndNutritionOK) Error() string {
	return fmt.Sprintf("[GET /scratch-recipe/nutrition][%d] getAllMyScratchRecipesAndNutritionOK  %+v", 200, o.Payload)
}

func (o *GetAllMyScratchRecipesAndNutritionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllMyScratchRecipesAndNutritionBadRequest creates a GetAllMyScratchRecipesAndNutritionBadRequest with default headers values
func NewGetAllMyScratchRecipesAndNutritionBadRequest() *GetAllMyScratchRecipesAndNutritionBadRequest {
	return &GetAllMyScratchRecipesAndNutritionBadRequest{}
}

/*GetAllMyScratchRecipesAndNutritionBadRequest handles this case with default header values.

Malformed JSON (this endpoint requires an array of strings
*/
type GetAllMyScratchRecipesAndNutritionBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *GetAllMyScratchRecipesAndNutritionBadRequest) Error() string {
	return fmt.Sprintf("[GET /scratch-recipe/nutrition][%d] getAllMyScratchRecipesAndNutritionBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllMyScratchRecipesAndNutritionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllMyScratchRecipesAndNutritionForbidden creates a GetAllMyScratchRecipesAndNutritionForbidden with default headers values
func NewGetAllMyScratchRecipesAndNutritionForbidden() *GetAllMyScratchRecipesAndNutritionForbidden {
	return &GetAllMyScratchRecipesAndNutritionForbidden{}
}

/*GetAllMyScratchRecipesAndNutritionForbidden handles this case with default header values.

Description was not specified
*/
type GetAllMyScratchRecipesAndNutritionForbidden struct {
	Payload *models.HTTPAPIError
}

func (o *GetAllMyScratchRecipesAndNutritionForbidden) Error() string {
	return fmt.Sprintf("[GET /scratch-recipe/nutrition][%d] getAllMyScratchRecipesAndNutritionForbidden  %+v", 403, o.Payload)
}

func (o *GetAllMyScratchRecipesAndNutritionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
