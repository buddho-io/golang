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
	"time"
)

// Throttle returns a sequence that yields elements from the given sequence at
// a maximum rate of one element per duration d.
func Throttle[T any](seq iter.Seq[T], d time.Duration) iter.Seq[T] {
	return func(yield func(T) bool) {
		t := time.NewTicker(d)
		defer t.Stop()

		for s := range seq {
			<-t.C // throttle

			if !yield(s) {
				return
			}
		}
	}
}

// Throttle2 returns a sequence that yields elements from the given sequence at
// a maximum rate of one element per duration d.
func Throttle2[K, V any](seq iter.Seq2[K, V], d time.Duration) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		t := time.NewTicker(d)
		defer t.Stop()

		for k, v := range seq {
			<-t.C // throttle

			if !yield(k, v) {
				return
			}
		}
	}
}
