package handlerAuth

import "github.com/gin-gonic/gin"

type AuthInterface interface {
	Execute() gin.HandlerFunc
}
