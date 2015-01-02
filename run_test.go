package slag

import "testing"

func checkRun(t *testing.T, fn interface{}, args ...string) {
	if err := Run(fn, args...); err != nil {
		t.Error("Run() failed:", err)
	}
}

func TestBool(t *testing.T) {
	verbose := false
	fn := func(o struct{ Verbose bool }, a ...string) error {
		verbose = o.Verbose
		return nil
	}
	checkRun(t, fn)
	if verbose != false {
		t.Error("verbose should be false: without --verbose")
	}
	checkRun(t, fn, "--verbose")
	if verbose != true {
		t.Error("verbose should be true: with --verbose")
	}
	checkRun(t, fn)
	if verbose != false {
		t.Error("verbose should be false: without --verbose")
	}
}
