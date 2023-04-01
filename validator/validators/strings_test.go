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

func TestOneOf(t *testing.T) {
	oneof := OneOf("a", "b", "c")
	expectResultNil(t, "oneof1", oneof.Validate("a"))
	unexpectResultNil(t, "oneof2", oneof.Validate("d"))
}

func TestIsNumber(t *testing.T) {
	isnumber := IsNumber()
	expectResultNil(t, "isnumber1", isnumber.Validate("-1.23"))
	unexpectResultNil(t, "isnumber2", isnumber.Validate("abc"))
}

func TestIsInteger(t *testing.T) {
	isinteger := IsInteger()
	expectResultNil(t, "isinteger1", isinteger.Validate("-123"))
	unexpectResultNil(t, "isinteger2", isinteger.Validate("1.2"))
	unexpectResultNil(t, "isinteger3", isinteger.Validate("abc"))
}
