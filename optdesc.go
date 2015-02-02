package slag

import "reflect"

type optDesc struct {
	name      string
	shortName string
	valueRef  *reflect.Value
	converter converter
}

func (o *optDesc) parseValue(args []string) (used int, err error) {
	return o.converter(o, args)
}

func (o *optDesc) displayName() string {
	if o.shortName != "" {
		if o.name != "" {
			return "-" + o.shortName + "/--" + o.name
		}
		return "-" + o.shortName
	}
	return "--" + o.name
}

func (o *optDesc) errorNeedArgument() error {
	// FIXME: better message.
	return ErrorSlag{
		message: o.displayName() + " need an argument",
	}
}
