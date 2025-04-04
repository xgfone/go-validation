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
	"fmt"
	"reflect"

	"github.com/xgfone/go-validation/validator"
)

// Array returns a new Validator to use the given validators to check
// each element of the array or slice.
//
// The validator name is "array(validators...)".
func Array(validators ...validator.Validator) validator.Validator {
	if len(validators) == 0 {
		panic("ArrayValidator: need at least one validator")
	}

	_validator, desc := composeValidators("array", validators...)
	return validator.NewValidator(desc, func(i any) error {
		switch vs := i.(type) {
		case []string:
			for i, s := range vs {
				if err := _validator.Validate(s); err != nil {
					return fmt.Errorf("%dth element is invalid: %v", i, err)
				}
			}

		default:
			vf := reflect.ValueOf(i)
			if vf.Kind() == reflect.Ptr {
				vf = vf.Elem()
			}
			switch vf.Kind() {
			case reflect.Slice, reflect.Array:
			default:
				return fmt.Errorf("expect the value is a slice or array, but got %T", i)
			}

			for i, _len := 0, vf.Len(); i < _len; i++ {
				vf.Index(i).Interface()
				if err := _validator.Validate(vf.Index(i).Interface()); err != nil {
					return fmt.Errorf("%dth element is invalid: %v", i, err)
				}
			}
		}

		return nil
	})
}
