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
	"testing"

	"github.com/xgfone/go-validation/validator"
)

func TestMapK(t *testing.T) {
	minint1 := func(i any) bool { return i.(int) >= 1 }
	maxint9 := func(i any) bool { return i.(int) <= 9 }
	mapk := MapK(
		validator.NewValidator("min", validator.BoolValidateFunc(minint1, errors.New("test"))),
		validator.NewValidator("max", validator.BoolValidateFunc(maxint9, errors.New("test"))),
	)

	expectResultNil(t, "mapk1", mapk.Validate(map[int]string{1: "a", 9: "b"}))
	unexpectResultNil(t, "mapk2", mapk.Validate(map[int]string{0: "a", 1: "b"}))
	unexpectResultNil(t, "mapk3", mapk.Validate(map[int]string{9: "a", 10: "b"}))
}

func TestMapV(t *testing.T) {
	minint1 := func(i any) bool { return i.(int) >= 1 }
	maxint9 := func(i any) bool { return i.(int) <= 9 }
	mapk := MapV(
		validator.NewValidator("min", validator.BoolValidateFunc(minint1, errors.New("test"))),
		validator.NewValidator("max", validator.BoolValidateFunc(maxint9, errors.New("test"))),
	)

	expectResultNil(t, "mapv1", mapk.Validate(map[string]int{"a": 1, "b": 9}))
	unexpectResultNil(t, "mapv2", mapk.Validate(map[string]int{"a": 0, "b": 1}))
	unexpectResultNil(t, "mapv3", mapk.Validate(map[string]int{"a": 9, "b": 10}))
}

func TestMapKV(t *testing.T) {
	kminint1 := func(i any) bool { return i.(KV).Key.(int) >= 1 }
	kmaxint9 := func(i any) bool { return i.(KV).Key.(int) <= 9 }
	vnotzero := func(i any) bool { return i.(KV).Value.(string) != "" }
	mapk := MapKV(
		validator.NewValidator("kmin", validator.BoolValidateFunc(kminint1, errors.New("test"))),
		validator.NewValidator("kmax", validator.BoolValidateFunc(kmaxint9, errors.New("test"))),
		validator.NewValidator("vnotzero", validator.BoolValidateFunc(vnotzero, errors.New("test"))),
	)

	expectResultNil(t, "mapkv1", mapk.Validate(map[int]string{1: "a", 9: "b"}))
	unexpectResultNil(t, "mapkv2", mapk.Validate(map[int]string{0: "a", 1: "b"}))
	unexpectResultNil(t, "mapkv3", mapk.Validate(map[int]string{9: "a", 10: "b"}))
	unexpectResultNil(t, "mapkv3", mapk.Validate(map[int]string{1: "a", 9: ""}))
}
