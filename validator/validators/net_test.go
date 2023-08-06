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

func TestAddr(t *testing.T) {
	addr := Addr()
	expectResultNil(t, "addr1", addr.Validate("1.2.3.4:80"))
	unexpectResultNil(t, "addr2", addr.Validate("1.2.3.4:"))
	unexpectResultNil(t, "addr3", addr.Validate("1.2.3.4"))
	unexpectResultNil(t, "addr4", addr.Validate(":80"))
	unexpectResultNil(t, "addr5", addr.Validate("80"))
	unexpectResultNil(t, "addr4", addr.Validate(":"))
}

func TestCIDR(t *testing.T) {
	cidr := Cidr()
	expectResultNil(t, "cidr1", cidr.Validate("1.2.3.4/24"))
	unexpectResultNil(t, "cidr2", cidr.Validate("1.2.3.4/128"))
	unexpectResultNil(t, "cidr3", cidr.Validate("1.2.3.4/"))
	unexpectResultNil(t, "cidr4", cidr.Validate("1.2.3.4"))
	unexpectResultNil(t, "cidr5", cidr.Validate("/32"))
	unexpectResultNil(t, "cidr6", cidr.Validate("32"))
}

func TestMac(t *testing.T) {
	mac := Mac()
	expectResultNil(t, "mac1", mac.Validate("11:22:33:44:55:66"))
	unexpectResultNil(t, "mac2", mac.Validate("11:22:33:44:55:zz"))
	unexpectResultNil(t, "mac3", mac.Validate("11:22:33:44:55:666"))
	unexpectResultNil(t, "mac4", mac.Validate("11:22:33:44:55:"))
	unexpectResultNil(t, "mac5", mac.Validate("11:22:33:44:55"))
}

func TestIP(t *testing.T) {
	ip := IP()
	expectResultNil(t, "ip1", ip.Validate("1.2.3.4"))
	expectResultNil(t, "ip2", ip.Validate("ff00::"))
	unexpectResultNil(t, "ip3", ip.Validate("1.2.3.4:80"))
	unexpectResultNil(t, "ip4", ip.Validate("1.2.3.4/24"))
}

func TestUrl(t *testing.T) {
	url := Url()
	expectResultNil(t, "url1", url.Validate("http://localhost"))
	unexpectResultNil(t, "url2", url.Validate("http://"))
	unexpectResultNil(t, "url3", url.Validate("://localhost"))
	unexpectResultNil(t, "url4", url.Validate("localhost"))
}
