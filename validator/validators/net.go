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
	"fmt"
	"net"

	"github.com/xgfone/go-validation/internal"
	"github.com/xgfone/go-validation/validator"
)

var (
	errInvalidMac        = errors.New("the string is not a valid mac")
	errInvalidStringAddr = errors.New("the string is not a valid address")

	errInvalidIP       = errors.New("invalid ip")
	errInvalidStringIP = errors.New("the string is not a valid ip")

	errInvalidCidr       = errors.New("invalid cidr")
	errInvalidStringCidr = errors.New("the string is not a valid cidr")
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
	return validator.NewValidator("mac", func(i interface{}) error {
		switch v := internal.Indirect(i).(type) {
		case string:
			if normalizeMac(v) == "" {
				return errInvalidMac
			}

		case fmt.Stringer:
			if normalizeMac(v.String()) == "" {
				return errInvalidMac
			}

		default:
			return fmt.Errorf("expect a string, but got %T", i)
		}

		return nil
	})
}

// IP returns a new Validator to chech whether the value is a valid IP.
//
// Support the types: string or net.IP.
//
// The validator rule is "ip".
func IP() validator.Validator {
	return validator.NewValidator("ip", func(i interface{}) error {
		switch v := internal.Indirect(i).(type) {
		case string:
			if net.ParseIP(v) == nil {
				return errInvalidStringIP
			}

		case net.IP:
			switch len(v) {
			case net.IPv4len, net.IPv6len:
			default:
				return errInvalidIP
			}

		case fmt.Stringer:
			if net.ParseIP(v.String()) == nil {
				return errInvalidStringIP
			}

		default:
			return fmt.Errorf("unsupported type %T", i)
		}

		return nil
	})
}

// Cidr returns a new Validator to chech whether the value is a valid cidr.
//
// Support the types: string or net.IPNet.
//
// The validator rule is "cidr".
func Cidr() validator.Validator {
	return validator.NewValidator("cidr", func(i interface{}) error {
		switch v := i.(type) {
		case string:
			if _, _, err := net.ParseCIDR(v); err != nil {
				return errInvalidStringCidr
			}

		case *net.IPNet:
			if v == nil {
				return errInvalidCidr
			}

		case fmt.Stringer:
			if _, _, err := net.ParseCIDR(v.String()); err != nil {
				return errInvalidStringCidr
			}

		default:
			return fmt.Errorf("unsupported type %T", i)
		}

		return nil
	})
}

// Addr returns a new Validator to chech whether the value is a valid HOST:PORT.
//
// Support the types: string.
//
// The validator rule is "addr".
func Addr() validator.Validator {
	return validator.NewValidator("addr", func(i interface{}) error {
		switch v := internal.Indirect(i).(type) {
		case string:
			if host, port, err := net.SplitHostPort(v); err != nil || host == "" || port == "" {
				return errInvalidStringAddr
			}

		case fmt.Stringer:
			if host, port, err := net.SplitHostPort(v.String()); err != nil || host == "" || port == "" {
				return errInvalidStringAddr
			}

		default:
			return fmt.Errorf("unsupported type %T", i)
		}

		return nil
	})
}

func normalizeMac(mac string) string {
	if ha, err := net.ParseMAC(mac); err == nil || len(ha) == 6 {
		return ha.String()
	}
	return ""
}
