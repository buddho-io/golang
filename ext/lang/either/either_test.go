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

package either

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRight(t *testing.T) {
	r := Right[error, int](42)

	require.True(t, r.IsRight())
	require.Equal(t, 42, r.Right())
	require.False(t, r.IsLeft())
	require.Nil(t, r.Left())
}

func TestLeft(t *testing.T) {
	err := errors.New("error")
	l := Left[error, int](err)

	require.True(t, l.IsLeft())
	require.Equal(t, err, l.Left())
	require.False(t, l.IsRight())
	require.Equal(t, 0, l.Right())
}
