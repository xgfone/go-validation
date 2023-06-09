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
	"fmt"
	"testing"
)

func ExampleExp() {
	valiator := Exp(2, 1, 4) // one of "2", "4", "8", "16"
	fmt.Println(valiator.String())

	fmt.Println(valiator.Validate(1))
	fmt.Println(valiator.Validate(2))
	fmt.Println(valiator.Validate(16))
	fmt.Println(valiator.Validate(32))

	// Output:
	// exp(2,1,4)
	// the integer is not in range [2, 4, 8, 16]
	// <nil>
	// <nil>
	// the integer is not in range [2, 4, 8, 16]
}

func TestMin(t *testing.T) {
	min := Min(10)
	unexpectResultNil(t, "min1", min.Validate(9))
	expectResultNil(t, "min2", min.Validate(10))
	expectResultNil(t, "min3", min.Validate(11))
}

func TestMax(t *testing.T) {
	max := Max(10)
	expectResultNil(t, "max1", max.Validate(9))
	expectResultNil(t, "max2", max.Validate(10))
	unexpectResultNil(t, "max3", max.Validate(11))
}

func TestRanger(t *testing.T) {
	ranger := Ranger(1, 10)
	unexpectResultNil(t, "ranger1", ranger.Validate(0))
	expectResultNil(t, "ranger2", ranger.Validate(1))
	expectResultNil(t, "ranger3", ranger.Validate(10))
	unexpectResultNil(t, "ranger4", ranger.Validate(11))
}
