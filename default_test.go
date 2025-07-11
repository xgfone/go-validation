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

package validation

import (
	"errors"
	"testing"
)

func TestUrlValidation(t *testing.T) {
	url := "http://localhost"
	if err := DefaultBuilder.Validate(url, "url"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}

	url = "localhost"
	if err := DefaultBuilder.Validate(url, "url"); err == nil {
		t.Error("expect an error, but got nil")
	}

	url = "http://"
	if err := DefaultBuilder.Validate(url, "url"); err == nil {
		t.Error("expect an error, but got nil")
	}

	url = "http:///"
	if err := DefaultBuilder.Validate(url, "url"); err == nil {
		t.Error("expect an error, but got nil")
	}

	url = "http:///path"
	if err := DefaultBuilder.Validate(url, "url"); err == nil {
		t.Error("expect an error, but got nil")
	}

	url = "http://localhsot"
	if err := DefaultBuilder.Validate(url, "zero||(max(128) && url)"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}

	url = "/path/to"
	if err := DefaultBuilder.Validate(url, "zero||(max(128) && url)"); err == nil {
		t.Error("expect an error, but got nil")
	}

	urls := []string{"http://localhost/path1", "http://localhost/path2"}
	if err := DefaultBuilder.Validate(urls, "ranger(1,9) && array(url)"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}

	urls = []string{"/path1", "/path2"}
	if err := DefaultBuilder.Validate(urls, "ranger(1,9) && array(url)"); err == nil {
		t.Error("expect an error, but got nil")
	}
}

type _Validator string

func (v _Validator) Validate() error {
	if v == "" {
		return errors.New("must not be empty")
	}
	return nil
}

func TestSelfValidator(t *testing.T) {
	v1 := _Validator("")
	if err := Validate(v1, "self"); err == nil {
		t.Errorf("expect an error, but nil")
	}

	v2 := _Validator("abc")
	if err := Validate(v2, "self"); err != nil {
		t.Errorf("expect nil, but got an error: %v", err)
	}
}
