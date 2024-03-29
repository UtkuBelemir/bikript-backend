package bikriptAPI

import (
	"github.com/dgrijalva/jwt-go"
	"crypto/rsa"
	"time"
	"log"
	"io/ioutil"
	"net/http"
	"fmt"
	"errors"
)

//TODO : FIX PATHS
/*
export BACKENDPATH="$HOME/Borsa/bikript_backend"
export PRIVRSA="$BACKENDPATH/keys/superscretkey.rsa"
export PUBRSA="$BACKENDPATH/keys/superscretkey.rsa.pub"
 */
const (
	privateKeyPath = "/Users/utkuelmalioglu/Borsa/bikript-backend/keys/superscretkey.rsa"
	publicKeyPath  = "/Users/utkuelmalioglu/Borsa/bikript-backend/keys/superscretkey.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type TokenClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

func initializeKeys(keyType string) {
	if keyType == "priv" {
		signBytes, err := ioutil.ReadFile(privateKeyPath)
		if err != nil {
			log.Fatal("Error on reading private key ! : " + err.Error())
			return
		}
		signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
		if err != nil {
			log.Fatal("Error on parsing private key ! : " + err.Error())
			return
		}
		return
	}
	if keyType == "pub" {
		verifyBytes, err := ioutil.ReadFile(publicKeyPath)
		if err != nil {
			log.Fatal("Error on reading verify key ! : " + err.Error())
			return
		}
		verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			log.Fatal("Error on parsing verify key ! : " + err.Error())
			return
		}
		return
	}
}
func CreateJwtToken(userName string) (string, error) {
	initializeKeys("priv")
	claims := TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(), //TODO: ZAMANI AYARLA
		},
		userName,
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokent, err := rawToken.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokent, nil
}
func TokenVerifyMiddleware(wri http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if IsTokenAcceptable(req) {
		next(wri, req)
		return
	}
	fmt.Println("UNAUTHORIZED ACCESS!!! : ")
}
func IsTokenAcceptable(req *http.Request) bool {
	initializeKeys("pub")
	cook := req.Header.Get("Authorization")
	token, err := jwt.Parse(cook, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err == nil && token.Valid {
		return true
	}
	fmt.Println("KEY IS NOT ACCEPTABLE ! : " + err.Error())
	return false
}
func IsTokenAcceptableForSocket(req *http.Request) bool {
	initializeKeys("pub")
	cook := req.Header.Get("Sec-WebSocket-Protocol")
	token, err := jwt.Parse(cook, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err == nil && token.Valid {
		return true
	}
	fmt.Println("KEY IS NOT ACCEPTABLE ! : " + err.Error())
	return false
}
func JWTData(curToken string) (jwt.MapClaims,error) {
	initializeKeys("pub")
	bikriptToken, err := jwt.Parse(curToken, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err == nil && bikriptToken.Valid {
		return bikriptToken.Claims.(jwt.MapClaims),nil
	}
	return nil,errors.New("token is not valid")
}
