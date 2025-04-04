// Copyright 2023~2025 xgfone
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
	"reflect"
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
	Validate(value any) error
	String() string
}

// ValidateFunc represents a validation function.
type ValidateFunc func(value any) (err error)

// NewValidator returns the new Validator based on the validation rule
// and function.
func NewValidator(rule string, validate ValidateFunc) Validator {
	return validator{s: rule, f: validate}
}

type validator struct {
	s string
	f ValidateFunc
}

func (v validator) Validate(i any) error { return v.f(i) }
func (v validator) String() string       { return v.s }

// ************************************************************************* //

// NewErrorValidator is a convenient function to new a Validator
// with the rule and error validation function.
func NewErrorValidator[T any](rule string, validate func(T) error) Validator {
	return NewValidator(rule, ErrorValidateFunc(validate))
}

// NewBoolValidator is a convenient function to new a Validator
// with the rule and bool validation function.
func NewBoolValidator[T any](rule string, validate func(T) bool, err error) Validator {
	return NewValidator(rule, BoolValidateFunc(validate, err))
}

// ************************************************************************* //

func trystring(t, v any) (value any, ok bool) {
	if _, ok := t.(string); !ok {
		return nil, false
	}

	if _v, ok := v.(fmt.Stringer); ok {
		return _v.String(), true
	}

	switch _v := reflect.ValueOf(v); _v.Kind() {
	case reflect.String:
		return _v.String(), true

	case reflect.Ptr:
		if _v.Elem().Kind() == reflect.String {
			return _v.Elem().String(), true
		}
	}

	return nil, false
}

// ErrorValidateFunc converts a T validation function to ValidateFunc,
// which asserts the type of the validated value is T or *T.
func ErrorValidateFunc[T any](validate func(T) error) ValidateFunc {
	if validate == nil {
		panic("ErrorValidateFunc: the validation function must not be nil")
	}

	return func(value any) (err error) {
		switch v := value.(type) {
		case T:
			return validate(v)

		case *T:
			var _v T
			if v != nil {
				_v = *v
			}
			return validate(_v)

		case interface{ ValidatedValue() T }:
			return validate(v.ValidatedValue())

		case interface{ Value() T }:
			return validate(v.Value())

		default:
			var x T
			if _v, ok := trystring(x, value); ok {
				return validate(_v.(T))
			}
			return fmt.Errorf("ErrorValidateFunc[%T]: unsupported type %T", x, value)
		}
	}
}

// BoolValidateFunc converts a T bool validation function to ValidateFunc,
// which asserts the type of the validated value is T or *T.
func BoolValidateFunc[T any](validate func(T) bool, err error) ValidateFunc {
	if validate == nil {
		panic("BoolValidateFunc: the validation function must not be nil")
	}

	return func(value any) error {
		var ok bool
		switch v := value.(type) {
		case T:
			ok = validate(v)

		case *T:
			var _v T
			if v != nil {
				_v = *v
			}
			ok = validate(_v)

		case interface{ ValidatedValue() T }:
			ok = validate(v.ValidatedValue())

		case interface{ Value() T }:
			ok = validate(v.Value())

		default:
			var x T
			if _v, ok := trystring(x, value); ok {
				ok = validate(_v.(T))
			}
			return fmt.Errorf("BoolValidateFunc[%T]: unsupported type %T", x, value)
		}

		if !ok {
			return err
		}
		return nil
	}
}

// ************************************************************************* //

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
func (vs andValidator) Validate(v any) (err error) {
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
func (vs orValidator) Validate(v any) (err error) {
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
