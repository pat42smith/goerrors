// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

// Package goerrors provides a few simple functions for handling errors in Go.
package goerrors

import (
	"fmt"
	"os"

	"github.com/pat42smith/gotest"
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

// OrError(e)(r) calls r.Error(e) if e is not nil.
//
// This is intended for use in test cases; r will usually be a *testing.T.
func OrError(e error) func(r gotest.Reporter) {
	if e == nil {
		return func(gotest.Reporter) {}
	} else {
		return func(r gotest.Reporter) {
			r.Helper()
			r.Error(e)
		}
	}
}

// OrError1(t, e)(r) is like OrError(e)(r), but returns t.
func OrError1[T any](t T, e error) func(gotest.Reporter) T {
	return func(r gotest.Reporter) T {
		if e != nil {
			r.Helper()
			r.Error(e)
		}
		return t
	}
}

// OrError2(t1, t2, e)(r) is like OrError(e)(r), but returns (t1, t2).
func OrError2[T1, T2 any](t1 T1, t2 T2, e error) func(gotest.Reporter) (T1, T2) {
	return func(r gotest.Reporter) (T1, T2) {
		if e != nil {
			r.Helper()
			r.Error(e)
		}
		return t1, t2
	}
}

// OrFatal(e)(r) calls r.Fatal(e) if e is not nil.
//
// This is intended for use in test cases; r will usually be a *testing.T,
// in which case OrFatal(e)(r) will not return when e != nil.
func OrFatal(e error) func(r gotest.Reporter) {
	if e == nil {
		return func(gotest.Reporter) {}
	} else {
		return func(r gotest.Reporter) {
			r.Helper()
			r.Fatal(e)
		}
	}
}

// OrFatal1(t, e)(r) is like OrFatal(e)(r), but returns t when e != nil or r.Fatal() returns.
func OrFatal1[T any](t T, e error) func(gotest.Reporter) T {
	return func(r gotest.Reporter) T {
		if e != nil {
			r.Helper()
			r.Fatal(e)
		}
		return t
	}
}

// OrFatal2(t1, t2, e)(r) is like OrFatal(e)(r), but returns (t1, t2) when e != nil or r.Fatal() returns.
func OrFatal2[T1, T2 any](t1 T1, t2 T2, e error) func(gotest.Reporter) (T1, T2) {
	return func(r gotest.Reporter) (T1, T2) {
		if e != nil {
			r.Helper()
			r.Fatal(e)
		}
		return t1, t2
	}
}
