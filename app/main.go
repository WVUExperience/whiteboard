package main

import (
    "github.com/gorilla/mux"
    "net/http"
)

func init() {
    router := mux.NewRouter()
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/login", LoginHandler).Methods("GET")
    router.HandleFunc("/i/{blobKey}", ImageServeHandler).Methods("GET")
    router.HandleFunc("/issue/{slug}", IssueHandler).Methods("GET")

    router.HandleFunc("/staff/dashboard", StaffDashboardHandler).Methods("GET", "POST")

    http.Handle("/", router)
}