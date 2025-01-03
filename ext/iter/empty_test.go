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

func TestEmpty(t *testing.T) {
	s := Empty[int]()

	result := slices.Collect(s)

	require.Len(t, result, 0)
}

func TestEmpty2(t *testing.T) {
	s := Empty2[int, string]()

	got := maps.Collect(s)
	require.Len(t, got, 0)
	require.Equal(t, map[int]string{}, got)
}
