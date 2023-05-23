package commonerror

func NewBusinessError(message string, entity string, code int64) BusinessError {
	return &businessError{message, entity, code}
}
