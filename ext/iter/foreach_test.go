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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForEach(t *testing.T) {
	seq := Range(0, 10)
	r := make([]int, 0, 10)
	ForEach(seq, func(i int) {
		r = append(r, i)
	})

	require.Len(t, r, 10)
}
