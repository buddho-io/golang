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

import "iter"

// Filter returns a new iterator that yields only the elements of the input iterator for which the predicate function returns true.
func Filter[T any](m iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		m(func(a T) bool {
			if f(a) {
				return yield(a)
			}
			return true
		})
	}
}

// Filter2 returns a new iterator that yields only the elements of the input iterator for which the predicate function returns true.
func Filter2[K, V any](m iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m(func(k K, v V) bool {
			if f(k, v) {
				return yield(k, v)
			}
			return true
		})
	}
}
