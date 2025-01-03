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
	"io"
	"slices"
	"testing"

	"github.com/buddho-io/golang/ext/lang"
	"github.com/buddho-io/golang/ext/lang/either"
	"github.com/stretchr/testify/require"
)

func TestStream(t *testing.T) {
	s := Stream(NewMockStream[int](1, 2, 3))

	got := slices.Collect(s)

	want := []int{1, 2, 3}

	require.Equal(t, want, either.Sequence(got).Right())
}

type mockStream[T any] struct {
	index  int
	values []T
	err    error
}

func NewMockStream[T any](values ...T) StreamLike[T] {
	m := &mockStream[T]{
		values: values,
	}

	return m
}

func (m *mockStream[T]) Recv() (T, error) {
	if m.err != nil {
		var r T
		return r, m.err
	}

	if m.index >= len(m.values) {
		return lang.Zero[T](), io.EOF
	}

	v := m.values[m.index]
	m.index++

	return v, nil
}

var _ StreamLike[int] = &mockStream[int]{}
