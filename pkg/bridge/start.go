package bridge

import (
	"encoding/base64"
	"image/jpeg"
	"io"

	"github.com/slainless/digides-ogimage/pkg/ogimage"
)

// 0.314484976 op/s
// using parameters: [github.com/slainless/digides-ogimage/pkg/bridge_test.MockParameters]
func Start(raw map[string]any) (string, error) {
	params, err := LoadParameters(raw)
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
