package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// keySysBit 32 for AES 256
// keySysBit 16 for AES 128 ex.  "jb(HH}=#jA=%6QK7"
// source https://go.dev/src/crypto/cipher/example_test.go
func EncryptAES256(key string, planText string) string {
	plantTextInByte := []byte(planText)
	keyInByte := []byte(key)
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

func DecryptAES256(key string, cipherText string) string {
	keyInByte := []byte(key)
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

func HmacSha256(secret string, data string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	h.Write([]byte(data))
	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}

func HmacSha256InByte(secret string, data []byte) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	h.Write(data)
	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}
