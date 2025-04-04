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

package validator

import (
	"errors"
	"testing"
)

func TestOr(t *testing.T) {
	err := errors.New("test")
	validator := Or(
		NewValidator("bool", BoolValidateFunc(func(i any) bool { return true }, err)),
		NewValidator("str", BoolValidateFunc(func(s string) bool { return true }, err)),
	)

	const expect = "(bool || str)"
	if s := validator.String(); s != expect {
		t.Errorf("expect %s, but got %s", expect, s)
	}

	if err := validator.Validate("abc"); err != nil {
		t.Error(err)
	}
}

func TestAnd(t *testing.T) {
	err := errors.New("test")
	validator := And(
		NewValidator("bool", BoolValidateFunc(func(i any) bool { return true }, err)),
		NewValidator("str", BoolValidateFunc(func(s string) bool { return true }, err)),
	)

	const expect = "(bool && str)"
	if s := validator.String(); s != expect {
		t.Errorf("expect %s, but got %s", expect, s)
	}

	if err := validator.Validate("abc"); err != nil {
		t.Error(err)
	}
}
