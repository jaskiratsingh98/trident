// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSnaplockFingerprintOperationCollectionGetParams creates a new SnaplockFingerprintOperationCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockFingerprintOperationCollectionGetParams() *SnaplockFingerprintOperationCollectionGetParams {
	return &SnaplockFingerprintOperationCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockFingerprintOperationCollectionGetParamsWithTimeout creates a new SnaplockFingerprintOperationCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSnaplockFingerprintOperationCollectionGetParamsWithTimeout(timeout time.Duration) *SnaplockFingerprintOperationCollectionGetParams {
	return &SnaplockFingerprintOperationCollectionGetParams{
		timeout: timeout,
	}
}

// NewSnaplockFingerprintOperationCollectionGetParamsWithContext creates a new SnaplockFingerprintOperationCollectionGetParams object
// with the ability to set a context for a request.
func NewSnaplockFingerprintOperationCollectionGetParamsWithContext(ctx context.Context) *SnaplockFingerprintOperationCollectionGetParams {
	return &SnaplockFingerprintOperationCollectionGetParams{
		Context: ctx,
	}
}

// NewSnaplockFingerprintOperationCollectionGetParamsWithHTTPClient creates a new SnaplockFingerprintOperationCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockFingerprintOperationCollectionGetParamsWithHTTPClient(client *http.Client) *SnaplockFingerprintOperationCollectionGetParams {
	return &SnaplockFingerprintOperationCollectionGetParams{
		HTTPClient: client,
	}
}

/* SnaplockFingerprintOperationCollectionGetParams contains all the parameters to send to the API endpoint
   for the snaplock fingerprint operation collection get operation.

   Typically these are written to a http.Request.
*/
type SnaplockFingerprintOperationCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* SvmUUID.

	   SVM UUID
	*/
	SVMUUIDQueryParameter string

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUIDQueryParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock fingerprint operation collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockFingerprintOperationCollectionGetParams) WithDefaults() *SnaplockFingerprintOperationCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock fingerprint operation collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockFingerprintOperationCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SnaplockFingerprintOperationCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithTimeout(timeout time.Duration) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithContext(ctx context.Context) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithHTTPClient(client *http.Client) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithFieldsQueryParameter(fields []string) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMUUIDQueryParameter adds the svmUUID to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID string) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithVolumeUUIDQueryParameter adds the volumeUUID to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) WithVolumeUUIDQueryParameter(volumeUUID string) *SnaplockFingerprintOperationCollectionGetParams {
	o.SetVolumeUUIDQueryParameter(volumeUUID)
	return o
}

// SetVolumeUUIDQueryParameter adds the volumeUuid to the snaplock fingerprint operation collection get params
func (o *SnaplockFingerprintOperationCollectionGetParams) SetVolumeUUIDQueryParameter(volumeUUID string) {
	o.VolumeUUIDQueryParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockFingerprintOperationCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// query param svm.uuid
	qrSvmUUID := o.SVMUUIDQueryParameter
	qSvmUUID := qrSvmUUID
	if qSvmUUID != "" {

		if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
			return err
		}
	}

	// query param volume.uuid
	qrVolumeUUID := o.VolumeUUIDQueryParameter
	qVolumeUUID := qrVolumeUUID
	if qVolumeUUID != "" {

		if err := r.SetQueryParam("volume.uuid", qVolumeUUID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnaplockFingerprintOperationCollectionGet binds the parameter fields
func (o *SnaplockFingerprintOperationCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamSnaplockFingerprintOperationCollectionGet binds the parameter order_by
func (o *SnaplockFingerprintOperationCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}