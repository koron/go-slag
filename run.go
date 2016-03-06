package slag

// Run execute fn with parsed args.
func Run(fn interface{}, args ...string) error {
	fd, err := parseFuncDesc(fn)
	if err != nil {
		return err
	}
	if err := fd.parseArgs(args); err != nil {
		return err
	}
	return fd.call()
}
