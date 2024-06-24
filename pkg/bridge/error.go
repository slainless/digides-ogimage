package bridge

import (
	"errors"
	"syscall/js"
)

var ErrInvalidJSError = errors.New("invalid js error")

func FromJsError(v js.Value) (err error) {
	if v.Type() != js.TypeObject {
		return ErrInvalidJSError
	}

	if v.Get("message").Type() != js.TypeString {
		return ErrInvalidJSError
	}

	return errors.New(v.Get("message").String())
}

func ToJsError(v error) js.Value {
	return js.Global().Get("Error").New(v.Error())
}

func ToJsTypeError(v error) js.Value {
	return js.Global().Get("TypeError").New(v.Error())
}
