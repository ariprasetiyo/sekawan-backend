package util

import (
	"crypto/md5"
	"encoding/base64"
	"strings"
)

func GeneratePasswordHash(salt string, planTexts ...string) string {
	var planText strings.Builder
	for _, text := range planTexts {
		planText.WriteString(text)
	}
	return GenerateMD5(salt, planText.String())
}

func GenerateMD5(salt string, planText string) string {
	resultHash := md5.Sum([]byte(planTextWithSalt(planText, salt)))
	return base64.StdEncoding.EncodeToString(resultHash[:])
}

func passwordWithSalt(userId string, planText string, salt string) string {
	return salt + userId + planText
}

func planTextWithSalt(planText string, salt string) string {
	return salt + planText
}
