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

// Cbsd cbsd
//
// swagger:model cbsd
type Cbsd struct {

	// capabilities
	// Required: true
	Capabilities *Capabilities `json:"capabilities"`

	// is the radio type A (only) or B (also applies to A/B type radios)
	// Required: true
	// Enum: [a b]
	CbsdCategory string `json:"cbsd_category"`

	// id of cbsd in SAS
	// Example: some_cbsd_id
	CbsdID string `json:"cbsd_id,omitempty"`

	// desired state of cbsd in SAS
	// Required: true
	// Enum: [unregistered registered]
	DesiredState string `json:"desired_state"`

	// fcc id
	// Example: some_fcc_id
	// Required: true
	// Min Length: 1
	FccID string `json:"fcc_id"`

	// frequency preferences
	// Required: true
	FrequencyPreferences FrequencyPreferences `json:"frequency_preferences"`

	// grant
	Grant *Grant `json:"grant,omitempty"`

	// database id of cbsd
	// Required: true
	ID int64 `json:"id"`

	// installation param
	// Required: true
	InstallationParam *InstallationParam `json:"installation_param"`

	// false if cbsd have not contacted DP for certain amount of time
	// Required: true
	IsActive bool `json:"is_active"`

	// serial number
	// Example: some_serial_number
	// Required: true
	// Min Length: 1
	SerialNumber string `json:"serial_number"`

	// should the CBSD be registered in a single-step mode
	// Required: true
	SingleStepEnabled bool `json:"single_step_enabled"`

	// state of cbsd in SAS
	// Required: true
	// Enum: [unregistered registered]
	State string `json:"state"`

	// user id
	// Example: some_user_id
	// Required: true
	// Min Length: 1
	UserID string `json:"user_id"`
}

// Validate validates this cbsd
func (m *Cbsd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCbsdCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFccID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencyPreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallationParam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSingleStepEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cbsd) validateCapabilities(formats strfmt.Registry) error {

	if err := validate.Required("capabilities", "body", m.Capabilities); err != nil {
		return err
	}

	if m.Capabilities != nil {
		if err := m.Capabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

var cbsdTypeCbsdCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["a","b"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cbsdTypeCbsdCategoryPropEnum = append(cbsdTypeCbsdCategoryPropEnum, v)
	}
}

const (

	// CbsdCbsdCategoryA captures enum value "a"
	CbsdCbsdCategoryA string = "a"

	// CbsdCbsdCategoryB captures enum value "b"
	CbsdCbsdCategoryB string = "b"
)

// prop value enum
func (m *Cbsd) validateCbsdCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cbsdTypeCbsdCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cbsd) validateCbsdCategory(formats strfmt.Registry) error {

	if err := validate.RequiredString("cbsd_category", "body", m.CbsdCategory); err != nil {
		return err
	}

	// value enum
	if err := m.validateCbsdCategoryEnum("cbsd_category", "body", m.CbsdCategory); err != nil {
		return err
	}

	return nil
}

var cbsdTypeDesiredStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unregistered","registered"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cbsdTypeDesiredStatePropEnum = append(cbsdTypeDesiredStatePropEnum, v)
	}
}

const (

	// CbsdDesiredStateUnregistered captures enum value "unregistered"
	CbsdDesiredStateUnregistered string = "unregistered"

	// CbsdDesiredStateRegistered captures enum value "registered"
	CbsdDesiredStateRegistered string = "registered"
)

// prop value enum
func (m *Cbsd) validateDesiredStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cbsdTypeDesiredStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cbsd) validateDesiredState(formats strfmt.Registry) error {

	if err := validate.RequiredString("desired_state", "body", m.DesiredState); err != nil {
		return err
	}

	// value enum
	if err := m.validateDesiredStateEnum("desired_state", "body", m.DesiredState); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateFccID(formats strfmt.Registry) error {

	if err := validate.RequiredString("fcc_id", "body", m.FccID); err != nil {
		return err
	}

	if err := validate.MinLength("fcc_id", "body", m.FccID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateFrequencyPreferences(formats strfmt.Registry) error {

	if err := m.FrequencyPreferences.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequency_preferences")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frequency_preferences")
		}
		return err
	}

	return nil
}

func (m *Cbsd) validateGrant(formats strfmt.Registry) error {
	if swag.IsZero(m.Grant) { // not required
		return nil
	}

	if m.Grant != nil {
		if err := m.Grant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grant")
			}
			return err
		}
	}

	return nil
}

func (m *Cbsd) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateInstallationParam(formats strfmt.Registry) error {

	if err := validate.Required("installation_param", "body", m.InstallationParam); err != nil {
		return err
	}

	if m.InstallationParam != nil {
		if err := m.InstallationParam.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installation_param")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("installation_param")
			}
			return err
		}
	}

	return nil
}

func (m *Cbsd) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("is_active", "body", bool(m.IsActive)); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateSerialNumber(formats strfmt.Registry) error {

	if err := validate.RequiredString("serial_number", "body", m.SerialNumber); err != nil {
		return err
	}

	if err := validate.MinLength("serial_number", "body", m.SerialNumber, 1); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateSingleStepEnabled(formats strfmt.Registry) error {

	if err := validate.Required("single_step_enabled", "body", bool(m.SingleStepEnabled)); err != nil {
		return err
	}

	return nil
}

var cbsdTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unregistered","registered"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cbsdTypeStatePropEnum = append(cbsdTypeStatePropEnum, v)
	}
}

const (

	// CbsdStateUnregistered captures enum value "unregistered"
	CbsdStateUnregistered string = "unregistered"

	// CbsdStateRegistered captures enum value "registered"
	CbsdStateRegistered string = "registered"
)

// prop value enum
func (m *Cbsd) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cbsdTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cbsd) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateUserID(formats strfmt.Registry) error {

	if err := validate.RequiredString("user_id", "body", m.UserID); err != nil {
		return err
	}

	if err := validate.MinLength("user_id", "body", m.UserID, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cbsd based on the context it is used
func (m *Cbsd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrequencyPreferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallationParam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cbsd) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	if m.Capabilities != nil {
		if err := m.Capabilities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *Cbsd) contextValidateFrequencyPreferences(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FrequencyPreferences.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequency_preferences")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frequency_preferences")
		}
		return err
	}

	return nil
}

func (m *Cbsd) contextValidateGrant(ctx context.Context, formats strfmt.Registry) error {

	if m.Grant != nil {
		if err := m.Grant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grant")
			}
			return err
		}
	}

	return nil
}

func (m *Cbsd) contextValidateInstallationParam(ctx context.Context, formats strfmt.Registry) error {

	if m.InstallationParam != nil {
		if err := m.InstallationParam.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installation_param")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("installation_param")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cbsd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cbsd) UnmarshalBinary(b []byte) error {
	var res Cbsd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
