package ogimage

import (
	"bytes"
	"image"
)

func decode(data []byte) (image.Image, error) {
	reader := bytes.NewReader(data)
	image, _, err := image.Decode(reader)
	return image, err
}
