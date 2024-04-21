package main

import (
	"encoding/base64"
	"image/jpeg"
	"io"

	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/slainless/digides-ogimage/pkg/ogimage"
	"github.com/teamortix/golang-wasm/wasm"
)

func start(raw map[string]any) (string, error) {
	params, err := bridge.LoadParameters(raw)
	if err != nil {
		return "", err
	}

	result, err := ogimage.Draw(params)
	if err != nil {
		return "", err
	}

	reader, writer := io.Pipe()
	encoder := base64.NewEncoder(base64.RawStdEncoding, writer)
	go func() {
		err = jpeg.Encode(encoder, result, &jpeg.Options{
			Quality: 100,
		})
		writer.Close()
	}()

	data, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(data), err
}

func main() {
	wasm.Expose("draw", start)
	wasm.Ready()
	<-make(chan struct{})
}
