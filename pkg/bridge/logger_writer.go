package bridge

import (
	"io"
)

type JSLogWriter struct {
	logger func(args ...any)
}

func (w *JSLogWriter) Write(data []byte) (int, error) {
	w.logger(string(data))
	return len(data), nil
}

func NewWriterFrom(logger func(args ...any)) io.Writer {
	return &JSLogWriter{
		logger: logger,
	}
}
