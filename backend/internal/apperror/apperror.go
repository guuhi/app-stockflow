package apperror

import "fmt"

type Code string

type AppError struct {
	Code    Code
	Message string
	Fields  map[string]string
	Err     error
}

func New(code Code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func Internal(err error) *AppError {
	return &AppError{Code: CodeInternal, Message: "internal server error", Err: err}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Wrap(cause error) *AppError {
	copy := *e
	copy.Err = cause
	return &copy
}

func (e *AppError) WithFields(fields map[string]string) *AppError {
	copy := *e
	copy.Fields = fields
	return &copy
}
