// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// SaveUserInfoReader is a Reader for the SaveUserInfo structure.
type SaveUserInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveUserInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSaveUserInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewSaveUserInfoCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSaveUserInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSaveUserInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSaveUserInfoOK creates a SaveUserInfoOK with default headers values
func NewSaveUserInfoOK() *SaveUserInfoOK {
	return &SaveUserInfoOK{}
}

/*SaveUserInfoOK handles this case with default header values.

No response was specified
*/
type SaveUserInfoOK struct {
	Payload *models.User
}

func (o *SaveUserInfoOK) Error() string {
	return fmt.Sprintf("[PUT /user/{userId}][%d] saveUserInfoOK  %+v", 200, o.Payload)
}

func (o *SaveUserInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveUserInfoCreated creates a SaveUserInfoCreated with default headers values
func NewSaveUserInfoCreated() *SaveUserInfoCreated {
	return &SaveUserInfoCreated{}
}

/*SaveUserInfoCreated handles this case with default header values.

Description was not specified
*/
type SaveUserInfoCreated struct {
	Payload *models.User
}

func (o *SaveUserInfoCreated) Error() string {
	return fmt.Sprintf("[PUT /user/{userId}][%d] saveUserInfoCreated  %+v", 201, o.Payload)
}

func (o *SaveUserInfoCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveUserInfoBadRequest creates a SaveUserInfoBadRequest with default headers values
func NewSaveUserInfoBadRequest() *SaveUserInfoBadRequest {
	return &SaveUserInfoBadRequest{}
}

/*SaveUserInfoBadRequest handles this case with default header values.

Description was not specified
*/
type SaveUserInfoBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *SaveUserInfoBadRequest) Error() string {
	return fmt.Sprintf("[PUT /user/{userId}][%d] saveUserInfoBadRequest  %+v", 400, o.Payload)
}

func (o *SaveUserInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveUserInfoUnauthorized creates a SaveUserInfoUnauthorized with default headers values
func NewSaveUserInfoUnauthorized() *SaveUserInfoUnauthorized {
	return &SaveUserInfoUnauthorized{}
}

/*SaveUserInfoUnauthorized handles this case with default header values.

Not authorised
*/
type SaveUserInfoUnauthorized struct {
	Payload *models.HTTPAPIError
}

func (o *SaveUserInfoUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /user/{userId}][%d] saveUserInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *SaveUserInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}