// Code generated by go-swagger; DO NOT EDIT.

package food

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// FoodWeightNamesReader is a Reader for the FoodWeightNames structure.
type FoodWeightNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FoodWeightNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFoodWeightNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewFoodWeightNamesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFoodWeightNamesOK creates a FoodWeightNamesOK with default headers values
func NewFoodWeightNamesOK() *FoodWeightNamesOK {
	return &FoodWeightNamesOK{}
}

/*FoodWeightNamesOK handles this case with default header values.

Description was not specified
*/
type FoodWeightNamesOK struct {
	Payload []string
}

func (o *FoodWeightNamesOK) Error() string {
	return fmt.Sprintf("[GET /food/{foodId}/weightnames][%d] foodWeightNamesOK  %+v", 200, o.Payload)
}

func (o *FoodWeightNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFoodWeightNamesBadRequest creates a FoodWeightNamesBadRequest with default headers values
func NewFoodWeightNamesBadRequest() *FoodWeightNamesBadRequest {
	return &FoodWeightNamesBadRequest{}
}

/*FoodWeightNamesBadRequest handles this case with default header values.

Description was not specified
*/
type FoodWeightNamesBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *FoodWeightNamesBadRequest) Error() string {
	return fmt.Sprintf("[GET /food/{foodId}/weightnames][%d] foodWeightNamesBadRequest  %+v", 400, o.Payload)
}

func (o *FoodWeightNamesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
