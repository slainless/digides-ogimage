package bridge

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
)

func DecodeImages(icon string, background string) (_icon image.Image, _background image.Image, err error) {
	ch := make(chan error)
	go func() {
		data, err := base64.StdEncoding.DecodeString(icon)
		if err != nil {
			ch <- err
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		_icon = img
		ch <- err
	}()

	go func() {
		data, err := base64.StdEncoding.DecodeString(background)
		if err != nil {
			ch <- err
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		_background = img
		ch <- err
	}()

	errs := []error{}
	for i := 0; i < 2; i++ {
		errs = append(errs, <-ch)
	}

	close(ch)
	return _icon, _background, join(errs)
}

func join(errs []error) error {
	filteredErrs := []error{}
	for _, e := range errs {
		if e == nil {
			continue
		}

		filteredErrs = append(filteredErrs, e)
	}

	return errors.Join(filteredErrs...)
}
