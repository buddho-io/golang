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

// Range returns a sequence of integers from start to end.
func Range(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Range2 returns a sequence of integers from start to end.
func Range2(start, end int) iter.Seq2[int, struct{}] {
	return func(yield func(int, struct{}) bool) {
		for i := start; i < end; i++ {
			if !yield(i, struct{}{}) {
				return
			}
		}
	}
}
