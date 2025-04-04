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

package internal

import (
	"bytes"
	"encoding/json"
	"reflect"
)

func encodeJSON(value any) (s string, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, 32))
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err = enc.Encode(value); err == nil {
		s = buf.String()
		if _len := len(s); _len > 0 && s[_len-1] == '\n' {
			s = s[:_len-1]
		}
	}
	return
}

func indirect(value any) any {
	if value == nil {
		return nil
	}

	switch vf := reflect.ValueOf(value); vf.Kind() {
	case reflect.Ptr, reflect.Interface:
		if vf.IsNil() {
			return nil
		}
		return indirect(vf.Elem().Interface())

	default:
		return value
	}
}

// Indirect is exported.
func Indirect(value any) any { return indirect(value) }
