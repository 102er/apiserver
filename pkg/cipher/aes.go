package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var (
	key = "1234567890abcdef"
	iv  = "ae0494f008333659"
)

func AesEncryptNoBase64(data []byte) (string, error) {
	return newAESCBCCipher(key, iv).AESCBCBase64Encode(data)
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

// AESCBCBase64Encode 开始加密
func (a *AESCBCCipher) AESCBCBase64Encode(_data []byte) (string, error) {
	_key := []byte(a.Key)
	_iv := []byte(a.IV)

	_data = PKCS7Padding(_data)
	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	return base64.StdEncoding.EncodeToString(_data), nil
}

// AESCBCBase64Decode 开始解密
func (a *AESCBCCipher) AESCBCBase64Decode(data string) (string, error) {
	_data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
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
