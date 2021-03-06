// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/chow-go-client/models"
)

// AccountSignUpReader is a Reader for the AccountSignUp structure.
type AccountSignUpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountSignUpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAccountSignUpCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAccountSignUpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAccountSignUpCreated creates a AccountSignUpCreated with default headers values
func NewAccountSignUpCreated() *AccountSignUpCreated {
	return &AccountSignUpCreated{}
}

/*AccountSignUpCreated handles this case with default header values.

Description was not specified
*/
type AccountSignUpCreated struct {
	Payload *models.User
}

func (o *AccountSignUpCreated) Error() string {
	return fmt.Sprintf("[POST /signup][%d] accountSignUpCreated  %+v", 201, o.Payload)
}

func (o *AccountSignUpCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountSignUpBadRequest creates a AccountSignUpBadRequest with default headers values
func NewAccountSignUpBadRequest() *AccountSignUpBadRequest {
	return &AccountSignUpBadRequest{}
}

/*AccountSignUpBadRequest handles this case with default header values.

Description was not specified
*/
type AccountSignUpBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *AccountSignUpBadRequest) Error() string {
	return fmt.Sprintf("[POST /signup][%d] accountSignUpBadRequest  %+v", 400, o.Payload)
}

func (o *AccountSignUpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
