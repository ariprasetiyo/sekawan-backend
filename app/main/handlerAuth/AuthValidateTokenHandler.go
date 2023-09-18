package handlerAuth

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/util"

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

	clientId := c.GetHeader(util.HEADER_CLIENT_ID)
	signatureInReq := c.GetHeader(util.HEADER_SIGNATURE)
	httpMethod := c.Request.Method
	sourceUrl := c.Request.URL.String()
	authorization := c.GetHeader(util.HEADER_AUTHORIZATION)
	jsonRequestBody := getBodyRequest(c, clientId)

	if util.IsEmptyString(signatureInReq) || !isValidSignature(signatureInReq, jsonRequestBody) {
		unauthorized(c)
		logrus.Infoln("invalid siganture", clientId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)
	} else if util.IsEmptyString(clientId) && util.IsEmptyString(authorization) &&
		util.IsEmptyString(httpMethod) && util.IsEmptyString(sourceUrl) &&
		util.IsEmptyObject(jsonRequestBody) {
		logrus.Info("invalid request", clientId, "authorization:", authorization, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl, " request body:", jsonRequestBody)
		unauthorized(c)
		return
	}

	unauthorizeda(c)
}

func getBodyRequest(c *gin.Context, clientId string) []byte {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorln("error read request body", clientId, err.Error())
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
	return jsonData
}

func unauthorizeda(c *gin.Context) {
	c.Header("WWW-Authenticate", "Unauthorized")
	c.AbortWithStatus(http.StatusUnauthorized)
}
