package apierror

type Error struct {
	Code int
	Err  error
}

func New(code int, err ...error) *Error {
	e := &Error{Code: code}
	if len(err) > 0 {
		e.Err = err[0]
	}
	return e
}
