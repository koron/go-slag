package slag

import "testing"

type snakeTest struct {
	input    string
	expected string
}

func TestToSname(t *testing.T) {
	tests := []snakeTest{
		{"a", "a"},
		{"B", "b"},
		{"ID", "id"},
		{"id", "id"},
		{"MyID", "my_id"},
		{"MyIDBaz", "my_id_baz"},
		{"FooBar", "foo_bar"},
		{"fooBar", "foo_bar"},
		{"foo_bar", "foo_bar"},
	}
	for _, c := range tests {
		if v := toSnake(c.input); v != c.expected {
			t.Errorf(`toSnake(%#v) -> %#v != expected %#v`,
				c.input, v, c.expected)
		}
	}
}
