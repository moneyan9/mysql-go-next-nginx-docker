package controllers

type Error struct {
	message string
}

func NewError(err error) *Error {
	return &Error{message: err.Error()}
}
