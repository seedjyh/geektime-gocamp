package code

type Error struct {
	Code    int
	Message string // 错误描述，用于给用户阅读
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
