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

package either

import (
	"github.com/buddho-io/golang/ext/lang"
)

// Map is a right biased map that applies the given function to the right value of the Either instance.
func Map[L, R, R1 any](e lang.Either[L, R], f func(R) R1) lang.Either[L, R1] {
	if e.IsLeft() {
		return Left[L, R1](e.Left())
	}
	return Right[L, R1](f(e.Right()))
}

// MapLeft is a left biased map that applies the given function to the left value of the Either instance.
func MapLeft[L, R, L1 any](e lang.Either[L, R], f func(L) L1) lang.Either[L1, R] {
	if e.IsRight() {
		return Right[L1, R](e.Right())
	}
	return Left[L1, R](f(e.Left()))
}

// FlatMap is a right biased flat map that applies the given function to the right value of the Either instance.
func FlatMap[L, R, R1 any](e lang.Either[L, R], f func(R) lang.Either[L, R1]) lang.Either[L, R1] {
	if e.IsLeft() {
		return Left[L, R1](e.Left())
	}
	return f(e.Right())
}

// FlatMapLeft is a left biased flat map that applies the given function to the left value of the Either instance.
func FlatMapLeft[L, R, L1 any](e lang.Either[L, R], f func(L) lang.Either[L1, R]) lang.Either[L1, R] {
	if e.IsRight() {
		return Right[L1, R](e.Right())
	}
	return f(e.Left())
}

func Filter[L, R any](e lang.Either[L, R], f func(R) bool) lang.Either[L, R] {
	if e.IsLeft() || f(e.Right()) {
		return e
	}
	return Left[L, R](e.Left())
}

func ForEach[L, R any](e lang.Either[L, R], f func(R)) {
	if e.IsRight() {
		f(e.Right())
	}
}

// Flatten returns the right value of the Either instance if it is right. Otherwise, it returns the right value of the inner Either instance.
func Flatten[L, R any](e lang.Either[L, lang.Either[L, R]]) lang.Either[L, R] {
	if e.IsLeft() {
		return Left[L, R](e.Left())
	}
	return e.Right()
}

// FromOption converts an Option instance to an Either instance. If the Option instance is empty, the left value is returned.
// Otherwise, the right value is returned.
func FromOption[L, R any](o lang.Option[R], l L) lang.Either[L, R] {
	if o.IsEmpty() {
		return Left[L, R](l)
	}
	return Right[L, R](o.Get())
}

// OrElse returns the given Either if it is right. Otherwise, it returns the result of the given function.
func OrElse[L, R any](e lang.Either[L, R], f func() lang.Either[L, R]) lang.Either[L, R] {
	if e.IsRight() {
		return e
	}
	return f()
}

// Sequence returns an Either instance with the right value as a slice of the right values of the given Either instances.
// If any of the given Either instances is left, the left value is returned.
func Sequence[L, R any](es []lang.Either[L, R]) lang.Either[L, []R] {
	values := make([]R, 0, len(es))
	for _, e := range es {
		if e.IsLeft() {
			return Left[L, []R](e.Left())
		}
		values = append(values, e.Right())
	}
	return Right[L, []R](values)
}
