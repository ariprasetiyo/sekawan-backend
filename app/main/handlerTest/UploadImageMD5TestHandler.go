package handlerTest

import (
	"encoding/base64"
	"sekawan-backend/app/main/enum"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/handlerAuth"
	handlerTest "sekawan-backend/app/main/handlerTest/model"
	"sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/util"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract"
	"github.com/sirupsen/logrus"
)

func UploadTestHandler(db *repository.Database) handler.HandlerInterface {
	return &uploadTestHandler{databaseImpl: *db}
}

type uploadTestHandler struct {
	databaseImpl repository.Database
}

func (uploadImageMD5Test uploadTestHandler) Execute(c *gin.Context) {

	msgId := c.GetHeader(util.HEADER_MSG_ID)
	userId := c.GetHeader(util.HEADER_USER_ID)
	acl := c.GetHeader(util.HEADER_ACL)
	bodyRequest := uploadImageMD5Test.getBodyRequest(msgId, c)

	var response handlerTest.PostTestResponse
	if util.IsEmptyObject(bodyRequest) {
		response = uploadImageMD5Test.buildResponseFailed(msgId, enum.BAD_REQUEST, "body is empty")
	} else if util.IsEmptyObject(bodyRequest.Type) || bodyRequest.Type != enum.TYPE_REQUEST_HTTP_UPLOAD_MD5_IMAGE_TEST {
		response = uploadImageMD5Test.buildResponseFailed(msgId, enum.BAD_REQUEST, "type is not valid")
	} else if util.IsEmptyString(userId) {
		response = uploadImageMD5Test.buildResponseFailed(msgId, enum.BAD_REQUEST, "user id is empty")
	} else if util.IsEmptyString(acl) {
		response = uploadImageMD5Test.buildResponseFailed(msgId, enum.BAD_REQUEST, "acl is empty")
	} else {
		response = uploadImageMD5Test.buildResponse(msgId, bodyRequest, userId, acl)
	}

	c.JSON(200, response)
}

func (uploadImageMD5Test uploadTestHandler) doExecute(request handlerTest.PostTestUploadMD5Request) {
	imageInMD5 := request.Body.Base64Image
	decoding, err := base64.StdEncoding.DecodeString(imageInMD5)
	if err != nil {
		logrus.Errorln("error decoding image base64", request.RequestId, err)
	}
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImageFromBytes(decoding)
	text, _ := client.Text()
	logrus.Println("text : ", request.RequestId, text)

}

func (uploadImageMD5Test uploadTestHandler) buildResponse(msgId string, bodyRequest handlerTest.PostTestUploadMD5Request, userId string, acl string) handlerTest.PostTestResponse {
	respHeader := handler.Response{ResponseId: msgId, Type: bodyRequest.Type, ResponseCode: enum.SUCCESS, ResponseMessage: enum.SUCCESS.String()}
	respBody := handlerTest.PostTestBodyResponse{Name: bodyRequest.Body.Base64Image, UserId: userId, Acl: acl}
	resp := handlerTest.PostTestResponse{Response: respHeader, Body: &respBody}
	return resp
}

func (uploadImageMD5Test uploadTestHandler) buildResponseFailed(msgId string, responseCode enum.RESP_CODE, msgError string) handlerTest.PostTestResponse {
	respHeader := handler.Response{ResponseId: msgId, ResponseCode: responseCode, ResponseMessage: msgError}
	resp := handlerTest.PostTestResponse{Response: respHeader}
	return resp
}

func (uploadImageMD5Test uploadTestHandler) decodeJwtToken(msgId string, token string) handlerAuth.JWTToken {
	return handlerAuth.DecodeJWTToken(msgId, token)
}

func (uploadImageMD5Test uploadTestHandler) getBodyRequest(msgId string, c *gin.Context) handlerTest.PostTestUploadMD5Request {
	var postTestRequest handlerTest.PostTestUploadMD5Request
	error := c.ShouldBindJSON(&postTestRequest)
	if error != nil {
		logrus.Errorln("error umarshal json to object", c.Request.Body, error)
	}
	return postTestRequest
}
