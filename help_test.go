package slag

import (
	"bytes"
	"testing"
)

func assertHelp(t *testing.T, fn interface{}, exp string) {
	fd, err := parseFuncDesc(fn)
	if err != nil {
		t.Fatalf("parse error %v: %s", fn, err)
	}
	b := new(bytes.Buffer)
	fd.help(b)
	s := b.String()
	if s != exp {
		t.Errorf("help expected %q but %q actually", exp, s)
	}
}

func TestHelp0(t *testing.T) {
	fn := func(a ...string) error { return nil }
	assertHelp(t, fn, "(NO OPTIONS)\n")
}

func TestHelp1(t *testing.T) {
	fn := func(v0 struct {
		Help bool
	}, a ...string) error {
		return nil
	}
	assertHelp(t, fn, "OPTIONS:\n  -h/--help\n")
}

func TestHelp2(t *testing.T) {
	fn := func(v0 struct {
		Help bool `slag:"show this message"`
	}, a ...string) error {
		return nil
	}
	assertHelp(t, fn, "OPTIONS:\n  -h/--help\n    \tshow this message\n")
}

func TestHelp3(t *testing.T) {
	fn := func(v0 struct {
		H bool `slag:"show this message"`
	}, a ...string) error {
		return nil
	}
	assertHelp(t, fn, "OPTIONS:\n  -h\tshow this message\n")
}
