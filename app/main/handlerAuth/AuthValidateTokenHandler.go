package handlerAuth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewAuthValidateTokenHandler(db *repository.Database) handler.HandlerInterface {
	return &AuthValidateToken{databaseImpl: *db}
}

type AuthValidateToken struct {
	databaseImpl repository.Database
}

func (auth AuthValidateToken) Execute(c *gin.Context) {

	msgId := c.GetHeader(util.HEADER_MSG_ID)
	signatureInReq := c.GetHeader(util.HEADER_SIGNATURE)
	httpMethod := c.Request.Method
	sourceUrl := c.Request.URL.String()
	authorization := c.GetHeader(util.HEADER_AUTHORIZATION)
	jsonRequestBody := getBodyRequest(c, msgId)

	if util.IsEmptyObject(msgId) || util.IsEmptyString(msgId) {
		logrus.Infoln("invalid request msg id", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)
		unauthorized(c)
		return
	} else if util.IsEmptyString(signatureInReq) || !isValidSignature(signatureInReq, jsonRequestBody) {
		logrus.Infoln("invalid siganture", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)
		unauthorized(c)
		return
	} else if util.IsEmptyString(msgId) && util.IsEmptyString(authorization) &&
		util.IsEmptyString(httpMethod) && util.IsEmptyString(sourceUrl) &&
		util.IsEmptyObject(jsonRequestBody) {
		logrus.Info("invalid request", msgId, "authorization:", authorization, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl, " request body:", jsonRequestBody)
		unauthorized(c)
		return
	}

	jwtToken := decodeJWTToken(authorization)
	if !isValidToken(jwtToken) {
		logrus.Infoln("invalid JWT", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)
		unauthorized(c)
		return
	} else if jwtToken.Body.ExpiredTs < time.Now().UnixMilli() {
		logrus.Infoln("expired token authorization", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)
		unauthorized(c)
		return
	}

}

func getBodyRequest(c *gin.Context, msgId string) []byte {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorln("error read request body", msgId, err.Error())
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
	return jsonData
}

func decodeJWTToken(jwtToken string) JWTToken {
	decodeResult, err := base64.StdEncoding.DecodeString(jwtToken)
	util.IsErrorDoPrintWithMessage("error decoding jwt token", err)

	var jwtTokenUnmarshal JWTToken
	error := json.Unmarshal([]byte(decodeResult), &jwtTokenUnmarshal)
	util.IsErrorDoPrintWithMessage("error decoding jwt token", error)
	return jwtTokenUnmarshal
}

func isValidToken(jwtToken JWTToken) bool {
	jwtSignatureServer := generateJWTSignature(jwtToken.Body)
	if jwtSignatureServer != jwtToken.Siganture {
		return false
	}
	return true
}

func unauthorizeda(c *gin.Context) {
	c.Header("WWW-Authenticate", "Unauthorized")
	c.AbortWithStatus(http.StatusUnauthorized)
}

// func getJWTToken(userId string) string {

// 	expiredToken := getExpiredToken()
// 	id := generateId()
// 	jWTBody := JWTBody{userId, expiredToken, id, ADMIN_SUPER}
// 	jwtSignature := generateJWTSignature(jWTBody)
// 	jwtFormatInJson, _ := json.Marshal(buildJWTFormat(jWTBody, jwtSignature))

// 	return base64.StdEncoding.EncodeToString(jwtFormatInJson)
// }
