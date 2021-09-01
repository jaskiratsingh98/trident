// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PerformanceMetricResponse performance metric response
//
// swagger:model performance_metric_response
type PerformanceMetricResponse struct {

	// links
	Links *PerformanceMetricResponseLinks `json:"_links,omitempty"`

	// Number of records
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*PerformanceMetricResponseRecordsItems0 `json:"records,omitempty"`
}

// Validate validates this performance metric response
func (m *PerformanceMetricResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponse) validateRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this performance metric response based on the context it is used
func (m *PerformanceMetricResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponse) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {
			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricResponse) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricResponseLinks performance metric response links
//
// swagger:model PerformanceMetricResponseLinks
type PerformanceMetricResponseLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance metric response links
func (m *PerformanceMetricResponseLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricResponseLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponseLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance metric response links based on the context it is used
func (m *PerformanceMetricResponseLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricResponseLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponseLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricResponseLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricResponseLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricResponseLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricResponseRecordsItems0 Performance numbers, such as IOPS latency and throughput.
//
// swagger:model PerformanceMetricResponseRecordsItems0
type PerformanceMetricResponseRecordsItems0 struct {

	// links
	Links *PerformanceMetricResponseRecordsItems0Links `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceMetricResponseRecordsItems0Iops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceMetricResponseRecordsItems0Latency `json:"latency,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: [ok error partial_no_data partial_no_response partial_other_error negative_delta not_found backfilled_data inconsistent_delta_time inconsistent_old_data partial_no_uuid]
	Status string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceMetricResponseRecordsItems0Throughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25T11:20:13Z
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance metric response records items0
func (m *PerformanceMetricResponseRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

var performanceMetricResponseRecordsItems0TypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceMetricResponseRecordsItems0TypeDurationPropEnum = append(performanceMetricResponseRecordsItems0TypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0DurationPT15S captures enum value "PT15S"
	PerformanceMetricResponseRecordsItems0DurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0DurationPT4M captures enum value "PT4M"
	PerformanceMetricResponseRecordsItems0DurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0DurationPT30M captures enum value "PT30M"
	PerformanceMetricResponseRecordsItems0DurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0DurationPT2H captures enum value "PT2H"
	PerformanceMetricResponseRecordsItems0DurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0DurationP1D captures enum value "P1D"
	PerformanceMetricResponseRecordsItems0DurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0DurationPT5M captures enum value "PT5M"
	PerformanceMetricResponseRecordsItems0DurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceMetricResponseRecordsItems0) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceMetricResponseRecordsItems0TypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) validateLatency(formats strfmt.Registry) error {
	if swag.IsZero(m.Latency) { // not required
		return nil
	}

	if m.Latency != nil {
		if err := m.Latency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

var performanceMetricResponseRecordsItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceMetricResponseRecordsItems0TypeStatusPropEnum = append(performanceMetricResponseRecordsItems0TypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// ok
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusOk captures enum value "ok"
	PerformanceMetricResponseRecordsItems0StatusOk string = "ok"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// error
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusError captures enum value "error"
	PerformanceMetricResponseRecordsItems0StatusError string = "error"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusPartialNoData captures enum value "partial_no_data"
	PerformanceMetricResponseRecordsItems0StatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceMetricResponseRecordsItems0StatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusPartialOtherError captures enum value "partial_other_error"
	PerformanceMetricResponseRecordsItems0StatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusNegativeDelta captures enum value "negative_delta"
	PerformanceMetricResponseRecordsItems0StatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// not_found
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusNotFound captures enum value "not_found"
	PerformanceMetricResponseRecordsItems0StatusNotFound string = "not_found"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusBackfilledData captures enum value "backfilled_data"
	PerformanceMetricResponseRecordsItems0StatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceMetricResponseRecordsItems0StatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceMetricResponseRecordsItems0StatusInconsistentOldData string = "inconsistent_old_data"

	// BEGIN DEBUGGING
	// PerformanceMetricResponseRecordsItems0
	// PerformanceMetricResponseRecordsItems0
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PerformanceMetricResponseRecordsItems0StatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceMetricResponseRecordsItems0StatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceMetricResponseRecordsItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceMetricResponseRecordsItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance metric response records items0 based on the context it is used
func (m *PerformanceMetricResponseRecordsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", string(m.Duration)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {
		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

	if m.Latency != nil {
		if err := m.Latency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricResponseRecordsItems0) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricResponseRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricResponseRecordsItems0Iops The rate of I/O operations observed at the storage object.
//
// swagger:model PerformanceMetricResponseRecordsItems0Iops
type PerformanceMetricResponseRecordsItems0Iops struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write int64 `json:"write,omitempty"`
}

// Validate validates this performance metric response records items0 iops
func (m *PerformanceMetricResponseRecordsItems0Iops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric response records items0 iops based on the context it is used
func (m *PerformanceMetricResponseRecordsItems0Iops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0Iops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0Iops) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricResponseRecordsItems0Iops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricResponseRecordsItems0Latency The round trip latency in microseconds observed at the storage object.
//
// swagger:model PerformanceMetricResponseRecordsItems0Latency
type PerformanceMetricResponseRecordsItems0Latency struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write int64 `json:"write,omitempty"`
}

// Validate validates this performance metric response records items0 latency
func (m *PerformanceMetricResponseRecordsItems0Latency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric response records items0 latency based on the context it is used
func (m *PerformanceMetricResponseRecordsItems0Latency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0Latency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0Latency) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricResponseRecordsItems0Latency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricResponseRecordsItems0Links performance metric response records items0 links
//
// swagger:model PerformanceMetricResponseRecordsItems0Links
type PerformanceMetricResponseRecordsItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance metric response records items0 links
func (m *PerformanceMetricResponseRecordsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricResponseRecordsItems0Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance metric response records items0 links based on the context it is used
func (m *PerformanceMetricResponseRecordsItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricResponseRecordsItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0Links) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricResponseRecordsItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricResponseRecordsItems0Throughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model PerformanceMetricResponseRecordsItems0Throughput
type PerformanceMetricResponseRecordsItems0Throughput struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write int64 `json:"write,omitempty"`
}

// Validate validates this performance metric response records items0 throughput
func (m *PerformanceMetricResponseRecordsItems0Throughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric response records items0 throughput based on the context it is used
func (m *PerformanceMetricResponseRecordsItems0Throughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0Throughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricResponseRecordsItems0Throughput) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricResponseRecordsItems0Throughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}