package r2

import (
	"errors"
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
)

func NewMockBucket(assetMapping map[string][]byte) js.Value {
	return js.ValueOf(map[string]any{
		"get": js.FuncOf(func(this js.Value, args []js.Value) any {
			path := args[0].String()

			return js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) any {
				resolve := args[0]
				reject := args[1]

				data := assetMapping[path]
				if data == nil {
					return reject.Invoke(bridge.ToJsError(errors.New("asset not found")))
				}

				return resolve.Invoke(js.ValueOf(map[string]any{
					"arrayBuffer": js.FuncOf(func(this js.Value, args []js.Value) any {
						return js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) any {
							resolve := args[0]
							view := js.Global().Get("Uint8Array").New(js.ValueOf(len(data)))
							js.CopyBytesToJS(view, data)
							return resolve.Invoke(view)
						}))
					}),
				}))
			}))
		}),
	})
}
