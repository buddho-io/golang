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
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := slices.Values([]int{1, 2, 3, 4, 5})

	got := slices.Collect(Context(ctx, s))

	want := []int{1, 2, 3, 4, 5}

	require.Equal(t, want, got)
}

func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := slices.Values([]int{1, 2, 3, 4, 5})

	var got []int //nolint:prealloc
	m := Context(ctx, s)

	i := 0
	for x := range m {
		got = append(got, x)
		i++
		if i == 3 {
			cancel()
		}
	}

	want := []int{1, 2, 3}

	require.Equal(t, want, got)
}

func TestContext2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := maps.All(map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"})

	got := maps.Collect(Context2(ctx, s))

	want := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}

	require.Equal(t, want, got)
}

func TestCancelContext2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := maps.All(map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"})

	got := map[int]string{}
	m := Context2(ctx, s)

	i := 0
	for k, v := range m {
		got[k] = v
		i++
		if i == 3 {
			cancel()
		}
	}

	require.Len(t, got, 3)
}
