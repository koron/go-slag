package slag

import "testing"

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

func TestBoolPtr(t *testing.T) {
	var flag *bool
	fn := func(o struct{ Flag *bool }, a ...string) error {
		flag = o.Flag
		return nil
	}
	checkRun(t, fn)
	if flag != nil {
		t.Error("flag should be nil: without --flag")
	}
	checkRun(t, fn, "--flag", "t")
	if flag == nil || *flag == false {
		t.Error("flag should be pointer to true")
	}
	checkRun(t, fn, "--flag", "f")
	if flag == nil || *flag == true {
		t.Error("flag should be pointer to false")
	}
	checkRun(t, fn)
	if flag != nil {
		t.Error("flag should be nil: without --flag (2nd)")
	}
}

func TestBoolSlice(t *testing.T) {
	var flags []bool
	fn := func(o struct{ Flags []bool }, a ...string) error {
		flags = o.Flags
		return nil
	}
	checkRun(t, fn)
	if len(flags) != 0 {
		t.Error("flags should be empty: without --flags")
	}
	checkRun(t, fn, "-f", "t", "-f", "t")
	if len(flags) != 2 || flags[0] != true || flags[1] != true {
		t.Errorf("flags should be {true, true}: %#v", flags)
	}
	checkRun(t, fn, "-f", "t", "-f", "f")
	if len(flags) != 2 || flags[0] != true || flags[1] != false {
		t.Errorf("flags should be {true, false}: %#v", flags)
	}
	checkRun(t, fn, "-f", "f", "-f", "t")
	if len(flags) != 2 || flags[0] != false || flags[1] != true {
		t.Errorf("flags should be {true, false}: %#v", flags)
	}
	checkRun(t, fn, "-f", "f", "-f", "f")
	if len(flags) != 2 || flags[0] != false || flags[1] != false {
		t.Errorf("flags should be {true, false}: %#v", flags)
	}
}

func TestString(t *testing.T) {
	name := ""
	fn := func(o struct{ Name string }, a ...string) error {
		name = o.Name
		return nil
	}
	checkRun(t, fn)
	if name != "" {
		t.Error("name should be empty: without --name")
	}
	checkRun(t, fn, "--name", "foo")
	if name != "foo" {
		t.Error("name should be \"foo\"")
	}
	checkRun(t, fn, "--name", "bar")
	if name != "bar" {
		t.Error("name should be \"bar\"")
	}
}

func TestStringPtr(t *testing.T) {
	var name *string
	fn := func(o struct{ Name *string }, a ...string) error {
		name = o.Name
		return nil
	}
	checkRun(t, fn)
	if name != nil {
		t.Error("name should be nil: without --name")
	}
	checkRun(t, fn, "--name", "foo")
	if name == nil || *name != "foo" {
		t.Error("*name should be \"foo\"")
	}
	checkRun(t, fn, "--name", "bar")
	if name == nil || *name != "bar" {
		t.Error("*name should be \"bar\"")
	}
}

func TestInt(t *testing.T) {
	num := 0
	fn := func(o struct{ Number int }, a ...string) error {
		num = o.Number
		return nil
	}
	checkRun(t, fn)
	if num != 0 {
		t.Error("num should be zero: without --number")
	}
	checkRun(t, fn, "-n", "42")
	if num != 42 {
		t.Error("num should be 42")
	}
	checkRun(t, fn, "-n", "-273")
	if num != -273 {
		t.Error("num should be -273")
	}
}

func TestInt64(t *testing.T) {
	var num int64
	fn := func(o struct{ Number int64 }, a ...string) error {
		num = o.Number
		return nil
	}
	checkRun(t, fn)
	if num != 0 {
		t.Error("num should be zero: without --number")
	}
	checkRun(t, fn, "-n", "42")
	if num != 42 {
		t.Error("num should be 42")
	}
	checkRun(t, fn, "-n", "-273")
	if num != -273 {
		t.Error("num should be -273")
	}
}

// TODO: test uint conveter
// TODO: test int/uint pointer conveter
