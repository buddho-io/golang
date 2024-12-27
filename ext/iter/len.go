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

// Len returns the number of elements in the sequence. This requires iterating
// over the entire sequence so it is not efficient for large sequences.
func Len[T any](seq iter.Seq[T]) int {
	c := 0
	for range seq {
		c++
	}
	return c
}

// Len2 returns the number of elements in the sequence. This requires iterating
// over the entire sequence so it is not efficient for large sequences.
func Len2[K, V any](seq iter.Seq2[K, V]) int {
	c := 0
	for range seq {
		c++
	}
	return c
}
