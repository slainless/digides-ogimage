package ogimage

import (
	"image"
	"image/color"

	"github.com/disintegration/gift"
	"github.com/fogleman/gg"
)

const (
	CanvasBoundX = 1200
	CanvasBoundY = 630
)

func Resize(img image.Image, w int, h int) *image.RGBA {
	resizing := gift.New(gift.ResizeToFill(225, 225, gift.LanczosResampling, gift.CenterAnchor))
	icon := image.NewRGBA(resizing.Bounds(img.Bounds()))
	resizing.Draw(icon, img)

	return icon
}

func Draw(param *Parameters) (image.Image, error) {
	icon := Resize(param.Icon, 225, 225)
	background := Resize(param.Background, CanvasBoundX, CanvasBoundY)

	canvas := gg.NewContext(CanvasBoundX, CanvasBoundY)
	canvas.DrawImage(background, 0, 0)

	canvas.SetColor(color.RGBA{0, 0, 0, 89})
	canvas.DrawRectangle(95, 135, 1010, 360)
	canvas.Fill()

	canvas.DrawImage(icon, 168, 202)

	return canvas.Image(), nil
}
