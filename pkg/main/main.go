package main

import (
	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/teamortix/golang-wasm/wasm"
)

func main() {
	wasm.Expose("draw", bridge.Start)
	wasm.Ready()
	<-make(chan struct{})
}
