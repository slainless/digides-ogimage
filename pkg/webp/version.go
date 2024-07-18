package webp

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero/api"
)

func (m *Module) Version(ctx context.Context) (string, error) {
	result, err := m.module.ExportedFunction("version").Call(ctx)
	if err != nil {
		return "", err
	}

	version := api.DecodeI32(result[0])
	return fmt.Sprintf("v%v.%v.%v", version&0xFF0000>>16, version&0x00FF00>>8, version&0x0000FF), nil
}
