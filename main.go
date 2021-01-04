package main

import (
	"log"
	"time"
	// "tridentsk/streamcalc/utils"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

    r := mux.NewRouter()
	r.HandleFunc("/value", StoreValue).Methods("POST")

    srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}