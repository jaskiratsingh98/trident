// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewSnmpUsersDeleteParams creates a new SnmpUsersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnmpUsersDeleteParams() *SnmpUsersDeleteParams {
	return &SnmpUsersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnmpUsersDeleteParamsWithTimeout creates a new SnmpUsersDeleteParams object
// with the ability to set a timeout on a request.
func NewSnmpUsersDeleteParamsWithTimeout(timeout time.Duration) *SnmpUsersDeleteParams {
	return &SnmpUsersDeleteParams{
		timeout: timeout,
	}
}

// NewSnmpUsersDeleteParamsWithContext creates a new SnmpUsersDeleteParams object
// with the ability to set a context for a request.
func NewSnmpUsersDeleteParamsWithContext(ctx context.Context) *SnmpUsersDeleteParams {
	return &SnmpUsersDeleteParams{
		Context: ctx,
	}
}

// NewSnmpUsersDeleteParamsWithHTTPClient creates a new SnmpUsersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnmpUsersDeleteParamsWithHTTPClient(client *http.Client) *SnmpUsersDeleteParams {
	return &SnmpUsersDeleteParams{
		HTTPClient: client,
	}
}

/* SnmpUsersDeleteParams contains all the parameters to send to the API endpoint
   for the snmp users delete operation.

   Typically these are written to a http.Request.
*/
type SnmpUsersDeleteParams struct {

	/* EngineID.

	   Engine ID of owning SVM or remote switch.
	*/
	EngineIDPathParameter string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Name.

	   SNMP user name.
	*/
	NamePathParameter string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snmp users delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnmpUsersDeleteParams) WithDefaults() *SnmpUsersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snmp users delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnmpUsersDeleteParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := SnmpUsersDeleteParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snmp users delete params
func (o *SnmpUsersDeleteParams) WithTimeout(timeout time.Duration) *SnmpUsersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snmp users delete params
func (o *SnmpUsersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snmp users delete params
func (o *SnmpUsersDeleteParams) WithContext(ctx context.Context) *SnmpUsersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snmp users delete params
func (o *SnmpUsersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snmp users delete params
func (o *SnmpUsersDeleteParams) WithHTTPClient(client *http.Client) *SnmpUsersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snmp users delete params
func (o *SnmpUsersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEngineIDPathParameter adds the engineID to the snmp users delete params
func (o *SnmpUsersDeleteParams) WithEngineIDPathParameter(engineID string) *SnmpUsersDeleteParams {
	o.SetEngineIDPathParameter(engineID)
	return o
}

// SetEngineIDPathParameter adds the engineId to the snmp users delete params
func (o *SnmpUsersDeleteParams) SetEngineIDPathParameter(engineID string) {
	o.EngineIDPathParameter = engineID
}

// WithFieldsQueryParameter adds the fields to the snmp users delete params
func (o *SnmpUsersDeleteParams) WithFieldsQueryParameter(fields []string) *SnmpUsersDeleteParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the snmp users delete params
func (o *SnmpUsersDeleteParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithNamePathParameter adds the name to the snmp users delete params
func (o *SnmpUsersDeleteParams) WithNamePathParameter(name string) *SnmpUsersDeleteParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the snmp users delete params
func (o *SnmpUsersDeleteParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the snmp users delete params
func (o *SnmpUsersDeleteParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SnmpUsersDeleteParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the snmp users delete params
func (o *SnmpUsersDeleteParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *SnmpUsersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param engine_id
	if err := r.SetPathParam("engine_id", o.EngineIDPathParameter); err != nil {
		return err
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnmpUsersDelete binds the parameter fields
func (o *SnmpUsersDeleteParams) bindParamFields(formats strfmt.Registry) []string {
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