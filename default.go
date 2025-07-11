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
	"fmt"

	"github.com/xgfone/go-validation/validator"
	"github.com/xgfone/go-validation/validator/str"
	"github.com/xgfone/go-validation/validator/validators"
)

func init() { RegisterDefaultsForBuilder(DefaultBuilder) }

// RegisterDefaultsForBuilder registers the default symbols and validators
// building functions into the builder.
//
// The registered default symbols:
//
//	timelayout: 15:04:05
//	datelayout: 2006-01-02
//	datetimelayout: 2006-01-02 15:04:05
//
// The Signature of the registered validator functions as follow:
//
//	ip() or ip
//	mac() or mac
//	url() or url
//	addr() or addr
//	cidr() or cidr
//	zero() or zero
//	empty() or empty
//	notzero() or notzero
//	notempty() or notempty
//	isinteger() or isinteger
//	isnumber() or isnumber
//	duration() or duration
//	required() or required
//	structure() or structure
//	exp(base, startExp, endExp int)
//	min(float64)
//	max(float64)
//	ranger(min, max float64)
//	time(formatLayout string)
//	oneof(...string)
//	array(...Validator)
//	mapkv(...Validator)
//	mapk(...Validator)
//	mapv(...Validator)
//	timeformat() or timeformat => time(timelayout)
//	dateformat() or dateformat => time(datelayout)
//	datetimeformat() or datetimeformat => time(datetimelayout)
//	posixregexp(rule string)
//	regexp(rule string)
//	self() or self: the validated value must have implemented validator.ValueValidator.
func RegisterDefaultsForBuilder(b *Builder) {
	b.RegisterSymbol("timelayout", "15:04:05")
	b.RegisterSymbol("datelayout", "2006-01-02")
	b.RegisterSymbol("datetimelayout", "2006-01-02 15:04:05")
	registerTimeValidator(b, "timeformat", "15:04:05")
	registerTimeValidator(b, "dateformat", "2006-01-02")
	registerTimeValidator(b, "datetimeformat", "2006-01-02 15:04:05")

	b.RegisterFunction(NewFunctionWithoutArgs("zero", validators.Zero))
	b.RegisterFunction(NewFunctionWithoutArgs("empty", validators.Empty))
	b.RegisterFunction(NewFunctionWithoutArgs("notzero", validators.NotZero))
	b.RegisterFunction(NewFunctionWithoutArgs("notempty", validators.NotEmpty))
	b.RegisterFunction(NewFunctionWithoutArgs("required", validators.Required))
	b.RegisterFunction(NewFunctionWithoutArgs("isnumber", validators.IsNumber))
	b.RegisterFunction(NewFunctionWithoutArgs("isinteger", validators.IsInteger))

	b.RegisterFunction(NewFunctionWithoutArgs("ip", validators.IP))
	b.RegisterFunction(NewFunctionWithoutArgs("mac", validators.Mac))
	b.RegisterFunction(NewFunctionWithoutArgs("url", validators.Url))
	b.RegisterFunction(NewFunctionWithoutArgs("cidr", validators.Cidr))
	b.RegisterFunction(NewFunctionWithoutArgs("addr", validators.Addr))

	b.RegisterFunction(NewFunctionWithOneFloat("min", validators.Min))
	b.RegisterFunction(NewFunctionWithOneFloat("max", validators.Max))
	b.RegisterFunction(NewFunctionWithTwoFloats("ranger", validators.Ranger))
	b.RegisterFunction(NewFunctionWithThreeInts("exp", validators.Exp))

	b.RegisterFunction(NewFunctionWithOneString("time", validators.Time))
	b.RegisterFunction(NewFunctionWithoutArgs("duration", validators.Duration))

	b.RegisterFunction(NewFunctionWithOneString("regexp", validators.Regexp))
	b.RegisterFunction(NewFunctionWithOneString("posixregexp", validators.RegexpPOSIX))

	b.RegisterFunction(NewFunctionWithStrings("oneof", validators.OneOf))
	b.RegisterFunction(NewFunctionWithValidators("array", validators.Array))
	b.RegisterFunction(NewFunctionWithValidators("mapk", validators.MapK))
	b.RegisterFunction(NewFunctionWithValidators("mapv", validators.MapV))
	b.RegisterFunction(NewFunctionWithValidators("mapkv", validators.MapKV))

	b.RegisterValidatorFunc("self", func(value any) (err error) {
		return value.(validator.ValueValidator).Validate()
	})
}

func registerTimeValidator(b *Builder, name, layout string) {
	b.RegisterFunction(NewFunctionWithoutArgs(name, func() validator.Validator {
		return validators.Time(layout)
	}))
}

// RegisterStringValidatorsForBuilder registers some string validators,
// that's, the value is a specific string.
//
//	isascii: [\x00-\x7F]+
//	isalpha: [a-zA-Z]+
//	isalphanumeric: [a-zA-Z0-9]+
//	isbase64
//	iscrc32
//	iscrc64
//	isdnsname
//	isdatauri
//	ise164
//	isemail
//	isexistingemail
//	isfloat
//	ishexadecimal: [0-9a-fA-F]+
//	ishexcolor
//	ishost
//	isimei
//	isimsi
//	isipv4
//	isipv6
//	isisbn10
//	isisbn13
//	isint
//	isjson
//	islatitude
//	islongitude
//	islowercase
//	ismd4
//	ismd5
//	ismagneturi
//	ismongoid
//	isprintableascii
//	isrfc3390
//	isrfc3390withoutzone
//	isrgbcolor
//	isrequesturi
//	isrequesturl
//	isripemd128
//	isripemd160
//	issha1
//	issha256
//	issha3224
//	issha3256
//	issha3384
//	issha3512
//	issha384
//	issha512
//	isssn
//	issemver
//	istiger128
//	istiger160
//	istiger192
//	isulid
//	isurl
//	isutfdigit
//	isutfletter
//	isutfletternumeric
//	isutfnumeric
//	isuuid
//	isuuid3
//	isuuid4
//	isuuid5
//	isuppercase
func RegisterStringValidatorsForBuilder(b *Builder) {
	registerStrValidator(b, str.IsASCII, "ascii")
	registerStrValidator(b, str.IsAlpha, "alpha")
	registerStrValidator(b, str.IsAlphanumeric, "alphanumeric")
	registerStrValidator(b, str.IsBase64, "base64")
	registerStrValidator(b, str.IsCRC32, "crc32")
	registerStrValidator(b, str.IsCRC32b, "crc64")
	registerStrValidator(b, str.IsDNSName, "dnsname")
	registerStrValidator(b, str.IsDataURI, "datauri")
	registerStrValidator(b, str.IsE164, "e164")
	registerStrValidator(b, str.IsEmail, "email")
	registerStrValidator(b, str.IsExistingEmail, "existingemail")
	registerStrValidator(b, str.IsFloat, "float")
	registerStrValidator(b, str.IsHexadecimal, "hexadecimal")
	registerStrValidator(b, str.IsHexcolor, "hexcolor")
	registerStrValidator(b, str.IsHost, "host")
	registerStrValidator(b, str.IsIMEI, "imei")
	registerStrValidator(b, str.IsIMSI, "imsi")
	registerStrValidator(b, str.IsIPv4, "ipv4")
	registerStrValidator(b, str.IsIPv6, "ipv6")
	registerStrValidator(b, str.IsISBN10, "isbn10")
	registerStrValidator(b, str.IsISBN13, "isbn13")
	registerStrValidator(b, str.IsInt, "int")
	registerStrValidator(b, str.IsJSON, "json")
	registerStrValidator(b, str.IsLatitude, "latitude")
	registerStrValidator(b, str.IsLongitude, "longitude")
	registerStrValidator(b, str.IsLowerCase, "lowercase")
	registerStrValidator(b, str.IsMD4, "md4")
	registerStrValidator(b, str.IsMD5, "md5")
	registerStrValidator(b, str.IsMagnetURI, "magneturi")
	registerStrValidator(b, str.IsMongoID, "mongoid")
	registerStrValidator(b, str.IsPrintableASCII, "printableascii")
	registerStrValidator(b, str.IsRFC3339, "rfc3390")
	registerStrValidator(b, str.IsRFC3339WithoutZone, "rfc3390withoutzone")
	registerStrValidator(b, str.IsRGBcolor, "rgbcolor")
	registerStrValidator(b, str.IsRequestURI, "requesturi")
	registerStrValidator(b, str.IsRequestURL, "requesturl")
	registerStrValidator(b, str.IsRipeMD128, "ripemd128")
	registerStrValidator(b, str.IsRipeMD160, "ripemd160")
	registerStrValidator(b, str.IsSHA1, "sha1")
	registerStrValidator(b, str.IsSHA256, "sha256")
	registerStrValidator(b, str.IsSHA3224, "sha3224")
	registerStrValidator(b, str.IsSHA3256, "sha3256")
	registerStrValidator(b, str.IsSHA3384, "sha3384")
	registerStrValidator(b, str.IsSHA3512, "sha3512")
	registerStrValidator(b, str.IsSHA384, "sha384")
	registerStrValidator(b, str.IsSHA512, "sha512")
	registerStrValidator(b, str.IsSSN, "ssn")
	registerStrValidator(b, str.IsSemver, "semver")
	registerStrValidator(b, str.IsTiger128, "tiger128")
	registerStrValidator(b, str.IsTiger160, "tiger160")
	registerStrValidator(b, str.IsTiger192, "tiger192")
	registerStrValidator(b, str.IsULID, "ulid")
	registerStrValidator(b, str.IsURL, "url")
	registerStrValidator(b, str.IsUTFDigit, "utfdigit")
	registerStrValidator(b, str.IsUTFLetter, "utfletter")
	registerStrValidator(b, str.IsUTFLetterNumeric, "utfletternumeric")
	registerStrValidator(b, str.IsUTFNumeric, "utfnumeric")
	registerStrValidator(b, str.IsUUID, "uuid")
	registerStrValidator(b, str.IsUUIDv3, "uuid3")
	registerStrValidator(b, str.IsUUIDv4, "uuid4")
	registerStrValidator(b, str.IsUUIDv5, "uuid5")
	registerStrValidator(b, str.IsUpperCase, "uppercase")
}

func registerStrValidator(b *Builder, f func(string) bool, name string) {
	err := fmt.Errorf("the string is not %s", name)
	b.RegisterValidatorFunc("is"+name, validator.BoolValidateFunc(f, err))
}
