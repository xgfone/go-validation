# Go Validation [![Build Status](https://github.com/xgfone/go-validation/actions/workflows/go.yml/badge.svg)](https://github.com/xgfone/go-validation/actions/workflows/go.yml) [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/go-validation)](https://pkg.go.dev/github.com/xgfone/go-validation) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-validation/master/LICENSE)

Provide a validation framework based on the built rule, supporting `Go1.13+`.


## Install
```shell
$ go get -u github.com/xgfone/go-validation
```


## Example
For registering the validator and validating whether a value is valid, See [Builder](https://pkg.go.dev/github.com/xgfone/go-validation/#example-Builder).

```go
package main

import "github.com/xgfone/go-validation"

func main() {
	// Validate whether an integer is in [min, max].
	validation.Validate(123, `min(1) && max(200)`) // => <nil>
	validation.Validate(456, `min(1) && max(200)`) // => an error
	validation.Validate(123, `ranger(1, 200)`)     // => <nil>
	validation.Validate(456, `ranger(1, 200)`)     // => an error

	// Validate whether an string is one of a string list.
	validation.Validate("a", `oneof("a", "b", "c")`) // => <nil>
	validation.Validate("d", `oneof("a", "b", "c")`) // => an error

	// Validate whether an string is an integer string that can be parsed to an integer.
	validation.Validate("123", `isinteger`)  // => <nil>
	validation.Validate("+123", `isinteger`) // => <nil>
	validation.Validate("-123", `isinteger`) // => <nil>
	validation.Validate("12.3", `isinteger`) // => an error
	validation.Validate("abc", `isinteger`)  // => an error

	// Validate whether an string is an float string that can be parsed to an float.
	validation.Validate("123", `isnumber`)  // => <nil>
	validation.Validate("12.3", `isnumber`) // => <nil>
	validation.Validate("-1.2", `isnumber`) // => <nil>
	validation.Validate("123.", `isnumber`) // => <nil>
	validation.Validate(".123", `isnumber`) // => <nil>
	validation.Validate("abc", `isnumber`)  // => an error

	// Validate whether a value is ZERO.
	validation.Validate(0, `zero`)     // => <nil>
	validation.Validate(1, `zero`)     // => an error
	validation.Validate("", `zero`)    // => <nil>
	validation.Validate("0", `zero`)   // => an error
	validation.Validate(false, `zero`) // => <nil>
	validation.Validate(true, `zero`)  // => an error
	var p *int
	validation.Validate(p, `zero`) // => <nil>
	p = new(int)
	validation.Validate(p, `zero`) // => an error

	// Validate whether a value is not ZERO.
	validation.Validate(0, `required`)     // => an error
	validation.Validate(1, `required`)     // => <nil>
	validation.Validate("", `required`)    // => an error
	validation.Validate("0", `required`)   // => <nil>
	validation.Validate(false, `required`) // => an error
	validation.Validate(true, `required`)  // => <nil>
	p = nil
	validation.Validate(p, `required`) // => an error
	p = new(int)
	validation.Validate(p, `required`) // => <nil>

	// Validate whether a slice/array is valid.
	validation.Validate([]int{1, 2}, `array(min(1), max(10))`)  // slice, => <nil>
	validation.Validate([3]int{1, 2}, `array(min(1), max(10))`) // array, => an error

	// Validate whether a map(key, value or key-value) is valid.
	maps := map[int]string{0: "a", 1: "b"}
	validation.Validate(maps, `mapk(min(0), max(10))`) // => <nil>
	validation.Validate(maps, `mapk(min(1), max(10))`) // => an error
	validation.Validate(maps, `mapv(oneof("a", "b"))`) // => <nil>
	validation.Validate(maps, `mapv(oneof("a", "c"))`) // => an error

    // For the validation rule, it support the multi-level AND/OR. For example,
    // - "min(100) && max(200)"
    //    => Only a value, such as integer or length of string/slice, in [100, 200] is valid.
    // - "min(200) || max(100)"
    //    => Only a value, such as integer or length of string/slice, in (-∞, 100] or [200, +∞) is valid.
    // - "(min(200) || max(100)) && required)"
    //    => Same as "min(200) || max(100)", but also cannot be ZERO.
}
```
