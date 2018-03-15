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
	router.HandleFunc("/profiledetails",myHandlers.ProfileDetailsPUT)
	router.HandleFunc("/signup",myHandlers.SignUpPOST)
	fmt.Println("App started to listen PORT 4000")
	http.ListenAndServe(":4000",router)
}
