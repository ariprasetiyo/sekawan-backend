package handle_auth

import (
	"encoding/json"
	"net/http"
	"os"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	keyUserId string = "jb(HH}=#jA=%6QK7"
)

func NewAuthGenerateTokenHandler(db *repository.Database) handler.HandlerInterface {
	return &AuthGenerateToken{databaseImpl: *db}
}

type AuthGenerateToken struct {
	databaseImpl repository.Database
}

func (auth AuthGenerateToken) Execute(c *gin.Context) {
	clientId := c.GetHeader(util.HEADER_CLIENT_ID)
	signature := c.GetHeader(util.HEADER_SIGNATURE)
	requestTime := c.GetHeader(util.HEADER_REQUEST_TIME)
	authorization := c.GetHeader(util.HEADER_AUTHORIZATION)
	httpMethod := c.Request.Method
	sourceUrl := c.Request.URL.String()

	if util.IsEmptyString(clientId) && util.IsEmptyString(signature) &&
		util.IsEmptyString(requestTime) && util.IsEmptyString(authorization) &&
		util.IsEmptyString(httpMethod) && util.IsEmptyString(sourceUrl) {
		unauthorized(c)
		logrus.Info("invalid request", clientId, " signature:", signature, " requestTime:", requestTime, "authorization:", authorization, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl)
	}

	if !isValidAuth(auth) {
		unauthorized(c)
	}

	auth.doProcess(c, "", "", "", "")
}

func doDeserialize() {
}

/*
1. generateToken() :
2. getExpiredToken() :
3. JWT format : base64 ( body information , signature )
3. JWT body information : uuid, expired time , user_id encrypted, created time
4. JWT signature :  json JWT information with user_in in plantext
4. JWT validation : signature client in userId plantext vs signature BE in userId plantext
*/
func (auth AuthGenerateToken) doProcess(c *gin.Context, userId string, token string, created_at string, expired_at string) {
	auth.saveToken(c, userId, token, created_at, "123")

	expiredToken := getExpiredToken()
	id := generateId()
	userIdEncrypted := util.EncryptAES256(keyUserId, userId)

	jWTBodyUserIdEncrypted := JWTBody{userIdEncrypted, expiredToken, id}
	jWTBodyUserId := JWTBody{userId, expiredToken, id}

	jwtSignature := generateJWTSignature(jWTBodyUserId)
	buildJWTToken(jWTBodyUserIdEncrypted, jwtSignature)
}

func buildJWTToken(body JWTBody, signature string) {
	var jwtToken JWTToken
	jwtToken.Body = body
	jwtToken.Siganture = signature
}

func generateJWTSignature(jwtBody JWTBody) string {
	jwtBodyJson, err := json.Marshal(jwtBody)
	util.IsErrorDoPrintWithMessage("error generate jwt json", err)
	return util.HmacSha256InByte("", jwtBodyJson)
}

func generateId() string {
	return uuid.New().String()
}

func getExpiredToken() int64 {
	expiredInLong, error := strconv.ParseInt(os.Getenv(util.CONFIG_TOKEN_EXPIRED_IN_MINUTES), 10, 0)
	util.IsErrorDoPanicWithMessage("", error)
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
