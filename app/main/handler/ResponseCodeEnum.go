package handler

type RESP_CODE_ENUM int32

const (
	// generic response code
	SUCCESS       RESP_CODE_ENUM = 200
	FAILED        RESP_CODE_ENUM = 301
	SETTLED       RESP_CODE_ENUM = 302
	PAID          RESP_CODE_ENUM = 303
	UNKNOWN_ERROR RESP_CODE_ENUM = 304
	BAD_REQUEST   RESP_CODE_ENUM = 305
	UNAUTHORIZED  RESP_CODE_ENUM = 306

	// response code auth
	AUTH_ERROR_DESERIALIZE_JSON_REQUEST RESP_CODE_ENUM = 505
	AUTH_ERROR_INVALID_CLIENT_ID        RESP_CODE_ENUM = 506
)

func (respCode RESP_CODE_ENUM) String() string {
	switch respCode {
	case SUCCESS:
		return "success"
	case BAD_REQUEST:
		return "bad_request"
	case UNAUTHORIZED:
		return "unauthorized"
	case AUTH_ERROR_DESERIALIZE_JSON_REQUEST:
		return "auth error deserialize json request"
	case UNKNOWN_ERROR:
		return "unknow_error"
	default:
		return "unknow_error"
	}
}
