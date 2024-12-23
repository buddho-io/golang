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

// Empty returns an empty sequence.
func Empty[T any]() iter.Seq[T] {
	return func(_ func(T) bool) {}
}

// Empty2 returns an empty sequence.
func Empty2[K, V any]() iter.Seq2[K, V] {
	return func(_ func(K, V) bool) {}
}
