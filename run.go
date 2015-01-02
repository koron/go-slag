package slag

import (
	"fmt"
	"reflect"
	"strings"
)

type optDesc struct {
	name      string
	shortName string
	valueRef  *reflect.Value
}

func (o *optDesc) parseValue([]string) (used int, err error) {
	// TODO: implement me.
	k := o.valueRef.Kind()
	switch k {
	case reflect.Bool:
		o.valueRef.SetBool(true)
	default:
		return 0, ErrorSlag{message: "not supported kind: " + k.String()}
	}
	return 0, nil
}

type descriptor struct {
	funcValue reflect.Value
	descs     []optDesc
	argValues []reflect.Value
}

func (d *descriptor) appendRemain(s string) {
	d.argValues = append(d.argValues, reflect.ValueOf(s))
}

func (d *descriptor) findOptDesc(s string) (*optDesc, error) {
	if len(s) >= 2 && s[:2] == "--" {
		n := s[2:]
		for _, o := range d.descs {
			if n == o.name {
				return &o, nil
			}
		}
		return nil, ErrorSlag{message: "unknown option: " + s}
	} else {
		n := s[1:]
		// TODO: implement short name detector.
		return nil, ErrorSlag{message: "not implement short name yet: " + n}
	}
}

func (d *descriptor) call() error {
	_ = d.funcValue.Call(d.argValues)
	// TODO: parse r as error.
	return nil
}

func optionName(s string) string {
	// TODO: generate (regular) name.
	return strings.ToLower(s)
}

func optionShortName(od []optDesc, s string) string {
	// TODO: generate short name.
	return ""
}

func checkField(od []optDesc, f reflect.StructField) (name, shortName string, err error) {
	n := optionName(f.Name)
	for _, d := range od {
		if n == d.name {
			return "", "", ErrorSlag{message: "duplicated option: " + n}
		}
	}
	sn := optionShortName(od, strings.ToLower(f.Name))
	// TODO: check f.Type is supported type or not.
	return n, sn, nil
}

func parseFunc(fn interface{}) (*descriptor, error) {
	// Assure fn is a func.
	fv := reflect.ValueOf(fn)
	if fv.Kind() != reflect.Func {
		return nil, ErrorSlag{message: "required a function"}
	}
	// Check types of argument and return of function.
	ft := fv.Type()
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
	// Extract option descriptors.
	od := make([]optDesc, 0)
	av := make([]reflect.Value, 0, len(it))
	for _, t := range it[:il-1] {
		v := reflect.New(t).Elem()
		av = append(av, v)
		for j := 0; j < t.NumField(); j++ {
			f := t.Field(j)
			n, sn, err := checkField(od, f)
			if err != nil {
				return nil, err
			}
			vf := v.Field(j)
			d := optDesc{
				name:      n,
				shortName: sn,
				valueRef:  &vf,
			}
			od = append(od, d)
		}
	}
	return &descriptor{
		funcValue: fv,
		descs:     od,
		argValues: av,
	}, nil
}

// Run execute fn with parsed args.
func Run(fn interface{}, args ...string) error {
	d, err := parseFunc(fn)
	if err != nil {
		return err
	}
	// Parse args.
	i := 0
	for ; i < len(args); i++ {
		s := args[i]
		if len(s) > 0 && s[0] == '-' {
			// Parse remained as pure args after "--".
			if s == "--" {
				i += 1
				break
			}
			// Parse as an option.
			o, err := d.findOptDesc(s)
			if err != nil {
				return err
			} else if o == nil {
				continue
			}
			n, err := o.parseValue(args[i+1:])
			if err != nil {
				return err
			}
			if n > 0 {
				i += n
			}
		} else {
			d.appendRemain(s)
		}
	}
	for ; i < len(args); i++ {
		d.appendRemain(args[i])
	}
	return d.call()
}
