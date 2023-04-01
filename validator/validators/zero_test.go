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

import "testing"

func TestZero(t *testing.T) {
	zero := Zero()
	expectResultNil(t, "zero1", zero.Validate(""))
	expectResultNil(t, "zero2", zero.Validate(0))
	unexpectResultNil(t, "zero3", zero.Validate("abc"))
	unexpectResultNil(t, "zero4", zero.Validate(123))
}

func TestRequired(t *testing.T) {
	required := Required()
	unexpectResultNil(t, "required1", required.Validate(""))
	unexpectResultNil(t, "required2", required.Validate(0))
	expectResultNil(t, "required3", required.Validate("abc"))
	expectResultNil(t, "required4", required.Validate(123))
}
