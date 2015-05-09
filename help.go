package slag

import "os"

// Help prints help to os.Stdout
func Help(fn interface{}) error {
	fd, err := parseFunc(fn)
	if err != nil {
		return err
	}
	fd.help(os.Stdout)
	return nil
}
