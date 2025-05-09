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

// NewSecurityExternalRoleMappingCollectionGetParams creates a new SecurityExternalRoleMappingCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityExternalRoleMappingCollectionGetParams() *SecurityExternalRoleMappingCollectionGetParams {
	return &SecurityExternalRoleMappingCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityExternalRoleMappingCollectionGetParamsWithTimeout creates a new SecurityExternalRoleMappingCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSecurityExternalRoleMappingCollectionGetParamsWithTimeout(timeout time.Duration) *SecurityExternalRoleMappingCollectionGetParams {
	return &SecurityExternalRoleMappingCollectionGetParams{
		timeout: timeout,
	}
}

// NewSecurityExternalRoleMappingCollectionGetParamsWithContext creates a new SecurityExternalRoleMappingCollectionGetParams object
// with the ability to set a context for a request.
func NewSecurityExternalRoleMappingCollectionGetParamsWithContext(ctx context.Context) *SecurityExternalRoleMappingCollectionGetParams {
	return &SecurityExternalRoleMappingCollectionGetParams{
		Context: ctx,
	}
}

// NewSecurityExternalRoleMappingCollectionGetParamsWithHTTPClient creates a new SecurityExternalRoleMappingCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityExternalRoleMappingCollectionGetParamsWithHTTPClient(client *http.Client) *SecurityExternalRoleMappingCollectionGetParams {
	return &SecurityExternalRoleMappingCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SecurityExternalRoleMappingCollectionGetParams contains all the parameters to send to the API endpoint

	for the security external role mapping collection get operation.

	Typically these are written to a http.Request.
*/
type SecurityExternalRoleMappingCollectionGetParams struct {

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ExternalRole.

	   Filter by external_role
	*/
	ExternalRole *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OntapRoleName.

	   Filter by ontap_role.name
	*/
	OntapRoleName *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Provider.

	   Filter by provider
	*/
	Provider *string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* Timestamp.

	   Filter by timestamp
	*/
	Timestamp *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security external role mapping collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityExternalRoleMappingCollectionGetParams) WithDefaults() *SecurityExternalRoleMappingCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security external role mapping collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityExternalRoleMappingCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SecurityExternalRoleMappingCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithTimeout(timeout time.Duration) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithContext(ctx context.Context) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithHTTPClient(client *http.Client) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithComment(comment *string) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithExternalRole adds the externalRole to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithExternalRole(externalRole *string) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetExternalRole(externalRole)
	return o
}

// SetExternalRole adds the externalRole to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetExternalRole(externalRole *string) {
	o.ExternalRole = externalRole
}

// WithFields adds the fields to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithFields(fields []string) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithMaxRecords(maxRecords *int64) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOntapRoleName adds the ontapRoleName to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithOntapRoleName(ontapRoleName *string) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetOntapRoleName(ontapRoleName)
	return o
}

// SetOntapRoleName adds the ontapRoleName to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetOntapRoleName(ontapRoleName *string) {
	o.OntapRoleName = ontapRoleName
}

// WithOrderBy adds the orderBy to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithOrderBy(orderBy []string) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithProvider adds the provider to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithProvider(provider *string) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithReturnRecords adds the returnRecords to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithReturnRecords(returnRecords *bool) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithTimestamp adds the timestamp to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) WithTimestamp(timestamp *string) *SecurityExternalRoleMappingCollectionGetParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the security external role mapping collection get params
func (o *SecurityExternalRoleMappingCollectionGetParams) SetTimestamp(timestamp *string) {
	o.Timestamp = timestamp
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityExternalRoleMappingCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.ExternalRole != nil {

		// query param external_role
		var qrExternalRole string

		if o.ExternalRole != nil {
			qrExternalRole = *o.ExternalRole
		}
		qExternalRole := qrExternalRole
		if qExternalRole != "" {

			if err := r.SetQueryParam("external_role", qExternalRole); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OntapRoleName != nil {

		// query param ontap_role.name
		var qrOntapRoleName string

		if o.OntapRoleName != nil {
			qrOntapRoleName = *o.OntapRoleName
		}
		qOntapRoleName := qrOntapRoleName
		if qOntapRoleName != "" {

			if err := r.SetQueryParam("ontap_role.name", qOntapRoleName); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.Provider != nil {

		// query param provider
		var qrProvider string

		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {

			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.Timestamp != nil {

		// query param timestamp
		var qrTimestamp string

		if o.Timestamp != nil {
			qrTimestamp = *o.Timestamp
		}
		qTimestamp := qrTimestamp
		if qTimestamp != "" {

			if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityExternalRoleMappingCollectionGet binds the parameter fields
func (o *SecurityExternalRoleMappingCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamSecurityExternalRoleMappingCollectionGet binds the parameter order_by
func (o *SecurityExternalRoleMappingCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
