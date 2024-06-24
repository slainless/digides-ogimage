package r2

import (
	"bytes"
	"errors"
	"image"
	"sync"
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

func GetFile(r2Path string, bucket js.Value) ([]byte, error) {
	response, jsErr, _ := bridge.ResolvePromise(bucket.Call("get", r2Path))
	if jsErr != nil {
		return nil, bridge.FromJsError(*jsErr)
	}

	if bridge.IsNullish(*response) {
		return nil, ErrFileNotFound
	}

	// Legacy code using ReadableStream, after benchmarking
	// it seems that simpler is best...
	// ---
	// jsReadableStream := response.Get("body")
	// if !jsReadableStream.InstanceOf(js.Global().Get("ReadableStream")) {
	// 	return nil, ErrInvalidStream
	// }

	// readableStream, err := NewReadableStream(jsReadableStream)
	// if err != nil {
	// 	return nil, err
	// }

	// return readableStream, nil
	// ---

	jsArrayBuffer, jsErr, _ := bridge.ResolvePromise(response.Call("arrayBuffer"))
	if jsErr != nil {
		return nil, bridge.FromJsError(*jsErr)
	}

	jsBuffer := js.Global().Get("Uint8Array").New(*jsArrayBuffer)
	buffer := make([]byte, jsBuffer.Get("byteLength").Int())
	js.CopyBytesToGo(buffer, jsBuffer)
	return buffer, nil
}

type R2Parameters interface {
	R2PathIcon() string
	R2PathBackground() string
}

func GetImages(iconPath string, backgroundPath string, bucket js.Value) (icon image.Image, background image.Image, err error) {
	wg := sync.WaitGroup{}
	var errIcon, errBg error

	wg.Add(1)
	go func() {
		defer wg.Done()

		var data []byte
		data, errIcon = GetFile(iconPath, bucket)
		if errIcon != nil {
			return
		}

		icon, _, errIcon = image.Decode(bytes.NewReader(data))
		if errIcon != nil {
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		var data []byte
		data, errBg = GetFile(backgroundPath, bucket)
		if errBg != nil {
			return
		}

		background, _, errBg = image.Decode(bytes.NewReader(data))
		if errBg != nil {
			return
		}
	}()

	wg.Wait()
	if errIcon != nil || errBg != nil {
		return nil, nil, errors.Join(errIcon, errBg)
	}

	return icon, background, nil
}
