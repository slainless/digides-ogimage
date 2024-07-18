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
)

func NewModule(ctx context.Context, runtime wazero.Runtime, config wazero.ModuleConfig) (*Module, error) {
	mod, err := runtime.InstantiateWithConfig(ctx, webpWasm, config)
	if err != nil {
		return nil, err
	}

	switch {
	case
		mod.ExportedFunction("encode") == nil,
		mod.ExportedFunction("version") == nil,
		mod.ExportedFunction("create_buffer") == nil,
		mod.ExportedFunction("free_buffer") == nil,
		mod.ExportedFunction("free_result") == nil:
		return nil, ErrInvalidWeBPModule
	}

	return &Module{module: mod}, nil
}

type Module struct {
	module api.Module
}
