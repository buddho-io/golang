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

	"github.com/buddho-io/golang/ext/lang"
	"github.com/buddho-io/golang/ext/lang/either"
)

func TestMap(t *testing.T) {
	o := Some[int](42)

	f := func(a int) int {
		return a + 1
	}

	require.Equal(t, Some[int](43), Map(o, f))

	n := None[int]()
	require.Equal(t, n, Map(n, f))
}

func TestFlatMap(t *testing.T) {
	o := Some[int](42)

	f := func(a int) lang.Option[int] {
		return Some[int](a + 1)
	}

	require.Equal(t, Some[int](43), FlatMap(o, f))

	n := None[int]()
	require.Equal(t, n, FlatMap(n, f))
}

func TestFilter(t *testing.T) {
	o := Some[int](42)

	f := func(a int) bool {
		return a == 42
	}

	require.Equal(t, o, Filter(o, f))

	n := None[int]()
	require.Equal(t, n, Filter(n, f))

	f = func(a int) bool {
		return a != 42
	}

	require.Equal(t, None[int](), Filter(o, f))
}

func TestForEach(t *testing.T) {
	o := Some[int](42)

	ForEach(o, func(a int) {
		require.Equal(t, 42, a)
	})
}

func TestFlatten(t *testing.T) {
	o := Some[lang.Option[int]](Some[int](42))

	require.Equal(t, Some[int](42), Flatten[int](o))

	n := None[lang.Option[int]]()
	require.Equal(t, None[int](), Flatten[int](n))
}

func TestFromEither(t *testing.T) {
	r := either.Right[error, int](42)
	l := either.Left[error, int](error(nil))

	require.Equal(t, Some[int](42), FromEither(r))
	require.Equal(t, None[int](), FromEither(l))
}

func TestOrElse(t *testing.T) {
	o := Some[int](42)

	f := func() lang.Option[int] {
		return Some[int](43)
	}

	require.Equal(t, Some[int](42), OrElse(o, f))

	n := None[int]()
	require.Equal(t, Some[int](43), OrElse(n, f))
}

func TestSequence(t *testing.T) {
	n := []lang.Option[int]{
		Some(1),
		Some(2),
		Some(3),
		None[int](),
		Some(4),
	}

	no := Sequence(n)

	expect := None[[]int]()

	require.True(t, no.IsEmpty())
	require.Equal(t, expect, no)

	s := []lang.Option[int]{
		Some(1),
		Some(2),
		Some(3),
	}

	no = Sequence(s)

	expectS := []int{1, 2, 3}

	require.True(t, no.IsDefined())
	require.Equal(t, expectS, no.Get())
}
