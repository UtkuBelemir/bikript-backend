package bikriptAPI
import (
	bikriptDB "../bikriptDatabase"
	"net/http"
	"fmt"
)
type BikriptHandlers struct{
	DBConnection *bikriptDB.Connection
}
func NewBikriptHandlers() *BikriptHandlers{
	bkHand := new(BikriptHandlers)
	bkHand.DBConnection = bikriptDB.NewConnection()
	return bkHand
}
func SetCORS(wri http.ResponseWriter){
	//TODO : EN SON KULLANILMAYAN REQUEST TYPELARI Düzelt
	wri.Header().Set("Content-Type", "application/json; charset=utf-8")
	wri.Header().Set("Access-Control-Allow-Methods", "*")
	wri.Header().Set("Access-Control-Allow-Origin", "*")
	wri.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept,Authorization")
}
func (bkHand BikriptHandlers) Index(wri http.ResponseWriter,req *http.Request){
	fmt.Fprintln(wri,"SA")
}