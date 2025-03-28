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
	"errors"
	"github.com/buddho-io/golang/ext/lang"
	"github.com/buddho-io/golang/ext/lang/either"
	"github.com/stretchr/testify/require"
	"io"
	"maps"
	"slices"
	"sync"
	"testing"
)

func TestEmptyRows(t *testing.T) {
	empty := Rows[int](&mockRows[int]{}, testScanner)
	result := slices.Collect(empty)
	require.Empty(t, result)
}

func TestSingleRows(t *testing.T) {
	single := Rows[int](&mockRows[int]{values: []int{42}}, testScanner)
	result := slices.Collect(single)
	require.Equal(t, []lang.Either[error, int]{either.Right[error, int](42)}, result)
}

func TestRowsScanError(t *testing.T) {
	single := Rows[int](&mockRows[int]{values: []int{42}}, func(f func(dest ...any) error) (int, error) {
		return 0, errors.New("scan error")
	})
	var result lang.Either[error, int]
	single(func(e lang.Either[error, int]) bool {
		result = e
		return false
	})

	require.Equal(t, either.Left[error, int](errors.New("scan error")), result)
}

func TestRowsClose(t *testing.T) {
	m := &mockRows[int]{values: []int{42}}
	single := Rows[int](m, testScanner)
	m.Close()

	result := slices.Collect(single)
	require.Empty(t, result)
	require.True(t, m.closed)
}

func TestEmptyRows2(t *testing.T) {
	empty := Rows2[int](&mockRows[int]{}, testScanner)
	result := maps.Collect(empty)
	require.Empty(t, result)
}

func TestSingleRows2(t *testing.T) {
	single := Rows2[int](&mockRows[int]{values: []int{42}}, testScanner)
	result := maps.Collect(single)
	require.Equal(t, map[int]error{42: nil}, result)
}

func TestRowsScanError2(t *testing.T) {
	single := Rows2[int](&mockRows[int]{values: []int{42}}, func(f func(dest ...any) error) (int, error) {
		return 0, errors.New("scan error")
	})
	var result map[int]error
	single(func(k int, v error) bool {
		result = map[int]error{k: v}
		return false
	})

	require.Equal(t, map[int]error{0: errors.New("scan error")}, result)
}

func TestRowsClose2(t *testing.T) {
	m := &mockRows[int]{values: []int{42}}
	single := Rows2[int](m, testScanner)
	m.Close()

	result := maps.Collect(single)
	require.Empty(t, result)
	require.True(t, m.closed)
}

type mockRows[T any] struct {
	values []T

	mu     sync.Mutex
	index  int
	closed bool
}

func (m *mockRows[T]) Next() bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed || m.index >= len(m.values) {
		return false
	}

	return true
}

func (m *mockRows[T]) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.closed = true
}

func (m *mockRows[T]) Err() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.index >= len(m.values) {
		return io.EOF
	}

	return nil
}

func (m *mockRows[T]) Scan(dest ...any) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.index >= len(m.values) {
		return io.EOF
	}

	d := dest[0]
	switch x := d.(type) {
	case *T:
		*x = m.values[m.index]
	}

	m.index++
	return nil
}

var _ RowLike = (*mockRows[int])(nil)

func testScanner(f func(dest ...any) error) (int, error) {
	var i int
	err := f(&i)
	return i, err
}
