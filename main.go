package main

import (
	"context"
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/handlerAuth"
	"sekawan-backend/app/main/handlerOCR"
	repository "sekawan-backend/app/main/repository"
	"sekawan-backend/app/main/server"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/gin-contrib/cors"
)

var db = make(map[string]string)

func main() {

	server.InitConfig()
	server.InitLogrusFormat()

	// running open telemetry
	cleanup := server.InitTracer()
	defer cleanup(context.Background())

	gin.SetMode(server.GIN_MODE)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(server.RequestResponseLogger())
	router.Use(otelgin.Middleware(server.APP_NAME))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	router.Use(cors.New(corsConfig))

	initPosgresSQL := server.InitPostgreSQL()
	database := repository.NewDatabase(initPosgresSQL.DB, initPosgresSQL)

	authGenerateTokenHandler := handlerAuth.NewAuthGenerateTokenHandler(&database)
	authValidateTokenHandler := handlerAuth.NewAuthValidateTokenHandler(&database)
	ocrHandler := handlerOCR.NewOcrApiHandler(&database)

	router.GET("/public/health", handler.HealthCheckHanlder)
	router.POST("/public/token", authGenerateTokenHandler.Execute)
	authorizedRouter := router.Group("/", authValidateTokenHandler.Execute)
	authorizedRouter.POST("/api/v1/ocr", ocrHandler.Execute)
	authorizedRouter.GET("/api/v1/test/get", ocrHandler.Execute)

	router.Run(":" + server.HTTP_SERVER_PORT)
}
