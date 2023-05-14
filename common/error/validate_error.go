package commonerror

type ValidateError interface {
	Error() string
	Entity() string
	Code() int64
}

type validateError struct {
	message string
	entity  string
	code    int64
}

func (e *validateError) Error() string {
	return e.message
}

func (e *validateError) Entity() string {
	return e.entity
}

func (e *validateError) Code() int64 {
	return e.code
}
