package wasm

import (
	"image/jpeg"
	"io"
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/slainless/digides-ogimage/pkg/draw"
	"github.com/slainless/digides-ogimage/pkg/reader"
)

// signature:
// function draw(parameters: payload, bucketName: string, env: any): Promise<ReadableStream>
var JsDraw = js.FuncOf(func(this js.Value, args []js.Value) any {
	rawParameters := args[0]
	bucket := args[1]

	return js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) any {
		resolve := args[0]
		reject := args[1]

		go func() {
			parameters, err := LoadParameters(rawParameters, bucket)
			if err != nil {
				switch err {
				case ErrBucketNotFound:
				case ErrInvalidCloudflareEnv:
				case ErrParametersInvalid:
				case ErrParametersInvalidField:
					reject.Invoke(bridge.ToJsTypeError(err))
					return
				}

				reject.Invoke(bridge.ToJsError(err))
				return
			}

			result, err := draw.Draw(parameters)
			if err != nil {
				reject.Invoke(bridge.ToJsError(err))
				return
			}

			pr, pw := io.Pipe()
			go func() {
				err = jpeg.Encode(pw, result, &jpeg.Options{
					Quality: 100,
				})
				pw.Close()
			}()

			resolve.Invoke(reader.NewReadableStreamFrom(pr, 2048))
		}()

		return nil
	}))
})
