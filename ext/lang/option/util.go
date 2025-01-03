// Copyright 2025 BuddhoIO
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
	"github.com/buddho-io/golang/ext/lang"
)

func Map[A, B any](o lang.Option[A], f func(A) B) lang.Option[B] {
	if o.IsEmpty() {
		return None[B]()
	}
	return Some[B](f(o.Get()))
}

func FlatMap[A, B any](o lang.Option[A], f func(A) lang.Option[B]) lang.Option[B] {
	if o.IsEmpty() {
		return None[B]()
	}
	return f(o.Get())
}

func Filter[A any](o lang.Option[A], f func(A) bool) lang.Option[A] {
	if o.IsEmpty() || f(o.Get()) {
		return o
	}
	return None[A]()
}

func ForEach[A any](o lang.Option[A], f func(A)) {
	if o.IsDefined() {
		f(o.Get())
	}
}

func Flatten[T any](o lang.Option[lang.Option[T]]) lang.Option[T] {
	if o.IsEmpty() {
		return None[T]()
	}
	return o.Get()
}

// FromEither converts an Either instance to an Option instance. If the Either instance is a Left, an empty Option is returned.
// Otherwise, an Option with the right value is returned.
func FromEither[L, R any](e lang.Either[L, R]) lang.Option[R] {
	if e.IsLeft() {
		return None[R]()
	}
	return Some[R](e.Right())
}

// OrElse returns the given Option if it is defined. Otherwise, it returns the result of the given function.
func OrElse[A any](o lang.Option[A], f func() lang.Option[A]) lang.Option[A] {
	if o.IsDefined() {
		return o
	}
	return f()
}

// Sequence merges a list of Option instances into a single Option instance. If any of the Option instances is empty,
// an empty Option is returned. Otherwise, an Option with the list of values is returned.
func Sequence[A any](os []lang.Option[A]) lang.Option[[]A] {
	values := make([]A, 0, len(os))
	for _, o := range os {
		if o.IsEmpty() {
			return None[[]A]()
		}
		values = append(values, o.Get())
	}
	return Some(values)
}
