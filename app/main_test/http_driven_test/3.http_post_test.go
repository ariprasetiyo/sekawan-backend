package httpdriventest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sekawan-backend/app/main/enum"
	"sekawan-backend/app/main/util"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestHttpPost(t *testing.T) {
	godotenv.Load()
	err := godotenv.Load(filepath.Join("../../../", ".env.example"))

	//build req data
	url := "/api/v1/test/post"
	jwtToken := GetTokenHttp()
	clientIdServerSide := clientId
	secretKeySHA256 := clientIdServerSide + clientKeyPattern + jwtToken

	requestId := uuid.New().String()
	requestTime := time.Now().UnixMilli()
	name := "ari prasetiyo"
	var bodyRequest = "{\"requestId\":\"" + requestId + "\",\"type\":\"" + enum.TYPE_REQUEST_HTTP_POST_TEST.String() + "\",\"requestTime\":" + strconv.FormatInt(requestTime, 10) + ",\"body\":{\"name\":\"" + name + "\"}}"
	signature := util.HmacSha256(secretKeySHA256, bodyRequest)

	//http request
	client := &http.Client{}
	var data = strings.NewReader(bodyRequest)
	req, err := http.NewRequest("POST", "http://localhost:8083"+url, data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Msg-Id", requestId)
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
