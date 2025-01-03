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
	"context"
	"iter"
)

// Context returns a new Seq that will yield elements from the input Seq until the context is done.
func Context[T any](ctx context.Context, s iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		done := ctx.Done()
		s(func(a T) bool {
			select {
			case <-done:
				return false
			default:
				return yield(a)
			}
		})
	}
}

// Context2 returns a new Seq2 that will yield elements from the input Seq2 until the context is done.
func Context2[K, V any](ctx context.Context, s iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		done := ctx.Done()
		s(func(k K, v V) bool {
			select {
			case <-done:
				return false
			default:
				return yield(k, v)
			}
		})
	}
}
