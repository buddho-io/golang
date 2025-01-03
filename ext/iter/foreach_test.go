package iter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestForEach(t *testing.T) {
	seq := Range(0, 10)
	r := make([]int, 0, 10)
	ForEach(seq, func(i int) {
		r = append(r, i)
	})

	require.Len(t, r, 10)
}
