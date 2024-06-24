package reader_test

import (
	"io"
	"syscall/js"
	"testing"
)

func BenchmarkNativeDefaultReader512(b *testing.B) {
	buffer := randomBuffer(512)
	for i := 0; i < b.N; i++ {
		resultingJsBuffer, err := nativeReading(buffer)
		if err != nil {
			b.Fatal(err)
		}

		resultingBuffer := make([]byte, len(buffer))
		js.CopyBytesToGo(resultingBuffer, *resultingJsBuffer)
	}
}

func BenchmarkDefaultReader512(b *testing.B) {
	buffer := randomBuffer(512)
	for i := 0; i < b.N; i++ {
		reader, err := prepareReader(buffer)
		if err != nil {
			b.Fatal(err)
		}

		_, err = io.ReadAll(reader)
		if err != nil {
			b.Fatal(err)
		}
	}
}
