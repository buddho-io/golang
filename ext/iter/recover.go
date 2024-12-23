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

package iter

import (
	"iter"

	"github.com/llinder/golang/ext/lang"
)

// Predicate is a function that returns true if the given value satisfies a condition.
type Predicate[T any] func(T) bool

// RecoverFunc is a function that returns a sequence of values given an error.
type RecoverFunc[T any] func(error) iter.Seq[lang.Either[error, T]]

// RecoverIf applies a function to errors in a sequence and returns a new sequence with the results.
func RecoverIf[T any](s iter.Seq[lang.Either[error, T]], p Predicate[error], rec RecoverFunc[T]) iter.Seq[lang.Either[error, T]] {
	return func(yield func(lang.Either[error, T]) bool) {
		for x := range s {
			if x.IsLeft() && p(x.Left()) {
				f := rec(x.Left())
				for y := range f {
					if !yield(y) {
						return
					}
				}
			} else {
				if !yield(x) {
					return
				}
			}
		}
	}
}

// RecoverFunc2 is a function that returns a sequence of values given an error.
type RecoverFunc2[T any] func(error) iter.Seq2[T, error]

// RecoverIf2 applies a function to errors in a sequence and returns a new sequence with the results.
func RecoverIf2[T any](s iter.Seq2[T, error], p Predicate[error], rec RecoverFunc2[T]) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for x, e := range s {
			if e != nil && p(e) {
				f := rec(e)
				for y, e := range f {
					if !yield(y, e) {
						return
					}
				}
			} else {
				if !yield(x, e) {
					return
				}
			}
		}
	}
}
