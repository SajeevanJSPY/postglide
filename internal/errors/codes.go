package errors

type Code string

const (
	// Internal
	CodeInternal Code = "INTERNAL"
	CodeDatabase Code = "DB_ERROR"
	CodeTimeout  Code = "TIMEOUT"

	// Validation / Client
	CodeUnAuthorized = "UNAUTHORIZED"
	CodeInvalidInput = "INVALID_INPUT"
)
