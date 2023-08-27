package handle_auth

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewOcrApiHandler(db *repository.Database) handler.HandlerInterface {
	return &ocrApiHandler{databaseImpl: *db}
}

type ocrApiHandler struct {
	databaseImpl repository.Database
}

func (auth ocrApiHandler) Execute(c *gin.Context) {

	clientId := c.GetHeader(util.HEADER_CLIENT_ID)
	signature := c.GetHeader(util.HEADER_SIGNATURE)
	requestTime := c.GetHeader(util.HEADER_REQUEST_TIME)
	authorization := c.GetHeader(util.HEADER_AUTHORIZATION)
	httpMethod := c.Request.Method
	sourceUrl := c.Request.URL.String()
	jsonRequestBody := getBodyRequest(c, clientId)

	if util.IsEmptyString(clientId) && util.IsEmptyString(signature) &&
		util.IsEmptyString(requestTime) && util.IsEmptyString(authorization) &&
		util.IsEmptyString(httpMethod) && util.IsEmptyString(sourceUrl) &&
		util.IsEmptyString(jsonRequestBody) {
		unauthorized(c)
		logrus.Info("invalid request", clientId, " signature:", signature, " requestTime:", requestTime, "authorization:", authorization, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl, " request body:", jsonRequestBody)
	}
	logrus.Info("call ocrApiHandler")
}

func getBodyRequest(c *gin.Context, clientId string) string {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error("error read request body", clientId, err.Error())
		unauthorized(c)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
	jsonRequestBody := string(jsonData)
	if len(strings.TrimSpace(jsonRequestBody)) == 0 {
		jsonRequestBody = util.EMPTY_JSON_OBJECT
	}
	return jsonRequestBody
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", "Unauthorized")
	c.AbortWithStatus(http.StatusUnauthorized)
}

func isValidAuth(handler ocrApiHandler) bool {
	if 1 == 1 {
		return true
	}
	logrus.Infoln("invalid auth client id")
	return false
}
