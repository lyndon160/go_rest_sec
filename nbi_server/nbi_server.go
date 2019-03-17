package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type NodeInfo struct {
	Alerts       int    "json:'alerts'"
	MemoryMBFree int    "json:'memory'"
	DiskMBFree   int    "json:'disk'"
	Time         string "json:'time'"
}

var mySigningKey = []byte("notuploadedtogithub")
var nbiPort = "61002"

func getAlerts(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Alerts since last call: "+strconv.Itoa(rand.Intn(100))) // Return JSON object. Get CPU or something?
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(strconv.Itoa(rand.Intn(4)))

	info := NodeInfo{Alerts: rand.Intn(4), MemoryMBFree: rand.Intn(50), DiskMBFree: rand.Intn(500), Time: time.Now().String()}

	json.NewEncoder(w).Encode(info)
	log.Print("Authorised access from " + r.RemoteAddr)

}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
				log.Print(err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			http.Error(w, "HTTP Status 401 - Authentication Failed: Bad Credentials", 401)
			log.Print(r.RemoteAddr + " HTTP Status 401 - Authentication Failed: Bad Credentials")
		}
	})
}

func handleRequests() {
	log.Print("Serving requests on port :" + nbiPort)
	http.Handle("/", isAuthorized(getAlerts))
	log.Fatal(http.ListenAndServe(":"+nbiPort, nil))
}

func main() {
	handleRequests()
}
