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
	"strconv"
	"unicode/utf8"

	"github.com/xgfone/go-validation/internal"
	"github.com/xgfone/go-validation/validator"
)

// CountString is used to count the number of the characters in the string.
var CountString func(string) int = utf8.RuneCountInString

// OneOf is equal to OneOfWithName("oneof", values...).
//
// The validator rule is "oneof(values...)".
func OneOf(values ...string) validator.Validator {
	return OneOfWithName("oneof", values...)
}

// OneOfWithName returns a new Validator with the validator name
// to chech whether the string value is one of the given strings.
//
// The validator rule is "name(values...)".
func OneOfWithName(name string, values ...string) validator.Validator {
	return internal.NewOneOf(name, values...)
}

// IsNumber returns a new validator to check whether the string value is a number,
// such as an integer or float.
func IsNumber() validator.Validator {
	return validator.NewBoolValidator("isnumber", func(value string) bool {
		_, err := strconv.ParseFloat(value, 64)
		return err == nil
	}, errors.New("the string is not a number"))
}

// IsInteger returns a new validator to check whether the string value is an integer.
func IsInteger() validator.Validator {
	return validator.NewBoolValidator("isinteger", func(value string) bool {
		_, err := strconv.ParseInt(value, 10, 64)
		return err == nil
	}, errors.New("the string is not an integer"))
}
