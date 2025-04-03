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

package internal

import (
	"fmt"
)

// OneOf is used to check whether a value is one of the values.
type OneOf struct {
	name   string
	desc   string
	values []string
}

// NewOneOf returns a new OneOf with the name and the valid values.
func NewOneOf(name string, values ...string) OneOf {
	if len(values) == 0 {
		panic(fmt.Errorf("%s: the values must be empty", name))
	}

	s, err := encodeJSON(values)
	if err != nil {
		panic(err)
	}

	desc := fmt.Sprintf("%s(%s)", name, s[1:len(s)-1])
	return OneOf{name: name, desc: desc, values: values}
}

// Name returns the name.
func (o OneOf) Name() string { return o.name }

// String returns the description.
func (o OneOf) String() string { return o.desc }

// Validate validates the value i is valid.
func (o OneOf) Validate(i interface{}) error {
	switch v := i.(type) {
	case string:
		if !containString(o.values, v) {
			return fmt.Errorf("the string '%s' is not one of %v", v, o.values)
		}

	case *string:
		var s string
		if v != nil {
			s = *v
		}

		if !containString(o.values, s) {
			return fmt.Errorf("the string '%s' is not one of %v", s, o.values)
		}

	case fmt.Stringer:
		if s := v.String(); !containString(o.values, s) {
			return fmt.Errorf("the string '%s' is not one of %v", s, o.values)
		}

	default:
		return fmt.Errorf("expect a string, but got %T", i)
	}

	return nil
}

func containString(ss []string, s string) bool {
	for i, _len := 0, len(ss); i < _len; i++ {
		if ss[i] == s {
			return true
		}
	}
	return false
}
