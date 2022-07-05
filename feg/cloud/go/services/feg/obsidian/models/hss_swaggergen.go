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

// Hss HSS configuration
//
// swagger:model hss
type Hss struct {

	// default sub profile
	DefaultSubProfile *SubscriptionProfile `json:"default_sub_profile,omitempty"`

	// lte auth amf
	// Example: gAA=
	// Format: byte
	LteAuthAmf strfmt.Base64 `json:"lte_auth_amf,omitempty"`

	// lte auth op
	// Example: EREREREREREREREREREREQ==
	// Format: byte
	LteAuthOp strfmt.Base64 `json:"lte_auth_op,omitempty"`

	// server
	Server *DiameterServerConfigs `json:"server,omitempty"`

	// stream subscribers
	// Example: false
	StreamSubscribers bool `json:"stream_subscribers,omitempty"`

	// sub profiles
	SubProfiles map[string]SubscriptionProfile `json:"sub_profiles,omitempty"`
}

// Validate validates this hss
func (m *Hss) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultSubProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hss) validateDefaultSubProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultSubProfile) { // not required
		return nil
	}

	if m.DefaultSubProfile != nil {
		if err := m.DefaultSubProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_sub_profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_sub_profile")
			}
			return err
		}
	}

	return nil
}

func (m *Hss) validateServer(formats strfmt.Registry) error {
	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *Hss) validateSubProfiles(formats strfmt.Registry) error {
	if swag.IsZero(m.SubProfiles) { // not required
		return nil
	}

	for k := range m.SubProfiles {

		if err := validate.Required("sub_profiles"+"."+k, "body", m.SubProfiles[k]); err != nil {
			return err
		}
		if val, ok := m.SubProfiles[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sub_profiles" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sub_profiles" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hss based on the context it is used
func (m *Hss) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultSubProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubProfiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hss) contextValidateDefaultSubProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultSubProfile != nil {
		if err := m.DefaultSubProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_sub_profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_sub_profile")
			}
			return err
		}
	}

	return nil
}

func (m *Hss) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {
		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *Hss) contextValidateSubProfiles(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.SubProfiles {

		if val, ok := m.SubProfiles[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Hss) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hss) UnmarshalBinary(b []byte) error {
	var res Hss
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}