// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// FacebookSignupSigninReader is a Reader for the FacebookSignupSignin structure.
type FacebookSignupSigninReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FacebookSignupSigninReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewFacebookSignupSigninCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewFacebookSignupSigninBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFacebookSignupSigninCreated creates a FacebookSignupSigninCreated with default headers values
func NewFacebookSignupSigninCreated() *FacebookSignupSigninCreated {
	return &FacebookSignupSigninCreated{}
}

/*FacebookSignupSigninCreated handles this case with default header values.

Description was not specified
*/
type FacebookSignupSigninCreated struct {
	Payload *models.User
}

func (o *FacebookSignupSigninCreated) Error() string {
	return fmt.Sprintf("[POST /fb_signin][%d] facebookSignupSigninCreated  %+v", 201, o.Payload)
}

func (o *FacebookSignupSigninCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFacebookSignupSigninBadRequest creates a FacebookSignupSigninBadRequest with default headers values
func NewFacebookSignupSigninBadRequest() *FacebookSignupSigninBadRequest {
	return &FacebookSignupSigninBadRequest{}
}

/*FacebookSignupSigninBadRequest handles this case with default header values.

Description was not specified
*/
type FacebookSignupSigninBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *FacebookSignupSigninBadRequest) Error() string {
	return fmt.Sprintf("[POST /fb_signin][%d] facebookSignupSigninBadRequest  %+v", 400, o.Payload)
}

func (o *FacebookSignupSigninBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}