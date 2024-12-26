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
	"io"
	"iter"

	"github.com/buddho-io/golang/ext/lang"
	"github.com/buddho-io/golang/ext/lang/either"
)

// Stream converts a StreamLike into an iter.Seq that can be used with the iter package.
// If the StreamLike returns an error, the iter.Seq will return an either.Left with that error.
// Otherwise, the iter.Seq will return an either.Right with the value.
func Stream[T any](stream StreamLike[T]) iter.Seq[lang.Either[error, T]] {
	return func(f func(lang.Either[error, T]) bool) {
		for {
			v, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return
			}
			if err != nil {
				if !f(either.Left[error, T](err)) {
					return
				}

				continue
			}

			if !f(either.Right[error, T](v)) {
				return
			}
		}
	}
}

// Stream2 converts a StreamLike into an iter.Seq2 that can be used with the iter package.
// If the StreamLike returns an error, the iter.Seq2 will return that error.
func Stream2[T any](stream StreamLike[T]) iter.Seq2[T, error] {
	return func(f func(T, error) bool) {
		for {
			v, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return
			}
			if err != nil {
				if !f(v, err) {
					return
				}

				continue
			}

			if !f(v, nil) {
				return
			}
		}
	}
}

// StreamLike is an interface that matches the gRPC Stream interface for receiving messages.
// This is existing to avoid dependencies on gRPC yet still be able to implement an Iterator
// that can be used with gRPC.
type StreamLike[T any] interface {
	Recv() (T, error)
}
