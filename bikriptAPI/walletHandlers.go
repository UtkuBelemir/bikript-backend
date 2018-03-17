package bikriptAPI

import (
	"net/http"
	bikriptModels "../bikriptDatabase/models"
	"encoding/json"
	"fmt"
)
type WalletRequest struct {
	Email         string `json:"email"`
	WalletAddress string `json:"address"`
	CoinType      string `json:"coin_type"`
}
func (bkHand BikriptHandlers) CreateWalletAddressPOST(wri http.ResponseWriter, req *http.Request) {
	SetCORS(wri)
	//TODO : ERROR HANDLİNG
	if req.Method == http.MethodPost {
		var initialReq WalletRequest
		var tempWallet bikriptModels.WalletAddresses
		doesHaveAddress := 0
		json.NewDecoder(req.Body).Decode(&initialReq)
		tempWallet.UserId = initialReq.Email
		tempWallet.CoinType = initialReq.CoinType
		fmt.Println(tempWallet.TableName())
		bkHand.DBConnection.DBConneciton.Debug().Where("user_id = ?",tempWallet.UserId).Find(&tempWallet).Count(&doesHaveAddress)
		if doesHaveAddress > 0{
			json.NewEncoder(wri).Encode(&UserAlreadyHaveAnAddress)
			return
		}
		bkHand.DBConnection.DBConneciton.Debug().Where("user_id is null").First(&tempWallet)
		err := bkHand.DBConnection.DBUpdate(tempWallet)
		if err != nil{
			fmt.Println("ERROR",err)
		}
		initialReq.WalletAddress = tempWallet.PublicAddress
		json.NewEncoder(wri).Encode(&initialReq)
	}
}
