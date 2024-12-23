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

// FlatMap returns a new Seq that will yield elements from the input Seq that are the result of applying the function to each element.
func FlatMap[A, B any](s iter.Seq[A], f func(A) iter.Seq[B]) iter.Seq[iter.Seq[B]] {
	return func(yield func(iter.Seq[B]) bool) {
		s(func(a A) bool {
			return yield(f(a))
		})
	}
}

// FlatMapConcat returns a new Seq that will yield elements from the input Seq that are the result of applying the
// function to each element and concatenating the results.
func FlatMapConcat[A, B any](s iter.Seq[A], f func(A) iter.Seq[B]) iter.Seq[B] {
	return Flatten(FlatMap(s, f))
}
