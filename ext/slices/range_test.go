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

package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRange(t *testing.T) {
	got := Range(0, 5)
	want := []int{0, 1, 2, 3, 4}

	require.Equal(t, want, got)
}
