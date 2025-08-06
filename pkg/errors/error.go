package errors

import "fmt"

type AppError struct {
	Code Code
	Err  error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s]: %s", e.Code, e.Err)
}

func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	return ok && e.Code == t.Code
}

func Wrap(code Code, err error) error {
	return &AppError{Code: code, Err: err}
}

func (e *AppError) Unwrap() error {
	return e.Err
}
