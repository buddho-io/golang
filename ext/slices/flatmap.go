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

func FlatMap[T, U any](s []T, f func(T) []U) [][]U {
	r := make([][]U, 0, len(s))
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func FlatMapConcat[T, U any](s []T, f func(T) []U) []U {
	return Flatten(FlatMap(s, f))
}
