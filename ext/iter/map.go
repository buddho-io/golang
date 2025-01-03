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

// Map applies a function to each element of a sequence and returns a new sequence with the results.
func Map[A, B any](s iter.Seq[A], f func(A) B) iter.Seq[B] {
	return func(g func(B) bool) {
		s(func(a A) bool {
			return g(f(a))
		})
	}
}

// Map2 applies a function to each element of a sequence and returns a new sequence with the results.
func Map2[A, B, A1, B1 any](s iter.Seq2[A, B], f func(A, B) (A1, B1)) iter.Seq2[A1, B1] {
	return func(g func(A1, B1) bool) {
		s(func(a A, b B) bool {
			return g(f(a, b))
		})
	}
}
