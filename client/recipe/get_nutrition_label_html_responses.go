// Code generated by go-swagger; DO NOT EDIT.

package recipe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetNutritionLabelHTMLReader is a Reader for the GetNutritionLabelHTML structure.
type GetNutritionLabelHTMLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNutritionLabelHTMLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNutritionLabelHTMLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetNutritionLabelHTMLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetNutritionLabelHTMLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNutritionLabelHTMLOK creates a GetNutritionLabelHTMLOK with default headers values
func NewGetNutritionLabelHTMLOK() *GetNutritionLabelHTMLOK {
	return &GetNutritionLabelHTMLOK{}
}

/*GetNutritionLabelHTMLOK handles this case with default header values.

Description was not specified
*/
type GetNutritionLabelHTMLOK struct {
	Payload string
}

func (o *GetNutritionLabelHTMLOK) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/nutrition-label/html][%d] getNutritionLabelHtmlOK  %+v", 200, o.Payload)
}

func (o *GetNutritionLabelHTMLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNutritionLabelHTMLBadRequest creates a GetNutritionLabelHTMLBadRequest with default headers values
func NewGetNutritionLabelHTMLBadRequest() *GetNutritionLabelHTMLBadRequest {
	return &GetNutritionLabelHTMLBadRequest{}
}

/*GetNutritionLabelHTMLBadRequest handles this case with default header values.

Description was not specified
*/
type GetNutritionLabelHTMLBadRequest struct {
	Payload string
}

func (o *GetNutritionLabelHTMLBadRequest) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/nutrition-label/html][%d] getNutritionLabelHtmlBadRequest  %+v", 400, o.Payload)
}

func (o *GetNutritionLabelHTMLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNutritionLabelHTMLNotFound creates a GetNutritionLabelHTMLNotFound with default headers values
func NewGetNutritionLabelHTMLNotFound() *GetNutritionLabelHTMLNotFound {
	return &GetNutritionLabelHTMLNotFound{}
}

/*GetNutritionLabelHTMLNotFound handles this case with default header values.

Description was not specified
*/
type GetNutritionLabelHTMLNotFound struct {
	Payload string
}

func (o *GetNutritionLabelHTMLNotFound) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/nutrition-label/html][%d] getNutritionLabelHtmlNotFound  %+v", 404, o.Payload)
}

func (o *GetNutritionLabelHTMLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}