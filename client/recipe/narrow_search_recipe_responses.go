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

// NarrowSearchRecipeReader is a Reader for the NarrowSearchRecipe structure.
type NarrowSearchRecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NarrowSearchRecipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNarrowSearchRecipeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewNarrowSearchRecipeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNarrowSearchRecipeOK creates a NarrowSearchRecipeOK with default headers values
func NewNarrowSearchRecipeOK() *NarrowSearchRecipeOK {
	return &NarrowSearchRecipeOK{}
}

/*NarrowSearchRecipeOK handles this case with default header values.

Description was not specified
*/
type NarrowSearchRecipeOK struct {
	Payload *models.PagedRecipeResults
}

func (o *NarrowSearchRecipeOK) Error() string {
	return fmt.Sprintf("[GET /recipe/narrow-search][%d] narrowSearchRecipeOK  %+v", 200, o.Payload)
}

func (o *NarrowSearchRecipeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PagedRecipeResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNarrowSearchRecipeBadRequest creates a NarrowSearchRecipeBadRequest with default headers values
func NewNarrowSearchRecipeBadRequest() *NarrowSearchRecipeBadRequest {
	return &NarrowSearchRecipeBadRequest{}
}

/*NarrowSearchRecipeBadRequest handles this case with default header values.

Description was not specified
*/
type NarrowSearchRecipeBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *NarrowSearchRecipeBadRequest) Error() string {
	return fmt.Sprintf("[GET /recipe/narrow-search][%d] narrowSearchRecipeBadRequest  %+v", 400, o.Payload)
}

func (o *NarrowSearchRecipeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
