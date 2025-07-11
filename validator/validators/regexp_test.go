// Copyright 2025 xgfone
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

func TestRegexp(t *testing.T) {
	v := Regexp("[a-zA-Z0-9._]{1,16}")

	if s := v.String(); s != `regexp("^[a-zA-Z0-9._]{1,16}$")` {
		t.Errorf("expect rule '%s', but got '%s'", `regexp("^[a-zA-Z0-9._]{1,16}$")`, s)
	}

	if err := v.Validate("abc"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}
	if err := v.Validate("abc123XYZ"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}
	if err := v.Validate("abc.xyz_123"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}
	if err := v.Validate("abc-123"); err == nil {
		t.Errorf("expect an error, but got nil")
	}
	if err := v.Validate("abc01234567890xyz"); err == nil {
		t.Errorf("expect an error, but got nil")
	}
}

func TestRegexpPOSIX(t *testing.T) {
	v := RegexpPOSIX("[a-zA-Z0-9._]+")

	if s := v.String(); s != `posixregexp("^[a-zA-Z0-9._]+$")` {
		t.Errorf("expect rule '%s', but got '%s'", `posixregexp("^[a-zA-Z0-9._]+$")`, s)
	}

	if err := v.Validate("abc"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}
	if err := v.Validate("abc123XYZ"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}
	if err := v.Validate("abc.xyz_123"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}
	if err := v.Validate("abc-123"); err == nil {
		t.Errorf("expect an error, but got nil")
	}
}
