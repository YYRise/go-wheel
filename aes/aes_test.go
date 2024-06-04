package aes

import (
	"fmt"
	"testing"
)

const (
	key = "UVBxglr6QglPqLoe9mPy+A==" //"IgkibX71IEf382PT"
	iv  = "IgkibX71IEf382PT"
)

func TestEncrypt(t *testing.T) {
	//s, _ := base64.StdEncoding.DecodeString(key)
	fmt.Println("s = ", len(key))
	t.Log(New(key, iv).Encrypt("123456"))
}

func TestDecrypt(t *testing.T) {
	t.Log(New(key, iv).Decrypt("GO+ri84zevE+z1biJwfQPw=="))
}

func BenchmarkEncryptAndDecrypt(b *testing.B) {
	b.ResetTimer()
	aes := New(key, iv)
	for i := 0; i < b.N; i++ {
		encryptString, _ := aes.Encrypt("123456")
		aes.Decrypt(encryptString)
	}
}
