package reader

import (
	"io"
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
)

var console = bridge.Console

func writeToBYOBRequest(reader io.Reader, controller js.Value) error {
	view := controller.Get("byobRequest").Get("view")
	buf := make([]byte, view.Get("byteLength").Int())
	n, err := reader.Read(buf)

	if err != nil {
		if err == io.EOF {
			controller.Call("close")
		} else {
			controller.Call("error", bridge.ToJsError(err))
		}

		controller.Get("byobRequest").Call("respond", 0)
		return err
	}

	js.CopyBytesToJS(view, buf[:n])
	controller.Get("byobRequest").Call("respond", js.ValueOf(n))
	return nil
}

func writeToEnqueue(reader io.Reader, controller js.Value, chunkSize int) error {
	buf := make([]byte, chunkSize)
	n, err := reader.Read(buf)

	if err != nil {
		if err == io.EOF {
			controller.Call("close")
		} else {
			controller.Call("error", bridge.ToJsError(err))
		}
		return err
	}

	view := js.Global().Get("Uint8Array").New(js.ValueOf(chunkSize))
	js.CopyBytesToJS(view, buf[:n])
	controller.Call("enqueue", view)
	return nil
}

func isBYOBController(controller js.Value) bool {
	request := controller.Get("byobRequest")
	return !bridge.IsNullish(request)
}

func NewReadableStreamFrom(reader io.Reader, defaultChunkSize int) js.Value {
	return js.Global().Get("ReadableStream").New(js.ValueOf(map[string]any{
		"type": js.ValueOf("bytes"),
		// start-mode
		"start": js.FuncOf(func(this js.Value, args []js.Value) any {
			controller := args[0]
			// console().Log(js.ValueOf("[Start] Is BYOB Controller:"), js.ValueOf(isBYOBController(controller)))
			if isBYOBController(controller) {
				// console().Log(js.ValueOf("[Start] BYOB View:"), controller.Get("byobRequest").Get("view"))
				_ = writeToBYOBRequest(reader, controller)
			} else {
				for {
					// console().Log(js.ValueOf("[Start] Fired enqueue."))
					err := writeToEnqueue(reader, controller, defaultChunkSize)
					if err != nil {
						break
					}
				}
			}

			return nil
		}),
		// pull-mode
		"pull": js.FuncOf(func(this js.Value, args []js.Value) any {
			controller := args[0]
			// console().Log(js.ValueOf("[Pull] Is BYOB Controller:"), js.ValueOf(isBYOBController(controller)))
			if isBYOBController(controller) {
				// console().Log(js.ValueOf("[Pull] BYOB View:"), controller.Get("byobRequest").Get("view"))
				_ = writeToBYOBRequest(reader, controller)
			} else {
				// console().Log(js.ValueOf("[Pull] Fired enqueue."))
				_ = writeToEnqueue(reader, controller, defaultChunkSize)
			}

			return nil
		}),
	}))
}
