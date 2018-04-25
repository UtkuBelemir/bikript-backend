package main

import (
	"net/http"
	bikriptAPI "./bikriptAPI"
	"github.com/gorilla/mux"
	"fmt"
	"sync"
	bikriptEmail "./bikriptEmail"
)

func main() {
	myHandlers := bikriptAPI.NewBikriptHandlers()
	router := mux.NewRouter()
	var wg sync.WaitGroup
	wg.Add(1)
	go bikriptEmail.CheckEmail(myHandlers.DBConnection)
	//NOT PROTECTED ROUTES
	router.HandleFunc("/",myHandlers.Index)
	router.HandleFunc("/profiledetails",myHandlers.ProfileDetails)
	router.HandleFunc("/getwallet/{cointype}",myHandlers.GetUserWallet)
	router.HandleFunc("/login",myHandlers.LoginPOST)
	router.HandleFunc("/create",myHandlers.CreateWalletAddressPOST)
	router.HandleFunc("/signup",myHandlers.SignUpPOST)
	router.HandleFunc("/activation",myHandlers.EmailActivationPOST)
	router.HandleFunc("/isloggedin",myHandlers.IsLoggedInPOST)
	fmt.Println("App started to listen PORT 4000")
	http.ListenAndServe(":4000",router)
	wg.Wait()
}
