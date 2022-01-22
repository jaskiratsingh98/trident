// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsistencyGroupNamespaceStatus Status information about the NVMe namespace.
//
//
// swagger:model consistency_group_namespace_status
type ConsistencyGroupNamespaceStatus struct {

	// The state of the volume and aggregate that contain the NVMe namespace. Namespaces are only available when their containers are available.
	//
	// Read Only: true
	// Enum: [online aggregate_offline volume_offline]
	ContainerState string `json:"container_state,omitempty"`

	// Reports if the NVMe namespace is mapped to an NVMe subsystem.<br/>
	// There is an added cost to retrieving this property's value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter. See [`Requesting specific fields`](#Requesting_specific_fields) to learn more.
	//
	// Read Only: true
	Mapped *bool `json:"mapped,omitempty"`

	// Reports if the NVMe namespace allows only read access.
	//
	// Read Only: true
	ReadOnly *bool `json:"read_only,omitempty"`

	// The state of the NVMe namespace. Normal states for a namespace are _online_ and _offline_. Other states indicate errors.
	//
	// Example: online
	// Read Only: true
	// Enum: [nvfail offline online space_error]
	State string `json:"state,omitempty"`
}

// Validate validates this consistency group namespace status
func (m *ConsistencyGroupNamespaceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consistencyGroupNamespaceStatusTypeContainerStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","aggregate_offline","volume_offline"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupNamespaceStatusTypeContainerStatePropEnum = append(consistencyGroupNamespaceStatusTypeContainerStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// consistency_group_namespace_status
	// ConsistencyGroupNamespaceStatus
	// container_state
	// ContainerState
	// online
	// END DEBUGGING
	// ConsistencyGroupNamespaceStatusContainerStateOnline captures enum value "online"
	ConsistencyGroupNamespaceStatusContainerStateOnline string = "online"

	// BEGIN DEBUGGING
	// consistency_group_namespace_status
	// ConsistencyGroupNamespaceStatus
	// container_state
	// ContainerState
	// aggregate_offline
	// END DEBUGGING
	// ConsistencyGroupNamespaceStatusContainerStateAggregateOffline captures enum value "aggregate_offline"
	ConsistencyGroupNamespaceStatusContainerStateAggregateOffline string = "aggregate_offline"

	// BEGIN DEBUGGING
	// consistency_group_namespace_status
	// ConsistencyGroupNamespaceStatus
	// container_state
	// ContainerState
	// volume_offline
	// END DEBUGGING
	// ConsistencyGroupNamespaceStatusContainerStateVolumeOffline captures enum value "volume_offline"
	ConsistencyGroupNamespaceStatusContainerStateVolumeOffline string = "volume_offline"
)

// prop value enum
func (m *ConsistencyGroupNamespaceStatus) validateContainerStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupNamespaceStatusTypeContainerStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupNamespaceStatus) validateContainerState(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerState) { // not required
		return nil
	}

	// value enum
	if err := m.validateContainerStateEnum("container_state", "body", m.ContainerState); err != nil {
		return err
	}

	return nil
}

var consistencyGroupNamespaceStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nvfail","offline","online","space_error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupNamespaceStatusTypeStatePropEnum = append(consistencyGroupNamespaceStatusTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// consistency_group_namespace_status
	// ConsistencyGroupNamespaceStatus
	// state
	// State
	// nvfail
	// END DEBUGGING
	// ConsistencyGroupNamespaceStatusStateNvfail captures enum value "nvfail"
	ConsistencyGroupNamespaceStatusStateNvfail string = "nvfail"

	// BEGIN DEBUGGING
	// consistency_group_namespace_status
	// ConsistencyGroupNamespaceStatus
	// state
	// State
	// offline
	// END DEBUGGING
	// ConsistencyGroupNamespaceStatusStateOffline captures enum value "offline"
	ConsistencyGroupNamespaceStatusStateOffline string = "offline"

	// BEGIN DEBUGGING
	// consistency_group_namespace_status
	// ConsistencyGroupNamespaceStatus
	// state
	// State
	// online
	// END DEBUGGING
	// ConsistencyGroupNamespaceStatusStateOnline captures enum value "online"
	ConsistencyGroupNamespaceStatusStateOnline string = "online"

	// BEGIN DEBUGGING
	// consistency_group_namespace_status
	// ConsistencyGroupNamespaceStatus
	// state
	// State
	// space_error
	// END DEBUGGING
	// ConsistencyGroupNamespaceStatusStateSpaceError captures enum value "space_error"
	ConsistencyGroupNamespaceStatusStateSpaceError string = "space_error"
)

// prop value enum
func (m *ConsistencyGroupNamespaceStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupNamespaceStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupNamespaceStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this consistency group namespace status based on the context it is used
func (m *ConsistencyGroupNamespaceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMapped(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReadOnly(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupNamespaceStatus) contextValidateContainerState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "container_state", "body", string(m.ContainerState)); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNamespaceStatus) contextValidateMapped(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mapped", "body", m.Mapped); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNamespaceStatus) contextValidateReadOnly(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "read_only", "body", m.ReadOnly); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNamespaceStatus) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceStatus) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupNamespaceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}