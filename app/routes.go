package main

import (
    "appengine"
    "appengine/user"
    "fmt"
    "github.com/hoisie/mustache"
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
    if r.Method == "POST" {
        if f, v := UploadImage(c, r); f != nil {
            c.Infof("%#v", r.Body)
            p := &Post{
                Title: string(v["title"][0]),
                Description: string(v["description"][0]),
                Student: Student{
                    Name: string(v["name"][0]),
                    Tagline: string(v["tagline"][0]),
                },
                PostImage: f.BlobKey,
            }
            SubmitPost(c, p)
        }
    }
    if currentUser == nil {
        url, _ := user.LoginURL(c, "/staff/dashboard")
        http.Redirect(w, r, url, 301)
        return
    } else if IsWVUStudent(currentUser.String()) && IsCampaignStaff(currentUser.String()) {
        logoutURL, _ := user.LogoutURL(c, "/")
        data := map[string]interface{}{
            "email": currentUser.String(),
            "id": currentUser.ID,
            "logout": logoutURL,
            "uploadUrl": GetUploadURL(c, "/staff/dashboard"),
        }
        page := mustache.RenderFile(GetPath("dash.html"), data)
        fmt.Fprint(w, page)
    } else {
        fmt.Fprint(w, "This page is restricted to campaign staff only.")
    }
}