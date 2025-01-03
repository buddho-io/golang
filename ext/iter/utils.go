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

package iter

import "iter"

// ToSeq2 converts a Seq into a Seq2. The second type argument is ignored and will always be nil.
func ToSeq2[K, V any](it iter.Seq[K]) iter.Seq2[K, V] {
	return func(f func(K, V) bool) {
		it(func(k K) bool {
			var v V
			return f(k, v)
		})
	}
}
