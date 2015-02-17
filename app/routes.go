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
            fmt.Fprint(w, "Welcome, campaign staffer (" + currentUser.String() + ").")   
        } else {
            fmt.Fprint(w, "Welcome, " + currentUser.String())   
        }
    } else {
        fmt.Fprint(w, "Sorry, this page is only available for WVU students.")   
    }
}