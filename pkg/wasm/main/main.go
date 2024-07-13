package main

import (
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/slainless/digides-ogimage/pkg/wasm"
)

func main() {
	bridge.Console().Log("WASM Loaded.")
	bridge.Console().Log("Injecting go_draw to global")
	js.Global().Set("go_draw", wasm.JsDraw)
	bridge.Console().Log("Go draw injected!", js.Global().Get("go_draw"))
	<-make(chan struct{})
}
