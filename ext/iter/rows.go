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
	"io"
	"iter"
)

// Rows returns a sequence of values from a database RowsLike instance. If an error occurs reading the row
// it is returned as a Left value in the sequence. If the row is read successfully, the value is returned as a Right value.
func Rows[T any](rows RowLike, scanner Scanner[T]) iter.Seq[lang.Either[error, T]] {
	return func(f func(lang.Either[error, T]) bool) {
		defer rows.Close()

		for rows.Next() {
			if err := rows.Err(); err != nil {
				if errors.Is(err, io.EOF) {
					return
				}

				if !f(either.Left[error, T](err)) {
					return
				}
			}

			v, err := scanner(rows.Scan)
			if err != nil {
				if !f(either.Left[error, T](err)) {
					return
				}
			} else {
				if !f(either.Right[error, T](v)) {
					return
				}
			}
		}
	}
}

// Rows2 returns a sequence of values from a database RowsLike instance. If an error occurs reading the row
// it is returned as the second yield argument. If the row is read successfully, the value is returned as the first yield argument.
func Rows2[T any](rows RowLike, scanner Scanner[T]) iter.Seq2[T, error] {
	return func(f func(T, error) bool) {
		defer rows.Close()

		for rows.Next() {
			if err := rows.Err(); err != nil {
				if errors.Is(err, io.EOF) {
					return
				}

				if !f(lang.Zero[T](), err) {
					return
				}
			}

			v, err := scanner(rows.Scan)
			if err != nil {
				if !f(lang.Zero[T](), err) {
					return
				}
			} else {
				if !f(v, nil) {
					return
				}
			}
		}
	}
}

// Scanner is a function type for scanning values from database sql.Rows.
type Scanner[T any] func(func(dest ...any) error) (T, error)

// RowLike is an interface for sql.Rows. It is used to allow for easier testing.
type RowLike interface {
	// Next returns true if there are more rows to iterate over.
	Next() bool
	// Close closes the sql.Rows.
	Close()
	// Err returns the error encountered by the sql.Rows.
	Err() error
	// Scan scans the values from the sql.Rows into the given destination.
	Scan(dest ...any) error
}
