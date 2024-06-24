package wasm

import (
	"io"
	"os"
	"syscall/js"
	"testing"

	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/slainless/digides-ogimage/pkg/r2"
)

const (
	backgroundPath = "uploads/profil/73.11.02.2006/common/35ade0a022c7b566dbffdc934f4cb174.png"
	iconPath       = "uploads/online/73.11.02.2006/common/300_barru.png"
)

var payload = js.ValueOf(map[string]any{
	"title":      js.ValueOf("This is a title"),
	"subtitle":   js.ValueOf("This is subtitle"),
	"background": js.ValueOf(backgroundPath),
	"icon":       js.ValueOf(iconPath),
})

var bucket js.Value

func init() {
	js.Global().Set("go_draw", JsDraw)

	background, err := os.Open("../../assets/35ade0a022c7b566dbffdc934f4cb174.png")
	if err != nil {
		panic(err)
	}

	icon, err := os.Open("../../assets/300_barru.png")
	if err != nil {
		panic(err)
	}

	assetMapping := map[string][]byte{}
	assetMapping[backgroundPath], err = io.ReadAll(background)
	if err != nil {
		panic(err)
	}

	assetMapping[iconPath], err = io.ReadAll(icon)
	if err != nil {
		panic(err)
	}

	bucket = r2.NewMockBucket(assetMapping)
}

func BenchmarkWASMDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// TODO: change bucket to use true wrangler bucket instead of mocking bucket
		// by:
		// - editing wasm_exec_node.cjs
		// - load wrangler.getPlatformProxy()
		// - inject the env to global variable
		// - access the injected env here via js.Global().
		_, _, _ = bridge.ResolvePromise(js.Global().Get("go_draw").Invoke(payload, bucket))
	}
}
