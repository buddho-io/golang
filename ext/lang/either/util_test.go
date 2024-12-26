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

package either

import (
	"errors"
	"testing"

	"github.com/buddho-io/golang/ext/lang"
	"github.com/buddho-io/golang/ext/lang/option"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	r := Right[error, int](42)
	l := Left[error, int](error(nil))

	f := func(a int) int {
		return a + 1
	}

	require.Equal(t, Right[error, int](43), Map(r, f))
	require.Equal(t, l, Map(l, f))
}

func TestMapLeft(t *testing.T) {
	r := Right[error, int](42)
	l := Left[error, int](error(nil))

	f := func(_ error) error {
		return errors.New("error")
	}

	require.Equal(t, r, MapLeft(r, f))
	require.Equal(t, Left[error, int](errors.New("error")), MapLeft(l, f))
}

func TestFlatMap(t *testing.T) {
	r := Right[error, int](42)
	l := Left[error, int](error(nil))

	f := func(a int) lang.Either[error, int] {
		return Right[error, int](a + 1)
	}

	require.Equal(t, Right[error, int](43), FlatMap(r, f))
	require.Equal(t, l, FlatMap(l, f))
}

func TestFlatMapLeft(t *testing.T) {
	r := Right[error, int](42)
	l := Left[error, int](error(nil))

	f := func(_ error) lang.Either[error, int] {
		return Left[error, int](errors.New("error"))
	}

	require.Equal(t, r, FlatMapLeft(r, f))
	require.Equal(t, Left[error, int](errors.New("error")), FlatMapLeft(l, f))
}

func TestOrElse(t *testing.T) {
	r := Right[error, int](42)
	l := Left[error, int](error(nil))

	f := func() lang.Either[error, int] {
		return Right[error, int](43)
	}

	require.Equal(t, 42, OrElse(r, f).Right())
	require.Equal(t, 43, OrElse(l, f).Right())
}

func TestFilter(t *testing.T) {
	r := Right[error, int](42)
	l := Left[error, int](error(nil))

	f := func(a int) bool {
		return a == 42
	}

	require.Equal(t, r, Filter(r, f))
	require.Equal(t, l, Filter(l, f))
}

func TestForEach(t *testing.T) {
	r := Right[error, int](42)
	l := Left[error, int](error(nil))

	var result int
	f := func(a int) {
		result = a
	}

	ForEach(r, f)
	require.Equal(t, 42, result)

	ForEach(l, f)
	require.Equal(t, 42, result)
}

func TestFlatten(t *testing.T) {
	r := Right[error, lang.Either[error, int]](Right[error, int](42))
	require.Equal(t, Right[error, int](42), Flatten(r))
}

func TestFromOption(t *testing.T) {
	o := option.Some(42)
	require.Equal(t, Right[error, int](42), FromOption[error, int](o, nil))

	n := option.None[int]()
	err := error(nil)
	require.Equal(t, Left[error, int](err), FromOption[error, int](n, err))
}

func TestSequence(t *testing.T) {
	e := []lang.Either[error, int]{
		Right[error, int](1),
		Right[error, int](2),
		Right[error, int](3),
		Left[error, int](errors.New("test")),
		Right[error, int](4),
	}

	r := Sequence(e)

	expectL := Left[error, []int](errors.New("test"))

	require.Equal(t, expectL, r)

	e = []lang.Either[error, int]{
		Right[error, int](1),
		Right[error, int](2),
		Right[error, int](3),
	}

	r = Sequence(e)

	expectR := []int{1, 2, 3}

	require.True(t, r.IsRight())
	require.Equal(t, expectR, r.Right())
}
