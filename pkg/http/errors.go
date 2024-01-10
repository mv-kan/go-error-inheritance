package http

import (
	"errors"
	"fmt"
)

var (
	ErrorHttp           = errors.New("http error")
	ErrorNotImplemented = fmt.Errorf("%w: not implemented", ErrorHttp)
)

type ServerError struct {
	Msg string
}

func (e *ServerError) Unwrap() error {
	return ErrorHttp
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("%s: server err 501, msg=\"%s\"", ErrorHttp.Error(), e.Msg)
}

type BadGatewayError struct {
	msg string
}

func (e *BadGatewayError) Unwrap() error {
	return ErrorHttp
}

func (e *BadGatewayError) Error() string {
	return fmt.Sprintf("%s: bad gateway err 502, msg=\"%s\"", ErrorHttp.Error(), e.msg)
}
