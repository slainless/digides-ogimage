package main

import (
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/wasm"
)

func main() {
	js.Global().Set("go_draw", wasm.JsDraw)
	<-make(chan struct{})
}
