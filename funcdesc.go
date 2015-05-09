package slag

import (
	"io"
	"reflect"
)

type funcDesc struct {
	funcValue reflect.Value
	descs     []optDesc
	argValues []reflect.Value
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
