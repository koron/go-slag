package slag

import "reflect"

type converter interface {
	convert(args []string, dest *reflect.Value) (used int, err error)
}
