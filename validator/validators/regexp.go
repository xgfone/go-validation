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

import (
	"fmt"
	"regexp"

	"github.com/xgfone/go-validation/validator"
)

// Regexp returns a new validator to check whether the string value conforms
// with the given regexp rule, which uses regexp.MustCompile to compile the rule.
//
// The validator rule is `regexp("rule")`.
func Regexp(rule string) validator.Validator {
	if rule[0] != '^' && rule[len(rule)-1] != '$' {
		rule = fmt.Sprintf("^%s$", rule)
	}

	re := regexp.MustCompile(rule)
	_rule := fmt.Sprintf(`regexp("%s")`, rule)
	return validator.NewBoolValidator(_rule, func(value string) bool {
		return re.MatchString(value)
	}, fmt.Errorf("invalid string for the regexp: %s", rule))
}

// RegexpPOSIX is the same as Regexp, but use regexp.MustCompilePOSIX
// to compile the rule.
//
// The validator rule is `posixregexp("rule")`.
func RegexpPOSIX(rule string) validator.Validator {
	if rule[0] != '^' && rule[len(rule)-1] != '$' {
		rule = fmt.Sprintf("^%s$", rule)
	}

	re := regexp.MustCompilePOSIX(rule)
	_rule := fmt.Sprintf(`posixregexp("%s")`, rule)
	return validator.NewBoolValidator(_rule, func(value string) bool {
		return re.MatchString(value)
	}, fmt.Errorf("invalid string for the posix regexp: %s", rule))
}
