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
	"fmt"
	"time"

	"github.com/xgfone/go-validation/validator"
)

// Time returns a new validator to check whether the string value conforms
// with the given time format.
//
// The validator rule is "time(format)".
func Time(format string) validator.Validator {
	rule := fmt.Sprintf(`time("%s")`, format)
	return validator.NewBoolValidator(rule, func(value string) bool {
		_, err := time.Parse(format, value)
		return err == nil
	}, fmt.Errorf("invalid time for the format '%s'", format))
}

// Duration returns a new validator to check whether the string value is
// a valid duration validated by time.ParseDuration.
//
// The validator rule is "duration".
func Duration() validator.Validator {
	return validator.NewBoolValidator("duration", func(value string) bool {
		_, err := time.ParseDuration(value)
		return err == nil
	}, fmt.Errorf("invalid duration"))
}
