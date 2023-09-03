package handle_auth

import (
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

func (auth AuthGenerateToken) doProcess(c *gin.Context, userId string, token string, created_at string, expired_at string) {
	auth.saveToken(c, userId, token, created_at, "123")
}

func generateToken() string {
	uuid := uuid.New().String()
	epochTime := string(time.Now().UnixMilli())
	var token string = epochTime + "-" + uuid
	return token
}

func getExpiredToken() int64 {
	expiredInLong, error := strconv.ParseInt(os.Getenv(util.CONFIG_TOKEN_EXPIRED_IN_MINUTES), 10, 0)
	util.IsErrorDoPanicWithMessage("", error)
	expiredInepochTime := time.Now().UnixMilli() + expiredInLong
	return expiredInepochTime
}

func encryptedAES256() {

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
