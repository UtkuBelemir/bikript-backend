package bikriptAPI

import (
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	bikriptModels "../bikriptDatabase/models"
)

//ENDPOINT IS : www.bikript.com/bk-api/ {{ function-name-with-lowercase-without-method }}
/*-------------------- PROTECTED ROUTES --------------------*/
func (bkHand BikriptHandlers) ProfileDetails(wri http.ResponseWriter, req *http.Request) {
	SetCORS(wri)
	//TODO : ERROR HANDLİNG
	if req.Method == http.MethodPost {
		var tempUserData bikriptModels.UserInfo
		json.NewDecoder(req.Body).Decode(&tempUserData)
		if tempUserData.Email == "" {
			fmt.Fprintln(wri, "E-Mail is missing") // BK-ERROR
			return
		}
		err := bkHand.DBConnection.DBUpdate(tempUserData)
		if err != nil {
			if err.Error() == `pq: duplicate key value violates unique constraint "users_pkey"` {
				json.NewEncoder(wri).Encode(EmailInUse)
			} else {
				fmt.Println("UNKNOWN ERROR :", err)
			}
			return
		}
		json.NewEncoder(wri).Encode(UpdateSuccess)
	}
}
func (bkHand BikriptHandlers) ProfileDetailsPUT(wri http.ResponseWriter, req *http.Request) {

	SetCORS(wri)
	if req.Method != http.MethodPut {
		bkHand.MethodNotAllowed(wri, http.StatusMethodNotAllowed)
		return
	} else {

	}
	//Name,Surname,Phone,BirthDay
}
func (bkHand BikriptHandlers) AccountVerificationPOST(wri http.ResponseWriter, req *http.Request) {
	SetCORS(wri)
	if req.Method != http.MethodPost {
		bkHand.MethodNotAllowed(wri, http.StatusMethodNotAllowed)
		return
	} else {

	}
}

/*--------------------  PUBLIC  ROUTES  --------------------*/
func (bkHand BikriptHandlers) SignUpPOST(wri http.ResponseWriter, req *http.Request) {
	SetCORS(wri)
	//TODO : ERROR ORDERS !
	if req.Method != http.MethodPost {
		bkHand.MethodNotAllowed(wri, http.StatusMethodNotAllowed)
		return
	} else {
		var tempUser bikriptModels.UserInfo
		json.NewDecoder(req.Body).Decode(&tempUser)
		if len(tempUser.Password) < 8 {
			json.NewEncoder(wri).Encode(ShortPassword)
			return
		}
		//TODO : CHECK EMAIL
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(tempUser.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("webSocket/handlers.go - UserCreator,GenerateFromPassword : " + err.Error())
			return
		}
		tempUser.Password = string(encryptedPassword)
		err = bkHand.DBConnection.DBSave(tempUser)
		if err != nil {
			if err.Error() == `pq: duplicate key value violates unique constraint "users_pkey"` {
				json.NewEncoder(wri).Encode(EmailInUse)
			} else if err.Error() == `pq: null value in column "phone_number" violates not-null constraint` {
				json.NewEncoder(wri).Encode(PhoneNumberNotNull)
			} else if err.Error() == `pq: null value in column "email" violates not-null constraint` {
				json.NewEncoder(wri).Encode(EmailNotNull)
			} else if err.Error() == `pq: null value in column "password" violates not-null constraint` {
				json.NewEncoder(wri).Encode(PasswordNotNull)
			} else {
				fmt.Println("UNKNOWN ERROR :", err)
			}
			return
		}
		json.NewEncoder(wri).Encode(SignUpSuccess)
	}
}
func (bkHand BikriptHandlers) LoginPOST(wri http.ResponseWriter, req *http.Request) {
	//TODO : ERROR HANDLING
	SetCORS(wri)
	if req.Method != http.MethodPost {
		bkHand.MethodNotAllowed(wri, 405)
		return
	} else {
		var tempUserData, fetchedUser bikriptModels.UserInfo
		json.NewDecoder(req.Body).Decode(&tempUserData)
		if len(tempUserData.Email) <= 0 || len(tempUserData.Password) <= 8 {
			json.NewEncoder(wri).Encode(EmailOrPassWrong)
			return
		}
		bkHand.DBConnection.DBConneciton.First(&fetchedUser, "email = ?", tempUserData.Email)
		if fetchedUser.Email == "" {
			json.NewEncoder(wri).Encode(EmailOrPassWrong)
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(tempUserData.Password))
		if err != nil {
			json.NewEncoder(wri).Encode(EmailOrPassWrong)
			return
		}
		var rTok ReturnCredentials
		rTok.Token, err = CreateJwtToken(fetchedUser.Email)
		if err != nil {
			fmt.Println("webSocket/handlers.go - UserLogin,CreateJwtToken : " + err.Error())
			return
		}
		rTok.Email = fetchedUser.Email
		json.NewEncoder(wri).Encode(rTok)
		return
	}
}
func (bkHand BikriptHandlers) IsLoggedInPOST(wri http.ResponseWriter, req *http.Request) {
	//TODO : ERROR HANDLING
	//TODO : Return New Token
	SetCORS(wri)
	if req.Method == http.MethodPost {
		if IsTokenAcceptable(req) {
			json.NewEncoder(wri).Encode(&TokenValid)
		} else {
			json.NewEncoder(wri).Encode(&TokenIsNotValid)
		}
	}
	return
}
