package handlerTest

import (
	"sekawan-backend/app/main/enum"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/handlerAuth"
	handlerTest "sekawan-backend/app/main/handlerTest/model"
	"sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/util"

	"github.com/gin-gonic/gin"
)

func GetTestHandler(db *repository.Database) handler.HandlerInterface {
	return &getTestHandler{databaseImpl: *db}
}

type getTestHandler struct {
	databaseImpl repository.Database
}

/*
TODO
1. token expired not run after set value on the config DONE
2. response acl still 0, not admin_super DONE
3. req log tdk tercapture DONE
*/
func (getTest getTestHandler) Execute(c *gin.Context) {

	authorization := c.GetHeader(util.HEADER_AUTHORIZATION)
	msgId := c.GetHeader(util.HEADER_MSG_ID)
	userId := c.GetHeader(util.HEADER_USER_ID)
	acl := c.GetHeader(util.HEADER_ACL)
	name := c.Query(util.QUERY_STRING_NAME)

	decodeResult := getTest.decodeJwtToken(msgId, authorization)
	c.JSON(200, getTest.buildResponse(msgId, name, userId, acl, decodeResult))
}

func (getTest getTestHandler) buildResponse(msgId string, name string, userId string, acl string, jwtToken handlerAuth.JWTToken) handlerTest.GetTestResponse {
	respHeader := handler.Response{ResponseId: msgId, Type: enum.TYPE_REQUEST_HTTP_GET_TEST, ResponseCode: enum.SUCCESS, ResponseMessage: enum.SUCCESS.String()}
	respBody := handlerTest.GetTestBodyResponse{JwtDecoding: jwtToken, Name: name, UserId: userId, Acl: acl}
	resp := handlerTest.GetTestResponse{Response: respHeader, Body: respBody}
	return resp
}

func (getTest getTestHandler) decodeJwtToken(msgId string, token string) handlerAuth.JWTToken {
	return handlerAuth.DecodeJWTToken(msgId, token)
}
