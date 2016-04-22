package db

import "fmt"

var (
	ErrNoAffected = NewErrorf("db: no affected row at this execute")
	ErrNilPtr = NewErrorf("db: destination pointer is nil")
)

type Error struct {
	msg string
}

func NewErrorf(format string, args ...interface{}) Error {
	return Error{msg: fmt.Sprintf(format, args...)}
}

func (e Error) Error() string {
	return e.msg
}
