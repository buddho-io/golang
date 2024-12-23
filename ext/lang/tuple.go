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

package lang

// Tuple2 is a tuple of two elements.
type Tuple2[A any, B any] interface {
	// A returns the first element of the tuple.
	A() A
	// B returns the second element of the tuple.
	B() B
}

// Tuple3 is a tuple of three elements.
type Tuple3[A any, B any, C any] interface {
	// A returns the first element of the tuple.
	A() A
	// B returns the second element of the tuple.
	B() B
	// C returns the third element of the tuple.
	C() C
}
