package cerror

type Error struct {
	str string
}

func NewError(msg string) *Error {
	return &Error{
		str: msg,
	}
}
func (e *Error) Error() string {
	return e.str
}
