package util

import (
	"crypto/md5"
	"encoding/base64"
)

func GeneratePasswordHash(userId string, planText string, salt string) string {
	resultHash := md5.Sum([]byte(passwordWithSalt(userId, planText, salt)))
	return base64.StdEncoding.EncodeToString(resultHash[:])
}

func passwordWithSalt(userId string, planText string, salt string) string {
	return salt + userId + planText
}
