package ogimage_test

import (
	_ "embed"
	"image"
	"os"
	"testing"

	"github.com/slainless/digides-ogimage/pkg/ogimage"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func LoadParameters() *ogimage.Parameters {
	params := &ogimage.Parameters{
		Title:    "This is a title",
		Subtitle: "This is subtitle",
	}

	background, err := os.Open("../../assets/35ade0a022c7b566dbffdc934f4cb174.png")
	panicIf(err)

	icon, err := os.Open("../../assets/300_barru.png")
	panicIf(err)

	params.Background, _, _ = image.Decode(background)
	params.Icon, _, _ = image.Decode(icon)

	return params
}

var parameters = LoadParameters()

func BenchmarkDirectDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ogimage.Draw(parameters)
	}
}
