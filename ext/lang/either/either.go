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
	"github.com/buddho-io/golang/ext/lang"
)

// left is an implementation of Either that represents the left value.
type left[L, R any] struct {
	left L
}

// IsLeft implements Either.IsLeft and returns true.
func (l left[L, R]) IsLeft() bool {
	return true
}

// Left implements Either.Left and returns the left value.
func (l left[L, R]) Left() L {
	return l.left
}

// IsRight implements Either.IsRight and returns false.
func (l left[L, R]) IsRight() bool {
	return false
}

// Right implements Either.Right and returns the zero value of the right type.
func (l left[L, R]) Right() R {
	return lang.Zero[R]()
}

// Left returns an Either instance with the given left value.
func Left[L, R any](le L) lang.Either[L, R] {
	return left[L, R]{le}
}

// right is an implementation of Either that represents the right value.
type right[L, R any] struct {
	right R
}

// IsLeft implements Either.IsLeft and returns false.
func (r right[L, R]) IsLeft() bool {
	return false
}

// Left implements Either.Left and returns the zero value of the left type.
func (r right[L, R]) Left() L {
	return lang.Zero[L]()
}

// IsRight implements Either.IsRight and returns true.
func (r right[L, R]) IsRight() bool {
	return true
}

// Right implements Either.Right and returns the right value.
func (r right[L, R]) Right() R {
	return r.right
}

// Right returns an Either instance with the given right value.
func Right[L, R any](ri R) lang.Either[L, R] {
	return right[L, R]{ri}
}
