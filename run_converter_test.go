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

func TestUint(t *testing.T) {
	var num uint64
	fn := func(o struct{ Number uint64 }, a ...string) error {
		num = o.Number
		return nil
	}
	checkRun(t, fn)
	if num != 0 {
		t.Error("num should be zero without -n")
	}
	checkRun(t, fn, "-n", "42")
	if num != 42 {
		t.Error("num should be 42")
	}
	checkRun(t, fn, "-n", "999999")
	if num != 999999 {
		t.Error("num should be 999999")
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

func TestIntPtr(t *testing.T) {
	var value *int
	fn := func(o struct{ Value *int }, a ...string) error {
		value = o.Value
		return nil
	}
	checkRun(t, fn)
	if value != nil {
		t.Error("value should be nil without -v")
	}
	checkRun(t, fn, "-v", "123")
	if value == nil || *value != 123 {
		t.Error("*value should be 123")
	}
	checkRun(t, fn, "-v", "-999999999")
	if value == nil || *value != -999999999 {
		t.Error("*value should be -999999999")
	}
}

func TestUintPtr(t *testing.T) {
	var value *uint
	fn := func(o struct{ Value *uint }, a ...string) error {
		value = o.Value
		return nil
	}
	checkRun(t, fn)
	if value != nil {
		t.Error("value should be nil without -v")
	}
	checkRun(t, fn, "-v", "123")
	if value == nil || *value != 123 {
		t.Error("*value should be 123")
	}
	checkRun(t, fn, "-v", "999999999")
	if value == nil || *value != 999999999 {
		t.Error("*value should be 999999999")
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

func TestStringSlice(t *testing.T) {
	var s []string
	fn := func(o struct{ Values []string }, a ...string) error {
		s = o.Values
		return nil
	}
	checkRun(t, fn)
	if len(s) != 0 {
		t.Error("slice should be empty without -v")
	}
	checkRun(t, fn, "-v", "foo", "-v", "bar")
	if len(s) != 2 || s[0] != "foo" || s[1] != "bar" {
		t.Errorf("slice should be %#v: %#v", []string{"foo", "bar"}, s)
	}
	checkRun(t, fn, "-v", "aa", "-v", "bb", "-v", "cc")
	if len(s) != 3 || s[0] != "aa" || s[1] != "bb" || s[2] != "cc" {
		t.Errorf("slice should be %#v: %#v", []string{"aa", "bb", "cc"}, s)
	}
}

func TestIntSlice(t *testing.T) {
	var s []int
	fn := func(o struct{ Values []int }, a ...string) error {
		s = o.Values
		return nil
	}
	checkRun(t, fn)
	if len(s) != 0 {
		t.Error("slice should be empty without -v")
	}
	checkRun(t, fn, "-v", "42", "-v", "-999999999")
	if len(s) != 2 || s[0] != 42 || s[1] != -999999999 {
		t.Errorf("slice should be %#v: %#v", []int{42, -999999999}, s)
	}
}

func TestUintSlice(t *testing.T) {
	var s []uint
	fn := func(o struct{ Values []uint }, a ...string) error {
		s = o.Values
		return nil
	}
	checkRun(t, fn)
	if len(s) != 0 {
		t.Error("slice should be empty without -v")
	}
	checkRun(t, fn, "-v", "42", "-v", "999999999")
	if len(s) != 2 || s[0] != 42 || s[1] != 999999999 {
		t.Errorf("slice should be %#v: %#v", []uint{42, 999999999}, s)
	}
}
