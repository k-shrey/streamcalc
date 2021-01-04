package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)
type input struct {
    Value int `json:"value"`
    Expiry int `json:"expiry"` 
}

func StoreValue(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var t input
    err = json.Unmarshal(body, &t)
	if err != nil {
		log.Println(err)
	}
	log.Println(t)
	w.WriteHeader(200)
}