package slag

import "reflect"

var errorType = reflect.TypeOf((*error)(nil)).Elem()

func getTypes(t reflect.Type, n int, f func(int) reflect.Type) []reflect.Type {
	a := make([]reflect.Type, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, f(i))
	}
	return a
}

func inTypes(t reflect.Type) []reflect.Type {
	return getTypes(t, t.NumIn(), func(n int) reflect.Type {
		return t.In(n)
	})
}

func outTypes(t reflect.Type) []reflect.Type {
	return getTypes(t, t.NumOut(), func(n int) reflect.Type {
		return t.Out(n)
	})
}

func isErrorType(t reflect.Type) bool {
	return t == errorType
}

func isStringArray(t reflect.Type) bool {
	// TODO: implement me.
	return true
}
