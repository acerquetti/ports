package domain

import (
	"errors"
)

const (
	ErrInvalidModel = iota
)

type PortsError struct {
	err    error
	errNum int
}

func (a *PortsError) Error() string {
	return a.err.Error()
}

func (a *PortsError) Num() int {
	return a.errNum
}

func Error(num int, msg string) error {
	return &PortsError{
		err:    errors.New(msg),
		errNum: num,
	}
}
