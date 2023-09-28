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

	var jsonRequestBody string
	if httpMethod == util.POST {
		jsonRequestBody = getBodyRequest(c, msgId)
	} else if httpMethod == util.GET {
		jsonRequestBody = sourceUrl
	}

	if util.IsEmptyObject(msgId) || util.IsEmptyString(msgId) {
		logrus.Infoln("invalid request msgid:", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)
		unauthorized(c)
		return
	} else if util.IsEmptyString(signatureInReq) || !isValidSignature(signatureInReq, []byte(jsonRequestBody)) {
		logrus.Infoln("invalid siganture msgId:", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl, "jsonRequestBody:", jsonRequestBody)
		unauthorized(c)
		return
	} else if util.IsEmptyString(msgId) && util.IsEmptyString(authorization) &&
		util.IsEmptyString(httpMethod) && util.IsEmptyString(sourceUrl) &&
		util.IsEmptyObject(jsonRequestBody) {
		logrus.Info("invalid request msgid:", msgId, "authorization:", authorization, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl, " request body:", jsonRequestBody)
		unauthorized(c)
		return
	}

	jwtToken := DecodeJWTToken(msgId, authorization)
	nowInMs := time.Now().UnixMilli()
	if !isValidToken(msgId, jwtToken) {
		logrus.Infoln("invalid JWT msgid:", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl, "jwt:", authorization)
		unauthorized(c)
		return
	} else if jwtToken.Body.ExpiredTs < nowInMs {
		logrus.Infoln("expired token authorization msgid:", msgId, "httpMethod:", httpMethod, "sourceUrl:", sourceUrl, "expiredAt:", jwtToken.Body.ExpiredTs, "now:", nowInMs)
		unauthorized(c)
		return
	} else if util.IsEmptyString(jwtToken.Body.UserId) {
		logrus.Infoln("userId is empty msgid:", msgId, "httpMethod:", httpMethod, "sourceUrl:", sourceUrl, "jwtToken:", jwtToken)
		unauthorized(c)
		return
	} else if util.IsEmptyString(*jwtToken.Body.Acl.String()) {
		logrus.Infoln("acl is empty msgid:", msgId, "httpMethod:", httpMethod, "sourceUrl:", sourceUrl, "jwtToken:", jwtToken)
		unauthorized(c)
		return
	}

	c.Request.Header.Add(util.HEADER_USER_ID, jwtToken.Body.UserId)
	c.Request.Header.Add(util.HEADER_ACL, *jwtToken.Body.Acl.String())
}

func getBodyRequest(c *gin.Context, msgId string) string {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorln("error read request body", msgId, err.Error())
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
	return string(jsonData)
}

func DecodeJWTToken(msgId string, jwtToken string) JWTToken {

	decodeResult, err := base64.StdEncoding.DecodeString(jwtToken)
	util.IsErrorDoPrintWithMessage("error decoding jwt token msgId: "+msgId+" jwtToken:"+jwtToken+", decodeResult:"+string(decodeResult), err)

	var jwtTokenUnmarshal JWTToken
	error := json.Unmarshal(decodeResult, &jwtTokenUnmarshal)
	util.IsErrorDoPrintWithMessage("error Unmarshal jwt token msgId: "+msgId+" jwtToken:"+jwtToken+", decodeResult:"+string(decodeResult), error)
	return jwtTokenUnmarshal
}

func isValidToken(msgId string, jwtToken JWTToken) bool {
	jwtSignatureServer := generateJWTSignature(jwtToken.Body)
	if jwtSignatureServer != jwtToken.Siganture {
		logrus.Infoln("isValidToken JWT msgid:", msgId, " jwtSignatureServer:", jwtSignatureServer, " jwtToken.Signature:", jwtToken.Siganture, "jwtToken:", jwtToken)
		return false
	}
	return true
}

func unauthorizeda(c *gin.Context) {
	c.Header("WWW-Authenticate", "Unauthorized")
	c.AbortWithStatus(http.StatusUnauthorized)
}
