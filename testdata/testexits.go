// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

// This is a simple program to test the case in which the OrExitX functions actually exit.
package main

import (
	"errors"
	"os"

	"github.com/pat42smith/goerrors"
)

func main() {
	switch os.Args[1] {
	case "0":
		goerrors.OrExit(errors.New("zero"))
	case "1":
		goerrors.OrExit1(100, errors.New("one"))
	case "2":
		goerrors.OrExit2(100, 200, errors.New("two"))
	}
}
