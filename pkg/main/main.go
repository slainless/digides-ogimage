package main

import (
	"bytes"
	"image/jpeg"
	"io"

	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/slainless/digides-ogimage/pkg/ogimage"
	"github.com/teamortix/golang-wasm/wasm"
)

func start(raw map[string]any) ([]byte, error) {
	params := bridge.LoadParameters(raw)

	result, err := ogimage.Draw(params)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	jpeg.Encode(&buf, result, &jpeg.Options{
		Quality: 80,
	})
	return io.ReadAll(&buf)
}

func main() {
	wasm.Expose("draw", start)
	wasm.Ready()
	<-make(chan struct{})
}
