package unittest

import (
	"os"
	"path/filepath"
	"sekawan-backend/app/main/enum"
	"sekawan-backend/app/main/handlerAuth"
	"sekawan-backend/app/main/util"
	"strconv"
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type OS interface {
	Getenv(string) string
}

type defaultOS struct{}

func (defaultOS) Getenv(key string) string {
	return os.Getenv(key)
}

func TestCurlRequestGenerateToken(t *testing.T) {
	godotenv.Load()
	err := godotenv.Load(filepath.Join("../../", ".env.example"))
	util.IsErrorDoPrint(err)

	clientIdServerSide := os.Getenv(util.CONFIG_APP_CLIENT_ID)
	clientApiKeyServerSide := os.Getenv(util.CONFIG_APP_CLIENT_API_KEY_PASSWORD)
	secretKeySHA256 := clientIdServerSide + "::" + clientApiKeyServerSide
	encryptionKey := os.Getenv(util.CONFIG_APP_ENCRIPTION_KEY)

	phoneNo := "+6285600070411"
	password := "cobacoba1-="
	authCredRequest := handlerAuth.AuthCredRequest{PhoneNo: phoneNo, Password: password}
	authCredRequestInJson, err := json.Marshal(authCredRequest)
	util.IsErrorDoPrint(err)
	encryptionCredential := util.EncryptAES256(encryptionKey, string(authCredRequestInJson))
	requestId := uuid.New().String()
	requestTime := time.Now().UnixMilli()

	var bodyRequest = "{\"requestId\":\"" + requestId + "\",\"type\":\"" + enum.TYPE_GENERATE_TOKEN.String() + "\",\"requestTime\":" + strconv.FormatInt(requestTime, 10) + ",\"body\":{\"cred\":\"" + encryptionCredential + "\"}}"
	signature := util.HmacSha256(secretKeySHA256, bodyRequest)
	curl := "curl -X POST localhost:8083/public/token -H 'Msg-Id: " + requestId + "' -H 'Client-Id: " + clientIdServerSide + "' -H 'Signature: " + signature + "' -d '" + bodyRequest + "'"
	println(curl)
}
