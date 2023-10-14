package httpdriventest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sekawan-backend/app/main/enum"
	"sekawan-backend/app/main/util"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestHttpGet(t *testing.T) {

	godotenv.Load()
	err := godotenv.Load(filepath.Join("../../../", ".env.example"))

	url := "/api/v1/test/get?name=test&type=" + enum.TYPE_REQUEST_HTTP_GET_TEST.String()
	jwtToken := GetTokenHttp()
	clientIdServerSide := clientId
	secretKeySHA256 := clientIdServerSide + clientKeyPattern + jwtToken
	signature := util.HmacSha256(secretKeySHA256, url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8083"+url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Msg-Id", uuid.New().String())
	req.Header.Set("Client-Id", clientId)
	req.Header.Set("Signature", signature)
	req.Header.Set("Authorization", jwtToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
