// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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
)

// NewFcpServiceDeleteParams creates a new FcpServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcpServiceDeleteParams() *FcpServiceDeleteParams {
	return &FcpServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcpServiceDeleteParamsWithTimeout creates a new FcpServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewFcpServiceDeleteParamsWithTimeout(timeout time.Duration) *FcpServiceDeleteParams {
	return &FcpServiceDeleteParams{
		timeout: timeout,
	}
}

// NewFcpServiceDeleteParamsWithContext creates a new FcpServiceDeleteParams object
// with the ability to set a context for a request.
func NewFcpServiceDeleteParamsWithContext(ctx context.Context) *FcpServiceDeleteParams {
	return &FcpServiceDeleteParams{
		Context: ctx,
	}
}

// NewFcpServiceDeleteParamsWithHTTPClient creates a new FcpServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcpServiceDeleteParamsWithHTTPClient(client *http.Client) *FcpServiceDeleteParams {
	return &FcpServiceDeleteParams{
		HTTPClient: client,
	}
}

/* FcpServiceDeleteParams contains all the parameters to send to the API endpoint
   for the fcp service delete operation.

   Typically these are written to a http.Request.
*/
type FcpServiceDeleteParams struct {

	/* SvmUUID.

	   The unique identifier of the SVM for which to delete the FC Protocol service.

	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fcp service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcpServiceDeleteParams) WithDefaults() *FcpServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fcp service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcpServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fcp service delete params
func (o *FcpServiceDeleteParams) WithTimeout(timeout time.Duration) *FcpServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fcp service delete params
func (o *FcpServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fcp service delete params
func (o *FcpServiceDeleteParams) WithContext(ctx context.Context) *FcpServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fcp service delete params
func (o *FcpServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fcp service delete params
func (o *FcpServiceDeleteParams) WithHTTPClient(client *http.Client) *FcpServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fcp service delete params
func (o *FcpServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSVMUUIDPathParameter adds the svmUUID to the fcp service delete params
func (o *FcpServiceDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *FcpServiceDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the fcp service delete params
func (o *FcpServiceDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FcpServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}