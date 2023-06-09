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
	return validator.NewValidator("zero", func(i interface{}) error {
		if reflect.ValueOf(i).IsZero() {
			return nil
		}
		return errShouldEmpty
	})
}

// Required returns a new Validator to chech whether a value is ZERO,
// which returns an error if the value is ZERO.
//
// The validator name is "required".
func Required() validator.Validator {
	return validator.NewValidator("required", func(i interface{}) error {
		if reflect.ValueOf(i).IsZero() {
			return errCannotEmpty
		}
		return nil
	})
}
