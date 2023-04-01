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
	"testing"

	"github.com/xgfone/go-validation/validator"
)

func TestArray(t *testing.T) {
	validatefunc := validator.StringBoolValidateFunc(func(s string) bool { return s != "" }, errors.New("test"))
	array := Array(validator.NewValidator("bool", validatefunc))
	expectResultNil(t, "array1", array.Validate([]string{}))
	expectResultNil(t, "array2", array.Validate([]string{"a", "b"}))
	unexpectResultNil(t, "array3", array.Validate([]string{"a", ""}))
}
