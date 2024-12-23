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

package option

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNone(t *testing.T) {
	o := None[int]()

	require.True(t, o.IsEmpty())
	require.False(t, o.IsDefined())
	require.Equal(t, 0, o.Get())
}

func TestSome(t *testing.T) {
	o := Some[int](42)

	require.False(t, o.IsEmpty())
	require.True(t, o.IsDefined())
	require.Equal(t, 42, o.Get())
}

func TestOf(t *testing.T) {
	var i []int

	t.Log("check empty slice")
	o := Of(i)
	require.True(t, o.IsEmpty())
	require.False(t, o.IsDefined())

	t.Log("check empty struct")
	empty := Of(testOption{})
	require.True(t, empty.IsEmpty())
	require.False(t, empty.IsDefined())

	t.Log("check non empty struct")
	nonEmpty := Of(testOption{test: 2})
	require.False(t, nonEmpty.IsEmpty())
	require.True(t, nonEmpty.IsDefined())

	t.Log("check nil pointer")
	var ptr *testOption
	nilPtr := Of(ptr)
	require.True(t, nilPtr.IsEmpty())
	require.False(t, nilPtr.IsDefined())

	t.Log("check non nil pointer")
	ptr = &testOption{}
	nonNilPtr := Of(ptr)
	require.False(t, nonNilPtr.IsEmpty())
	require.True(t, nonNilPtr.IsDefined())
}

type testOption struct {
	test int
}
