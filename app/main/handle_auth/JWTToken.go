package handle_auth

type JWTToken struct {
	Body      JWTBody `json:"body"`
	Siganture string  `json:"signature"`
}
