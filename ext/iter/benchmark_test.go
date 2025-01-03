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
	"testing"
)

func BenchmarkFilter(b *testing.B) {
	count := 1000000
	filter := func(i int) bool {
		return i%2 == 0
	}

	b.Run("Filter", func(b *testing.B) {
		s := Filter(Range(0, count), filter)

		for i := 0; i < b.N; i++ {
			c := Len(s)

			if c < count/2 {
				b.Fatalf("expected result length %d, got %d", count/2, c)
			}
		}
	})

	b.Run("PlainFilter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// In order for the results to be useful for further, we need to collect them.
			// Rational is to make it similar to iter.Seq which can have further processing
			// applied to the results.
			r := make([]int, 0, count/2)
			for j := 0; j < count; j++ {
				if filter(j) {
					r = append(r, j)
				}
			}
			if len(r) < count/2 {
				b.Fatalf("expected result length %d, got %d", count/2, len(r))
			}
		}
	})
}

func BenchmarkMap(b *testing.B) {
	count := 1000000
	m := func(i int) int {
		return i * 2
	}

	b.Run("Map", func(b *testing.B) {
		s := Map(Range(0, count), m)

		for i := 0; i < b.N; i++ {
			c := Len(s)
			if c < count {
				b.Fatalf("expected result length %d, got %d", count, c)
			}
		}
	})

	b.Run("PlainMap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// In order for the results to be useful for further, we need to collect them.
			// Rational is to make it similar to iter.Seq which can have further processing
			// applied to the results.
			r := make([]int, 0, count)
			for j := 0; j < count; j++ {
				r = append(r, m(j))
			}
			if len(r) < count {
				b.Fatalf("expected result length %d, got %d", count, len(r))
			}
		}
	})
}

func BenchmarkChannel(b *testing.B) {
	count := 1000000

	b.Run("Channel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ch := make(chan int)
			go func() {
				for i := 0; i < count; i++ {
					ch <- i
				}
				close(ch)
			}()

			c := Len(Channel(ch))
			if c < count {
				b.Fatalf("expected result length %d, got %d", count, c)
			}
		}
	})

	b.Run("PlainChannel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ch := make(chan int)
			go func() {
				for i := 0; i < count; i++ {
					ch <- i
				}
				close(ch)
			}()

			// In order for the results to be useful for further, we need to collect them.
			// Rational is to make it similar to iter.Seq which can have further processing
			// applied to the results.
			r := make([]int, 0, count)
			for v := range ch {
				r = append(r, v)
			}
			if len(r) < count {
				b.Fatalf("expected result length %d, got %d", count, len(r))
			}
		}
	})
}

func BenchmarkChain(b *testing.B) {
	count := 1000000

	b.Run("Chain", func(b *testing.B) {
		seq1 := Range(0, count/2)
		seq2 := Range(count/2, count)
		s := Filter(
			Map(Concat(seq1, seq2), func(i int) int { return i * 2 }),
			func(i int) bool { return i%2 == 0 },
		)

		for i := 0; i < b.N; i++ {
			c := 0
			for range s {
				c++
			}
			if c < count {
				b.Fatalf("expected result length %d, got %d", count, c)
			}
		}
	})

	b.Run("PlainChain", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r := make([]int, 0, count)
			// concat
			for j := 0; j < count/2; j++ {
				r = append(r, j)
			}
			for j := count / 2; j < count; j++ {
				r = append(r, j)
			}

			// map
			for j := 0; j < count; j++ {
				r[j] *= 2
			}

			// filter
			r2 := make([]int, 0, count)
			for j := 0; j < count; j++ {
				if j%2 == 0 {
					r2 = append(r2, j)
				}
			}

			if len(r2) < count/2 {
				b.Fatalf("expected result length %d, got %d", count, len(r2))
			}
		}
	})
}
