package handlerAuth

/*
user can login with phone number or user id or email
*/
type AuthModel struct {
	phoneNumber string `json:"phoneNumber"`
	userId      string `json:"userId"`
	email       string `json:"email"`
	password    string `json:"password"`
}
