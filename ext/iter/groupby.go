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
	"slices"
)

// GroupBy returns a sequence that groups elements from the given sequence by the
// given key function. The key function is used to determine the group of each
// element. The resulting sequence yields key-value pairs where the key is the
// group and the value is a sequence of elements in that group.
// This function must consume the entire input sequence to group the elements so it's
// not efficient for large sequences.
func GroupBy[T any, K comparable](seq iter.Seq[T], key func(T) K) iter.Seq2[K, iter.Seq[T]] {
	return func(yield func(K, iter.Seq[T]) bool) {
		m := make(map[K][]T)
		seq(func(a T) bool {
			k := key(a)
			m[k] = append(m[k], a)
			return true
		})
		for k, v := range m {
			if !yield(k, slices.Values(v)) {
				return
			}
		}
	}
}
