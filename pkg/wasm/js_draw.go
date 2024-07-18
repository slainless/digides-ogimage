package wasm

import (
	"errors"
	"image/jpeg"
	"io"
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/slainless/digides-ogimage/pkg/draw"
	"github.com/slainless/digides-ogimage/pkg/reader"
)

func parseQuality(quality js.Value) (int, error) {
	if quality.Type() != js.TypeNumber {
		return 0, errors.New("quality must be a number")
	}

	if quality.Int() < 0 || quality.Int() > 100 {
		return 0, errors.New("quality must be between 0 and 100")
	}

	if js.Global().Get("Number").Call("isNaN", quality).Truthy() {
		return 0, errors.New("quality must be a number")
	}

	return quality.Int(), nil
}

// signature:
// function draw(parameters: payload, bucketName: string, env: any): Promise<ReadableStream>
var JsDraw = js.FuncOf(func(this js.Value, args []js.Value) any {
	rawParameters := args[0]
	bucket := args[1]
	quality := args[2]

	return js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) any {
		resolve := args[0]
		reject := args[1]

		q, err := parseQuality(quality)
		if err != nil {
			reject.Invoke(bridge.ToJsError(err))
			return nil
		}

		go func() {
			parameters, err := LoadParameters(rawParameters, bucket)
			if err != nil {
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
					Quality: q,
				})
				defer pw.Close()
				if err != nil {
					reject.Invoke(bridge.ToJsError(err))
					return
				}
			}()

			resolve.Invoke(reader.NewReadableStreamFrom(pr, 2048))
		}()

		return nil
	}))
})
