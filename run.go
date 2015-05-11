package slag

import (
	"fmt"
	"reflect"
	"strings"
)

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

// validateFunc checks type of arguments and return value of function.
func validateFunc(funcValue reflect.Value) (argTypes []reflect.Type, err error) {
	ft := funcValue.Type()
	it := inTypes(ft)
	ot := outTypes(ft)
	il := len(it)
	if il < 1 || !isStringArray(it[il-1]) {
		return nil, ErrorSlag{message: "required []string as last argument"}
	}
	for i, t := range it[:il-1] {
		if t.Kind() != reflect.Struct {
			return nil, ErrorSlag{
				message: fmt.Sprintf("argument #%d must be struct", i),
			}
		}
	}
	if len(ot) != 1 || !isErrorType(ot[0]) {
		return nil, ErrorSlag{message: "required to return an error"}
	}
	return it, nil
}

func parseFunc(fn interface{}) (*funcDesc, error) {
	// Assure fn is a func.
	fv := reflect.ValueOf(fn)
	if fv.Kind() != reflect.Func {
		return nil, ErrorSlag{message: "required a function"}
	}
	// Check types of arguments and return value of function.
	it, err := validateFunc(fv)
	if err != nil {
		return nil, err
	}
	// Extract option descriptors.
	descs := &optDescs{
		shortNames: make(map[string]bool),
	}
	av := make([]reflect.Value, 0, len(it))
	for _, t := range it[:len(it)-1] {
		v := reflect.New(t).Elem()
		av = append(av, v)
		// collect optDesc to descs.
		for j := 0; j < t.NumField(); j++ {
			f := t.Field(j)
			n, sn, err := descs.check(f)
			if err != nil {
				return nil, err
			}
			// check f.Type is supported or not.
			c, err := findConverter(f.Type)
			if err != nil {
				return nil, err
			}
			// compose option descriptor
			vf := v.Field(j)
			// TODO: parse f.Tag and extract info: desc, default or so.
			d := optDesc{
				name:      n,
				shortName: sn,
				valueRef:  &vf,
				converter: c,
			}
			descs.add(d)
		}
	}
	return &funcDesc{
		funcValue: fv,
		descs:     descs.descs,
		argValues: av,
	}, nil
}

// Run execute fn with parsed args.
func Run(fn interface{}, args ...string) error {
	fd, err := parseFunc(fn)
	if err != nil {
		return err
	}
	if err := fd.parseArgs(args); err != nil {
		return err
	}
	return fd.call()
}
