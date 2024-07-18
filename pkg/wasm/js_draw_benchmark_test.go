package wasm

import (
	"syscall/js"
	"testing"

	"github.com/slainless/digides-ogimage/pkg/bridge"
)

const (
	backgroundPath = "uploads/profil/73.11.02.2006/common/35ade0a022c7b566dbffdc934f4cb174.png"
	iconPath       = "uploads/online/73.11.02.2006/common/300_barru.png"
)

var payload = js.ValueOf(map[string]any{
	"title":      "This is a title",
	"subtitle":   "This is subtitle",
	"background": backgroundPath,
	"icon":       iconPath,
})

var bucket js.Value

func init() {
	js.Global().Set("go_draw", NewJSDraw(nil))

	platform, jsErr, err := bridge.ResolvePromise(
		js.Global().Get("require").Invoke("wrangler").Get("getPlatformProxy").
			Invoke(js.ValueOf(map[string]any{
				"persist": js.ValueOf(map[string]any{
					"path": "../../.wrangler/state/v3",
				}),
			})))

	if err != nil {
		panic(err)
	}

	if jsErr != nil {
		panic(bridge.FromJsError(*jsErr))
	}

	bucket = platform.Get("env").Get("R2_ASSETS")
}

func BenchmarkWASMDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = bridge.ResolvePromise(js.Global().Get("go_draw").Invoke(payload, bucket))
	}
}
