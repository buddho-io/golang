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
	"strconv"
	"testing"

	"github.com/buddho-io/golang/ext/lang"
	"github.com/buddho-io/golang/ext/lang/tuple"
	"github.com/stretchr/testify/require"
)

func TestDeferred(t *testing.T) {
	count := 0
	f := func() (int, bool) {
		if count < 3 {
			count++
			return count, true
		}
		return 0, false
	}

	got := slices.Collect(Deferred(f))

	want := []int{1, 2, 3}

	require.Equal(t, want, got)
}

func TestDeferred2(t *testing.T) {
	count := 0
	f := func() (lang.Tuple2[int, string], bool) {
		if count < 3 {
			count++
			return tuple.Two(count, strconv.Itoa(count)), true
		}
		return tuple.Two(0, ""), false
	}

	got := maps.Collect(Deferred2(f))

	want := map[int]string{1: "1", 2: "2", 3: "3"}

	require.Equal(t, want, got)
}
