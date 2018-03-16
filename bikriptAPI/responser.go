package bikriptAPI
type ResponseModel struct {
	Type    string	`json:"type"`
	Message string	`json:"message"`
}
type ReturnCredentials struct {
	Token 		string	`json:"token"`
	Email		string	`json:"email"`
}
var(
	ShortPassword = ResponseModel{Message:"short_password",Type:"error"}
	EmailTypo = ResponseModel{Message:"wrong_email",Type:"error"}
	EmailInUse = ResponseModel{Message:"email_or_phone_in_use",Type:"error"}
	PhoneNumberInUse = ResponseModel{Message:"phone_number_in_use",Type:"error"}
	SignUpSuccess = ResponseModel{Message:"signup_success",Type:"success"}
	EmailOrPassWrong = ResponseModel{Message:"email_or_password_wrong",Type:"error"}
	PasswordNotNull = ResponseModel{Message:"password_cannot_be_null",Type:"error"}
	EmailNotNull = ResponseModel{Message:"email_cannot_be_null",Type:"error"}
	PhoneNumberNotNull = ResponseModel{Message:"phone_number_cannot_be_null",Type:"error"}
)