package errors

type xError struct {
	preErr error
	code   int
	msg    string
	*stack
}

func New(code int, msg string) error {
	return &xError{
		preErr: nil,
		code:   code,
		msg:    msg,
		stack:  callers(),
	}
}

func (w *xError) Error() string { return w.msg }

func (w *xError) Unwrap() error { return w.preErr }

func (w *xError) Code() int { return w.code }

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*xError); ok {
		return &xError{
			preErr: e,
			code:   e.code,
			msg:    msg,
			stack:  callers(),
		}
	} else {
		return New(-1, e.Error())
	}
}

func Unwrap(err error) error {
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}
