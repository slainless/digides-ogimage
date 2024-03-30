package bridge

import (
	"github.com/slainless/digides-ogimage/pkg/ogimage"
	"github.com/spf13/cast"
)

func LoadParameters(input map[string]any) *ogimage.Parameters {
	return &ogimage.Parameters{
		Title:    cast.ToString(input["title"]),
		Subtitle: cast.ToString(input["subtitle"]),
	}
}
