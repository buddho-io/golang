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
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	m := Filter(slices.Values(v), func(x int) bool { return x%2 == 0 })

	got := slices.Collect(m)

	want := []int{2, 4}

	require.Equal(t, want, got)
}

func TestFilter2(t *testing.T) {
	v := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
	m := Filter2(maps.All(v), func(k int, _ string) bool { return k%2 == 0 })

	got := maps.Collect(m)

	want := map[int]string{2: "two", 4: "four"}

	require.Equal(t, want, got)
}
