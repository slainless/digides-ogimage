package reader

import "syscall/js"

func getReader(stream js.Value) (reader *js.Value, isBYOB bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			_reader := stream.Call("getReader")
			if !_reader.InstanceOf(js.Global().Get("ReadableStreamDefaultReader")) {
				reader = nil
				isBYOB = false
				err = ErrInvalidStream
				return
			}

			reader = &_reader
			isBYOB = false
			err = nil
		}
	}()

	_reader := stream.Call("getReader", js.ValueOf(map[string]any{
		"mode": "byob",
	}))

	if !_reader.InstanceOf(js.Global().Get("ReadableStreamBYOBReader")) {
		return nil, false, ErrInvalidStream
	}

	return reader, true, nil
}
