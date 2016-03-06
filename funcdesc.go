package slag

import (
	"fmt"
	"io"
	"reflect"
)

type funcDesc struct {
	funcValue reflect.Value
	descs     []optDesc
	argValues []reflect.Value
}

func parseFuncDesc(fn interface{}) (*funcDesc, error) {
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

func (fd *funcDesc) appendRemain(s string) {
	fd.argValues = append(fd.argValues, reflect.ValueOf(s))
}

func (fd *funcDesc) findOptDesc(s string) (*optDesc, error) {
	if len(s) >= 2 && s[:2] == "--" {
		// Find optDesc with (long) name.
		n := s[2:]
		if n != "" {
			for _, o := range fd.descs {
				if n == o.name {
					return &o, nil
				}
			}
		}
		return nil, ErrorSlag{message: "unknown option: " + s}
	}
	// Find option with short name.
	n := s[1:]
	if n != "" {
		for _, o := range fd.descs {
			if n == o.shortName {
				return &o, nil
			}
		}
	}
	return nil, ErrorSlag{message: "unknown option (short): " + s}
}

func (fd *funcDesc) parseArgs(args []string) error {
	i := 0
	for ; i < len(args); i++ {
		s := args[i]
		if len(s) > 0 && s[0] == '-' {
			// Parse remained as pure args after "--".
			if s == "--" {
				i++
				break
			}
			// Parse as an option.
			o, err := fd.findOptDesc(s)
			if err != nil {
				return err
			}
			n, err := o.parseValue(args[i+1:])
			if err != nil {
				return err
			}
			if n > 0 {
				i += n
			}
		} else {
			fd.appendRemain(s)
		}
	}
	for ; i < len(args); i++ {
		fd.appendRemain(args[i])
	}
	return nil
}

func (fd *funcDesc) call() error {
	rv := fd.funcValue.Call(fd.argValues)
	// Parse returned as error.
	v := rv[0].Interface()
	if v == nil {
		return nil
	}
	return v.(error)
}

func (fd *funcDesc) help(w io.Writer) {
	// TODO: show help
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
