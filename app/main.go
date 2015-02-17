package main

import (
    "github.com/gorilla/mux"
    "net/http"
)

func init() {
    router := mux.NewRouter()
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/login", LoginHandler).Methods("GET", "POST")

    http.Handle("/", router)
}