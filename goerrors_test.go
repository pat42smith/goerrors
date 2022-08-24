// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package goerrors

import (
	"fmt"
	"testing"

	"github.com/pat42smith/gotest"
)

func TestNilOrPanic(t *testing.T) {
	NilOrPanic(nil)

	e := gotest.MustPanic(t, func() {
		NilOrPanic(fmt.Errorf("oops!"))
	})
	gotest.Expect(t, "oops!", e.(error).Error())
}

func TestNilOrPanic1(t *testing.T) {
	var x byte = 99
	gotest.Expect(t, x, NilOrPanic1(x, nil))

	e := gotest.MustPanic(t, func() {
		NilOrPanic1(x, fmt.Errorf("broken"))
	})
	gotest.Expect(t, "broken", e.(error).Error())

	foo := func() (byte, error) {
		return x, nil
	}
	gotest.Expect(t, x, NilOrPanic1(foo()))
}

func TestNilOrPanic2(t *testing.T) {
	x := "seventeen"
	y := 3.7
	u, v := NilOrPanic2(x, y, nil)
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)

	e := gotest.MustPanic(t, func() {
		NilOrPanic2(x, y, fmt.Errorf("sabotaged"))
	})
	gotest.Expect(t, "sabotaged", e.(error).Error())

	foo := func() (string, float64, error) {
		return x, y, nil
	}
	u, v = NilOrPanic2(foo())
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)
}
