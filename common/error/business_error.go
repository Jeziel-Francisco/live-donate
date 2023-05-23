package commonerror

type BusinessError interface {
	Error() string
	Entity() string
	Code() int64
}

type businessError struct {
	message string
	entity  string
	code    int64
}

func (e *businessError) Error() string {
	return e.message
}

func (e *businessError) Entity() string {
	return e.entity
}

func (e *businessError) Code() int64 {
	return e.code
}
