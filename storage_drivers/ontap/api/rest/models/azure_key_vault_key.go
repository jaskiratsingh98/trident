// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AzureKeyVaultKey azure key vault key
//
// swagger:model azure_key_vault_key
type AzureKeyVaultKey struct {

	// Key identifier of the AKV key encryption key.
	// Example: https://keyvault1.vault.azure.net/keys/key1/12345678901234567890123456789012
	// Format: uri
	KeyID *strfmt.URI `json:"key_id,omitempty"`
}

// Validate validates this azure key vault key
func (m *AzureKeyVaultKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureKeyVaultKey) validateKeyID(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyID) { // not required
		return nil
	}

	if err := validate.FormatOf("key_id", "body", "uri", m.KeyID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this azure key vault key based on context it is used
func (m *AzureKeyVaultKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureKeyVaultKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureKeyVaultKey) UnmarshalBinary(b []byte) error {
	var res AzureKeyVaultKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
