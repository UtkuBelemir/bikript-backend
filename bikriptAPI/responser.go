package bikriptAPI
type ResponseModel struct {
	Type    string	`json:"type"`
	Message string	`json:"message"`
}
var(
	ShortPassword = ResponseModel{Message:"short_password",Type:"error"}
	WrongEmail = ResponseModel{Message:"wrong_email",Type:"error"}
	EmailInUse = ResponseModel{Message:"email_or_phone_in_use",Type:"error"}
	PhoneNumber = ResponseModel{Message:"phone_number_in_use",Type:"error"}
	SignUpSuccess = ResponseModel{Message:"signup_success",Type:"success"}
)