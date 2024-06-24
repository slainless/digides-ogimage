package wasm

import (
	"errors"
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
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

	if bridge.IsString(title) == false || bridge.Falsey(title) {
		return nil, errors.Join(ErrParametersInvalidField, errors.New("title must not be empty"))
	}

	if bridge.IsString(subtitle) == false || bridge.Falsey(subtitle) {
		return nil, errors.Join(ErrParametersInvalidField, errors.New("subtitle must not be empty"))
	}

	if bridge.IsString(icon) == false || bridge.Falsey(icon) {
		return nil, errors.Join(ErrParametersInvalidField, errors.New("icon must not be empty"))
	}

	if bridge.IsString(background) == false || bridge.Falsey(background) {
		return nil, errors.Join(ErrParametersInvalidField, errors.New("background must not be empty"))
	}

	if bridge.IsNullish(fontTitle) == false && bridge.IsString(fontTitle) == false {
		return nil, errors.Join(ErrParametersInvalidField, errors.New("titleFont must be a string"))
	}

	if bridge.IsNullish(fontSubtitle) == false && bridge.IsString(fontSubtitle) == false {
		return nil, errors.Join(ErrParametersInvalidField, errors.New("subtitleFont must be a string"))
	}

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
