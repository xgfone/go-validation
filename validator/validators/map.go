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

func composeValidators(name string, validators ...validator.Validator) (validator.Validator, string) {
	validator := validator.And(validators...)
	desc := validator.String()
	if desc[0] == '(' {
		desc = name + desc
	} else {
		desc = fmt.Sprintf("%s(%s)", name, desc)
	}
	return validator, desc
}

// MapK returns a new Validator to use the given validators to check
// each key of the map.
//
// The validator name is "mapk(validators...)".
func MapK(validators ...validator.Validator) validator.Validator {
	if len(validators) == 0 {
		panic("MapKValidator: need at least one validator")
	}

	_validator, desc := composeValidators("mapk", validators...)
	return validator.NewValidator(desc, func(i any) error {
		switch vs := i.(type) {
		case map[string]string:
			for key := range vs {
				if err := _validator.Validate(key); err != nil {
					return fmt.Errorf("map key '%s' is invalid: %v", key, err)
				}
			}

		case map[string]any:
			for key := range vs {
				if err := _validator.Validate(key); err != nil {
					return fmt.Errorf("map key '%s' is invalid: %v", key, err)
				}
			}

		default:
			vf := reflect.ValueOf(i)
			if vf.Kind() != reflect.Map {
				return fmt.Errorf("expect the value is a map, but got %T", i)
			}

			for _, key := range vf.MapKeys() {
				if err := _validator.Validate(key.Interface()); err != nil {
					return fmt.Errorf("map key '%v' is invalid: %v", key.Interface(), err)
				}
			}
		}

		return nil
	})
}

// MapV returns a new Validator to use the given validators to check
// each value of the map.
//
// The validator rule is "mapv(validators...)".
func MapV(validators ...validator.Validator) validator.Validator {
	if len(validators) == 0 {
		panic("MapVValidator: need at least one validator")
	}

	_validator, desc := composeValidators("mapv", validators...)
	return validator.NewValidator(desc, func(i any) error {
		switch vs := i.(type) {
		case map[string]string:
			for _, value := range vs {
				if err := _validator.Validate(value); err != nil {
					return fmt.Errorf("map value '%s' is invalid: %v", value, err)
				}
			}

		case map[string]any:
			for _, value := range vs {
				if err := _validator.Validate(value); err != nil {
					return fmt.Errorf("map value '%v' is invalid: %v", value, err)
				}
			}

		default:
			vf := reflect.ValueOf(i)
			if vf.Kind() != reflect.Map {
				return fmt.Errorf("expect the value is a map, but got %T", i)
			}

			for iter := vf.MapRange(); iter.Next(); {
				value := iter.Value().Interface()
				if err := _validator.Validate(value); err != nil {
					return fmt.Errorf("map value '%v' is invalid: %v", value, err)
				}
			}
		}

		return nil
	})
}

// KV represents a key-value pair.
type KV struct {
	Key   any
	Value any
}

// MapKV returns a new Validator to use the given validators to check
// each key-value pair of the map.
//
// The value validated by the validators is a KV.
//
// The validator rule is "mapkv(validators...)".
func MapKV(validators ...validator.Validator) validator.Validator {
	if len(validators) == 0 {
		panic("MapKVValidator: need at least one validator")
	}

	_validator, desc := composeValidators("mapkv", validators...)
	return validator.NewValidator(desc, func(i any) error {
		switch vs := i.(type) {
		case map[string]string:
			for key, value := range vs {
				if err := _validator.Validate(KV{Key: key, Value: value}); err != nil {
					return fmt.Errorf("map from key '%v' is invalid: %v", key, err)
				}
			}

		case map[string]any:
			for key, value := range vs {
				if err := _validator.Validate(KV{Key: key, Value: value}); err != nil {
					return fmt.Errorf("map from key '%v' is invalid: %v", key, err)
				}
			}

		default:
			vf := reflect.ValueOf(i)
			if vf.Kind() != reflect.Map {
				return fmt.Errorf("expect the value is a map, but got %T", i)
			}

			for iter := vf.MapRange(); iter.Next(); {
				key := iter.Key().Interface()
				value := iter.Value().Interface()
				if err := _validator.Validate(KV{Key: key, Value: value}); err != nil {
					return fmt.Errorf("map from key '%v' is invalid: %v", key, err)
				}
			}
		}

		return nil
	})
}
