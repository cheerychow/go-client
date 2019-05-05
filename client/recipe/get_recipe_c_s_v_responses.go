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

// GetRecipeCSVReader is a Reader for the GetRecipeCSV structure.
type GetRecipeCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecipeCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRecipeCSVOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRecipeCSVBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRecipeCSVNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecipeCSVOK creates a GetRecipeCSVOK with default headers values
func NewGetRecipeCSVOK() *GetRecipeCSVOK {
	return &GetRecipeCSVOK{}
}

/*GetRecipeCSVOK handles this case with default header values.

Description was not specified
*/
type GetRecipeCSVOK struct {
	Payload string
}

func (o *GetRecipeCSVOK) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/csv][%d] getRecipeCSVOK  %+v", 200, o.Payload)
}

func (o *GetRecipeCSVOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipeCSVBadRequest creates a GetRecipeCSVBadRequest with default headers values
func NewGetRecipeCSVBadRequest() *GetRecipeCSVBadRequest {
	return &GetRecipeCSVBadRequest{}
}

/*GetRecipeCSVBadRequest handles this case with default header values.

Description was not specified
*/
type GetRecipeCSVBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *GetRecipeCSVBadRequest) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/csv][%d] getRecipeCSVBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecipeCSVBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipeCSVNotFound creates a GetRecipeCSVNotFound with default headers values
func NewGetRecipeCSVNotFound() *GetRecipeCSVNotFound {
	return &GetRecipeCSVNotFound{}
}

/*GetRecipeCSVNotFound handles this case with default header values.

Description was not specified
*/
type GetRecipeCSVNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *GetRecipeCSVNotFound) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/csv][%d] getRecipeCSVNotFound  %+v", 404, o.Payload)
}

func (o *GetRecipeCSVNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
