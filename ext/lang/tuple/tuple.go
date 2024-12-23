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

package tuple

import (
	"github.com/llinder/golang/ext/lang"
)

// tuple2 implements lang.Tuple2.
type tuple2[A any, B any] struct {
	a A
	b B
}

// Two creates a new lang.Tuple2 instance with the given elements.
func Two[A, B any](a A, b B) lang.Tuple2[A, B] {
	return tuple2[A, B]{a, b}
}

// A returns the first element of the tuple.
func (t tuple2[A, B]) A() A {
	return t.a
}

// B returns the second element of the tuple.
func (t tuple2[A, B]) B() B {
	return t.b
}

var _ lang.Tuple2[int, string] = tuple2[int, string]{}

// tuple3 implements lang.Tuple3.
type tuple3[A any, B any, C any] struct {
	a A
	b B
	c C
}

func (t tuple3[A, B, C]) A() A {
	return t.a
}

func (t tuple3[A, B, C]) B() B {
	return t.b
}

func (t tuple3[A, B, C]) C() C {
	return t.c
}

// Three creates a new lang.Tuple3 instance with the given elements.
func Three[A, B, C any](a A, b B, c C) lang.Tuple3[A, B, C] {
	return tuple3[A, B, C]{a, b, c}
}
