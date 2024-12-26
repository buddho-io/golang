package iter

import (
	"slices"
	"testing"
)

func BenchmarkFilter(b *testing.B) {
	count := 10000

	b.Run("Filter", func(b *testing.B) {
		s := Filter(Range(0, count), func(i int) bool {
			return i%2 == 0
		})

		var r []int
		for i := 0; i < b.N; i++ {
			r = slices.Collect(s)
		}

		if len(r) < count/2 {
			b.Fatalf("expected result length %d, got %d", count/2, len(r))
		}
	})

	b.Run("PlainFilter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r := make([]int, 0, count/2)
			for j := 0; j < count; j++ {
				if j%2 == 0 {
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
	count := 10000

	b.Run("Map", func(b *testing.B) {
		s := Map(Range(0, count), func(i int) int {
			return i * 2
		})

		for i := 0; i < b.N; i++ {
			r := slices.Collect(s)

			if len(r) < count {
				b.Fatalf("expected result length %d, got %d", count, len(r))
			}
		}
	})

	b.Run("PlainMap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r := make([]int, 0, count)
			for j := 0; j < count; j++ {
				r = append(r, j*2)
			}
			if len(r) < count {
				b.Fatalf("expected result length %d, got %d", count, len(r))
			}
		}
	})
}

func BenchmarkChannel(b *testing.B) {
	count := 10000

	b.Run("Channel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ch := make(chan int)
			go func() {
				for i := 0; i < count; i++ {
					ch <- i
				}
				close(ch)
			}()

			r := slices.Collect(Channel(ch))
			if len(r) < count {
				b.Fatalf("expected result length %d, got %d", count, len(r))
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
