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

// GetRecipeByIDReader is a Reader for the GetRecipeByID structure.
type GetRecipeByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecipeByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRecipeByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRecipeByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRecipeByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecipeByIDOK creates a GetRecipeByIDOK with default headers values
func NewGetRecipeByIDOK() *GetRecipeByIDOK {
	return &GetRecipeByIDOK{}
}

/*GetRecipeByIDOK handles this case with default header values.

Description was not specified
*/
type GetRecipeByIDOK struct {
	Payload *models.Recipe
}

func (o *GetRecipeByIDOK) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}][%d] getRecipeByIdOK  %+v", 200, o.Payload)
}

func (o *GetRecipeByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recipe)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipeByIDBadRequest creates a GetRecipeByIDBadRequest with default headers values
func NewGetRecipeByIDBadRequest() *GetRecipeByIDBadRequest {
	return &GetRecipeByIDBadRequest{}
}

/*GetRecipeByIDBadRequest handles this case with default header values.

Description was not specified
*/
type GetRecipeByIDBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *GetRecipeByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}][%d] getRecipeByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecipeByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipeByIDNotFound creates a GetRecipeByIDNotFound with default headers values
func NewGetRecipeByIDNotFound() *GetRecipeByIDNotFound {
	return &GetRecipeByIDNotFound{}
}

/*GetRecipeByIDNotFound handles this case with default header values.

Description was not specified
*/
type GetRecipeByIDNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *GetRecipeByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}][%d] getRecipeByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetRecipeByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
