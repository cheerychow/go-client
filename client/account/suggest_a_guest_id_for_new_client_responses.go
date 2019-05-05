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

// SuggestAGuestIDForNewClientReader is a Reader for the SuggestAGuestIDForNewClient structure.
type SuggestAGuestIDForNewClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuggestAGuestIDForNewClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSuggestAGuestIDForNewClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSuggestAGuestIDForNewClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSuggestAGuestIDForNewClientOK creates a SuggestAGuestIDForNewClientOK with default headers values
func NewSuggestAGuestIDForNewClientOK() *SuggestAGuestIDForNewClientOK {
	return &SuggestAGuestIDForNewClientOK{}
}

/*SuggestAGuestIDForNewClientOK handles this case with default header values.

Description was not specified
*/
type SuggestAGuestIDForNewClientOK struct {
	Payload *models.User
}

func (o *SuggestAGuestIDForNewClientOK) Error() string {
	return fmt.Sprintf("[POST /request-guestid][%d] suggestAGuestIdForNewClientOK  %+v", 200, o.Payload)
}

func (o *SuggestAGuestIDForNewClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestAGuestIDForNewClientBadRequest creates a SuggestAGuestIDForNewClientBadRequest with default headers values
func NewSuggestAGuestIDForNewClientBadRequest() *SuggestAGuestIDForNewClientBadRequest {
	return &SuggestAGuestIDForNewClientBadRequest{}
}

/*SuggestAGuestIDForNewClientBadRequest handles this case with default header values.

Not signed in.
*/
type SuggestAGuestIDForNewClientBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *SuggestAGuestIDForNewClientBadRequest) Error() string {
	return fmt.Sprintf("[POST /request-guestid][%d] suggestAGuestIdForNewClientBadRequest  %+v", 400, o.Payload)
}

func (o *SuggestAGuestIDForNewClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
