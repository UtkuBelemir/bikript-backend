package bikriptEmail
import (
	bikriptDB "../bikriptDatabase"
	"sync"
)
func main(){
	dbCnn := bikriptDB.NewConnection()
	var wg sync.WaitGroup
	wg.Add(1)
	go CheckEmail(dbCnn)
	wg.Wait()


}