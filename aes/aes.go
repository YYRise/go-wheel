package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var _ Aes = (*aesCbcPkcs7)(nil)

type Aes interface {
	// Encrypt 加密
	Encrypt(encryptStr string) (string, error)

	// Decrypt 解密
	Decrypt(decryptStr string) (string, error)
}

type aesCbcPkcs7 struct {
	key string // key must be 16, 24 or 32 bytes
	iv  string // iv must be 16 bytes (ECB mode doesn't require setting iv)
}

func New(key, iv string) Aes {
	return &aesCbcPkcs7{
		key: key,
		iv:  iv,
	}
}

func (a *aesCbcPkcs7) getIv(blockSize int) string {
	if len(a.iv) != blockSize && len(a.key) >= blockSize {
		a.iv = a.key[:blockSize]
	}
	return a.iv
}

func (a *aesCbcPkcs7) Encrypt(encryptStr string) (string, error) {
	encryptBytes := []byte(encryptStr)
	block, err := aes.NewCipher([]byte(a.key))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encryptBytes = pkcs7Padding(encryptBytes, blockSize)
	iv := a.getIv(blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(encryptBytes))
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (a *aesCbcPkcs7) Decrypt(decryptStr string) (string, error) {
	decryptBytes, err := base64.StdEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(a.key))
	if err != nil {
		return "", err
	}
	iv := a.getIv(block.BlockSize())
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	decrypted := make([]byte, len(decryptBytes))

	blockMode.CryptBlocks(decrypted, decryptBytes)
	decrypted = pkcs7UnPadding(decrypted)
	return string(decrypted), nil
}

// PKCS7Padding 补充到blockSize
func pkcs7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS7UnPadding 移除填充
func pkcs7UnPadding(decrypted []byte) []byte {
	length := len(decrypted)
	unPadding := int(decrypted[length-1])
	return decrypted[:(length - unPadding)]
}
