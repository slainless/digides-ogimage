package ogimage

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/valyala/fastjson"
)

func DecryptCommand(key []byte, message []byte) (*Parameters, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := message[:12]
	plaintext, err := aesgcm.Open(nil, nonce, message, nil)
	if err != nil {
		return nil, err
	}

	value, err := fastjson.Parse(string(plaintext))
	if err != nil {
		return nil, err
	}

	return &Parameters{
		Title:    string(value.GetStringBytes("title")),
		Subtitle: string(value.GetStringBytes("subtitle")),
	}, nil
}
