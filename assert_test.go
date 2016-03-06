package slag

import "testing"

func assertBools(t *testing.T, act, exp []bool) {
	match := true
	if len(act) == len(exp) {
		for i, b := range act {
			if b != exp[i] {
				match = false
				break
			}
		}
	} else {
		match = false
	}
	if !match {
		t.Errorf("[]bool should be %#v but actually %#v", exp, act)
	}
}

func assertStrings(t *testing.T, act, exp []string) {
	match := true
	if len(act) == len(exp) {
		for i, b := range act {
			if b != exp[i] {
				match = false
				break
			}
		}
	} else {
		match = false
	}
	if !match {
		t.Errorf("[]string should be %#v but actually %#v", exp, act)
	}
}
