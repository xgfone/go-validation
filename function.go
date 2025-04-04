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
	"github.com/xgfone/predicate"
)

// Function is called by the builder to parse and build the validator.
type Function interface {
	Call(c *Context, args ...any) error
	Name() string
}

type functionImpl struct {
	call func(*Context, ...any) error
	name string
}

func (f functionImpl) Name() string { return f.name }
func (f functionImpl) Call(c *Context, args ...any) error {
	return f.call(c, args...)
}

func toBuilderFunction(f Function) predicate.BuilderFunction {
	return func(context predicate.BuilderContext, args ...any) error {
		return f.Call(context.(*Context), args...)
	}
}

// NewFunction returns a new Function.
func NewFunction(name string, call func(*Context, ...any) error) Function {
	return functionImpl{name: name, call: call}
}

// ValidatorFunction converts a validator to a Function with the name,
// which is equal to
//
//	NewFunctionWithoutArgs(name, func() validator.Validator { return v })
func ValidatorFunction(name string, v validator.Validator) Function {
	return NewFunctionWithoutArgs(name, func() validator.Validator { return v })
}

// ************************************************************************* //

// NewFunctionWithoutArgs returns a new Function which parses and builds
// the validator without any arguments.
func NewFunctionWithoutArgs(name string, newf func() validator.Validator) Function {
	return NewFunction(name, func(c *Context, args ...any) (err error) {
		if len(args) > 0 {
			err = fmt.Errorf("%s must not have any arguments", name)
		} else {
			c.AppendValidators(newf())
		}
		return
	})
}

// NewFunctionWithOneFloat returns a new Function which parses and builds
// the validator with only one float64 argument.
func NewFunctionWithOneFloat(name string, newf func(float64) validator.Validator) Function {
	return NewFunction(name, func(c *Context, args ...any) (err error) {
		if len(args) != 1 {
			return fmt.Errorf("%s must have and only have one argument", name)
		}

		v, err := getFloat(name, -1, args[0])
		if err == nil {
			c.AppendValidators(newf(v))
		}
		return
	})
}

func getFloat(name string, index int, i any) (f float64, err error) {
	switch v := i.(type) {
	case int:
		f = float64(v)

	case float64:
		f = v

	default:
		if index < 0 {
			err = fmt.Errorf("%s does not support the argument type %T", name, i)
		} else {
			err = fmt.Errorf("%s expects %dth argument is an int or float, but got %T", name, index, i)
		}
	}

	return
}

// NewFunctionWithTwoFloats returns a new Function which parses and builds
// the validator with only two float64 arguments.
func NewFunctionWithTwoFloats(name string, newf func(float64, float64) validator.Validator) Function {
	return NewFunction(name, func(c *Context, args ...any) (err error) {
		if len(args) != 2 {
			return fmt.Errorf("%s must have and only have two arguments", name)
		}

		first, err := getFloat(name, 0, args[0])
		if err != nil {
			return
		}

		second, err := getFloat(name, 1, args[1])
		if err != nil {
			return
		}

		c.AppendValidators(newf(first, second))
		return
	})
}

// NewFunctionWithFloats returns a new Function which parses and builds
// the validator with any float64 arguments.
func NewFunctionWithFloats(name string, newf func(...float64) validator.Validator) Function {
	return NewFunction(name, func(c *Context, args ...any) (err error) {
		vs := make([]float64, len(args))
		for i, v := range args {
			if vs[i], err = getFloat(name, i, v); err != nil {
				return err
			}
		}
		c.AppendValidators(newf(vs...))
		return
	})
}

// NewFunctionWithOneString returns a new Function which parses and builds
// the validator with only one string argument.
func NewFunctionWithOneString(name string, newf func(string) validator.Validator) Function {
	return NewFunction(name, func(c *Context, args ...any) (err error) {
		if len(args) != 1 {
			return fmt.Errorf("%s must have and only have one argument", name)
		}

		if s, ok := args[0].(string); ok {
			c.AppendValidators(newf(s))
		} else {
			err = fmt.Errorf("%s expects a string, but got %T", name, args[0])
		}

		return
	})
}

// NewFunctionWithStrings returns a new Function which parses and builds
// the validator with any string arguments.
func NewFunctionWithStrings(name string, newf func(...string) validator.Validator) Function {
	return NewFunction(name, func(c *Context, args ...any) (err error) {
		var ok bool
		vs := make([]string, len(args))
		for i, v := range args {
			if vs[i], ok = v.(string); !ok {
				return fmt.Errorf("%s expects %dth argument is a string, but got %T", name, i, v)
			}
		}
		c.AppendValidators(newf(vs...))
		return
	})
}

// NewFunctionWithValidators returns a new Function which parses and builds
// the validator with any Validator arguments but at least one.
//
// Notice: the parsed validators is composed to a new Valiator by And.
func NewFunctionWithValidators(name string, newf func(...validator.Validator) validator.Validator) Function {
	return NewFunction(name, func(c *Context, args ...any) (err error) {
		if len(args) == 0 {
			return fmt.Errorf("%s validator has no arguments", name)
		}

		ac := c.New()
		for i, arg := range args {
			b, ok := arg.(predicate.ContextBuilder)
			if !ok {
				return fmt.Errorf("%s expects %dth argument is a validator, but got %T", name, i, arg)
			}

			nc := ac.New()
			if err := b.Build(nc); err != nil {
				return err
			}
			ac.And(nc)
		}

		c.AppendValidators(newf(ac.(*Context).Validators()...))
		return
	})
}

// NewFunctionWithThreeInts returns a new Function which parses and builds
// the validator with only three int arguments.
func NewFunctionWithThreeInts(name string, newf func(int, int, int) validator.Validator) Function {
	return NewFunction(name, func(c *Context, args ...any) (err error) {
		if len(args) != 3 {
			return fmt.Errorf("%s must have and only have three arguments", name)
		}

		first, err := getInt(name, 0, args[0])
		if err != nil {
			return
		}

		second, err := getInt(name, 1, args[1])
		if err != nil {
			return
		}

		third, err := getInt(name, 2, args[2])
		if err != nil {
			return
		}

		c.AppendValidators(newf(first, second, third))
		return
	})
}

func getInt(name string, index int, i any) (v int, err error) {
	v, ok := i.(int)
	if ok {
		return
	}

	if index < 0 {
		err = fmt.Errorf("%s does not support the argument type %T", name, i)
	} else {
		err = fmt.Errorf("%s expects %dth argument is an int, but got %T", name, index, i)
	}

	return
}
