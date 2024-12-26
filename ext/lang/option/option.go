// Copyright 2024 BuddhoIO
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

package option

import (
	"reflect"

	"github.com/buddho-io/golang/ext/lang"
)

// none represents an empty Option.
type none[T any] struct{}

// IsDefined implements Option.IsDefined and returns false for an empty Option.
func (n none[T]) IsDefined() bool {
	return false
}

// IsEmpty implements Option.IsEmpty and returns true for an empty Option.
func (n none[T]) IsEmpty() bool {
	return true
}

// Get implements Option.Get and returns the zero value for the type T.
func (n none[T]) Get() T {
	return lang.Zero[T]()
}

func (n none[T]) GetOrElse(defaultValue T) T {
	return defaultValue
}

var _ lang.Option[int] = none[int]{}

// some represents an Option with a value.
type some[T any] struct {
	value T
}

// IsDefined implements Option.IsDefined and returns true for an Option with a value.
func (s some[T]) IsDefined() bool {
	return true
}

// IsEmpty implements Option.IsEmpty and returns false for an Option with a value.
func (s some[T]) IsEmpty() bool {
	return false
}

// Get implements Option.Get and returns the value of the Option.
func (s some[T]) Get() T {
	return s.value
}

func (s some[T]) GetOrElse(_ T) T { return s.Get() }

var _ lang.Option[int] = some[int]{}

// None returns an empty Option.
func None[T any]() lang.Option[T] {
	return none[T]{}
}

// Some returns an Option with the given value.
func Some[T any](t T) lang.Option[T] {
	return some[T]{t}
}

func Of[T any](v T) lang.Option[T] {
	if v := reflect.ValueOf(v); (
	// Only these types can be nil
	v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Chan ||
		v.Kind() == reflect.Func) && v.IsNil() {
		return None[T]()
	}

	var r T
	if any(v) == any(r) {
		return None[T]()
	}

	return Some(v)
}
