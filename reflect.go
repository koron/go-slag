package slag

import "reflect"

// errorType represents a reflection type for Error.
var errorType = reflect.TypeOf((*error)(nil)).Elem()

func getTypes(t reflect.Type, n int, f func(int) reflect.Type) []reflect.Type {
	a := make([]reflect.Type, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, f(i))
	}
	return a
}

// inTypes returns "argument types" of a function.
func inTypes(t reflect.Type) []reflect.Type {
	return getTypes(t, t.NumIn(), func(n int) reflect.Type {
		return t.In(n)
	})
}

// outTypes returns "return types" of a function.
func outTypes(t reflect.Type) []reflect.Type {
	return getTypes(t, t.NumOut(), func(n int) reflect.Type {
		return t.Out(n)
	})
}

func isErrorType(t reflect.Type) bool {
	return t == errorType
}

func isStringArray(t reflect.Type) bool {
	if t.Kind() != reflect.Slice {
		// Not a slice
		return false
	}
	if t.Elem().Kind() != reflect.String {
		// Not a string in slice.
		return false
	}
	return true
}
