package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/tick", TickHandler).Methods("POST")
	r.HandleFunc("/stats/{instrument}", GetStatsHandler).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
