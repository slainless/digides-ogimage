package bridge

import (
	"errors"

	"github.com/slainless/digides-ogimage/pkg/ogimage"
	"github.com/spf13/cast"
)

func LoadParameters(input map[string]any) (*ogimage.Parameters, error) {
	if input["background"] == nil || input["icon"] == nil {
		return nil, errors.New("No background or icon given")
	}

	icon, background, err := DecodeImages(cast.ToString(input["icon"]), cast.ToString(input["background"]))
	if err != nil {
		return nil, err
	}

	parameters := &ogimage.Parameters{
		Title:      cast.ToString(input["title"]),
		Subtitle:   cast.ToString(input["subtitle"]),
		Icon:       icon,
		Background: background,
	}

	return parameters, nil
}
