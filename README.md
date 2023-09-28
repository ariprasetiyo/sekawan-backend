# sekawan-backend
golang project

# config & db
1. change .env.example to .env
2. create database sekawan_web
3. create ACL for access db sekawan_web
4. grant all that user for access sekawan_web

# init go modul
go mod init sekawan-backend

# manage unused dependency
go mod tidy : removes unused dependencies automaticly

# init gin gonic
go get -u github.com/gin-gonic/gin
go get -u github.com/sirupsen/logrus
go get -u github.com/joho/godotenv

# init gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/pkg/errors
go get -u github.com/uptrace/opentelemetry-go-extra/otelgorm
go get -u github.com/gin-contrib/cors
go get -t github.com/otiai10/gosseract/v2

go get -u github.com/valyala/fasthttp
https://davidbacisin.com/writing/using-fasthttp-for-api-requests-golang?fbclid=IwAR2OLczNMA5-ENa6-blY7x0HpoLkfBlbWislUOxG6Qy-OH4dQxFMFBHxEWA

# Telemetry
https://signoz.io/opentelemetry/go/
go get -u go.opentelemetry.io/otel
go get -u go.opentelemetry.io/otel/exporters/otlp/otlptrace
go get -u github.com/uptrace/opentelemetry-go-extra/otelgorm

go get -u "gocv.io/x/gocv"
go get -u -d gocv.io/x/gocv

# shortcut command
1. mac : command + f12 = call implementation code
2. mac : option + f12 = call reference code

# rum spesific unit test
go test app/main_test/curl_generator_test.go -v
go test -run TestCurlRequestGenerateToken -v
go test -run TestCurlRequestGenerateHTTPGet -v
go test -run TestCurlRequestGenerateHTTPPost -v

# curl
curl -X POST localhost:8083/public/token -H 'Client-Id: xxxxx' "Signature: " -d ""