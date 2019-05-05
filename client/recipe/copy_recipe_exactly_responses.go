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

// CopyRecipeExactlyReader is a Reader for the CopyRecipeExactly structure.
type CopyRecipeExactlyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyRecipeExactlyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCopyRecipeExactlyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCopyRecipeExactlyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCopyRecipeExactlyOK creates a CopyRecipeExactlyOK with default headers values
func NewCopyRecipeExactlyOK() *CopyRecipeExactlyOK {
	return &CopyRecipeExactlyOK{}
}

/*CopyRecipeExactlyOK handles this case with default header values.

Description was not specified
*/
type CopyRecipeExactlyOK struct {
	Payload *models.Recipe
}

func (o *CopyRecipeExactlyOK) Error() string {
	return fmt.Sprintf("[GET /recipe/copy/{recipeId}][%d] copyRecipeExactlyOK  %+v", 200, o.Payload)
}

func (o *CopyRecipeExactlyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recipe)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyRecipeExactlyBadRequest creates a CopyRecipeExactlyBadRequest with default headers values
func NewCopyRecipeExactlyBadRequest() *CopyRecipeExactlyBadRequest {
	return &CopyRecipeExactlyBadRequest{}
}

/*CopyRecipeExactlyBadRequest handles this case with default header values.

Description was not specified
*/
type CopyRecipeExactlyBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *CopyRecipeExactlyBadRequest) Error() string {
	return fmt.Sprintf("[GET /recipe/copy/{recipeId}][%d] copyRecipeExactlyBadRequest  %+v", 400, o.Payload)
}

func (o *CopyRecipeExactlyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
