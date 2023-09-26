package handlerAuth

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sekawan-backend/app/main/enum"
	"sekawan-backend/app/main/handler"
	handlerAuth_model "sekawan-backend/app/main/handlerAuth/model"
	"sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var secretKey string
var expiredToken string
var secretKeySHA256 string

func NewAuthGenerateTokenHandler(db *repository.Database) handler.HandlerInterface {

	secretKey = os.Getenv(util.HEADER_USER_AGENT)
	expiredToken = os.Getenv(util.CONFIG_APP_TOKEN_EXPIRED_IN_MINUTES)
	clientIdServerSide := os.Getenv(util.CONFIG_APP_CLIENT_ID)
	clientApiKeyServerSide := os.Getenv(util.CONFIG_APP_CLIENT_API_KEY_PASSWORD)
	secretKeySHA256 = clientIdServerSide + "::" + clientApiKeyServerSide
	return &AuthGenerateToken{databaseImpl: *db}
}

type AuthGenerateToken struct {
	databaseImpl repository.Database
}

/*
clientID :
signature :
*/
func (auth AuthGenerateToken) Execute(c *gin.Context) {

	msgId := c.GetHeader(util.HEADER_MSG_ID)
	signatureInReq := c.GetHeader(util.HEADER_SIGNATURE)
	httpMethod := c.Request.Method
	sourceUrl := c.Request.URL.String()

	requestBody := auth.getRequestBody(c)
	var request AuthRequest
	var response AuthResponse
	defaultRequestType := enum.TYPE_GENERATE_TOKEN
	err := json.Unmarshal(requestBody, &request)

	if util.IsEmptyObject(msgId) || util.IsEmptyString(msgId) {

		responseCode := enum.AUTH_ERROR_INVALID_MSG_ID
		reponseHeader := handler.Response{ResponseId: request.RequestId, Type: defaultRequestType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		logrus.Infoln("invalid request msg id", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)
	} else if util.IsEmptyObject(requestBody) {

		responseCode := enum.AUTH_ERROR_DESERIALIZE_JSON_REQUEST
		reponseHeader := handler.Response{ResponseId: request.RequestId, Type: defaultRequestType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		logrus.Infoln("empty request body", msgId, signatureInReq)

	} else if err != nil {

		responseCode := enum.AUTH_ERROR_DESERIALIZE_JSON_REQUEST
		reponseHeader := handler.Response{ResponseId: request.RequestId, Type: defaultRequestType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		util.IsErrorDoPrintWithMessage("error unmarshal auth request body", err)

	} else if util.IsEmptyString(signatureInReq) || !isValidSignature(signatureInReq, requestBody) {

		responseCode := enum.UNAUTHORIZED
		reponseHeader := handler.Response{ResponseId: request.RequestId, Type: defaultRequestType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		logrus.Infoln("invalid siganture", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)

	} else if util.IsEmptyString(msgId) &&
		util.IsEmptyString(httpMethod) && util.IsEmptyString(sourceUrl) {

		responseCode := enum.BAD_REQUEST
		reponseHeader := handler.Response{ResponseId: request.RequestId, Type: defaultRequestType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		logrus.Infoln("invalid request", msgId, " signature:", signatureInReq, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)

	} else if util.IsEmptyString(request.RequestId) &&
		util.IsEmptyObject(request.Type) &&
		util.IsEmptyObject(request.Body) &&
		util.IsEmptyString(request.Body.Cred) {

		responseCode := enum.BAD_REQUEST
		reponseHeader := handler.Response{ResponseId: request.RequestId, Type: defaultRequestType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		logrus.Infoln("invalid request", msgId, " request.Type:", request.Type, " request.Body:", request.Body, " request.Body.Cred:", request.Body.Cred)

	} else {

		encryptionKey := os.Getenv(util.CONFIG_APP_ENCRIPTION_KEY)
		var authCredRequest AuthCredRequest
		resultDecrypted := util.DecryptAES256(encryptionKey, request.Body.Cred)
		err := json.Unmarshal([]byte(resultDecrypted), &authCredRequest)
		util.IsErrorDoPrintWithMessage("error unmarshal auth request body", err)

		emailMd5 := auth.getMd5(authCredRequest.Email)
		phoneNoMd5 := auth.getMd5(authCredRequest.PhoneNo)
		passwordMd5 := auth.getMd5(authCredRequest.Password)
		users := auth.getUsers(c, phoneNoMd5, emailMd5, passwordMd5)
		if util.IsEmptyString(users.UserId) {
			responseCode := enum.UNAUTHORIZED
			reponseHeader := handler.Response{ResponseId: request.RequestId, Type: request.Type, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
			response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		} else {
			token := getJWTToken(users.UserId)
			response = auth.buildResponse(c, users.UserId, request, token)
		}
	}

	c.JSON(200, response)
}

func (auth AuthGenerateToken) getRequestBody(c *gin.Context) []byte {
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorln("error read request body", c.Request.Body, err.Error())
	}
	return requestBody
}

func isValidSignature(signatureInReq string, requestBody []byte) bool {
	signatureInServer := util.HmacSha256InByte(secretKeySHA256, requestBody)
	return signatureInReq == signatureInServer
}

func (auth AuthGenerateToken) getMd5(value string) string {
	salt := os.Getenv(util.CONFIG_APP_SALT_MD5)
	phoneNoMd5 := util.GenerateMD5(salt, value)
	return phoneNoMd5
}

func (auth AuthGenerateToken) buildResponse(c *gin.Context, userId string, request AuthRequest, token string) AuthResponse {

	responseCode := enum.SUCCESS
	reponseHeader := handler.Response{ResponseId: request.RequestId, Type: request.Type, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
	response := AuthResponse{reponseHeader, AuthBodyResponse{UserId: userId, Token: token}}

	return response
}

/*
jwtBody :

  - expiredToken : expired token
  - id : uuid
  - userId : user id from db
  - acl : access control list

JWT signature : sha256 json JWT information ( jwtBody )

1. JWT format : base64 ( body information , signature )
2. JWT validation : signature client in userId plantext vs signature BE in userId plantext
*/
func getJWTToken(userId string) string {

	expiredToken := getExpiredToken()
	jWTBody := JWTBody{userId, expiredToken, ADMIN_SUPER}
	jwtSignature := generateJWTSignature(jWTBody)
	jwtFormatInJson, _ := json.Marshal(buildJWTFormat(jWTBody, jwtSignature))

	return base64.StdEncoding.EncodeToString(jwtFormatInJson)
}

func buildJWTFormat(body JWTBody, signature string) JWTToken {
	var jwtToken JWTToken
	jwtToken.Body = body
	jwtToken.Siganture = signature
	return jwtToken
}

func generateJWTSignature(jwtBody JWTBody) string {
	jwtBodyJson, err := json.Marshal(jwtBody)
	util.IsErrorDoPrintWithMessage("error generate jwt json", err)
	return util.HmacSha256InByte(secretKey, jwtBodyJson)
}

func generateId() string {
	return uuid.New().String()
}

func getExpiredToken() int64 {
	expiredInMinutes, error := strconv.ParseInt(expiredToken, 10, 0)
	expiredInMs := expiredInMinutes * 60000
	util.IsErrorDoPanicWithMessage("error get expired jwt token", error)
	expiredInepochTime := time.Now().UnixMilli() + expiredInMs
	return expiredInepochTime
}

func (auth AuthGenerateToken) saveToken(c *gin.Context, userId string, token string, created_at string, expired_at string) {
	auth.databaseImpl.SaveToken(c, userId, token, created_at, expired_at)
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", "Unauthorized")
	c.AbortWithStatus(http.StatusUnauthorized)
}

func (auth AuthGenerateToken) getUsers(c *gin.Context, phoneNumbeMd5 string, emailMd5 string, passwordMd5 string) handlerAuth_model.AuthModel {
	if !util.IsEmptyString(phoneNumbeMd5) {
		return auth.databaseImpl.GetUserIdByPhoneNo(c, phoneNumbeMd5, passwordMd5)
	} else if !util.IsEmptyString(emailMd5) {
		return auth.databaseImpl.GetUserIdByEmail(c, emailMd5, passwordMd5)
	}
	return handlerAuth_model.AuthModel{}
}
