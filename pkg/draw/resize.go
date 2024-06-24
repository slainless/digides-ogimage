package draw

import (
	"image"

	"github.com/disintegration/gift"
)

func ResizeToFill(img image.Image, w int, h int) *image.RGBA {
	resizing := gift.New(gift.ResizeToFill(w, h, gift.LanczosResampling, gift.CenterAnchor))
	icon := image.NewRGBA(resizing.Bounds(img.Bounds()))
	resizing.Draw(icon, img)

	return icon
}

func ResizeToFit(img image.Image, w int, h int) *image.RGBA {
	resizing := gift.New(gift.ResizeToFit(w, h, gift.LanczosResampling))
	icon := image.NewRGBA(resizing.Bounds(img.Bounds()))
	resizing.Draw(icon, img)

	return icon
}
