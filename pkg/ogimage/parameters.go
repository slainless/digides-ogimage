package ogimage

import (
	"image"
)

type Parameters struct {
	Title    string
	Subtitle string

	Icon       image.Image
	Background image.Image

	RawFontTitle    []byte
	RawFontSubtitle []byte
}
