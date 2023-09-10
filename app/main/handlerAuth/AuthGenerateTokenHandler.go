package handlerAuth

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"sekawan-backend/app/main/handler"
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
	defaultType := handler.TYPE_GENERATE_TOKEN

	if util.IsEmptyString(clientId) && util.IsEmptyString(signature) &&
		util.IsEmptyString(httpMethod) && util.IsEmptyString(sourceUrl) {
		responseCode := handler.BAD_REQUEST
		reponseHeader := handler.Response{ResponseId: defaultResponseId, Type: defaultType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
		logrus.Info("invalid request", clientId, " signature:", signature, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)
	} else if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		responseCode := handler.AUTH_ERROR_DESERIALIZE_JSON_REQUEST
		reponseHeader := handler.Response{ResponseId: defaultResponseId, Type: defaultType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
	} else if !isValidAuth(auth) {
		responseCode := handler.UNAUTHORIZED
		reponseHeader := handler.Response{ResponseId: defaultResponseId, Type: defaultType, ResponseCode: responseCode, ResponseMessage: responseCode.String()}
		response = AuthResponse{reponseHeader, AuthBodyResponse{}}
	}

	c.JSON(200, response)
}

func doDeserialize() {
}

/*
1. JWT format : base64 ( body information , signature )
2. JWT validation : signature client in userId plantext vs signature BE in userId plantext
*/
func (auth AuthGenerateToken) doProcess(c *gin.Context, userId string) {
	//todo here
	getJWTToken(userId)
}

/*
jwtBody :

  - expiredToken : expired token
  - id : uuid
  - userId : user id from db
  - acl : access control list

JWT signature : sha256 json JWT information ( jwtBody )
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

func isValidAuth(cerberus AuthGenerateToken) bool {
	if 1 == 1 {
		return true
	}
	logrus.Infoln("invalid auth client id")
	return false
}
