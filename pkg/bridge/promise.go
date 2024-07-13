package bridge

import (
	"syscall/js"
)

var (
	ErrNotPromise = NewJSTypeError("PromiseError", "not a valid promise")
)

func ResolvePromise(promise js.Value) (result *js.Value, err *js.Value, goErr error) {
	defer func() {
		if r := recover(); r != nil {
			goErr = r.(error)
		}
	}()

	if promise.InstanceOf(js.Global().Get("Promise")) == false {
		return nil, nil, ErrNotPromise
	}

	valChan := make(chan js.Value)
	errChan := make(chan js.Value)

	promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		valChan <- args[0]
		return nil
	}))
	promise.Call("catch", js.FuncOf(func(this js.Value, args []js.Value) any {
		errChan <- args[0]
		return nil
	}))

	select {
	case val := <-valChan:
		return &val, nil, nil
	case err := <-errChan:
		return nil, &err, nil
	}
}
