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
)

// NewClusterNisDeleteParams creates a new ClusterNisDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterNisDeleteParams() *ClusterNisDeleteParams {
	return &ClusterNisDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterNisDeleteParamsWithTimeout creates a new ClusterNisDeleteParams object
// with the ability to set a timeout on a request.
func NewClusterNisDeleteParamsWithTimeout(timeout time.Duration) *ClusterNisDeleteParams {
	return &ClusterNisDeleteParams{
		timeout: timeout,
	}
}

// NewClusterNisDeleteParamsWithContext creates a new ClusterNisDeleteParams object
// with the ability to set a context for a request.
func NewClusterNisDeleteParamsWithContext(ctx context.Context) *ClusterNisDeleteParams {
	return &ClusterNisDeleteParams{
		Context: ctx,
	}
}

// NewClusterNisDeleteParamsWithHTTPClient creates a new ClusterNisDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterNisDeleteParamsWithHTTPClient(client *http.Client) *ClusterNisDeleteParams {
	return &ClusterNisDeleteParams{
		HTTPClient: client,
	}
}

/* ClusterNisDeleteParams contains all the parameters to send to the API endpoint
   for the cluster nis delete operation.

   Typically these are written to a http.Request.
*/
type ClusterNisDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster nis delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNisDeleteParams) WithDefaults() *ClusterNisDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster nis delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNisDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster nis delete params
func (o *ClusterNisDeleteParams) WithTimeout(timeout time.Duration) *ClusterNisDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster nis delete params
func (o *ClusterNisDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster nis delete params
func (o *ClusterNisDeleteParams) WithContext(ctx context.Context) *ClusterNisDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster nis delete params
func (o *ClusterNisDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster nis delete params
func (o *ClusterNisDeleteParams) WithHTTPClient(client *http.Client) *ClusterNisDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster nis delete params
func (o *ClusterNisDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterNisDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}