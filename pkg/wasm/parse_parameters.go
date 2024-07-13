package wasm

import (
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/fonts"
	"github.com/slainless/digides-ogimage/pkg/r2"
)

func parseParameters(input js.Value) (*Parameters, error) {
	if input.Type() != js.TypeObject {
		return nil, ErrParametersInvalid
	}

	title := input.Get("title")
	subtitle := input.Get("subtitle")
	icon := input.Get("icon")
	background := input.Get("background")
	fontTitle := input.Get("titleFont")
	fontSubtitle := input.Get("subtitleFont")

	return &Parameters{
		title:            title.String(),
		subtitle:         subtitle.String(),
		R2PathIcon:       icon.String(),
		R2PathBackground: background.String(),
		FontTitle:        fontTitle.String(),
		FontSubtitle:     fontSubtitle.String(),
	}, nil
}

func LoadParameters(parameters js.Value, bucket js.Value) (*Parameters, error) {
	if bucket.Type() != js.TypeObject {
		return nil, ErrBucketNotFound
	}

	params, err := parseParameters(parameters)
	if err != nil {
		return nil, err
	}

	params.icon, params.background, err = r2.GetImages(params.R2PathIcon, params.R2PathBackground, bucket)
	if err != nil {
		return nil, err
	}

	params.fontFaceTitle = fonts.OutfitRegularFont
	params.fontFaceSubtitle = fonts.OutfitRegularFont

	return params, nil
}
