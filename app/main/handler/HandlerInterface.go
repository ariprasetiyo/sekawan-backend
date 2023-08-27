package handler

import (
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	Execute(c *gin.Context)
}
