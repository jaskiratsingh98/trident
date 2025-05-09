// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NfsCollectionPerformanceMetricsGetReader is a Reader for the NfsCollectionPerformanceMetricsGet structure.
type NfsCollectionPerformanceMetricsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsCollectionPerformanceMetricsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsCollectionPerformanceMetricsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsCollectionPerformanceMetricsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsCollectionPerformanceMetricsGetOK creates a NfsCollectionPerformanceMetricsGetOK with default headers values
func NewNfsCollectionPerformanceMetricsGetOK() *NfsCollectionPerformanceMetricsGetOK {
	return &NfsCollectionPerformanceMetricsGetOK{}
}

/*
NfsCollectionPerformanceMetricsGetOK describes a response with status code 200, with default header values.

OK
*/
type NfsCollectionPerformanceMetricsGetOK struct {
	Payload *models.PerformanceSvmNfsResponse
}

// IsSuccess returns true when this nfs collection performance metrics get o k response has a 2xx status code
func (o *NfsCollectionPerformanceMetricsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nfs collection performance metrics get o k response has a 3xx status code
func (o *NfsCollectionPerformanceMetricsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nfs collection performance metrics get o k response has a 4xx status code
func (o *NfsCollectionPerformanceMetricsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nfs collection performance metrics get o k response has a 5xx status code
func (o *NfsCollectionPerformanceMetricsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nfs collection performance metrics get o k response a status code equal to that given
func (o *NfsCollectionPerformanceMetricsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nfs collection performance metrics get o k response
func (o *NfsCollectionPerformanceMetricsGetOK) Code() int {
	return 200
}

func (o *NfsCollectionPerformanceMetricsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/services/{svm.uuid}/metrics][%d] nfsCollectionPerformanceMetricsGetOK %s", 200, payload)
}

func (o *NfsCollectionPerformanceMetricsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/services/{svm.uuid}/metrics][%d] nfsCollectionPerformanceMetricsGetOK %s", 200, payload)
}

func (o *NfsCollectionPerformanceMetricsGetOK) GetPayload() *models.PerformanceSvmNfsResponse {
	return o.Payload
}

func (o *NfsCollectionPerformanceMetricsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceSvmNfsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNfsCollectionPerformanceMetricsGetDefault creates a NfsCollectionPerformanceMetricsGetDefault with default headers values
func NewNfsCollectionPerformanceMetricsGetDefault(code int) *NfsCollectionPerformanceMetricsGetDefault {
	return &NfsCollectionPerformanceMetricsGetDefault{
		_statusCode: code,
	}
}

/*
NfsCollectionPerformanceMetricsGetDefault describes a response with status code -1, with default header values.

Error
*/
type NfsCollectionPerformanceMetricsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nfs collection performance metrics get default response has a 2xx status code
func (o *NfsCollectionPerformanceMetricsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nfs collection performance metrics get default response has a 3xx status code
func (o *NfsCollectionPerformanceMetricsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nfs collection performance metrics get default response has a 4xx status code
func (o *NfsCollectionPerformanceMetricsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nfs collection performance metrics get default response has a 5xx status code
func (o *NfsCollectionPerformanceMetricsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nfs collection performance metrics get default response a status code equal to that given
func (o *NfsCollectionPerformanceMetricsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nfs collection performance metrics get default response
func (o *NfsCollectionPerformanceMetricsGetDefault) Code() int {
	return o._statusCode
}

func (o *NfsCollectionPerformanceMetricsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/services/{svm.uuid}/metrics][%d] nfs_collection_performance_metrics_get default %s", o._statusCode, payload)
}

func (o *NfsCollectionPerformanceMetricsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/services/{svm.uuid}/metrics][%d] nfs_collection_performance_metrics_get default %s", o._statusCode, payload)
}

func (o *NfsCollectionPerformanceMetricsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsCollectionPerformanceMetricsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
