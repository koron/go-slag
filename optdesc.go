package slag

import (
	"reflect"
	"strings"
)

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

func (o *optDesc) errorParseFailure(err error) error {
	return ErrorSlag{
		message: o.displayName() + " parse failure: " + err.Error(),
	}
}

type optDescs struct {
	descs      []optDesc
	shortNames map[string]bool
}

func (ds *optDescs) add(d optDesc) {
	ds.descs = append(ds.descs, d)
	if d.shortName != "" {
		ds.shortNames[d.shortName] = true
	}
}

func (ds *optDescs) toShort(s string) string {
	for _, r := range []rune(strings.ToLower(s)) {
		l := string(r)
		if _, ok := ds.shortNames[l]; !ok {
			return l
		}
		u := strings.ToUpper(l)
		if _, ok := ds.shortNames[u]; !ok {
			return u
		}
	}
	return ""
}

func (ds *optDescs) check(f reflect.StructField) (name, shortName string, err error) {
	n := toSnake(f.Name)
	for _, d := range ds.descs {
		if n == d.name {
			return "", "", ErrorSlag{message: "duplicated option: " + n}
		}
	}
	sn := ds.toShort(f.Name)
	return n, sn, nil
}
