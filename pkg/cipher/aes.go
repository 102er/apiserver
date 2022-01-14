package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var (
	key = "1234567890abcdef"
	iv  = "xxxxxxxxxxxxxxxx"
)

func AesEncryptBase64(data []byte) (string, error) {
	_data, err := newAESCBCCipher(key, iv).encode(data)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(_data), nil
}

func AesDecodeBase64(data []byte) (string, error) {
	_data, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return "", err
	}
	return newAESCBCCipher(key, iv).decode(_data)
}

type AESCBCCipher struct {
	Key string
	IV  string
}

func newAESCBCCipher(key, iv string) *AESCBCCipher {
	return &AESCBCCipher{
		key, iv,
	}
}

// Encode 开始加密
func (a *AESCBCCipher) encode(_data []byte) ([]byte, error) {
	_key := []byte(a.Key)
	_iv := []byte(a.IV)

	_data = PKCS7Padding(_data)
	block, err := aes.NewCipher(_key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	return _data, nil
}

// AESCBCBase64Decode 开始解密
func (a *AESCBCCipher) decode(_data []byte) (string, error) {
	_key := []byte(a.Key)
	_iv := []byte(a.IV)

	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	_data = PKCS7UnPadding(_data)

	return string(_data), nil
}

func PKCS7Padding(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
