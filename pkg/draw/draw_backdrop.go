package draw

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/slainless/digides-ogimage/pkg/bridge"
)

func drawBackdrop(canvas *gg.Context, bound image.Rectangle) {
	canvas.SetColor(color.NRGBA{0, 0, 0, 89})
	canvas.DrawRectangle(
		float64(bound.Min.X),
		float64(bound.Min.Y),
		float64(bound.Dx()),
		float64(bound.Dy()),
	)
	bridge.Console().Log(bound.Max.X, bound.Max.Y)
	canvas.Fill()
}
