package main

import (
	"net/http"
	bikriptAPI "./bikriptAPI"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	myHandlers := bikriptAPI.NewBikriptHandlers()
	router := mux.NewRouter()
	//NOT PROTECTED ROUTES
	router.HandleFunc("/",myHandlers.Index)
	router.HandleFunc("/profiledetails",myHandlers.ProfileDetails)
	router.HandleFunc("/login",myHandlers.LoginPOST)
	router.HandleFunc("/create",myHandlers.CreateWalletAddressPOST)
	router.HandleFunc("/signup",myHandlers.SignUpPOST)
	router.HandleFunc("/isloggedin",myHandlers.IsLoggedInPOST)
	fmt.Println("App started to listen PORT 4000")
	http.ListenAndServe(":4000",router)
}
