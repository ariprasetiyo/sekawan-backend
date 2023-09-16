package handlerAuth

import (
	"encoding/base64"
	"encoding/json"
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
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var secretKey = os.Getenv(util.HEADER_USER_AGENT)
var expiredToken = os.Getenv(util.CONFIG_TOKEN_EXPIRED_IN_MINUTES)

func NewAuthGenerateTokenHandler(db *repository.Database) handler.HandlerInterface {
	return &AuthGenerateToken{databaseImpl: *db}
}

type AuthGenerateToken struct {
	databaseImpl repository.Database
}

func (auth AuthGenerateToken) Execute(c *gin.Context) {
	clientId := c.GetHeader(util.HEADER_CLIENT_ID)
	signature := c.GetHeader(util.HEADER_SIGNATURE)
	httpMethod := c.Request.Method
	sourceUrl := c.Request.URL.String()

	var request AuthRequest
	var response AuthResponse
	defaultResponseId := uuid.New().String()
	defaultType := enum.TYPE_GENERATE_TOKEN

	if util.IsEmptyString(clientId) && util.IsEmptyString(signature) &&
		util.IsEmptyString(httpMethod) && util.IsEmptyString(sourceUrl) {

		responseCode := enum.BAD_REQUEST
		reponseHeader := handler.Response{ResponseId: defaultResponseId, Type: defaultType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		logrus.Info("invalid request", clientId, " signature:", signature, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)

	} else if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {

		responseCode := enum.AUTH_ERROR_DESERIALIZE_JSON_REQUEST
		reponseHeader := handler.Response{ResponseId: defaultResponseId, Type: defaultType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}

	} else if util.IsEmptyString(request.RequestId) &&
		util.IsEmptyString(request.Type) &&
		util.IsEmptyObject(request.Body) &&
		util.IsEmptyString(request.Body.Cred) {

		responseCode := enum.BAD_REQUEST
		reponseHeader := handler.Response{ResponseId: defaultResponseId, Type: defaultType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		logrus.Info("invalid request", clientId, " request.Type:", request.Type, " request.Body:", request.Body, " request.Body.Cred:", request.Body.Cred)
	} else {

		encryptionKey := os.Getenv(util.CONFIG_APP_ENCRIPTION_KEY)
		var authCredRequest AuthCredRequest
		resultDecrypted := util.DecryptAES256(encryptionKey, request.Body.Cred)
		err := json.Unmarshal([]byte(resultDecrypted), authCredRequest)
		util.IsErrorDoPrintWithMessage("error unmarshal auth request body", err)

		emailMd5 := auth.getMd5(authCredRequest.Email)
		phoneNoMd5 := auth.getMd5(authCredRequest.PhoneNo)
		passwordMd5 := auth.getPasswordMd5(authCredRequest.PhoneNo, authCredRequest.Password)
		users := auth.getUsers(c, phoneNoMd5, emailMd5, passwordMd5)
		if users == nil {
			responseCode := enum.UNAUTHORIZED
			reponseHeader := handler.Response{ResponseId: request.RequestId, Type: enum.STRING_TO_REQ_TYPE[request.Type], ResponseCode: responseCode, ResponseMessage: responseCode.String()}
			response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		} else {
			token := getJWTToken(users.UserId)
			auth.buildResponse(c, users.UserId, request, token)
		}
	}

	c.JSON(200, response)
}

func (auth AuthGenerateToken) getMd5(value string) string {
	salt := os.Getenv(util.CONFIG_APP_SALT_MD5)
	phoneNoMd5 := util.GenerateMD5(salt, value)
	return phoneNoMd5
}

func (auth AuthGenerateToken) getPasswordMd5(phoneNo string, password string) string {
	salt := os.Getenv(util.CONFIG_APP_SALT_MD5)
	apiKey := os.Getenv(util.CONFIG_APP_API_KEY_PASSWORD)
	passwordMd5 := util.GeneratePasswordHash(salt, apiKey, phoneNo, password)
	return passwordMd5
}

func (auth AuthGenerateToken) buildResponse(c *gin.Context, userId string, request AuthRequest, token string) AuthResponse {

	responseCode := enum.SUCCESS
	reponseHeader := handler.Response{ResponseId: request.RequestId, Type: enum.STRING_TO_REQ_TYPE[request.Type], ResponseCode: responseCode, ResponseMessage: responseCode.String()}
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
	id := generateId()
	jWTBody := JWTBody{userId, expiredToken, id, ADMIN_SUPER}
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
	expiredInLong, error := strconv.ParseInt(expiredToken, 10, 0)
	util.IsErrorDoPanicWithMessage("error get expired jwt token", error)
	expiredInepochTime := time.Now().UnixMilli() + expiredInLong
	return expiredInepochTime
}

func (auth AuthGenerateToken) saveToken(c *gin.Context, userId string, token string, created_at string, expired_at string) {
	auth.databaseImpl.SaveToken(c, userId, token, created_at, expired_at)
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", "Unauthorized")
	c.AbortWithStatus(http.StatusUnauthorized)
}

func (auth AuthGenerateToken) getUsers(c *gin.Context, phoneNumbeMd5 string, emailMd5 string, passwordMd5 string) *handlerAuth_model.AuthModel {
	if !util.IsEmptyString(phoneNumbeMd5) {
		return auth.databaseImpl.GetUserIdByPhoneNo(c, phoneNumbeMd5, passwordMd5)
	} else if !util.IsEmptyString(emailMd5) {
		return auth.databaseImpl.GetUserIdByEmail(c, emailMd5, passwordMd5)
	}
	return nil
}
