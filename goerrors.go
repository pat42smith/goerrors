// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

// Package goerrors provides a few simple functions for handling errors in Go.
package goerrors

// NilOrPanic panics with e if e is not nil.
func NilOrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// NilOrPanic1 panics with e if e is not nil. Otherwise, it returns t.
func NilOrPanic1[T any](t T, e error) T {
	NilOrPanic(e)
	return t
}

// NilOrPanic2 panics with e if e is not nil. Otherwise, it returns (t1, t2).
func NilOrPanic2[T1, T2 any](t1 T1, t2 T2, e error) (T1, T2) {
	NilOrPanic(e)
	return t1, t2
}
