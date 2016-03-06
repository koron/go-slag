package slag

import "os"

// Help prints help to os.Stdout
func Help(fn interface{}) error {
	fd, err := parseFuncDesc(fn)
	if err != nil {
		return err
	}
	fd.help(os.Stdout)
	return nil
}
