package slag

import (
	"errors"
	"testing"
)

func isErr(err error, msg string) bool {
	if err != nil && err.Error() == msg {
		return true
	}
	return false
}

func checkRun(t *testing.T, fn interface{}, args ...string) {
	if err := Run(fn, args...); err != nil {
		t.Error("Run() failed:", err)
	}
}

func TestError(t *testing.T) {
	fn := func(a ...string) error {
		return errors.New("PLANNED ERROR")
	}
	if err := Run(fn); err == nil || err.Error() != "PLANNED ERROR" {
		t.Error("error didn't occurred")
	}
}

func TestInvalidReturnType(t *testing.T) {
	fn1 := func(a ...string) int {
		return 0
	}
	if !isErr(Run(fn1), "required to return an error") {
		t.Error("a func which returns int won't be accepted")
	}

	fn2 := func(a ...string) {
	}
	if !isErr(Run(fn2), "required to return an error") {
		t.Error("a func which returns none won't be accepted")
	}
}

func TestInvalidLastArgType(t *testing.T) {
	fn1 := func(a ...int) error {
		return nil
	}
	if !isErr(Run(fn1), "required []string as last argument") {
		t.Error("a func w/ []int as last arg, won't be accepted")
	}

	fn2 := func(a string) error {
		return nil
	}
	if !isErr(Run(fn2), "required []string as last argument") {
		t.Error("a func w/ string as last arg, won't be accepted")
	}
}
