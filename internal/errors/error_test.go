package errors_test

import (
	stderrors "errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"postglide.io/postglide/internal/errors"
)

func TestWrapAndErrorMessage(t *testing.T) {
	assert := assert.New(t)

	appErrWithNilError := errors.Wrap(errors.CodeInternal, nil)
	assert.Nil(appErrWithNilError)

	errMessage := "authentication failed"
	authenticationError := stderrors.New(errMessage)
	appErr := errors.Wrap(errors.CodeUnAuthorized, authenticationError)
	assert.NotNil(appErr)

	want := "[UNAUTHORIZED]: " + errMessage
	got := appErr.Error()
	assert.Equal(got, want)
}

func TestStackTraceContainsCaller(t *testing.T) {
	err := stderrors.New("db failed")
	appErr := errors.Wrap(errors.CodeInternal, err)

	stack := appErr.StackTrace()

	assert.True(t, strings.Contains(stack, "TestStackTraceContainsCaller"))
}
