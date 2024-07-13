package reader

import (
	"io"
	"syscall/js"

	"github.com/slainless/digides-ogimage/pkg/bridge"
)

var (
	ErrInvalidStream        = bridge.NewJSTypeError("StreamError", "not a valid readable stream")
	ErrInvalidReadingResult = bridge.NewJSTypeError("StreamError", "invalid reading result: either data returned as null, non-object, or invalid object")
)

type ReadableStream struct {
	reader js.Value
	isBYOB bool
	buffer []byte
}

func NewReadableStream(jsReadableStream js.Value) (*ReadableStream, error) {
	// TODO: add recovery mechanism

	if jsReadableStream.Type() != js.TypeObject {
		return nil, ErrInvalidStream
	}

	if jsReadableStream.Get("getReader").Type() != js.TypeFunction {
		return nil, ErrInvalidStream
	}

	reader, isBYOB, err := getReader(jsReadableStream)
	if err != nil {
		return nil, err
	}

	return &ReadableStream{reader: *reader, isBYOB: isBYOB}, nil
}

func (r *ReadableStream) readBYOB(p []byte) (n int, err error) {
	value, jsErr, _ := bridge.ResolvePromise(r.reader.Call("read", js.Global().Get("Uint8Array").New(len(p))))
	if jsErr != nil {
		return 0, bridge.FromJsError(*jsErr)
	}

	if value == nil {
		return 0, ErrInvalidReadingResult
	}

	if value.Type() != js.TypeObject {
		return 0, ErrInvalidReadingResult
	}

	if value.Get("done").Truthy() {
		return 0, io.EOF
	}

	bytes := value.Get("value")
	if !bytes.InstanceOf(js.Global().Get("Uint8Array")) {
		return 0, ErrInvalidReadingResult
	}

	n = bytes.Get("byteLength").Int()
	js.CopyBytesToGo(p, bytes)
	return n, nil
}

func (r *ReadableStream) pull() (bytes *js.Value, n int, err error) {
	value, jsErr, _ := bridge.ResolvePromise(r.reader.Call("read"))
	if jsErr != nil {
		return nil, 0, bridge.FromJsError(*jsErr)
	}

	if value == nil {
		return nil, 0, ErrInvalidReadingResult
	}

	if value.Type() != js.TypeObject {
		return nil, 0, ErrInvalidReadingResult
	}

	if value.Get("done").Truthy() {
		return nil, 0, io.EOF
	}

	_bytes := value.Get("value")
	if _bytes.Type() == js.TypeNumber {
		return &_bytes, 1, nil
	}

	if !_bytes.InstanceOf(js.Global().Get("Uint8Array")) {
		return nil, 0, ErrInvalidReadingResult
	}

	n = _bytes.Get("byteLength").Int()
	return &_bytes, n, nil
}

func (r *ReadableStream) pullUntil(desiredSize int) ([]byte, error) {
	if len(r.buffer) >= desiredSize {
		view := r.buffer[:desiredSize]
		r.buffer = r.buffer[desiredSize:]
		return view, nil
	}

	for {
		bytes, n, err := r.pull()
		if err != nil {

			if err == io.EOF {
				// if buffer is not empty, then return it
				// and let the next pull to return an EOF
				if r.buffer != nil {
					view := r.buffer
					r.buffer = nil
					return view, nil
					// if buffer is empty, then return an EOF
				} else {
					return nil, err
				}
			}

			return nil, err
		}

		newBuffer := make([]byte, len(r.buffer)+n)
		copy(newBuffer, r.buffer)

		receivedBytes := *bytes
		if receivedBytes.Type() == js.TypeNumber {
			newBuffer[len(r.buffer)] = byte(receivedBytes.Int())
		} else {
			js.CopyBytesToGo(newBuffer[len(r.buffer):], *bytes)
		}

		// if newBuffer size beyond desiredSize
		// then return the desiredSize bytes from newBuffer
		// and set the buffer to be the rest of the newBuffer
		if len(newBuffer) > desiredSize {
			r.buffer = newBuffer[:desiredSize]
			return newBuffer[:desiredSize], nil
			// else if its equal, then return the newBuffer directly and
			// nil the buffer
		} else if len(r.buffer) == desiredSize {
			r.buffer = nil
			return newBuffer, nil
			// else, set the new buffer then pull again
		} else {
			r.buffer = newBuffer
		}
	}
}

func (r *ReadableStream) read(p []byte) (n int, err error) {
	bytes, err := r.pullUntil(len(p))
	if err != nil {
		return 0, err
	}

	copy(p, bytes)
	return len(bytes), nil
}

func (r *ReadableStream) Read(p []byte) (n int, err error) {
	if r.isBYOB {
		return r.readBYOB(p)
	} else {
		return r.read(p)
	}
}
