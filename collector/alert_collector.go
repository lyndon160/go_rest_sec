package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//These must to be set as envrionment variables.
var signKey = []byte("notuploadedtogithub")
var remoteIp = "127.0.0.1"
var remotePort = "61002"

func getRemoteAlerts() {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest("GET", "http://"+remoteIp+":"+remotePort+"/", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		log.Print("Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Print(string(body))

}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Lyndon Fawcett"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleCollection() {
	getRemoteAlerts()
}

func main() {
	log.Print("Collecting alerts from " + remoteIp + ":" + remotePort)
	handleCollection()
}
