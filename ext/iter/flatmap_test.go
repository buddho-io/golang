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
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlatMap(t *testing.T) {
	s := slices.Values([]int{1, 2, 3, 4, 5})

	fm := FlatMap(s, func(a int) iter.Seq[string] {
		return slices.Values([]string{fmt.Sprintf("%d_a", a), fmt.Sprintf("%d_b", a)})
	})

	m := Flatten(fm)

	var got []string //nolint:prealloc
	for x := range m {
		got = append(got, x)
	}

	want := []string{"1_a", "1_b", "2_a", "2_b", "3_a", "3_b", "4_a", "4_b", "5_a", "5_b"}

	require.Equal(t, want, got)
}

func TestFlatMapConcat(t *testing.T) {
	s := slices.Values([]int{1, 2, 3, 4, 5})

	got := slices.Collect(
		FlatMapConcat(s, func(a int) iter.Seq[string] {
			return slices.Values([]string{fmt.Sprintf("%d_a", a), fmt.Sprintf("%d_b", a)})
		}),
	)

	want := []string{"1_a", "1_b", "2_a", "2_b", "3_a", "3_b", "4_a", "4_b", "5_a", "5_b"}

	require.Equal(t, want, got)
}
