package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type appError struct {
	code       Code
	err        error
	stackTrace string
}

func Wrap(code Code, err error) *appError {
	if err == nil {
		return nil
	}
	stack := captureStack(3)
	return &appError{code: code, err: err, stackTrace: stack}
}

func (e *appError) Code() Code {
	return e.code
}

func (e *appError) Error() string {
	return fmt.Sprintf("[%s]: %s", e.code, e.err)
}

func (e *appError) StackTrace() string {
	return e.stackTrace
}

func (e *appError) Unwrap() error {
	return e.err
}

func captureStack(skip int) string {
	const depth = 32
	pcs := make([]uintptr, depth)
	n := runtime.Callers(skip, pcs)
	if n == 0 {
		return "no pcs captured"
	}

	frames := runtime.CallersFrames(pcs[:n])
	var b strings.Builder
	for {
		frame, more := frames.Next()
		b.WriteString(fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}
	return b.String()
}

func (e *appError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "[%s]: %s\nStacktrace:\n%s", e.code, e.err, e.stackTrace)
			return
		}
		fallthrough
	case 's':
		fmt.Fprint(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}
