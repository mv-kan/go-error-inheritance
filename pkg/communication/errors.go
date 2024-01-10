package communication

import (
	"errors"
	"fmt"
)

var (
	ErrorCommunication = errors.New("communication error")
)

type TimeoutError struct {
	Timeout string
}

func (t *TimeoutError) Unwrap() error {
	return ErrorCommunication
}

func (t *TimeoutError) Error() string {
	return fmt.Sprintf("%s: timeout (timeout=%s)", ErrorCommunication.Error(), t.Timeout)
}

type UnexpectedError struct {
	Msg string
}

func (e *UnexpectedError) Unwrap() error {
	return ErrorCommunication
}

func (e *UnexpectedError) Error() string {
	return fmt.Sprintf("%s: %s", ErrorCommunication.Error(), e.Msg)
}
