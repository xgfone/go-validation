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
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/xgfone/go-validation/validator"
	"github.com/xgfone/go-validation/validator/validators"
	"github.com/xgfone/predicate"
)

// DefaultBuilder is the global default validation rule builder,
// which will register some default validator building functions.
// See RegisterDefaults.
var DefaultBuilder = NewBuilder()

// RegisterSymbol is equal to DefaultBuilder.RegisterSymbol(name, value).
func RegisterSymbol(name string, value any) {
	DefaultBuilder.RegisterSymbol(name, value)
}

// RegisterValidator is equal to DefaultBuilder.RegisterValidator(name, v).
func RegisterValidator(name string, v validator.Validator) {
	DefaultBuilder.RegisterValidator(name, v)
}

// RegisterValidatorFunc is equal to DefaultBuilder.RegisterValidatorFunc(name, f).
func RegisterValidatorFunc(name string, f validator.ValidateFunc) {
	DefaultBuilder.RegisterValidatorFunc(name, f)
}

// RegisterValidatorOneof is equal to DefaultBuilder.RegisterValidatorOneof(name, values...).
func RegisterValidatorOneof(name string, values ...string) {
	DefaultBuilder.RegisterValidatorOneof(name, values...)
}

// Validate is equal to DefaultBuilder.Validate(v, rule).
func Validate(v any, rule string) error {
	return DefaultBuilder.Validate(v, rule)
}

// Builder is used to build the validator based on the rule.
type Builder struct {
	// Symbols is used to define the global symbols,
	// which is used by the default of GetIdentifier.
	Symbols map[string]any

	*predicate.Builder
	validators atomic.Value
	vcacheLock sync.Mutex
	vcacheMap  map[string]validator.Validator
}

// NewBuilder returns a new validation rule builder.
func NewBuilder() *Builder {
	builder := &Builder{
		vcacheMap: make(map[string]validator.Validator),
		Symbols:   make(map[string]any),
	}

	builder.Builder = predicate.NewBuilder()
	builder.Builder.GetIdentifier = builder.getIdentifier
	builder.Builder.EQ = builder.eq

	builder.updateValidators()
	return builder
}

func (b *Builder) getIdentifier(selector []string) (any, error) {
	// Support the format "zero" instead of "zero()"

	// First, lookup the function table.
	if f := b.GetFunc(selector[0]); f != nil {
		return f, nil
	}

	// Second, lookup the symbol table.
	if v, ok := b.Symbols[selector[0]]; ok {
		return v, nil
	}

	// We find no the identifier.
	return nil, fmt.Errorf("%s is not defined", selector[0])
}

func (b *Builder) eq(ctx predicate.BuilderContext, left, right any) error {
	// Support the format "min == 123" or "123 == min"
	if f, ok := left.(predicate.BuilderFunction); ok {
		return f(ctx, right)
	}
	if f, ok := right.(predicate.BuilderFunction); ok {
		return f(ctx, left)
	}
	return fmt.Errorf("left or right is not BuilderFunction: %T, %T", left, right)
}

// Validators returns the names of all the validators.
func (b *Builder) ValidatorNames() []string {
	return b.Builder.GetAllFuncNames()
}

// RegisterSymbol registers the symbol with the name and value.
func (b *Builder) RegisterSymbol(name string, value any) {
	if name == "" {
		panic("the symbol name must not be empty")
	}
	if value == nil {
		panic("the symbol value must not be nil")
	}
	b.Symbols[name] = value
}

// RegisterSymbolNames registers a set of symbols with the names,
// the value of whose are equal to their name.
func (b *Builder) RegisterSymbolNames(names ...string) {
	for _, name := range names {
		b.RegisterSymbol(name, name)
	}
}

// RegisterFunction registers the builder function.
//
// If the function has existed, reset it to the new function.
func (b *Builder) RegisterFunction(function Function) {
	b.Builder.RegisterFunc(function.Name(), toBuilderFunction(function))
}

// RegisterValidator is the convenient method to convert the validator
// to the builder function, which is equal to
//
//	b.RegisterFunction(ValidatorFunction(name, validator))
func (b *Builder) RegisterValidator(name string, validator validator.Validator) {
	b.RegisterFunction(ValidatorFunction(name, validator))
}

// RegisterValidatorFunc is a convenient method, which is equal to
//
//	b.RegisterValidator(name, validator.NewValidator(name, f))
func (b *Builder) RegisterValidatorFunc(name string, f validator.ValidateFunc) {
	b.RegisterValidator(name, validator.NewValidator(name, f))
}

// RegisterValidatorOneof is a convenient method to register a oneof validator
// as the builder Function, which is equal to
//
//	b.RegisterValidator(name, validators.OneOfWithName(name, values...))
func (b *Builder) RegisterValidatorOneof(name string, values ...string) {
	b.RegisterValidator(name, validators.OneOfWithName(name, values...))
}

// Build parses and builds the validation rule into the context.
func (b *Builder) Build(c *Context, rule string) error {
	return b.Builder.Build(c, rule)
}

// BuildValidator builds a validator from the validation rule.
//
// If the rule has been built, returns it from the caches.
func (b *Builder) BuildValidator(rule string) (validator.Validator, error) {
	if rule == "" {
		return nil, errors.New("the validation rule must not be empty")
	}

	if validator, ok := b.loadValidator(rule); ok {
		return validator, nil
	}

	b.vcacheLock.Lock()
	defer b.vcacheLock.Unlock()

	if validator, ok := b.loadValidator(rule); ok {
		return validator, nil
	}

	c := NewContext()
	if err := b.Build(c, rule); err != nil {
		return nil, err
	}

	validator := c.Validator()
	b.vcacheMap[rule] = validator
	b.updateValidators()

	return validator, nil
}

func (b *Builder) loadValidator(rule string) (v validator.Validator, ok bool) {
	v, ok = b.validators.Load().(map[string]validator.Validator)[rule]
	return
}

func (b *Builder) updateValidators() {
	validators := make(map[string]validator.Validator, len(b.vcacheMap))
	for rule, validator := range b.vcacheMap {
		validators[rule] = validator
	}
	b.validators.Store(validators)
}

// Validate validates whether the value v is valid by the rule.
//
// If failing to build the rule to the validator, panic with the error.
func (b *Builder) Validate(v any, rule string) (err error) {
	if rule == "" {
		return nil
	}

	validator, err := b.BuildValidator(rule)
	if err != nil {
		panic(err)
	}
	return validator.Validate(v)
}
