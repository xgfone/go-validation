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
	"errors"
	"net"
	"net/url"

	"github.com/xgfone/go-validation/validator"
)

// Mac returns a new Validator to chech whether a string is a valid 48-bit MAC.
//
// Support the mac format:
//   - xx:xx:xx:xx:xx:xx
//   - XX:XX:XX:XX:XX:XX
//   - Xx:Xx:Xx:Xx:Xx:Xx
//   - xx-xx-xx-xx-xx-xx
//   - XX-XX-XX-XX-XX-XX
//   - Xx-Xx-Xx-Xx-Xx-Xx
//   - xxxx.xxxx.xxxx
//   - XXXX.XXXX.XXXX
//   - XxXx.XxXx.XxXx
//
// The validator rule is "mac".
func Mac() validator.Validator {
	return validator.NewBoolValidator("mac", func(value string) bool {
		ha, err := net.ParseMAC(value)
		return err == nil && len(ha) == 6
	}, errors.New("the string is not a valid mac"))
}

// IP returns a new Validator to chech whether the value is a valid IP.
//
// The validator rule is "ip".
func IP() validator.Validator {
	return validator.NewBoolValidator("ip", func(value string) bool {
		return net.ParseIP(value) != nil
	}, errors.New("the string is not a valid ip"))
}

// Cidr returns a new Validator to chech whether the value is a valid cidr.
//
// The validator rule is "cidr".
func Cidr() validator.Validator {
	return validator.NewBoolValidator("cidr", func(value string) bool {
		_, _, err := net.ParseCIDR(value)
		return err == nil
	}, errors.New("the string is not a valid cidr"))
}

// Addr returns a new Validator to chech whether the value is a valid HOST:PORT.
//
// The validator rule is "addr".
func Addr() validator.Validator {
	return validator.NewBoolValidator("cidr", func(value string) bool {
		host, port, err := net.SplitHostPort(value)
		return err == nil && host != "" && port != ""
	}, errors.New("the string is not a valid address"))
}

func Url() validator.Validator {
	return validator.NewBoolValidator("url", func(value string) bool {
		u, err := url.Parse(value)
		return err == nil && u.Scheme != "" && u.Host != ""
	}, errors.New("the string is not a valid url"))
}
