package zero

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func NewCFWasmRuntime() (wazero.Runtime, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	runtime := wazero.NewRuntimeWithConfig(
		context.TODO(),
		wazero.NewRuntimeConfigInterpreter(),
	)
	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

	return runtime, ctx, cancel
}
