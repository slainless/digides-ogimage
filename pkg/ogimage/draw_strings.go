package ogimage

import (
	"errors"
	"image"
	"image/color"
	"math"

	"github.com/disintegration/gift"
	"github.com/fogleman/gg"
	"github.com/goki/freetype/truetype"
	"github.com/slainless/digides-ogimage/pkg/fonts"
)

func drawStrings(param *Parameters) (image.Image, error) {
	const (
		debug = false

		canvasRatio = 2.5

		rootEm float64 = 16

		titleFontSize    = rootEm * 6
		subtitleFontSize = rootEm * 3
	)

	var (
		titleFontFace = truetype.NewFace(fonts.OutfitRegularFont, &truetype.Options{
			Size: titleFontSize,
		})
		subtitleFontFace = truetype.NewFace(fonts.OutfitRegularFont, &truetype.Options{
			Size: subtitleFontSize,
		})

		titleAscent    = float64(titleFontFace.Metrics().Ascent.Round())
		subtitleHeight = float64(subtitleFontFace.Metrics().Height.Round())

		canvasHeight = titleAscent + (subtitleHeight * 2)
		canvasWidth  = canvasHeight * canvasRatio
	)

	canvas := gg.NewContext(int(math.Round(canvasWidth)), int(math.Round(canvasHeight)))

	canvas.SetFontFace(titleFontFace)
	titleWidth, _ := canvas.MeasureString(param.Title)
	if titleWidth >= canvasWidth {
		return nil, errors.New("title surpasses maximum text length limit, consider using another smaller font-face or reduce string length")
	}

	canvas.SetFontFace(subtitleFontFace)
	_, _subtitleHeight := canvas.MeasureMultilineString(param.Subtitle, 0)
	if _subtitleHeight > subtitleHeight*2 {
		return nil, errors.New("subtitle surpasses maximum text length/height limit, consider using another smaller font-face or reduce string length")
	}

	if debug {
		canvas.SetColor(color.NRGBA{255, 0, 0, 20})
		canvas.DrawRectangle(0, 0, canvasWidth, titleAscent)
		canvas.Fill()
	}

	canvas.SetColor(color.NRGBA{255, 255, 255, 255})
	canvas.SetFontFace(titleFontFace)
	canvas.DrawString(param.Title, 0, titleAscent)

	if debug {
		canvas.SetColor(color.NRGBA{0, 255, 0, 20})
		canvas.DrawRectangle(0, titleFontSize, canvasWidth, subtitleHeight*2)
		canvas.Fill()
	}

	canvas.SetColor(color.NRGBA{255, 255, 255, 255})
	canvas.SetFontFace(subtitleFontFace)
	canvas.DrawStringWrapped(param.Subtitle, 0, titleAscent, 0, 0, float64(canvas.Width()), 0, gg.AlignLeft)

	// return canvas.Image(), nil

	if math.Round(_subtitleHeight) > subtitleHeight {
		return canvas.Image(), nil
	}

	result := canvas.Image()

	resizing := gift.New()
	resizing.Add(gift.CropToSize(canvas.Width(), int(math.Round(titleAscent+_subtitleHeight)+1), gift.TopLeftAnchor))
	resizing.Add(gift.Resize(canvas.Width()*2, 0, gift.LanczosResampling))

	resized := image.NewRGBA(resizing.Bounds(result.Bounds()))
	resizing.Draw(resized, result)

	return resized, nil
}
