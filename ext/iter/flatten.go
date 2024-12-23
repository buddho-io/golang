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

// Flatten returns a new Seq that will yield elements from the nested input Seq.
func Flatten[T any](s iter.Seq[iter.Seq[T]]) iter.Seq[T] {
	return func(g func(T) bool) {
		s(func(a iter.Seq[T]) bool {
			a(func(b T) bool {
				return g(b)
			})
			return true
		})
	}
}
