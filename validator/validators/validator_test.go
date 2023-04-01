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

func expectResultNil(t *testing.T, flag string, result error) {
	if result != nil {
		t.Errorf("%s: expect '<nil>', but got '%v'", flag, result)
	}
}

func unexpectResultNil(t *testing.T, flag string, result error) {
	if result == nil {
		t.Errorf("%s: unexpect '<nil>'", flag)
	}
}
