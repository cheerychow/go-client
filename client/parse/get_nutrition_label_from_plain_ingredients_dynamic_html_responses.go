// Code generated by go-swagger; DO NOT EDIT.

package parse

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetNutritionLabelFromPlainIngredientsDynamicHTMLReader is a Reader for the GetNutritionLabelFromPlainIngredientsDynamicHTML structure.
type GetNutritionLabelFromPlainIngredientsDynamicHTMLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNutritionLabelFromPlainIngredientsDynamicHTMLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNutritionLabelFromPlainIngredientsDynamicHTMLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNutritionLabelFromPlainIngredientsDynamicHTMLOK creates a GetNutritionLabelFromPlainIngredientsDynamicHTMLOK with default headers values
func NewGetNutritionLabelFromPlainIngredientsDynamicHTMLOK() *GetNutritionLabelFromPlainIngredientsDynamicHTMLOK {
	return &GetNutritionLabelFromPlainIngredientsDynamicHTMLOK{}
}

/*GetNutritionLabelFromPlainIngredientsDynamicHTMLOK handles this case with default header values.

Description was not specified
*/
type GetNutritionLabelFromPlainIngredientsDynamicHTMLOK struct {
	Payload string
}

func (o *GetNutritionLabelFromPlainIngredientsDynamicHTMLOK) Error() string {
	return fmt.Sprintf("[POST /parse/nutrition-label/dynamic-html][%d] getNutritionLabelFromPlainIngredientsDynamicHtmlOK  %+v", 200, o.Payload)
}

func (o *GetNutritionLabelFromPlainIngredientsDynamicHTMLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest creates a GetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest with default headers values
func NewGetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest() *GetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest {
	return &GetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest{}
}

/*GetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest handles this case with default header values.

Description was not specified
*/
type GetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest struct {
	Payload string
}

func (o *GetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest) Error() string {
	return fmt.Sprintf("[POST /parse/nutrition-label/dynamic-html][%d] getNutritionLabelFromPlainIngredientsDynamicHtmlBadRequest  %+v", 400, o.Payload)
}

func (o *GetNutritionLabelFromPlainIngredientsDynamicHTMLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound creates a GetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound with default headers values
func NewGetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound() *GetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound {
	return &GetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound{}
}

/*GetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound handles this case with default header values.

Description was not specified
*/
type GetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound struct {
	Payload string
}

func (o *GetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound) Error() string {
	return fmt.Sprintf("[POST /parse/nutrition-label/dynamic-html][%d] getNutritionLabelFromPlainIngredientsDynamicHtmlNotFound  %+v", 404, o.Payload)
}

func (o *GetNutritionLabelFromPlainIngredientsDynamicHTMLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
