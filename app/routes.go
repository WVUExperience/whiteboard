package main

import (
    "appengine"
    "appengine/user"
    "fmt"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world!")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    currentUser := user.Current(c)
    if currentUser == nil {
        url, _ := user.LoginURL(c, "/login")
        http.Redirect(w, r, url, 301)
        return
    } else if IsWVUStudent(currentUser.String()) {
        if IsCampaignStaff(currentUser.String()) {
            fmt.Fprint(w, "<a href='/staff/dashboard'>Campaign Staff Dashboard</a><br>")
        }

        fmt.Fprint(w, "Welcome, " + currentUser.String())
    } else {
        fmt.Fprint(w, "Sorry, this page is only available for WVU students.")   
    }
}

func StaffDashboardHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    currentUser := user.Current(c)
    if currentUser == nil {
        url, _ := user.LoginURL(c, "/staff/dashboard")
        http.Redirect(w, r, url, 301)
        return
    } else if IsWVUStudent(currentUser.String()) && IsCampaignStaff(currentUser.String()) {
        fmt.Fprintf(w, "%#v", currentUser)
    } else {
        fmt.Fprint(w, "This page is restricted to campaign staff only.")
    }
}