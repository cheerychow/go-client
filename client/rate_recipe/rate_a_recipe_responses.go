// Code generated by go-swagger; DO NOT EDIT.

package rate_recipe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/chow-go-client/models"
)

// RateARecipeReader is a Reader for the RateARecipe structure.
type RateARecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RateARecipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewRateARecipeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRateARecipeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRateARecipeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRateARecipeCreated creates a RateARecipeCreated with default headers values
func NewRateARecipeCreated() *RateARecipeCreated {
	return &RateARecipeCreated{}
}

/*RateARecipeCreated handles this case with default header values.

The recipe rating.
*/
type RateARecipeCreated struct {
	Payload *models.RecipeRatingResult
}

func (o *RateARecipeCreated) Error() string {
	return fmt.Sprintf("[POST /rate-recipe][%d] rateARecipeCreated  %+v", 201, o.Payload)
}

func (o *RateARecipeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecipeRatingResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRateARecipeBadRequest creates a RateARecipeBadRequest with default headers values
func NewRateARecipeBadRequest() *RateARecipeBadRequest {
	return &RateARecipeBadRequest{}
}

/*RateARecipeBadRequest handles this case with default header values.

Description was not specified
*/
type RateARecipeBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *RateARecipeBadRequest) Error() string {
	return fmt.Sprintf("[POST /rate-recipe][%d] rateARecipeBadRequest  %+v", 400, o.Payload)
}

func (o *RateARecipeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRateARecipeNotFound creates a RateARecipeNotFound with default headers values
func NewRateARecipeNotFound() *RateARecipeNotFound {
	return &RateARecipeNotFound{}
}

/*RateARecipeNotFound handles this case with default header values.

Description was not specified
*/
type RateARecipeNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *RateARecipeNotFound) Error() string {
	return fmt.Sprintf("[POST /rate-recipe][%d] rateARecipeNotFound  %+v", 404, o.Payload)
}

func (o *RateARecipeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
