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

func TestSingle(t *testing.T) {
	s := Single(1)

	got := slices.Collect(s)

	require.Len(t, got, 1)
	require.Equal(t, 1, got[0])
}

func TestSingle2(t *testing.T) {
	s := Single2(1, "one")

	got := maps.Collect(s)
	require.Len(t, got, 1)
	require.Equal(t, map[int]string{1: "one"}, got)
}
