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

package validators

import (
	"errors"
	"reflect"

	"github.com/xgfone/go-validation/validator"
)

var (
	errShouldEmpty = errors.New("the value should be empty")
	errCannotEmpty = errors.New("the value cannot be empty")
)

// Zero returns a new Validator to chech whether the value is ZERO,
// which returns an error if the value is not ZERO.
//
// The validator rule is "zero".
func Zero() validator.Validator {
	return zeroValidator("zero", true, errShouldEmpty)
}

// Empty is equal to Zero, which is the alias of Zero.
//
// The validator rule is "empty".
func Empty() validator.Validator {
	return zeroValidator("empty", true, errShouldEmpty)
}

// NotZero returns a new Validator to chech whether a value is not ZERO,
// which returns an error if the value is ZERO.
//
// The validator name is "notzero".
func NotZero() validator.Validator {
	return zeroValidator("notzero", false, errCannotEmpty)
}

// NotEmpty is equal to NotZero, which is the alias of NotZero.
//
// The validator name is "notempty".
func NotEmpty() validator.Validator {
	return zeroValidator("notempty", false, errCannotEmpty)
}

// Required is equal to NotZero, which is the alias of NotZero.
//
// The validator name is "required".
func Required() validator.Validator {
	return zeroValidator("required", false, errCannotEmpty)
}

func zeroValidator(name string, zero bool, err error) validator.Validator {
	return validator.NewValidator(name, func(i any) error {
		if iszero(i) == zero {
			return nil
		}
		return err
	})
}

func iszero(v any) bool {
	if i, ok := v.(interface{ IsZero() bool }); ok && i.IsZero() {
		return true
	}
	if reflect.ValueOf(v).IsZero() {
		return true
	}
	return false
}
