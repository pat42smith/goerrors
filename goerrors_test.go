// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package goerrors

import (
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"strings"
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

func TestNilOrExit(t *testing.T) {
	NilOrExit(nil)
}

func TestNilOrExit1(t *testing.T) {
	var x uint16 = 331
	gotest.Expect(t, x, NilOrExit1(x, nil))
}

func TestNilOrExit2(t *testing.T) {
	x := true
	y := "never"
	u, v := NilOrPanic2(x, y, nil)
	gotest.Expect(t, x, u)
	gotest.Expect(t, y, v)
}

func TestExits(t *testing.T) {
	tmp := t.TempDir()
	testexits := filepath.Join(tmp, "testexits.exe")
	gocmd := NilOrPanic1(exec.LookPath("go"))
	cmd := exec.Command(gocmd, "build", "-o", testexits, "testdata/testexits.go")
	out, e := cmd.CombinedOutput()
	if e != nil || len(out) > 0 {
		t.Log(string(out))
		t.Fatal(e)
	}

	for _, s := range []string{"0 zero\n", "1 one\n", "2 two\n"} {
		arg, result, _ := strings.Cut(s, " ")
		cmd = exec.Command(testexits, arg)
		cmd.Stdin = nil
		out := NilOrPanic1(cmd.StdoutPipe())
		err := NilOrPanic1(cmd.StderrPipe())
		NilOrPanic(cmd.Start())
		gotest.Expect(t, 0, len(NilOrPanic1(io.ReadAll(out))))
		gotest.Expect(t, result, string(NilOrPanic1(io.ReadAll(err))))
		switch e := cmd.Wait().(type) {
		case *exec.ExitError:
			gotest.Expect(t, 1, e.ExitCode())
		default:
			t.Fatal(e)
		}
	}
}
