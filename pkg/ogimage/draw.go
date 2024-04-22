package ogimage

import (
	"image"

	"github.com/fogleman/gg"
)

// 0.262145631 s/op
// using parameters: [github.com/slainless/digides-ogimage/pkg/ogimage_test.LoadParameters]
func Draw(param *Parameters) (image.Image, error) {
	const (
		canvasBoundX = 1200
		canvasBoundY = 630

		margin = 95
	)

	canvas := gg.NewContext(canvasBoundX, canvasBoundY)

	elements, err := drawElements(param)
	if err != nil {
		return nil, err
	}

	background := ResizeToFill(param.Background, canvasBoundX, canvasBoundY)
	canvas.DrawImage(background, 0, 0)

	resizedElements := ResizeToFit(elements, canvasBoundX-(margin*2), canvasBoundY-(margin*2))
	canvas.DrawImageAnchored(resizedElements, margin, int(canvasBoundY/2), 0, 0.5)

	return canvas.Image(), nil
}
