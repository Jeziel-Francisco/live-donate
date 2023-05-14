package commonerror

func NewValidateError(message string, entity string, code int64) ValidateError {
	return &validateError{message, entity, code}
}
