package wasm

import (
	"errors"
	"image"

	"github.com/goki/freetype/truetype"
)

var (
	ErrParametersInvalid      = errors.New("Invalid parameters")
	ErrParametersInvalidField = errors.New("Invalid parameters field")

	ErrInvalidCloudflareEnv = errors.New("Invalid cloudflare env")
	ErrBucketNotFound       = errors.New("Bucket not found")
)

type Parameters struct {
	title            string
	subtitle         string
	icon             image.Image
	background       image.Image
	fontFaceTitle    *truetype.Font
	fontFaceSubtitle *truetype.Font

	R2PathIcon       string
	R2PathBackground string

	FontTitle    string
	FontSubtitle string
}

func (p *Parameters) Title() string {
	return p.title
}

func (p *Parameters) Subtitle() string {
	return p.subtitle
}

func (p *Parameters) Icon() image.Image {
	return p.icon
}

func (p *Parameters) Background() image.Image {
	return p.background
}

func (p *Parameters) FontFaceTitle() *truetype.Font {
	return p.fontFaceTitle
}

func (p *Parameters) FontFaceSubtitle() *truetype.Font {
	return p.fontFaceSubtitle
}

func (p *Parameters) SetTitle(title string) *Parameters {
	p.title = title
	return p
}

func (p *Parameters) SetSubtitle(subtitle string) *Parameters {
	p.subtitle = subtitle
	return p
}

func (p *Parameters) SetIcon(icon image.Image) *Parameters {
	p.icon = icon
	return p
}

func (p *Parameters) SetBackground(background image.Image) *Parameters {
	p.background = background
	return p
}

func (p *Parameters) SetFontFaceTitle(fontFaceTitle *truetype.Font) *Parameters {
	p.fontFaceTitle = fontFaceTitle
	return p
}

func (p *Parameters) SetFontFaceSubtitle(fontFaceSubtitle *truetype.Font) *Parameters {
	p.fontFaceSubtitle = fontFaceSubtitle
	return p
}
