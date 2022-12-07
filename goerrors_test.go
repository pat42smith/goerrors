// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package goerrors

import (
	"errors"
	"io"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pat42smith/gotest"
)

func TestOrPanic(t *testing.T) {
	OrPanic(nil)

	e := gotest.MustPanic(t, func() {
		OrPanic(errors.New("oops!"))
	})
	gotest.Expect(t, "oops!", e.(error).Error())
}

func TestOrPanic1(t *testing.T) {
	var x byte = 99
	gotest.Expect(t, x, OrPanic1(x, nil))

	e := gotest.MustPanic(t, func() {
		OrPanic1(x, errors.New("broken"))
	})
	gotest.Expect(t, "broken", e.(error).Error())

	foo := func() (byte, error) {
		return x, nil
	}
	gotest.Expect(t, x, OrPanic1(foo()))
}

func TestOrPanic2(t *testing.T) {
	x := "seventeen"
	y := 3.7
	u, v := OrPanic2(x, y, nil)
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)

	e := gotest.MustPanic(t, func() {
		OrPanic2(x, y, errors.New("sabotaged"))
	})
	gotest.Expect(t, "sabotaged", e.(error).Error())

	foo := func() (string, float64, error) {
		return x, y, nil
	}
	u, v = OrPanic2(foo())
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)
}

func TestOrExit(t *testing.T) {
	OrExit(nil)
}

func TestOrExit1(t *testing.T) {
	var x uint16 = 331
	gotest.Expect(t, x, OrExit1(x, nil))
}

func TestOrExit2(t *testing.T) {
	x := true
	y := "never"
	u, v := OrPanic2(x, y, nil)
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)
}

func TestExits(t *testing.T) {
	tmp := t.TempDir()
	testexits := filepath.Join(tmp, "testexits.exe")
	cmd := exec.Command("go", "build", "-o", testexits, "testdata/testexits.go")
	out, e := cmd.CombinedOutput()
	if e != nil || len(out) > 0 {
		t.Log(string(out))
		t.Fatal(e)
	}

	for _, s := range []string{"0 zero\n", "1 one\n", "2 two\n"} {
		arg, result, _ := strings.Cut(s, " ")
		cmd = exec.Command(testexits, arg)
		cmd.Stdin = nil
		out := OrPanic1(cmd.StdoutPipe())
		err := OrPanic1(cmd.StderrPipe())
		OrPanic(cmd.Start())
		gotest.Expect(t, 0, len(OrPanic1(io.ReadAll(out))))
		gotest.Expect(t, result, string(OrPanic1(io.ReadAll(err))))
		switch e := cmd.Wait().(type) {
		case *exec.ExitError:
			gotest.Expect(t, 1, e.ExitCode())
		default:
			t.Fatal(e)
		}
	}
}

func TestOrError(t *testing.T) {
	var sr gotest.StubReporter
	OrError(nil)(&sr)
	sr.Expect(t, false, false, "", "")

	OrError(errors.New("problem"))(&sr)
	sr.Expect(t, true, false, "problem\n", "")
}

func TestOrError1(t *testing.T) {
	var sr gotest.StubReporter
	x := OrError1(99.0, nil)(&sr)
	sr.Expect(t, false, false, "", "")
	gotest.Expect(t, 99.0, x)

	c := OrError1(rune('日'), errors.New("disaster"))(&sr)
	sr.Expect(t, true, false, "disaster\n", "")
	gotest.Expect(t, rune('日'), c)

	var sr1 gotest.StubReporter
	foo := func() (string, error) {
		return "doughnuts", errors.New("donuts")
	}
	s := OrError1(foo())(&sr1)
	sr1.Expect(t, true, false, "donuts\n", "")
	gotest.Expect(t, "doughnuts", s)
}

func TestOrError2(t *testing.T) {
	x := "seventeen"
	y := 3.7
	var sr gotest.StubReporter
	u, v := OrError2(x, y, nil)(&sr)
	sr.Expect(t, false, false, "", "")
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)

	u, v = OrError2(x, y, errors.New("falling"))(&sr)
	sr.Expect(t, true, false, "falling\n", "")
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)

	var sr1 gotest.StubReporter
	foo := func() (string, float64, error) {
		return x, y, nil
	}
	u, v = OrError2(foo())(&sr1)
	sr1.Expect(t, false, false, "", "")
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)
}

func TestOrFatal(t *testing.T) {
	var sr gotest.StubReporter
	OrFatal(nil)(&sr)
	sr.Expect(t, false, false, "", "")

	OrFatal(errors.New("problem"))(&sr)
	sr.Expect(t, true, true, "problem\n", "")
}

func TestOrFatal1(t *testing.T) {
	var sr gotest.StubReporter
	x := OrFatal1(99.0, nil)(&sr)
	sr.Expect(t, false, false, "", "")
	gotest.Expect(t, 99.0, x)

	c := OrFatal1(rune('日'), errors.New("disaster"))(&sr)
	sr.Expect(t, true, true, "disaster\n", "")
	gotest.Expect(t, rune('日'), c)

	var sr1 gotest.StubReporter
	foo := func() (string, error) {
		return "doughnuts", errors.New("donuts")
	}
	s := OrFatal1(foo())(&sr1)
	sr1.Expect(t, true, true, "donuts\n", "")
	gotest.Expect(t, "doughnuts", s)
}

func TestOrFatal2(t *testing.T) {
	x := "seventeen"
	y := 3.7
	var sr gotest.StubReporter
	u, v := OrFatal2(x, y, nil)(&sr)
	sr.Expect(t, false, false, "", "")
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)

	u, v = OrFatal2(x, y, errors.New("falling"))(&sr)
	sr.Expect(t, true, true, "falling\n", "")
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)

	var sr1 gotest.StubReporter
	foo := func() (string, float64, error) {
		return x, y, nil
	}
	u, v = OrFatal2(foo())(&sr1)
	sr1.Expect(t, false, false, "", "")
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)
}
