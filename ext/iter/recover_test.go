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
	"errors"
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/llinder/golang/ext/lang"
	"github.com/llinder/golang/ext/lang/either"
)

func TestRecoverIf(t *testing.T) {
	s := slices.Values([]lang.Either[error, int]{
		either.Right[error, int](0),
		either.Right[error, int](1),
		either.Left[error, int](errors.New("error")),
	})

	rec := func(_ error) iter.Seq[lang.Either[error, int]] {
		return slices.Values([]lang.Either[error, int]{
			either.Right[error, int](2),
			either.Right[error, int](3),
		})
	}

	r := func(e error) bool {
		return e.Error() == "error"
	}

	var got []lang.Either[error, int] //nolint:prealloc
	m := RecoverIf(s, r, rec)
	for x := range m {
		got = append(got, x)
	}

	result := either.Sequence(got)

	want := either.Right[error, []int]([]int{0, 1, 2, 3})

	require.Equal(t, want, result)
}
