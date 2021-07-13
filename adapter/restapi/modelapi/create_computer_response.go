// Code generated by go-swagger; DO NOT EDIT.

package modelapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateComputerResponse Created computer and computer api key.
//
// swagger:model CreateComputerResponse
type CreateComputerResponse struct {

	// comp key
	// Example: xxxx.yyyy.zzzz
	CompKey string `json:"comp_key,omitempty"`

	// computer
	Computer *Computer `json:"computer,omitempty"`
}

// Validate validates this create computer response
func (m *CreateComputerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateComputerResponse) validateComputer(formats strfmt.Registry) error {
	if swag.IsZero(m.Computer) { // not required
		return nil
	}

	if m.Computer != nil {
		if err := m.Computer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create computer response based on the context it is used
func (m *CreateComputerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateComputerResponse) contextValidateComputer(ctx context.Context, formats strfmt.Registry) error {

	if m.Computer != nil {
		if err := m.Computer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateComputerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateComputerResponse) UnmarshalBinary(b []byte) error {
	var res CreateComputerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}