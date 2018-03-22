package bikriptAPI
type ResponseError struct {
	Error    string	`json:"err"`
}
type ReturnCredentials struct {
	Token 		string	`json:"token"`
	Email		string	`json:"email"`
}
var(
	ShortPassword = ResponseError{Error:"short_password"}
	EmailTypo = ResponseError{Error:"wrong_email"}
	EmailInUse = ResponseError{Error:"email_or_phone_in_use"}
	PhoneNumberInUse = ResponseError{Error:"phone_number_in_use"}
	SignUpSuccess = make(map[string]string)
	UpdateSuccess = make(map[string]string)
	EmailOrPassWrong = ResponseError{Error:"email_or_password_wrong"}
	PasswordNotNull = ResponseError{Error:"password_cannot_be_null"}
	EmailNotNull = ResponseError{Error:"email_cannot_be_null"}
	PhoneNumberNotNull = ResponseError{Error:"phone_number_cannot_be_null"}
	UserAlreadyHaveAnAddress = ResponseError{Error:"user_have_an_address"}
	TokenIsNotValid = ResponseError{Error:"token_expired_or_not_valid"}
	TokenValid = make(map[string]string)
)