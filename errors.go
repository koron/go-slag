package slag

// ErrorSlag represents an error in parsing slag.
type ErrorSlag struct {
	message string
}

func (e ErrorSlag) Error() string {
	return e.message
}
