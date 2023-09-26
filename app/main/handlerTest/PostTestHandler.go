package handlerTest

import (
	"bytes"
	"io/ioutil"
	"sekawan-backend/app/main/enum"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/handlerAuth"
	handlerTest "sekawan-backend/app/main/handlerTest/model"
	"sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostTestHandler(db *repository.Database) handler.HandlerInterface {
	return &postTestHandler{databaseImpl: *db}
}

type postTestHandler struct {
	databaseImpl repository.Database
}

func (postTest postTestHandler) Execute(c *gin.Context) {

	httpMethod := c.Request.Method
	authorization := c.GetHeader(util.HEADER_AUTHORIZATION)
	msgId := c.GetHeader(util.HEADER_MSG_ID)
	jsonRequestBody := postTest.getBodyRequest(c)

	if util.IsEmptyString(authorization) &&
		util.IsEmptyString(jsonRequestBody) {
		logrus.Infoln("invalid request", "authorization:", authorization, " httpMethod:", httpMethod, " request body:", jsonRequestBody)
	}
	decodeResult := postTest.decodeJwtToken(msgId, authorization)
	c.JSON(200, postTest.buildResponse(msgId, "name", decodeResult))
}

func (getTest postTestHandler) buildResponse(msgId string, name string, jwtToken handlerAuth.JWTToken) handlerTest.GetTestResponse {
	respHeader := handler.Response{ResponseId: msgId, Type: enum.TYPE_REQUEST_HTTP_GET_TEST, ResponseCode: enum.SUCCESS, ResponseMessage: enum.SUCCESS.String()}
	respBody := handlerTest.GetTestBodyResponse{JwtDecoding: jwtToken, Name: name}
	resp := handlerTest.GetTestResponse{Response: respHeader, Body: respBody}
	return resp
}

func (postTest postTestHandler) decodeJwtToken(msgId string, token string) handlerAuth.JWTToken {
	return handlerAuth.DecodeJWTToken(msgId, token)
}

func (postTest postTestHandler) getBodyRequest(c *gin.Context) string {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error("error read request body", err.Error())
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
	jsonRequestBody := string(jsonData)
	if len(strings.TrimSpace(jsonRequestBody)) == 0 {
		jsonRequestBody = util.EMPTY_JSON_OBJECT
	}
	return jsonRequestBody
}
