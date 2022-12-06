// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

// Package goerrors provides a few simple functions for handling errors in Go.
package goerrors

import (
	"fmt"
	"os"
)

// OrPanic panics with e if e is not nil.
func OrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// OrPanic1 is like OrPanic, but returns t when e is nil.
func OrPanic1[T any](t T, e error) T {
	OrPanic(e)
	return t
}

// OrPanic2 is like OrPanic, but returns (t1, t2) when e is nil.
func OrPanic2[T1, T2 any](t1 T1, t2 T2, e error) (T1, T2) {
	OrPanic(e)
	return t1, t2
}

// If e is not nil, OrExit prints e to stderr and terminates the process.
func OrExit(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

// OrExit1 is like OrExit, but returns t when e is nil.
func OrExit1[T any](t T, e error) T {
	OrExit(e)
	return t
}

// OrExit2 is like OrExit, but returns (t1, t2) when e is nil.
func OrExit2[T1, T2 any](t1 T1, t2 T2, e error) (T1, T2) {
	OrExit(e)
	return t1, t2
}
