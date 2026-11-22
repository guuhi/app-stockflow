package apperror

const (
	CodeValidation Code = "VALIDATION_ERROR"
	CodeNotFound   Code = "NOT_FOUND"
	CodeInternal   Code = "INTERNAL"
)

var (
	ErrValidation          = New(CodeValidation, "validation error")
	ErrNotFound            = New(CodeNotFound, "resource not found")
	ErrDatabaseQueryFailed = New(CodeInternal, "failed to execute database query")
)
