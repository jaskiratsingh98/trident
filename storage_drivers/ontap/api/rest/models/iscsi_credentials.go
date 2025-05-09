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

// IscsiCredentials iscsi credentials
//
// swagger:model iscsi_credentials
type IscsiCredentials struct {

	// links
	Links *IscsiCredentialsInlineLinks `json:"_links,omitempty"`

	// The iSCSI authentication type. Required in POST; optional in PATCH.
	//
	// Enum: ["chap","none","deny"]
	AuthenticationType *string `json:"authentication_type,omitempty"`

	// chap
	Chap *IscsiCredentialsInlineChap `json:"chap,omitempty"`

	// The iSCSI initiator to which the credentials apply. Required in POST.
	//
	// Example: iqn.1998-01.com.corp.iscsi:name1
	Initiator *string `json:"initiator,omitempty"`

	// initiator address
	InitiatorAddress *IscsiCredentialsInlineInitiatorAddress `json:"initiator_address,omitempty"`

	// svm
	Svm *IscsiCredentialsInlineSvm `json:"svm,omitempty"`
}

// Validate validates this iscsi credentials
func (m *IscsiCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentials) validateLinks(formats strfmt.Registry) error {
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

var iscsiCredentialsTypeAuthenticationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chap","none","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iscsiCredentialsTypeAuthenticationTypePropEnum = append(iscsiCredentialsTypeAuthenticationTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// iscsi_credentials
	// IscsiCredentials
	// authentication_type
	// AuthenticationType
	// chap
	// END DEBUGGING
	// IscsiCredentialsAuthenticationTypeChap captures enum value "chap"
	IscsiCredentialsAuthenticationTypeChap string = "chap"

	// BEGIN DEBUGGING
	// iscsi_credentials
	// IscsiCredentials
	// authentication_type
	// AuthenticationType
	// none
	// END DEBUGGING
	// IscsiCredentialsAuthenticationTypeNone captures enum value "none"
	IscsiCredentialsAuthenticationTypeNone string = "none"

	// BEGIN DEBUGGING
	// iscsi_credentials
	// IscsiCredentials
	// authentication_type
	// AuthenticationType
	// deny
	// END DEBUGGING
	// IscsiCredentialsAuthenticationTypeDeny captures enum value "deny"
	IscsiCredentialsAuthenticationTypeDeny string = "deny"
)

// prop value enum
func (m *IscsiCredentials) validateAuthenticationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, iscsiCredentialsTypeAuthenticationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IscsiCredentials) validateAuthenticationType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationTypeEnum("authentication_type", "body", *m.AuthenticationType); err != nil {
		return err
	}

	return nil
}

func (m *IscsiCredentials) validateChap(formats strfmt.Registry) error {
	if swag.IsZero(m.Chap) { // not required
		return nil
	}

	if m.Chap != nil {
		if err := m.Chap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chap")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiCredentials) validateInitiatorAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.InitiatorAddress) { // not required
		return nil
	}

	if m.InitiatorAddress != nil {
		if err := m.InitiatorAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator_address")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiCredentials) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi credentials based on the context it is used
func (m *IscsiCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiatorAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentials) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiCredentials) contextValidateChap(ctx context.Context, formats strfmt.Registry) error {

	if m.Chap != nil {
		if err := m.Chap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chap")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiCredentials) contextValidateInitiatorAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.InitiatorAddress != nil {
		if err := m.InitiatorAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator_address")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiCredentials) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiCredentials) UnmarshalBinary(b []byte) error {
	var res IscsiCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiCredentialsInlineChap Challenge-Handshake Authentication Protocol (CHAP) credentials.
//
// swagger:model iscsi_credentials_inline_chap
type IscsiCredentialsInlineChap struct {

	// inbound
	Inbound *IscsiCredentialsInlineChapInlineInbound `json:"inbound,omitempty"`

	// outbound
	Outbound *IscsiCredentialsInlineChapInlineOutbound `json:"outbound,omitempty"`
}

// Validate validates this iscsi credentials inline chap
func (m *IscsiCredentialsInlineChap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInbound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutbound(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineChap) validateInbound(formats strfmt.Registry) error {
	if swag.IsZero(m.Inbound) { // not required
		return nil
	}

	if m.Inbound != nil {
		if err := m.Inbound.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chap" + "." + "inbound")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiCredentialsInlineChap) validateOutbound(formats strfmt.Registry) error {
	if swag.IsZero(m.Outbound) { // not required
		return nil
	}

	if m.Outbound != nil {
		if err := m.Outbound.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chap" + "." + "outbound")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi credentials inline chap based on the context it is used
func (m *IscsiCredentialsInlineChap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInbound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutbound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineChap) contextValidateInbound(ctx context.Context, formats strfmt.Registry) error {

	if m.Inbound != nil {
		if err := m.Inbound.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chap" + "." + "inbound")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiCredentialsInlineChap) contextValidateOutbound(ctx context.Context, formats strfmt.Registry) error {

	if m.Outbound != nil {
		if err := m.Outbound.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chap" + "." + "outbound")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiCredentialsInlineChap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiCredentialsInlineChap) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineChap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiCredentialsInlineChapInlineInbound Inbound CHAP credentials.
//
// swagger:model iscsi_credentials_inline_chap_inline_inbound
type IscsiCredentialsInlineChapInlineInbound struct {

	// The inbound CHAP password. Write-only; optional in POST and PATCH.
	//
	// Max Length: 512
	// Min Length: 1
	Password *string `json:"password,omitempty"`

	// The inbound CHAP user name. Optional in POST and PATCH.
	//
	// Max Length: 128
	// Min Length: 1
	User *string `json:"user,omitempty"`
}

// Validate validates this iscsi credentials inline chap inline inbound
func (m *IscsiCredentialsInlineChapInlineInbound) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineChapInlineInbound) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("chap"+"."+"inbound"+"."+"password", "body", *m.Password, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("chap"+"."+"inbound"+"."+"password", "body", *m.Password, 512); err != nil {
		return err
	}

	return nil
}

func (m *IscsiCredentialsInlineChapInlineInbound) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if err := validate.MinLength("chap"+"."+"inbound"+"."+"user", "body", *m.User, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("chap"+"."+"inbound"+"."+"user", "body", *m.User, 128); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi credentials inline chap inline inbound based on context it is used
func (m *IscsiCredentialsInlineChapInlineInbound) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IscsiCredentialsInlineChapInlineInbound) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiCredentialsInlineChapInlineInbound) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineChapInlineInbound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiCredentialsInlineChapInlineOutbound Output CHAP credentials.</br>
// To clear previously set outbound CHAP credentials, set property `chap.outbound.user` to an empty string in PATCH.
//
// swagger:model iscsi_credentials_inline_chap_inline_outbound
type IscsiCredentialsInlineChapInlineOutbound struct {

	// The outbound CHAP password. Write-only; optional in POST and PATCH.
	//
	// Max Length: 512
	// Min Length: 1
	Password *string `json:"password,omitempty"`

	// The outbound CHAP user name. Optional in POST and PATCH.</br>
	// To clear previously set outbound CHAP credentials, set this property to an empty string in PATCH.
	//
	// Max Length: 128
	// Min Length: 1
	User *string `json:"user,omitempty"`
}

// Validate validates this iscsi credentials inline chap inline outbound
func (m *IscsiCredentialsInlineChapInlineOutbound) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineChapInlineOutbound) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("chap"+"."+"outbound"+"."+"password", "body", *m.Password, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("chap"+"."+"outbound"+"."+"password", "body", *m.Password, 512); err != nil {
		return err
	}

	return nil
}

func (m *IscsiCredentialsInlineChapInlineOutbound) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if err := validate.MinLength("chap"+"."+"outbound"+"."+"user", "body", *m.User, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("chap"+"."+"outbound"+"."+"user", "body", *m.User, 128); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi credentials inline chap inline outbound based on context it is used
func (m *IscsiCredentialsInlineChapInlineOutbound) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IscsiCredentialsInlineChapInlineOutbound) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiCredentialsInlineChapInlineOutbound) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineChapInlineOutbound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiCredentialsInlineInitiatorAddress Initiator address ranges.
//
// swagger:model iscsi_credentials_inline_initiator_address
type IscsiCredentialsInlineInitiatorAddress struct {

	// masks
	Masks []*IPInfo `json:"masks,omitempty"`

	// ranges
	Ranges []*IPAddressRange `json:"ranges,omitempty"`
}

// Validate validates this iscsi credentials inline initiator address
func (m *IscsiCredentialsInlineInitiatorAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMasks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineInitiatorAddress) validateMasks(formats strfmt.Registry) error {
	if swag.IsZero(m.Masks) { // not required
		return nil
	}

	for i := 0; i < len(m.Masks); i++ {
		if swag.IsZero(m.Masks[i]) { // not required
			continue
		}

		if m.Masks[i] != nil {
			if err := m.Masks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_address" + "." + "masks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiCredentialsInlineInitiatorAddress) validateRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.Ranges) { // not required
		return nil
	}

	for i := 0; i < len(m.Ranges); i++ {
		if swag.IsZero(m.Ranges[i]) { // not required
			continue
		}

		if m.Ranges[i] != nil {
			if err := m.Ranges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_address" + "." + "ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this iscsi credentials inline initiator address based on the context it is used
func (m *IscsiCredentialsInlineInitiatorAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineInitiatorAddress) contextValidateMasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Masks); i++ {

		if m.Masks[i] != nil {
			if err := m.Masks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_address" + "." + "masks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiCredentialsInlineInitiatorAddress) contextValidateRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ranges); i++ {

		if m.Ranges[i] != nil {
			if err := m.Ranges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_address" + "." + "ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiCredentialsInlineInitiatorAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiCredentialsInlineInitiatorAddress) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineInitiatorAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiCredentialsInlineLinks iscsi credentials inline links
//
// swagger:model iscsi_credentials_inline__links
type IscsiCredentialsInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this iscsi credentials inline links
func (m *IscsiCredentialsInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi credentials inline links based on the context it is used
func (m *IscsiCredentialsInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IscsiCredentialsInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiCredentialsInlineLinks) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiCredentialsInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model iscsi_credentials_inline_svm
type IscsiCredentialsInlineSvm struct {

	// links
	Links *IscsiCredentialsInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this iscsi credentials inline svm
func (m *IscsiCredentialsInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi credentials inline svm based on the context it is used
func (m *IscsiCredentialsInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiCredentialsInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiCredentialsInlineSvm) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiCredentialsInlineSvmInlineLinks iscsi credentials inline svm inline links
//
// swagger:model iscsi_credentials_inline_svm_inline__links
type IscsiCredentialsInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this iscsi credentials inline svm inline links
func (m *IscsiCredentialsInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi credentials inline svm inline links based on the context it is used
func (m *IscsiCredentialsInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiCredentialsInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiCredentialsInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiCredentialsInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
