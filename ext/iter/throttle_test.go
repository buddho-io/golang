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
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	// Given
	seq := Range(1, 4)

	// When
	s := Throttle(seq, 100*time.Millisecond)

	// Then
	start := time.Now()
	c := Len(s)
	elapsed := time.Since(start)

	if c != 3 {
		t.Errorf("Expected 3 but got %d", c)
	}

	if elapsed < 300*time.Millisecond {
		t.Errorf("Expected at least 300ms but got %v", elapsed)
	}
}

func TestThrottle2(t *testing.T) {
	// Given
	seq := Range2(1, 4)

	// When
	s := Throttle2(seq, 100*time.Millisecond)

	// Then
	start := time.Now()
	c := Len2(s)
	elapsed := time.Since(start)

	if c != 3 {
		t.Errorf("Expected 3 but got %d", c)
	}

	if elapsed < 300*time.Millisecond {
		t.Errorf("Expected at least 300ms but got %v", elapsed)
	}
}
