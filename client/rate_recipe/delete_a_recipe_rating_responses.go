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

// DeleteARecipeRatingReader is a Reader for the DeleteARecipeRating structure.
type DeleteARecipeRatingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteARecipeRatingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteARecipeRatingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteARecipeRatingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteARecipeRatingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteARecipeRatingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteARecipeRatingOK creates a DeleteARecipeRatingOK with default headers values
func NewDeleteARecipeRatingOK() *DeleteARecipeRatingOK {
	return &DeleteARecipeRatingOK{}
}

/*DeleteARecipeRatingOK handles this case with default header values.

Description was not specified
*/
type DeleteARecipeRatingOK struct {
	Payload *models.HTTPAPIResponse
}

func (o *DeleteARecipeRatingOK) Error() string {
	return fmt.Sprintf("[DELETE /rate-recipe/{ratingId}][%d] deleteARecipeRatingOK  %+v", 200, o.Payload)
}

func (o *DeleteARecipeRatingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteARecipeRatingBadRequest creates a DeleteARecipeRatingBadRequest with default headers values
func NewDeleteARecipeRatingBadRequest() *DeleteARecipeRatingBadRequest {
	return &DeleteARecipeRatingBadRequest{}
}

/*DeleteARecipeRatingBadRequest handles this case with default header values.

Description was not specified
*/
type DeleteARecipeRatingBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *DeleteARecipeRatingBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /rate-recipe/{ratingId}][%d] deleteARecipeRatingBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteARecipeRatingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteARecipeRatingForbidden creates a DeleteARecipeRatingForbidden with default headers values
func NewDeleteARecipeRatingForbidden() *DeleteARecipeRatingForbidden {
	return &DeleteARecipeRatingForbidden{}
}

/*DeleteARecipeRatingForbidden handles this case with default header values.

You're not the owner of this rating.
*/
type DeleteARecipeRatingForbidden struct {
	Payload *models.HTTPAPIError
}

func (o *DeleteARecipeRatingForbidden) Error() string {
	return fmt.Sprintf("[DELETE /rate-recipe/{ratingId}][%d] deleteARecipeRatingForbidden  %+v", 403, o.Payload)
}

func (o *DeleteARecipeRatingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteARecipeRatingNotFound creates a DeleteARecipeRatingNotFound with default headers values
func NewDeleteARecipeRatingNotFound() *DeleteARecipeRatingNotFound {
	return &DeleteARecipeRatingNotFound{}
}

/*DeleteARecipeRatingNotFound handles this case with default header values.

Rating not found.
*/
type DeleteARecipeRatingNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *DeleteARecipeRatingNotFound) Error() string {
	return fmt.Sprintf("[DELETE /rate-recipe/{ratingId}][%d] deleteARecipeRatingNotFound  %+v", 404, o.Payload)
}

func (o *DeleteARecipeRatingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
