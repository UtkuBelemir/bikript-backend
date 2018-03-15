package bikriptAPI

import (
	"fmt"
	"net/http"
)

func (bkHand BikriptHandlers) MethodNotAllowed(wri http.ResponseWriter,status int){
	fmt.Fprintln(wri,"METHOD NOT ALLOWED ROUTE : " + string(status)) //TODO : METHOD NOT ALLOWED PAGE MUST PLANNED AGAIN
}