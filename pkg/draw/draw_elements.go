package draw

import (
	"image"
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

func drawElements(param Parameters) (image.Image, error) {
	const (
		debug = false

		canvasRatio = 2.805
		iconSize    = 255

		canvasMargin = 65
		gap          = 60
	)

	var (
		canvasHeight float64 = iconSize + (canvasMargin * 2)
		canvasWidth          = canvasRatio * canvasHeight
	)

	strings, err := drawStrings(param)
	if err != nil {
		return nil, err
	}

	canvas := gg.NewContext(int(math.Round(canvasWidth)), int(math.Round(canvasHeight)))
	canvas.DrawRectangle(0, 0, canvasWidth, canvasHeight)
	canvas.SetColor(color.NRGBA{0, 0, 0, 89})
	canvas.Fill()

	if debug {
		canvas.DrawRectangle(canvasMargin, canvasMargin, iconSize, iconSize)
		canvas.SetColor(color.NRGBA{0, 255, 0, 100})
		canvas.Fill()
	}

	icon := ResizeToFit(param.Icon(), iconSize, iconSize)
	canvas.DrawImage(icon, canvasMargin, canvasMargin)

	if debug {
		canvas.DrawRectangle(canvasMargin+iconSize+gap, canvasMargin, canvasWidth-iconSize-gap-(canvasMargin*2), iconSize)
		canvas.SetColor(color.NRGBA{0, 255, 255, 100})
		canvas.Fill()
	}

	_strings := ResizeToFit(strings, int(canvasWidth-iconSize-gap-(canvasMargin*2)), iconSize)
	canvas.DrawImageAnchored(_strings, int(canvasMargin+iconSize+gap+(canvasWidth-iconSize-gap-(canvasMargin*2))/2), int(canvasHeight/2), 0.5, 0.5)

	return canvas.Image(), nil
}
