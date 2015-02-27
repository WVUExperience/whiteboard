package main

import (
    "github.com/gorilla/mux"
    "net/http"
)

func init() {
    router := mux.NewRouter()
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/about", AboutHandler)
    router.HandleFunc("/login", LoginHandler).Methods("GET").Queries("returnUrl", "")
    router.HandleFunc("/i/{blobKey}", ImageServeHandler).Methods("GET")
    router.HandleFunc("/issue/{slug}", IssueHandler).Methods("GET", "DELETE")
    router.HandleFunc("/vote/{slug}", VoteHandler).Methods("POST")

    router.HandleFunc("/staff/dashboard", StaffDashboardHandler).Methods("GET", "POST")

    http.Handle("/", router)
}