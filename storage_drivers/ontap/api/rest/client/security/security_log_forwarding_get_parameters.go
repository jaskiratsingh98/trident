// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityLogForwardingGetParams creates a new SecurityLogForwardingGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityLogForwardingGetParams() *SecurityLogForwardingGetParams {
	return &SecurityLogForwardingGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityLogForwardingGetParamsWithTimeout creates a new SecurityLogForwardingGetParams object
// with the ability to set a timeout on a request.
func NewSecurityLogForwardingGetParamsWithTimeout(timeout time.Duration) *SecurityLogForwardingGetParams {
	return &SecurityLogForwardingGetParams{
		timeout: timeout,
	}
}

// NewSecurityLogForwardingGetParamsWithContext creates a new SecurityLogForwardingGetParams object
// with the ability to set a context for a request.
func NewSecurityLogForwardingGetParamsWithContext(ctx context.Context) *SecurityLogForwardingGetParams {
	return &SecurityLogForwardingGetParams{
		Context: ctx,
	}
}

// NewSecurityLogForwardingGetParamsWithHTTPClient creates a new SecurityLogForwardingGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityLogForwardingGetParamsWithHTTPClient(client *http.Client) *SecurityLogForwardingGetParams {
	return &SecurityLogForwardingGetParams{
		HTTPClient: client,
	}
}

/* SecurityLogForwardingGetParams contains all the parameters to send to the API endpoint
   for the security log forwarding get operation.

   Typically these are written to a http.Request.
*/
type SecurityLogForwardingGetParams struct {

	/* Address.

	   IP address of remote syslog/splunk server.
	*/
	AddressPathParameter string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Port.

	   Port number of remote syslog/splunk server.
	*/
	PortPathParameter int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security log forwarding get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityLogForwardingGetParams) WithDefaults() *SecurityLogForwardingGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security log forwarding get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityLogForwardingGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) WithTimeout(timeout time.Duration) *SecurityLogForwardingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) WithContext(ctx context.Context) *SecurityLogForwardingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) WithHTTPClient(client *http.Client) *SecurityLogForwardingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressPathParameter adds the address to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) WithAddressPathParameter(address string) *SecurityLogForwardingGetParams {
	o.SetAddressPathParameter(address)
	return o
}

// SetAddressPathParameter adds the address to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) SetAddressPathParameter(address string) {
	o.AddressPathParameter = address
}

// WithFieldsQueryParameter adds the fields to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) WithFieldsQueryParameter(fields []string) *SecurityLogForwardingGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithPortPathParameter adds the port to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) WithPortPathParameter(port int64) *SecurityLogForwardingGetParams {
	o.SetPortPathParameter(port)
	return o
}

// SetPortPathParameter adds the port to the security log forwarding get params
func (o *SecurityLogForwardingGetParams) SetPortPathParameter(port int64) {
	o.PortPathParameter = port
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityLogForwardingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.AddressPathParameter); err != nil {
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

	// path param port
	if err := r.SetPathParam("port", swag.FormatInt64(o.PortPathParameter)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityLogForwardingGet binds the parameter fields
func (o *SecurityLogForwardingGetParams) bindParamFields(formats strfmt.Registry) []string {
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