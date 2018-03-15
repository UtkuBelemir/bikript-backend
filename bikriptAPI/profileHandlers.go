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
func (bkHand BikriptHandlers) ProfileDetailsGET(wri http.ResponseWriter, req *http.Request) {
	SetCORS(wri)
	if req.Method != http.MethodGet {
		bkHand.MethodNotAllowed(wri, http.StatusMethodNotAllowed)
		return
	} else {

	}
}
func (bkHand BikriptHandlers) ProfileDetailsPUT(wri http.ResponseWriter, req *http.Request) {
	//TODO : ERROR HANDLİNG
	SetCORS(wri)
	if req.Method != http.MethodPut {
		bkHand.MethodNotAllowed(wri, http.StatusMethodNotAllowed)
		return
	} else {
		var tempUserData bikriptModels.UserInfo
		json.NewDecoder(req.Body).Decode(&tempUserData)
		if tempUserData.Email == "" {
			fmt.Fprintln(wri, "E-Mail is missing") // BK-ERROR
			return
		}
		bkHand.DBConnection.DBUpdate(tempUserData)
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
		if err != nil{
			if err.Error() == `pq: duplicate key value violates unique constraint "users_pkey"`{
				json.NewEncoder(wri).Encode(EmailInUse)
			}
			fmt.Println("ERROR :",err)
			return
		}
		json.NewEncoder(wri).Encode(SignUpSuccess)
	}
}
func (bkHand BikriptHandlers) LoginPOST(wri http.ResponseWriter, req *http.Request) {
	SetCORS(wri)
	if req.Method != http.MethodPost {
		bkHand.MethodNotAllowed(wri, 405)
		return
	} else {

	}
}
