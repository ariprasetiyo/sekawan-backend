package handle_auth

import (
	"net/http"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/repository"

	"github.com/gin-gonic/gin"
)

func NewAuthValidateTokenHandler(db *repository.Database) handler.HandlerInterface {
	return &AuthValidateToken{databaseImpl: *db}
}

type AuthValidateToken struct {
	databaseImpl repository.Database
}

func (auth AuthValidateToken) Execute(c *gin.Context) {
	unauthorizeda(c)
}

func unauthorizeda(c *gin.Context) {
	c.Header("WWW-Authenticate", "Unauthorized")
	c.AbortWithStatus(http.StatusUnauthorized)
}
