package slag

import "reflect"

type converter func(d *optDesc, args []string) (used int, err error)

func findConverter(t reflect.Type) (c converter, err error) {
	k := t.Kind()
	// TODO: support other types.
	switch k {
	case reflect.Bool:
		return boolConverter, nil
	case reflect.String:
		return stringConverter, nil
	default:
		return nil, ErrorSlag{message: "not supported kind: " + k.String()}
	}
	return
}

func boolConverter(d *optDesc, args []string) (used int, err error) {
	d.valueRef.SetBool(true)
	return 0, nil
}

func stringConverter(d *optDesc, args []string) (used int, err error) {
	if len(args) < 1 {
		return 0, d.errorNeedArgument()
	}
	d.valueRef.SetString(args[0])
	return 1, nil
}
