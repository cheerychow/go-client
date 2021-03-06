// Code generated by go-swagger; DO NOT EDIT.

package nutrition_tips

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetNutritionTipsLabelHTMLReader is a Reader for the GetNutritionTipsLabelHTML structure.
type GetNutritionTipsLabelHTMLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNutritionTipsLabelHTMLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNutritionTipsLabelHTMLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetNutritionTipsLabelHTMLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetNutritionTipsLabelHTMLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNutritionTipsLabelHTMLOK creates a GetNutritionTipsLabelHTMLOK with default headers values
func NewGetNutritionTipsLabelHTMLOK() *GetNutritionTipsLabelHTMLOK {
	return &GetNutritionTipsLabelHTMLOK{}
}

/*GetNutritionTipsLabelHTMLOK handles this case with default header values.

Description was not specified
*/
type GetNutritionTipsLabelHTMLOK struct {
	Payload string
}

func (o *GetNutritionTipsLabelHTMLOK) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/nutrition-tips-label/html][%d] getNutritionTipsLabelHtmlOK  %+v", 200, o.Payload)
}

func (o *GetNutritionTipsLabelHTMLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNutritionTipsLabelHTMLBadRequest creates a GetNutritionTipsLabelHTMLBadRequest with default headers values
func NewGetNutritionTipsLabelHTMLBadRequest() *GetNutritionTipsLabelHTMLBadRequest {
	return &GetNutritionTipsLabelHTMLBadRequest{}
}

/*GetNutritionTipsLabelHTMLBadRequest handles this case with default header values.

Description was not specified
*/
type GetNutritionTipsLabelHTMLBadRequest struct {
	Payload string
}

func (o *GetNutritionTipsLabelHTMLBadRequest) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/nutrition-tips-label/html][%d] getNutritionTipsLabelHtmlBadRequest  %+v", 400, o.Payload)
}

func (o *GetNutritionTipsLabelHTMLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNutritionTipsLabelHTMLNotFound creates a GetNutritionTipsLabelHTMLNotFound with default headers values
func NewGetNutritionTipsLabelHTMLNotFound() *GetNutritionTipsLabelHTMLNotFound {
	return &GetNutritionTipsLabelHTMLNotFound{}
}

/*GetNutritionTipsLabelHTMLNotFound handles this case with default header values.

Description was not specified
*/
type GetNutritionTipsLabelHTMLNotFound struct {
	Payload string
}

func (o *GetNutritionTipsLabelHTMLNotFound) Error() string {
	return fmt.Sprintf("[GET /recipe/{recipeId}/nutrition-tips-label/html][%d] getNutritionTipsLabelHtmlNotFound  %+v", 404, o.Payload)
}

func (o *GetNutritionTipsLabelHTMLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
