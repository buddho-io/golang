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

import (
	"iter"

	"github.com/buddho-io/golang/ext/lang"
)

// Channel returns a sequence that yields values from a channel.
// If the channel is closed, the sequence is complete.
func Channel[T any](ch <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range ch {
			if !yield(value) {
				break
			}
		}
	}
}

// Channel2 returns a sequence that yields key/value pairs from a channel.
// If the channel is closed, the sequence is complete.
func Channel2[K, V any](ch <-chan lang.Tuple2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for pair := range ch {
			if !yield(pair.A(), pair.B()) {
				break
			}
		}
	}
}
