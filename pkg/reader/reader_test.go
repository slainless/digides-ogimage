package reader_test

import (
	"crypto/rand"
	"io"
	"syscall/js"
	"testing"

	"github.com/slainless/digides-ogimage/pkg/bridge"
	"github.com/slainless/digides-ogimage/pkg/reader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var Console = bridge.Console

func randomBuffer(size int) []byte {
	res := make([]byte, size)
	_, err := rand.Read(res)
	if err != nil {
		panic(err)
	}
	return res
}

func prepareReader(buffer []byte) (io.Reader, error) {
	jsBuffer := js.Global().Get("Uint8Array").New(len(buffer))
	js.CopyBytesToJS(jsBuffer, buffer)

	jsReader := js.Global().Get("ReadableStream").Call("from", jsBuffer)
	return reader.NewReadableStream(jsReader)
}

func nativeReading(buffer []byte) (*js.Value, error) {
	jsBuffer := js.Global().Get("Uint8Array").New(js.ValueOf(len(buffer)))
	js.CopyBytesToJS(jsBuffer, buffer)

	stream := js.Global().Get("ReadableStream").Call("from", jsBuffer)
	reader := stream.Call("getReader")

	resultingBuffer := js.Global().Get("Uint8Array").New(js.ValueOf(len(buffer)))

	for i := 0; i < len(buffer); i++ {
		value, jsErr, _ := bridge.ResolvePromise(reader.Call("read", jsBuffer))
		if jsErr != nil {
			return nil, bridge.FromJsError(*jsErr)
		}

		if value.Get("done").Truthy() {
			break
		}

		resultingBuffer.SetIndex(i, value.Get("value"))
	}

	return &resultingBuffer, nil
}

func TestDefaultReader64(t *testing.T) {
	buffer := randomBuffer(64)
	reader, err := prepareReader(buffer)
	require.NoError(t, err)

	returnedBuffer, err := io.ReadAll(reader)
	require.NoError(t, err)
	assert.Equal(t, buffer, returnedBuffer)
}

func TestDefaultReader256(t *testing.T) {
	buffer := randomBuffer(256)
	reader, err := prepareReader(buffer)
	require.NoError(t, err)

	returnedBuffer, err := io.ReadAll(reader)
	require.NoError(t, err)
	assert.Equal(t, buffer, returnedBuffer)
}

func TestDefaultReader512(t *testing.T) {
	buffer := randomBuffer(512)
	reader, err := prepareReader(buffer)
	require.NoError(t, err)

	returnedBuffer, err := io.ReadAll(reader)
	require.NoError(t, err)
	assert.Equal(t, buffer, returnedBuffer)
}

func TestNativeDefaultReader512(t *testing.T) {
	buffer := randomBuffer(512)
	resultingJsBuffer, err := nativeReading(buffer)
	require.NoError(t, err)

	resultingBuffer := make([]byte, len(buffer))
	js.CopyBytesToGo(resultingBuffer, *resultingJsBuffer)
	assert.Equal(t, buffer, resultingBuffer)
}

func TestDefaultReader1024(t *testing.T) {
	buffer := randomBuffer(1024)
	reader, err := prepareReader(buffer)
	require.NoError(t, err)

	returnedBuffer, err := io.ReadAll(reader)
	require.NoError(t, err)
	assert.Equal(t, buffer, returnedBuffer)
}
