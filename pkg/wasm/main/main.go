package main

import (
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/slainless/digides-ogimage/pkg/r2"
	"github.com/slainless/digides-ogimage/pkg/reader"
	"github.com/slainless/digides-ogimage/pkg/wasm"
)

func main() {
	bridge.Console().Log("WASM Loaded.")
	js.Global().Set("godrawer", js.ValueOf(map[string]any{
		"draw": wasm.JsDraw,
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
