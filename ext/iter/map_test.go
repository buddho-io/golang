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
	"maps"
	"slices"
	"strconv"
	"testing"

	"github.com/buddho-io/golang/ext/lang/option"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	s := Range(0, 10)

	r := Map(s, func(i int) int {
		return i * 2
	})

	result := slices.Collect(r)

	require.Equal(t, []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}, result)
}

func TestMap2(t *testing.T) {
	m := map[int]error{
		0: nil,
		1: nil,
		2: nil,
		3: errors.New("error"),
		4: nil,
	}

	s := maps.All(m)

	r := Map2[int, error, string, string](s, func(k int, v error) (string, string) {
		ks := strconv.Itoa(k)
		vs := option.Of(v)

		return ks, option.Map(vs, func(v error) string { return v.Error() }).GetOrElse("")
	})

	result := maps.Collect(r)

	expect := map[string]string{
		"0": "",
		"1": "",
		"2": "",
		"3": "error",
		"4": "",
	}

	require.Equal(t, expect, result)
}
