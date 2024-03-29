// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetKeysResult get keys result
//
// swagger:model GetKeysResult
type GetKeysResult struct {

	// Custom attributes (optional)
	Attributes interface{} `json:"attributes,omitempty"`

	// Database record id
	ID string `json:"id,omitempty"`

	// Masked API key (asterisks in the middle)
	Key string `json:"key,omitempty"`

	// Organization ID (optional)
	OrganizationID *string `json:"organizationId,omitempty"`

	// Project ID (optional)
	ProjectID *string `json:"projectId,omitempty"`

	// User ID (optional)
	UserID *string `json:"userId,omitempty"`
}

// Validate validates this get keys result
func (m *GetKeysResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get keys result based on context it is used
func (m *GetKeysResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetKeysResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKeysResult) UnmarshalBinary(b []byte) error {
	var res GetKeysResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
