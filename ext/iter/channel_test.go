// Copyright 2024 BuddhoIO
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

	"github.com/buddho-io/golang/ext/lang"
	"github.com/buddho-io/golang/ext/lang/tuple"
	"github.com/stretchr/testify/require"
)

func TestChannel(t *testing.T) {
	v := []int{1, 2, 3}
	ch := make(chan int)

	go func() {
		for _, x := range v {
			ch <- x
		}
		close(ch)
	}()

	got := slices.Collect(Channel(ch))

	want := []int{1, 2, 3}

	require.Equal(t, want, got)
}

func TestCloseChannel(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	ch := make(chan int)

	go func() {
		for _, x := range v {
			if x > 2 {
				close(ch)
				return
			}
			ch <- x
		}
	}()

	got := slices.Collect(Channel(ch))

	want := []int{1, 2}

	require.Equal(t, want, got)
}

func TestChannel2(t *testing.T) {
	v := map[int]string{1: "one", 2: "two", 3: "three"}
	ch := make(chan lang.Tuple2[int, string])

	go func() {
		for k, v := range v {
			ch <- tuple.Two(k, v)
		}
		close(ch)
	}()

	got := maps.Collect(Channel2(ch))

	want := map[int]string{1: "one", 2: "two", 3: "three"}
	require.Equal(t, want, got)
}

func TestCloseChannel2(t *testing.T) {
	v := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
	ch := make(chan lang.Tuple2[int, string], 1)

	go func() {
		count := 0
		for k, v := range v {
			if count > 2 {
				close(ch)
				return
			}
			ch <- tuple.Two(k, v)
			count++
		}
	}()

	got := maps.Collect(Channel2(ch))

	require.Len(t, got, 3)
}
