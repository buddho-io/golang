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

	"github.com/buddho-io/golang/ext/lang"
)

// DeferredF is a function that returns a value and a boolean indicating if the sequence is complete.
type DeferredF[T any] func() (T, bool)

// Deferred returns a sequence that yields values from a function.
func Deferred[T any](f DeferredF[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			v, ok := f()
			if !ok {
				return
			}

			if !yield(v) {
				return
			}
		}
	}
}

// Deferred2F is a function that returns a tuple and a boolean indicating if the sequence is complete.
type Deferred2F[K, V any] func() (lang.Tuple2[K, V], bool)

// Deferred2 returns a sequence that yields key/value pairs from a function.
func Deferred2[K, V any](f Deferred2F[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for {
			t, ok := f()
			if !ok {
				return
			}

			if !yield(t.A(), t.B()) {
				return
			}
		}
	}
}
