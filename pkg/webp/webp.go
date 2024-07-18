package webp

import (
	"context"
	_ "embed"
	"errors"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

var (
	ErrInvalidWeBPModule = errors.New("invalid webp wasm module")
)

var (
	//go:embed webp.generated.wasm
	webpWasm []byte
	// compiledWebpWasm wazero.CompiledModule
)

func NewModule(ctx context.Context, runtime wazero.Runtime) (*Module, error) {
	// mod, err := zero.Runtime.InstantiateModule(zero.RuntimeCtx, compiledWebpWasm, wazero.NewModuleConfig())
	mod, err := runtime.InstantiateWithConfig(ctx, webpWasm, wazero.NewModuleConfig())
	if err != nil {
		return nil, err
	}

	switch {
	case
		mod.ExportedFunction("encode") == nil,
		mod.ExportedFunction("version") == nil,
		mod.ExportedFunction("free_result") == nil:
		return nil, ErrInvalidWeBPModule
	}

	return &Module{module: mod}, nil
}

type Module struct {
	module api.Module
}
