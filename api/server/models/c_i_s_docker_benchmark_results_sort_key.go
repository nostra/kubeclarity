// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CISDockerBenchmarkResultsSortKey c i s docker benchmark results sort key
//
// swagger:model CISDockerBenchmarkResultsSortKey
type CISDockerBenchmarkResultsSortKey string

func NewCISDockerBenchmarkResultsSortKey(value CISDockerBenchmarkResultsSortKey) *CISDockerBenchmarkResultsSortKey {
	v := value
	return &v
}

const (

	// CISDockerBenchmarkResultsSortKeyCode captures enum value "code"
	CISDockerBenchmarkResultsSortKeyCode CISDockerBenchmarkResultsSortKey = "code"

	// CISDockerBenchmarkResultsSortKeyLevel captures enum value "level"
	CISDockerBenchmarkResultsSortKeyLevel CISDockerBenchmarkResultsSortKey = "level"
)

// for schema
var cISDockerBenchmarkResultsSortKeyEnum []interface{}

func init() {
	var res []CISDockerBenchmarkResultsSortKey
	if err := json.Unmarshal([]byte(`["code","level"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cISDockerBenchmarkResultsSortKeyEnum = append(cISDockerBenchmarkResultsSortKeyEnum, v)
	}
}

func (m CISDockerBenchmarkResultsSortKey) validateCISDockerBenchmarkResultsSortKeyEnum(path, location string, value CISDockerBenchmarkResultsSortKey) error {
	if err := validate.EnumCase(path, location, value, cISDockerBenchmarkResultsSortKeyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this c i s docker benchmark results sort key
func (m CISDockerBenchmarkResultsSortKey) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCISDockerBenchmarkResultsSortKeyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this c i s docker benchmark results sort key based on context it is used
func (m CISDockerBenchmarkResultsSortKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}