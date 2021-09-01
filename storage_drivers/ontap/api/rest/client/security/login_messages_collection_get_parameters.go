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

// NewLoginMessagesCollectionGetParams creates a new LoginMessagesCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLoginMessagesCollectionGetParams() *LoginMessagesCollectionGetParams {
	return &LoginMessagesCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLoginMessagesCollectionGetParamsWithTimeout creates a new LoginMessagesCollectionGetParams object
// with the ability to set a timeout on a request.
func NewLoginMessagesCollectionGetParamsWithTimeout(timeout time.Duration) *LoginMessagesCollectionGetParams {
	return &LoginMessagesCollectionGetParams{
		timeout: timeout,
	}
}

// NewLoginMessagesCollectionGetParamsWithContext creates a new LoginMessagesCollectionGetParams object
// with the ability to set a context for a request.
func NewLoginMessagesCollectionGetParamsWithContext(ctx context.Context) *LoginMessagesCollectionGetParams {
	return &LoginMessagesCollectionGetParams{
		Context: ctx,
	}
}

// NewLoginMessagesCollectionGetParamsWithHTTPClient creates a new LoginMessagesCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLoginMessagesCollectionGetParamsWithHTTPClient(client *http.Client) *LoginMessagesCollectionGetParams {
	return &LoginMessagesCollectionGetParams{
		HTTPClient: client,
	}
}

/* LoginMessagesCollectionGetParams contains all the parameters to send to the API endpoint
   for the login messages collection get operation.

   Typically these are written to a http.Request.
*/
type LoginMessagesCollectionGetParams struct {

	/* Banner.

	   Filter by banner
	*/
	BannerQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Message.

	   Filter by message
	*/
	MessageQueryParameter *string

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

	/* Scope.

	   Filter by scope
	*/
	ScopeQueryParameter *string

	/* ShowClusterMessage.

	   Filter by show_cluster_message
	*/
	ShowClusterMessageQueryParameter *bool

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the login messages collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginMessagesCollectionGetParams) WithDefaults() *LoginMessagesCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the login messages collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginMessagesCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := LoginMessagesCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithTimeout(timeout time.Duration) *LoginMessagesCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithContext(ctx context.Context) *LoginMessagesCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithHTTPClient(client *http.Client) *LoginMessagesCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBannerQueryParameter adds the banner to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithBannerQueryParameter(banner *string) *LoginMessagesCollectionGetParams {
	o.SetBannerQueryParameter(banner)
	return o
}

// SetBannerQueryParameter adds the banner to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetBannerQueryParameter(banner *string) {
	o.BannerQueryParameter = banner
}

// WithFieldsQueryParameter adds the fields to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithFieldsQueryParameter(fields []string) *LoginMessagesCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *LoginMessagesCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithMessageQueryParameter adds the message to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithMessageQueryParameter(message *string) *LoginMessagesCollectionGetParams {
	o.SetMessageQueryParameter(message)
	return o
}

// SetMessageQueryParameter adds the message to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetMessageQueryParameter(message *string) {
	o.MessageQueryParameter = message
}

// WithOrderByQueryParameter adds the orderBy to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *LoginMessagesCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *LoginMessagesCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *LoginMessagesCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithScopeQueryParameter adds the scope to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithScopeQueryParameter(scope *string) *LoginMessagesCollectionGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithShowClusterMessageQueryParameter adds the showClusterMessage to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithShowClusterMessageQueryParameter(showClusterMessage *bool) *LoginMessagesCollectionGetParams {
	o.SetShowClusterMessageQueryParameter(showClusterMessage)
	return o
}

// SetShowClusterMessageQueryParameter adds the showClusterMessage to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetShowClusterMessageQueryParameter(showClusterMessage *bool) {
	o.ShowClusterMessageQueryParameter = showClusterMessage
}

// WithSVMNameQueryParameter adds the svmName to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *LoginMessagesCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *LoginMessagesCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUUIDQueryParameter adds the uuid to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) WithUUIDQueryParameter(uuid *string) *LoginMessagesCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the login messages collection get params
func (o *LoginMessagesCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *LoginMessagesCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BannerQueryParameter != nil {

		// query param banner
		var qrBanner string

		if o.BannerQueryParameter != nil {
			qrBanner = *o.BannerQueryParameter
		}
		qBanner := qrBanner
		if qBanner != "" {

			if err := r.SetQueryParam("banner", qBanner); err != nil {
				return err
			}
		}
	}

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

	if o.MessageQueryParameter != nil {

		// query param message
		var qrMessage string

		if o.MessageQueryParameter != nil {
			qrMessage = *o.MessageQueryParameter
		}
		qMessage := qrMessage
		if qMessage != "" {

			if err := r.SetQueryParam("message", qMessage); err != nil {
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

	if o.ScopeQueryParameter != nil {

		// query param scope
		var qrScope string

		if o.ScopeQueryParameter != nil {
			qrScope = *o.ScopeQueryParameter
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if o.ShowClusterMessageQueryParameter != nil {

		// query param show_cluster_message
		var qrShowClusterMessage bool

		if o.ShowClusterMessageQueryParameter != nil {
			qrShowClusterMessage = *o.ShowClusterMessageQueryParameter
		}
		qShowClusterMessage := swag.FormatBool(qrShowClusterMessage)
		if qShowClusterMessage != "" {

			if err := r.SetQueryParam("show_cluster_message", qShowClusterMessage); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLoginMessagesCollectionGet binds the parameter fields
func (o *LoginMessagesCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamLoginMessagesCollectionGet binds the parameter order_by
func (o *LoginMessagesCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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