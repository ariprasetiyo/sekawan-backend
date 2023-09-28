package handlerTest

import (
	"sekawan-backend/app/main/enum"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/handlerAuth"
	handlerTest "sekawan-backend/app/main/handlerTest/model"
	"sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/util"

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

	authorization := c.GetHeader(util.HEADER_AUTHORIZATION)
	msgId := c.GetHeader(util.HEADER_MSG_ID)
	userId := c.GetHeader(util.HEADER_USER_ID)
	acl := c.GetHeader(util.HEADER_ACL)
	bodyRequest := postTest.getBodyRequest(msgId, c)

	var response handlerTest.PostTestResponse
	if util.IsEmptyObject(bodyRequest) {
		response = postTest.buildResponseFailed(msgId, enum.BAD_REQUEST, "body is empty")
	} else if util.IsEmptyObject(bodyRequest.Type) {
		response = postTest.buildResponseFailed(msgId, enum.BAD_REQUEST, "type is empty")
	} else if util.IsEmptyString(bodyRequest.Body.Name) {
		response = postTest.buildResponseFailed(msgId, enum.BAD_REQUEST, "body name is empty")
	} else if util.IsEmptyString(userId) {
		response = postTest.buildResponseFailed(msgId, enum.BAD_REQUEST, "used id is empty")
	} else if util.IsEmptyString(acl) {
		response = postTest.buildResponseFailed(msgId, enum.BAD_REQUEST, "acl is empty")
	} else {
		decodeResult := postTest.decodeJwtToken(msgId, authorization)
		response = postTest.buildResponse(msgId, bodyRequest, userId, acl, decodeResult)
	}

	c.JSON(200, response)
}

func (postTest postTestHandler) buildResponse(msgId string, bodyRequest handlerTest.PostTestRequest, userId string, acl string, jwtToken handlerAuth.JWTToken) handlerTest.PostTestResponse {
	respHeader := handler.Response{ResponseId: msgId, Type: bodyRequest.Type, ResponseCode: enum.SUCCESS, ResponseMessage: enum.SUCCESS.String()}
	respBody := handlerTest.PostTestBodyResponse{JwtDecoding: &jwtToken, Name: bodyRequest.Body.Name, UserId: userId, Acl: acl}
	resp := handlerTest.PostTestResponse{Response: respHeader, Body: &respBody}
	return resp
}

func (postTest postTestHandler) buildResponseFailed(msgId string, responseCode enum.RESP_CODE, msgError string) handlerTest.PostTestResponse {
	respHeader := handler.Response{ResponseId: msgId, ResponseCode: responseCode, ResponseMessage: msgError}
	resp := handlerTest.PostTestResponse{Response: respHeader}
	return resp
}

func (postTest postTestHandler) decodeJwtToken(msgId string, token string) handlerAuth.JWTToken {
	return handlerAuth.DecodeJWTToken(msgId, token)
}

func (postTest postTestHandler) getBodyRequest(msgId string, c *gin.Context) handlerTest.PostTestRequest {
	var postTestRequest handlerTest.PostTestRequest
	error := c.ShouldBindJSON(&postTestRequest)
	if error != nil {
		logrus.Errorln("error umarshal json to object", c.Request.Body, error)
	}
	return postTestRequest
}
