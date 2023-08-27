package handle_auth

import "github.com/gin-gonic/gin"

type AuthInterface interface {
	Execute() gin.HandlerFunc
}
