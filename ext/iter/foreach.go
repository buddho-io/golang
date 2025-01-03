package iter

import "iter"

// ForEach applies the function f to each element of the sequence seq.
func ForEach[T any](seq iter.Seq[T], f func(T)) {
	for s := range seq {
		f(s)
	}
}
