// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NvmeNamespaceModifyReader is a Reader for the NvmeNamespaceModify structure.
type NvmeNamespaceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeNamespaceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeNamespaceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewNvmeNamespaceModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeNamespaceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeNamespaceModifyOK creates a NvmeNamespaceModifyOK with default headers values
func NewNvmeNamespaceModifyOK() *NvmeNamespaceModifyOK {
	return &NvmeNamespaceModifyOK{}
}

/*
NvmeNamespaceModifyOK describes a response with status code 200, with default header values.

OK
*/
type NvmeNamespaceModifyOK struct {
}

// IsSuccess returns true when this nvme namespace modify o k response has a 2xx status code
func (o *NvmeNamespaceModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme namespace modify o k response has a 3xx status code
func (o *NvmeNamespaceModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme namespace modify o k response has a 4xx status code
func (o *NvmeNamespaceModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme namespace modify o k response has a 5xx status code
func (o *NvmeNamespaceModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme namespace modify o k response a status code equal to that given
func (o *NvmeNamespaceModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme namespace modify o k response
func (o *NvmeNamespaceModifyOK) Code() int {
	return 200
}

func (o *NvmeNamespaceModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/namespaces/{uuid}][%d] nvmeNamespaceModifyOK", 200)
}

func (o *NvmeNamespaceModifyOK) String() string {
	return fmt.Sprintf("[PATCH /storage/namespaces/{uuid}][%d] nvmeNamespaceModifyOK", 200)
}

func (o *NvmeNamespaceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNvmeNamespaceModifyAccepted creates a NvmeNamespaceModifyAccepted with default headers values
func NewNvmeNamespaceModifyAccepted() *NvmeNamespaceModifyAccepted {
	return &NvmeNamespaceModifyAccepted{}
}

/*
NvmeNamespaceModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type NvmeNamespaceModifyAccepted struct {
	Payload *models.NvmeNamespaceJobLinkResponse
}

// IsSuccess returns true when this nvme namespace modify accepted response has a 2xx status code
func (o *NvmeNamespaceModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme namespace modify accepted response has a 3xx status code
func (o *NvmeNamespaceModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme namespace modify accepted response has a 4xx status code
func (o *NvmeNamespaceModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme namespace modify accepted response has a 5xx status code
func (o *NvmeNamespaceModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme namespace modify accepted response a status code equal to that given
func (o *NvmeNamespaceModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the nvme namespace modify accepted response
func (o *NvmeNamespaceModifyAccepted) Code() int {
	return 202
}

func (o *NvmeNamespaceModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/namespaces/{uuid}][%d] nvmeNamespaceModifyAccepted %s", 202, payload)
}

func (o *NvmeNamespaceModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/namespaces/{uuid}][%d] nvmeNamespaceModifyAccepted %s", 202, payload)
}

func (o *NvmeNamespaceModifyAccepted) GetPayload() *models.NvmeNamespaceJobLinkResponse {
	return o.Payload
}

func (o *NvmeNamespaceModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeNamespaceJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeNamespaceModifyDefault creates a NvmeNamespaceModifyDefault with default headers values
func NewNvmeNamespaceModifyDefault(code int) *NvmeNamespaceModifyDefault {
	return &NvmeNamespaceModifyDefault{
		_statusCode: code,
	}
}

/*
	NvmeNamespaceModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374127 | The specified namespace name is invalid. |
| 5376461 | The specified namespace name is invalid. |
| 5376462 | The specified namespace name is too long. |
| 5376463 | The snapshot portion of the specified namespace name is too long. |
| 5376466 | An attempt was made to rename an NVMe namespace to a snapshot name. |
| 5376467 | An attempt was made to rename a primary NVMe namespace to a secondary name. |
| 5376468 | An attempt was made to rename an NVMe namespace to a reserved name. |
| 13565952 | The namespace clone request failed. |
| 72089724 | The specified namespace size is too large. |
| 72089730 | The specified namespace cannot be updated as it resides in a snapshot. |
| 72090005 | The specified `clone.source.uuid` and `clone.source.name` do not refer to the same LUN. |
| 72090006 | The specified namespace was not found. This can apply to `clone.source` or the target namespace. The `target` property of the error object identifies the property. |
| 72090007 | The specified namespace was not found. This can apply to `clone.source` or the target namespace. The `target` property of the error object identifies the property. |
| 72090010 | An error occurred after successfully overwriting data for the namespace as a clone. Some properties were not modified. |
| 72090011 | An error occurred after successfully modifying some of the properties of the namespace. Some properties were not modified. |
| 72090016 | The namespace's aggregate is offline. The aggregate must be online to modify or remove the namespace. |
| 72090017 | The namespace's volume is offline. The volume must be online to modify or remove the namespace. |
| 72090038 | An attempt was made to reduce the size of the specified namespace. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeNamespaceModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme namespace modify default response has a 2xx status code
func (o *NvmeNamespaceModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme namespace modify default response has a 3xx status code
func (o *NvmeNamespaceModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme namespace modify default response has a 4xx status code
func (o *NvmeNamespaceModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme namespace modify default response has a 5xx status code
func (o *NvmeNamespaceModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme namespace modify default response a status code equal to that given
func (o *NvmeNamespaceModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme namespace modify default response
func (o *NvmeNamespaceModifyDefault) Code() int {
	return o._statusCode
}

func (o *NvmeNamespaceModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/namespaces/{uuid}][%d] nvme_namespace_modify default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/namespaces/{uuid}][%d] nvme_namespace_modify default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeNamespaceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
