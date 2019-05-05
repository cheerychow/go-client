// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/chow-go-client/models"
)

// GetPublicAccountInfoReader is a Reader for the GetPublicAccountInfo structure.
type GetPublicAccountInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicAccountInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicAccountInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPublicAccountInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPublicAccountInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicAccountInfoOK creates a GetPublicAccountInfoOK with default headers values
func NewGetPublicAccountInfoOK() *GetPublicAccountInfoOK {
	return &GetPublicAccountInfoOK{}
}

/*GetPublicAccountInfoOK handles this case with default header values.

Description was not specified
*/
type GetPublicAccountInfoOK struct {
	Payload *models.User
}

func (o *GetPublicAccountInfoOK) Error() string {
	return fmt.Sprintf("[GET /account][%d] getPublicAccountInfoOK  %+v", 200, o.Payload)
}

func (o *GetPublicAccountInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicAccountInfoBadRequest creates a GetPublicAccountInfoBadRequest with default headers values
func NewGetPublicAccountInfoBadRequest() *GetPublicAccountInfoBadRequest {
	return &GetPublicAccountInfoBadRequest{}
}

/*GetPublicAccountInfoBadRequest handles this case with default header values.

Description was not specified
*/
type GetPublicAccountInfoBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *GetPublicAccountInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /account][%d] getPublicAccountInfoBadRequest  %+v", 400, o.Payload)
}

func (o *GetPublicAccountInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicAccountInfoNotFound creates a GetPublicAccountInfoNotFound with default headers values
func NewGetPublicAccountInfoNotFound() *GetPublicAccountInfoNotFound {
	return &GetPublicAccountInfoNotFound{}
}

/*GetPublicAccountInfoNotFound handles this case with default header values.

Description was not specified
*/
type GetPublicAccountInfoNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *GetPublicAccountInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /account][%d] getPublicAccountInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetPublicAccountInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
