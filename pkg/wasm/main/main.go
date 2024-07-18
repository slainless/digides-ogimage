package main

import (
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/slainless/digides-ogimage/pkg/r2"
	"github.com/slainless/digides-ogimage/pkg/reader"
	"github.com/slainless/digides-ogimage/pkg/wasm"
	"github.com/slainless/digides-ogimage/pkg/webp"
	"github.com/slainless/digides-ogimage/pkg/zero"
	"github.com/tetratelabs/wazero"
)

func main() {
	bridge.Console().Log("WASM Loaded.")
	wasmRuntime, ctx, cancel := zero.NewCFWasmRuntime()
	defer cancel()
	defer wasmRuntime.Close(ctx)

	webpModule, err := webp.NewModule(ctx, wasmRuntime,
		wazero.NewModuleConfig().
			WithStderr(bridge.NewWriterFrom(bridge.Console().Error)).
			WithStdout(bridge.NewWriterFrom(bridge.Console().Log)),
	)
	if err != nil {
		bridge.Console().Error(bridge.ToJsError(err))
		return
	}

	js.Global().Set("godrawer", js.ValueOf(map[string]any{
		"draw": wasm.NewJSDraw(webpModule),
		"errors": map[string]any{
			"ErrInvalidStream":          reader.ErrInvalidStream.ToJS(),
			"ErrInvalidReadingResult":   reader.ErrInvalidReadingResult.ToJS(),
			"ErrFileNotFound":           r2.ErrFileNotFound.ToJS(),
			"ErrBucketNotFound":         wasm.ErrBucketNotFound.ToJS(),
			"ErrInvalidCloudflareEnv":   wasm.ErrInvalidCloudflareEnv.ToJS(),
			"ErrParametersInvalid":      wasm.ErrParametersInvalid.ToJS(),
			"ErrParametersInvalidField": wasm.ErrParametersInvalidField.ToJS(),
		},
	}))
	<-make(chan struct{})
}
