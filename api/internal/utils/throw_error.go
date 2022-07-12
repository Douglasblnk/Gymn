package utils

type Error struct {
	Error      string
	StatusCode int
}

func Throw(msg string, status int) *Error {
	return &Error{
		Error:      msg,
		StatusCode: status,
	}
}
