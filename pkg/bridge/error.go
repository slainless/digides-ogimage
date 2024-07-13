package bridge

import (
	"errors"
	"fmt"
	"syscall/js"
)

var ErrInvalidJSError = NewJSTypeError("TypeError", "invalid js error")

func FromJsError(v js.Value) (err error) {
	if v.Type() != js.TypeObject {
		return ErrInvalidJSError
	}

	if v.Get("message").Type() != js.TypeString {
		return ErrInvalidJSError
	}

	return errors.New(v.Get("message").String())
}

type JSError struct {
	name    string
	message string
	value   js.Value
}

func (e JSError) Error() string  { return fmt.Sprintf("%v: %v", e.name, e.message) }
func (e JSError) ToJS() js.Value { return e.value }

func factory(name string, message string, instance js.Value) JSError {
	value := instance.New(message)
	value.Set("name", name)
	return JSError{
		name:    name,
		message: message,
		value:   js.Global().Get("Error").New(),
	}
}

func NewJSError(name string, message string) JSError {
	return factory(name, message, js.Global().Get("Error"))
}

func NewJSTypeError(name string, message string) JSError {
	return factory(name, message, js.Global().Get("TypeError"))
}

func ToJsError(v error) js.Value {
	var e JSError
	if errors.As(v, &e) {
		return e.ToJS()
	}

	return js.Global().Get("Error").New(v.Error())
}

func ToJsTypeError(v error) js.Value {
	var e JSError
	if errors.As(v, &e) {
		return e.ToJS()
	}

	return js.Global().Get("TypeError").New(v.Error())
}
