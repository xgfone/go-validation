// Copyright 2023 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package validator provides a validator interface.
package validator

import (
	"fmt"
	"strings"
)

// ValueValidator represents the interface implemented by the value.
//
// If a value has implemented the interface, it can validate itself.
type ValueValidator interface {
	Validate() error
}

// Validator is a validator to check whether the given value is valid.
type Validator interface {
	Validate(value interface{}) error
	String() string
}

// ValidateFunc represents a validation function.
type ValidateFunc func(value interface{}) error

// NewValidator returns the new Validator based on the validation rule
// and function.
func NewValidator(rule string, validate ValidateFunc) Validator {
	return validator{s: rule, f: validate}
}

type validator struct {
	s string
	f ValidateFunc
}

func (v validator) Validate(i interface{}) error { return v.f(i) }
func (v validator) String() string               { return v.s }

// BoolValidateFunc converts a bool validation function to ValidateFunc,
// which returns err if validate returns false, or nil if true.
func BoolValidateFunc(validate func(interface{}) bool, err error) ValidateFunc {
	if err == nil {
		panic("BoolValidateFunc: the error must not be nil")
	}
	if validate == nil {
		panic("BoolValidateFunc: the validation function must not be nil")
	}

	return func(i interface{}) error {
		if validate(i) {
			return nil
		}
		return err
	}
}

// StringBoolValidateFunc converts a bool validation function to ValidateFunc,
// which returns err if validate returns false, or nil if true.
func StringBoolValidateFunc(validate func(string) bool, err error) ValidateFunc {
	if err == nil {
		panic("StringBoolValidateFunc: the error must not be nil")
	}
	if validate == nil {
		panic("StringBoolValidateFunc: the validation function must not be nil")
	}

	return func(i interface{}) error {
		var ok bool
		switch t := i.(type) {
		case string:
			ok = validate(t)

		case *string:
			ok = t != nil && validate(*t)

		case fmt.Stringer:
			ok = validate(t.String())

		default:
			return fmt.Errorf("unsupported type '%T'", i)
		}

		if ok {
			return nil
		}
		return err
	}
}

func formatValidators(sep string, validators []Validator) string {
	switch len(validators) {
	case 0:
		return ""
	case 1:
		return validators[0].String()
	}

	var b strings.Builder
	b.Grow(32)

	b.WriteByte('(')
	for i, validator := range validators {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(validator.String())
	}
	b.WriteByte(')')

	return b.String()
}

// ************************************************************************* //

// AndValidator is a And validator based on a set of the validators.
type andValidator []Validator

// Validate implements the interface Validator.
func (vs andValidator) Validate(v interface{}) (err error) {
	for i, _len := 0, len(vs); i < _len; i++ {
		if err = vs[i].Validate(v); err != nil {
			return
		}
	}
	return
}

func (vs andValidator) String() string {
	return formatValidators(" && ", []Validator(vs))
}

// And returns a new And Validator.
func And(validators ...Validator) Validator {
	switch len(validators) {
	case 0:
		panic("AndValidator: no validators")
	case 1:
		return validators[0]
	}

	vs := make(andValidator, 0, len(validators))
	for _, v := range validators {
		if andv, ok := v.(andValidator); ok {
			vs = append(vs, []Validator(andv)...)
		} else {
			vs = append(vs, v)
		}
	}

	return andValidator(vs)
}

// ************************************************************************* //

// OrValidator is a OR validator based on a set of the validators.
type orValidator []Validator

// Validate implements the interface Validator.
func (vs orValidator) Validate(v interface{}) (err error) {
	for i, _len := 0, len(vs); i < _len; i++ {
		if err = vs[i].Validate(v); err == nil {
			return nil
		}
	}
	return
}

func (vs orValidator) String() string {
	return formatValidators(" || ", []Validator(vs))
}

// Or returns a new OR Validator.
func Or(validators ...Validator) Validator {
	switch len(validators) {
	case 0:
		panic("OrValidator: no validators")
	case 1:
		return validators[0]
	}

	vs := make(orValidator, 0, len(validators))
	for _, v := range validators {
		if orv, ok := v.(orValidator); ok {
			vs = append(vs, []Validator(orv)...)
		} else {
			vs = append(vs, v)
		}
	}

	return orValidator(vs)
}
