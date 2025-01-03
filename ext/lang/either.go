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

package lang

// Either is an interface that represents a value of one of two possible types.
type Either[L, R any] interface {
	// IsLeft returns true if the Either instance is a Left.
	IsLeft() bool
	// Left returns the left value.
	Left() L
	// IsRight returns true if the Either instance is a Right.
	IsRight() bool
	// Right returns the right value.
	Right() R
}
