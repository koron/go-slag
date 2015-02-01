package slag

import "reflect"

type converter func(args []string, dest *reflect.Value) (used int, err error)

func findConverter(t reflect.Type) (c converter, err error) {
	k := t.Kind()
	// TODO: support other types.
	switch k {
	case reflect.Bool:
		return boolConverter, nil
	default:
		return nil, ErrorSlag{message: "not supported kind: " + k.String()}
	}
	return
}

func boolConverter(args []string, dest *reflect.Value) (used int, err error) {
	dest.SetBool(true)
	return 0, nil
}
