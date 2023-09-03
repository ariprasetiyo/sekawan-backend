package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

const (
	// keySysBit 32 for AES 256
	// keySysBit 16 for AES 128
	// source https://go.dev/src/crypto/cipher/example_test.go
	keySysBit string = "jb(HH}=#jA=%6QK7"
)

func EncryptAES256(planText string) string {
	plantTextInByte := []byte(planText)
	keyInByte := []byte(keySysBit)
	c, err := aes.NewCipher(keyInByte)
	IsErrorDoPrintWithMessage("error encryption new chiper "+planText, err)
	gcm, err := cipher.NewGCM(c)
	IsErrorDoPrintWithMessage("error encryption new gcm "+planText, err)
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	encodeResult := base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, plantTextInByte, nil))
	return encodeResult
}

func DecryptAES256(cipherText string) string {
	keyInByte := []byte(keySysBit)
	ciphertext, err := base64.URLEncoding.DecodeString((cipherText))
	IsErrorDoPrintWithMessage("error encryption new gcm "+cipherText, err)

	c, err := aes.NewCipher(keyInByte)
	IsErrorDoPrintWithMessage("error decrypt new NewCipher "+cipherText, err)
	gcm, err := cipher.NewGCM(c)
	IsErrorDoPrintWithMessage("error decrypt new NewGCM "+cipherText, err)

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		IsErrorDoPrintWithMessage("error decrypt nonceSize "+cipherText, err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		IsErrorDoPrintWithMessage("error decrypt gcm.Open "+cipherText, err)
	}
	return string(plaintext)
}
