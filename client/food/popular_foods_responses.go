// Code generated by go-swagger; DO NOT EDIT.

package food

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/chow-go-client/models"
)

// PopularFoodsReader is a Reader for the PopularFoods structure.
type PopularFoodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PopularFoodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPopularFoodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPopularFoodsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPopularFoodsOK creates a PopularFoodsOK with default headers values
func NewPopularFoodsOK() *PopularFoodsOK {
	return &PopularFoodsOK{}
}

/*PopularFoodsOK handles this case with default header values.

Description was not specified
*/
type PopularFoodsOK struct {
	Payload []*models.PopularFood
}

func (o *PopularFoodsOK) Error() string {
	return fmt.Sprintf("[GET /food/popular][%d] popularFoodsOK  %+v", 200, o.Payload)
}

func (o *PopularFoodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPopularFoodsBadRequest creates a PopularFoodsBadRequest with default headers values
func NewPopularFoodsBadRequest() *PopularFoodsBadRequest {
	return &PopularFoodsBadRequest{}
}

/*PopularFoodsBadRequest handles this case with default header values.

Description was not specified
*/
type PopularFoodsBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *PopularFoodsBadRequest) Error() string {
	return fmt.Sprintf("[GET /food/popular][%d] popularFoodsBadRequest  %+v", 400, o.Payload)
}

func (o *PopularFoodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
