package util

const (
	EMPTY_JSON_OBJECT = "{}"
	URI_HEALTH_CHECK  = "/health"
	SLASH             = "/"
	COLON             = ":"
	SEMICOLON         = ";"
	CUSTOM            = "custom "
	EMPTY_STRING      = ""

	POST = "POST"
	GET  = "GET"

	HEADER_USER_AGENT    = "User-Agent"
	HEADER_MSG_ID        = "Msg-Id"
	HEADER_CLIENT_ID     = "Client-Id"
	HEADER_CLIENT_USER   = "Client-User"
	HEADER_AUTHORIZATION = "Authorization"
	HEADER_SIGNATURE     = "Signature"
	HEADER_REQUEST_TIME  = "Request-Time"
	HEADER_SOURCE_URL    = "Source-Url"
	HEADER_METHOD        = "Method"
	HEADER_REQUEST_ID    = "Request-id"
	HEADER_USER_ID       = "User-Id"
	HEADER_ACL           = "ACL"

	QUERY_STRING_NAME       = "name"
	QUERY_STRING_TYPE       = "type"
	QUERY_STRING_REQUEST_ID = "request_id"

	HEADER_CONTENT_TYPE           = "Content-Type"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"

	CONFIG_HTTP_SERVER_PORT = "HTTP_SERVER_PORT"
	CONFIG_APP_ENV          = "APP_ENV"
	CONFIG_GIN_MODE         = "GIN_MODE"

	CONFIG_DB_HOST                    = "DB_HOST"
	CONFIG_DB_PORT                    = "DB_PORT"
	CONFIG_DB_NAME                    = "DB_NAME"
	CONFIG_DB_USERNAME                = "DB_USERNAME"
	CONFIG_DB_PASSWORD                = "DB_PASSWORD"
	CONFIG_DB_IS_SHOW_QUERY           = "DB_IS_SHOW_QUERY"
	CONFIG_DB_MAX_IDLE_CONNECTION     = "DB_MAX_IDLE_CONNECTION"
	CONFIG_DB_MAX_OPEN_CONNECTION     = "DB_MAX_OPEN_CONNECTION"
	CONFIG_DB_MAX_LIFETIME_CONNECTION = "DB_MAX_LIFETIME_CONNECTION"

	CONFIG_APP_SALT_MD5                = "APP_SALT_MD5"
	CONFIG_APP_CLIENT_API_KEY_PASSWORD = "APP_CLIENT_API_KEY_PASSWORD"
	CONFIG_APP_CLIENT_ID               = "APP_CLIENT_ID"
	CONFIG_APP_ENCRIPTION_KEY          = "APP_API_ENCRIPTION_KEY"

	CONFIG_APP_TOKEN_EXPIRED_IN_MINUTES = "APP_TOKEN_EXPIRED_IN_MINUTES"

	CONFIG_APP_NAME                    = "APP_NAME"
	CONFIG_OTEL_EXPORTER_OTLP_ENDPOINT = "OTEL_EXPORTER_OTLP_ENDPOINT"
	CONFIG_OTEL_INSECURE_MODE          = "OTEL_INSECURE_MODE"

	PARAM_MERCHANT_ID = "merchantId"

	REQUEST_TYPE_BY_REQUEST_ID             = "requestId"
	REQUEST_TYPE_BY_PARTNER_TRANSACTION_ID = "partnerId"
)
