// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SoftwarePackageCreateReader is a Reader for the SoftwarePackageCreate structure.
type SoftwarePackageCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwarePackageCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSoftwarePackageCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwarePackageCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwarePackageCreateAccepted creates a SoftwarePackageCreateAccepted with default headers values
func NewSoftwarePackageCreateAccepted() *SoftwarePackageCreateAccepted {
	return &SoftwarePackageCreateAccepted{}
}

/* SoftwarePackageCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SoftwarePackageCreateAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SoftwarePackageCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /cluster/software/download][%d] softwarePackageCreateAccepted  %+v", 202, o.Payload)
}
func (o *SoftwarePackageCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SoftwarePackageCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwarePackageCreateDefault creates a SoftwarePackageCreateDefault with default headers values
func NewSoftwarePackageCreateDefault(code int) *SoftwarePackageCreateDefault {
	return &SoftwarePackageCreateDefault{
		_statusCode: code,
	}
}

/* SoftwarePackageCreateDefault describes a response with status code -1, with default header values.

Error
*/
type SoftwarePackageCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the software package create default response
func (o *SoftwarePackageCreateDefault) Code() int {
	return o._statusCode
}

func (o *SoftwarePackageCreateDefault) Error() string {
	return fmt.Sprintf("[POST /cluster/software/download][%d] software_package_create default  %+v", o._statusCode, o.Payload)
}
func (o *SoftwarePackageCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SoftwarePackageCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}