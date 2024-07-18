package webp

import (
	"context"
	"errors"
	"image"

	"github.com/tetratelabs/wazero/api"
)

var (
	ErrMemoryException = errors.New("memory I/O failed")
)

func (m *Module) createBuffer(ctx context.Context, data []byte, width, height int) (uintptr, error) {
	buf, err := m.module.ExportedFunction("create_buffer").Call(
		ctx,
		api.EncodeI64(int64(width)),
		api.EncodeI64(int64(height)),
	)
	if err != nil {
		return 0, err
	}

	start := api.DecodeExternref(buf[0])
	ok := m.module.Memory().Write(uint32(start), data)
	if !ok {
		return 0, ErrMemoryException
	}

	return start, nil
}

func (m *Module) freeBuffer(ctx context.Context, buf uintptr) error {
	_, err := m.module.ExportedFunction("free_buffer").Call(ctx, api.EncodeExternref(buf))
	return err
}

func (m *Module) freeResult(ctx context.Context, result uintptr) error {
	_, err := m.module.ExportedFunction("free_result").Call(ctx, api.EncodeExternref(result))
	return err
}

func (m *Module) encode(ctx context.Context, buffer uintptr, width, height int, quality float32) (buf []byte, err error) {
	result, err := m.module.ExportedFunction("encode").Call(
		ctx,
		api.EncodeExternref(buffer),
		api.EncodeU32(uint32(width)),
		api.EncodeU32(uint32(height)),
		api.EncodeF32(quality),
	)
	defer m.freeBuffer(ctx, buffer)

	start := api.DecodeExternref(result[0])
	defer m.freeResult(ctx, start)

	bufPtr, ok := m.module.Memory().ReadUint32Le(uint32(start))
	if !ok {
		return nil, ErrMemoryException
	}

	sz, ok := m.module.Memory().ReadUint32Le(uint32(start) + 32)
	if !ok {
		return nil, ErrMemoryException
	}

	resultBuffer, ok := m.module.Memory().Read(bufPtr, sz)
	if !ok {
		return nil, ErrMemoryException
	}

	return resultBuffer, nil
}

func (m *Module) Encode(ctx context.Context, img *image.RGBA, quality float32) ([]byte, error) {
	bound := img.Bounds()
	width := bound.Dx()
	height := bound.Dy()

	buf, err := m.createBuffer(ctx, img.Pix, bound.Dx(), bound.Dy())
	if err != nil {
		return nil, err
	}

	// buffer should be freed by encode
	result, err := m.encode(ctx, buf, width, height, 1)
	if err != nil {
		return nil, err
	}

	return result, nil
}
