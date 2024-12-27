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

	"github.com/stretchr/testify/require"
)

func TestConcat(t *testing.T) {
	// Given
	seq1 := Range(1, 4)
	seq2 := Range(4, 7)
	seq3 := Range(7, 10)

	// When
	s := Concat(seq1, seq2, seq3)

	// Then
	r := slices.Collect(s)

	require.Len(t, r, 9)
}

func TestConcat2(t *testing.T) {
	// Given
	seq1 := Range2(1, 4)
	seq2 := Range2(4, 7)
	seq3 := Range2(7, 10)

	// When
	s := Concat2(seq1, seq2, seq3)

	// Then
	r := maps.Collect(s)

	require.Len(t, r, 9)
}
