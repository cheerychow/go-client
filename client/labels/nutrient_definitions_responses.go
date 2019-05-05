// Code generated by go-swagger; DO NOT EDIT.

package labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/chow-go-client/models"
)

// NutrientDefinitionsReader is a Reader for the NutrientDefinitions structure.
type NutrientDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NutrientDefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNutrientDefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewNutrientDefinitionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNutrientDefinitionsOK creates a NutrientDefinitionsOK with default headers values
func NewNutrientDefinitionsOK() *NutrientDefinitionsOK {
	return &NutrientDefinitionsOK{}
}

/*NutrientDefinitionsOK handles this case with default header values.

Description was not specified
*/
type NutrientDefinitionsOK struct {
	Payload []*models.NutrDef
}

func (o *NutrientDefinitionsOK) Error() string {
	return fmt.Sprintf("[GET /labels/nutrient-definitions][%d] nutrientDefinitionsOK  %+v", 200, o.Payload)
}

func (o *NutrientDefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNutrientDefinitionsBadRequest creates a NutrientDefinitionsBadRequest with default headers values
func NewNutrientDefinitionsBadRequest() *NutrientDefinitionsBadRequest {
	return &NutrientDefinitionsBadRequest{}
}

/*NutrientDefinitionsBadRequest handles this case with default header values.

Description was not specified
*/
type NutrientDefinitionsBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *NutrientDefinitionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /labels/nutrient-definitions][%d] nutrientDefinitionsBadRequest  %+v", 400, o.Payload)
}

func (o *NutrientDefinitionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
