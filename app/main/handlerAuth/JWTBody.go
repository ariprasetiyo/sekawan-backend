package handlerAuth

type JWTBody struct {
	UserId    string   `json:"userId"`
	ExpiredTs int64    `json:"expiredTs"`
	Acl       ACL_ENUM `json:"acl"`
}
