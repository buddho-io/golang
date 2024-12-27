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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupBy(t *testing.T) {
	// Given
	seq := Range(1, 10)

	// When
	s := GroupBy(seq, func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})

	// Then
	r := maps.Collect(s)

	require.Len(t, r, 2)
	require.Equal(t, 4, Len(r["even"]))
	require.Equal(t, 5, Len(r["odd"]))
}
