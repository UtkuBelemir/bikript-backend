package bikriptAPI

import (
	"net/http"
	bikriptModels "../bikriptDatabase/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

type WalletRequest struct {
	Email         string `json:"email"`
	WalletAddress string `json:"address"`
	CoinType      string `json:"coin_type"`
}

func (bkHand BikriptHandlers) GetUserWallet(wri http.ResponseWriter, req *http.Request) {
	//TODO : ERROR HANDLİNG
	SetCORS(wri)
	if req.Method == http.MethodGet {
		var tempWallet bikriptModels.WalletAddresses
		vars := mux.Vars(req)
		bikriptToken := req.Header.Get("Authorization")
		tokenData, err := JWTData(bikriptToken)
		if err != nil {
			fmt.Println(err)
			return
		}
		tempWallet.CoinType = vars["cointype"]
		tempWallet.UserId = fmt.Sprintf("%s", tokenData["email"])
		bkHand.DBConnection.DBConneciton.Debug().Where("user_id = ?", tempWallet.UserId).Find(&tempWallet)
		json.NewEncoder(wri).Encode(&tempWallet)
	}
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
		bkHand.DBConnection.DBConneciton.Debug().Where("user_id = ?", tempWallet.UserId).Find(&tempWallet).Count(&doesHaveAddress)
		if doesHaveAddress > 0 {
			json.NewEncoder(wri).Encode(&UserAlreadyHaveAnAddress)
			return
		}
		bkHand.DBConnection.DBConneciton.Debug().Where("user_id is null").First(&tempWallet)
		err := bkHand.DBConnection.DBUpdate(tempWallet)
		if err != nil {
			fmt.Println("ERROR", err)
		}
		initialReq.WalletAddress = tempWallet.PublicAddress
		json.NewEncoder(wri).Encode(&initialReq)
	}
}
