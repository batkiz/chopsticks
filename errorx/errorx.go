package errorx

import (
	"errors"
	"fmt"
)

type Error struct {
	message string
	err     error
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Is(target error) bool {
	if target == nil || e.err == nil {
		return e.err == target
	}

	return errors.Is(e.err, target)
}

func (e *Error) Unwrap() error {
	u, ok := e.err.(interface {
		Unwrap() error
	})
	if !ok {
		return e.err
	}

	return u.Unwrap()
}

func makeMessage(err error, layer, function, msg string) string {
	var message string
	var e Error
	if errors.As(err, &e) {
		message = fmt.Sprintf("[%s:%s] %s => %s", layer, function, msg, err.Error())
	} else {
		message = fmt.Sprintf("[%s:%s] %s => [Raw:Error] %v", layer, function, msg, err.Error())
	}

	return message
}

func Wrapf(err error, layer string, function string, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(format, args...)

	return Error{
		message: makeMessage(err, layer, function, msg),
		err:     err,
	}
}

type WrapfFuncWithLayerFunction func(err error, format string, args ...interface{}) error

func NewLayerFunctionErrorWrapf(layer string, function string) WrapfFuncWithLayerFunction {
	return func(err error, format string, args ...interface{}) error {
		return Wrapf(err, layer, function, format, args...)
	}
}
