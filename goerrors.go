// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

// Package goerrors provides a few simple functions for handling errors in Go.
package goerrors

import (
	"fmt"
	"os"
)

// NilOrPanic panics with e if e is not nil.
func NilOrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// NilOrPanic1 is like NilOrPanic, but returns t when e is nil.
func NilOrPanic1[T any](t T, e error) T {
	NilOrPanic(e)
	return t
}

// NilOrPanic2 is like NilOrPanic, but returns (t1, t2) when e is nil.
func NilOrPanic2[T1, T2 any](t1 T1, t2 T2, e error) (T1, T2) {
	NilOrPanic(e)
	return t1, t2
}

// If e is not nil, NilOrExit prints e to stderr and terminates the process.
func NilOrExit(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

// NilOrExit1 is like NilOrExit, but returns t when e is nil.
func NilOrExit1[T any](t T, e error) T {
	NilOrExit(e)
	return t
}

// NilOrExit2 is like NilOrExit, but returns (t1, t2) when e is nil.
func NilOrExit2[T1, T2 any](t1 T1, t2 T2, e error) (T1, T2) {
	NilOrExit(e)
	return t1, t2
}
