package slag

import (
	"reflect"
	"strconv"
)

type converter func(d *optDesc, args []string) (used int, err error)

func findConverter(t reflect.Type) (c converter, err error) {
	switch k := t.Kind(); k {
	case reflect.Bool:
		return boolConverter, nil
	case reflect.String:
		return stringConverter, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return intConverter, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return uintConverter, nil
	case reflect.Float32, reflect.Float64:
		return floatConverter, nil

	case reflect.Ptr:
		t2 := t.Elem()
		switch k2 := t2.Kind(); k2 {
		case reflect.Bool:
			return boolPtrConverter, nil
		case reflect.String:
			return stringPtrConverter, nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
			reflect.Int64:
			return intPtrConverter, nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
			reflect.Uint64:
			return uintPtrConverter, nil
		default:
			// FIXME: better message.
			return nil, ErrorSlag{message: "not supported kind: " + k2.String()}
		}

	// TODO: support array types.

	default:
		// FIXME: better message.
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

func intConverter(d *optDesc, args []string) (used int, err error) {
	if len(args) < 1 {
		return 0, d.errorNeedArgument()
	}
	v, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		return 0, d.errorParseFailure(err)
	}
	d.valueRef.SetInt(v)
	return 1, nil
}

func uintConverter(d *optDesc, args []string) (used int, err error) {
	if len(args) < 1 {
		return 0, d.errorNeedArgument()
	}
	v, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		return 0, d.errorParseFailure(err)
	}
	d.valueRef.SetUint(v)
	return 1, nil
}

func floatConverter(d *optDesc, args []string) (used int, err error) {
	if len(args) < 1 {
		return 0, d.errorNeedArgument()
	}
	v, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, d.errorParseFailure(err)
	}
	d.valueRef.SetFloat(v)
	return 1, nil
}

func boolPtrConverter(d *optDesc, args []string) (used int, err error) {
	if len(args) < 1 {
		return 0, d.errorNeedArgument()
	}
	v, err := strconv.ParseBool(args[0])
	if err != nil {
		return 0, d.errorParseFailure(err)
	}
	pv := reflect.New(d.valueRef.Type().Elem())
	pv.Elem().SetBool(v)
	d.valueRef.Set(pv)
	return 1, nil
}

func stringPtrConverter(d *optDesc, args []string) (used int, err error) {
	if len(args) < 1 {
		return 0, d.errorNeedArgument()
	}
	v := args[0]
	if err != nil {
		return 0, d.errorParseFailure(err)
	}
	pv := reflect.New(d.valueRef.Type().Elem())
	pv.Elem().SetString(v)
	d.valueRef.Set(pv)
	return 1, nil
}

func intPtrConverter(d *optDesc, args []string) (used int, err error) {
	if len(args) < 1 {
		return 0, d.errorNeedArgument()
	}
	v, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		return 0, d.errorParseFailure(err)
	}
	pv := reflect.New(d.valueRef.Type().Elem())
	pv.Elem().SetInt(v)
	d.valueRef.Set(pv)
	return 1, nil
}

func uintPtrConverter(d *optDesc, args []string) (used int, err error) {
	if len(args) < 1 {
		return 0, d.errorNeedArgument()
	}
	v, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		return 0, d.errorParseFailure(err)
	}
	pv := reflect.New(d.valueRef.Type().Elem())
	pv.Elem().SetUint(v)
	d.valueRef.Set(pv)
	return 1, nil
}
