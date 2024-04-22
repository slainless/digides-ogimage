package bridge_test

import (
	"encoding/base64"
	"io"
	"os"
	"testing"

	"github.com/slainless/digides-ogimage/pkg/bridge"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func MockParameters() map[string]any {
	params := map[string]any{
		"title":    "This is a title",
		"subtitle": "This is subtitle",
	}

	background, err := os.Open("../../assets/35ade0a022c7b566dbffdc934f4cb174.png")
	panicIf(err)

	icon, err := os.Open("../../assets/300_barru.png")
	panicIf(err)

	_icon, err := io.ReadAll(icon)
	panicIf(err)

	_background, err := io.ReadAll(background)
	panicIf(err)

	params["icon"] = base64.StdEncoding.EncodeToString(_icon)
	params["background"] = base64.StdEncoding.EncodeToString(_background)

	return params
}

var parameters = MockParameters()

func BenchmarkStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = bridge.Start(parameters)
	}
}
