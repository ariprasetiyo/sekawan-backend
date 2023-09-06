package auth

type JWTBody struct {
	UserId    string   `json:"userId"`
	ExpiredTs int64    `json:"expiredTs"`
	Id        string   `json:"id"`
	Acl       ACL_ENUM `json:"acl"`
}
