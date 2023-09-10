package handlerAuth

type JWTToken struct {
	Body      JWTBody `json:"body"`
	Siganture string  `json:"signature"`
}
