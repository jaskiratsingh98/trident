// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// UnixGroupUsersGetReader is a Reader for the UnixGroupUsersGet structure.
type UnixGroupUsersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupUsersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixGroupUsersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupUsersGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupUsersGetOK creates a UnixGroupUsersGetOK with default headers values
func NewUnixGroupUsersGetOK() *UnixGroupUsersGetOK {
	return &UnixGroupUsersGetOK{}
}

/* UnixGroupUsersGetOK describes a response with status code 200, with default header values.

OK
*/
type UnixGroupUsersGetOK struct {
	Payload *models.UnixGroupUsers
}

func (o *UnixGroupUsersGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{unix_group.name}/users/{name}][%d] unixGroupUsersGetOK  %+v", 200, o.Payload)
}
func (o *UnixGroupUsersGetOK) GetPayload() *models.UnixGroupUsers {
	return o.Payload
}

func (o *UnixGroupUsersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnixGroupUsers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnixGroupUsersGetDefault creates a UnixGroupUsersGetDefault with default headers values
func NewUnixGroupUsersGetDefault(code int) *UnixGroupUsersGetDefault {
	return &UnixGroupUsersGetDefault{
		_statusCode: code,
	}
}

/* UnixGroupUsersGetDefault describes a response with status code -1, with default header values.

Error
*/
type UnixGroupUsersGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the unix group users get default response
func (o *UnixGroupUsersGetDefault) Code() int {
	return o._statusCode
}

func (o *UnixGroupUsersGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{unix_group.name}/users/{name}][%d] unix_group_users_get default  %+v", o._statusCode, o.Payload)
}
func (o *UnixGroupUsersGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupUsersGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}