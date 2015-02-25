package main

import (
    "appengine"
    "appengine/user"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/hoisie/mustache"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    data := map[string]interface{}{
        "posts": GetAllPosts(c),
    }
    page := mustache.RenderFile(GetPath("index.html"), data)
    fmt.Fprint(w, page)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, _ := user.LoginURL(c, r.FormValue("returnUrl"))
        http.Redirect(w, r, url, 301)
    } else {
        http.Redirect(w, r, r.FormValue("returnUrl"), 301)
    }
}

func IssueHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    vars := mux.Vars(r)
    p := GetPost(c, vars["slug"])
    page := mustache.RenderFile(GetPath("issue.html"), p)
    fmt.Fprint(w, page)
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    vars := mux.Vars(r)
    data := map[string]interface{}{
        "error": nil,
        "success": false,
    }
    if u == nil {
        data["error"] = "You must be logged in to perform this action."
        data["success"] = false
    } else if IsWVUStudent(u.Email) {
        p := GetPost(c, vars["slug"])
        if !p.HasVoted(u.Email) {
            data["success"] = true
            p.SubmitVote(c, u.Email)
        } else {
            data["success"] = false
            data["error"] = "You cannot vote twice on an issue."
        }
    } else {
        data["success"] = false
        data["error"] = "You do not have permission to vote on this issue."
    }
    WriteJSON(w, data)
}

func StaffDashboardHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    logoutURL, _ := user.LogoutURL(c, "/")
    if u == nil {
        url, _ := user.LoginURL(c, "/staff/dashboard")
        http.Redirect(w, r, url, 301)
        return
    }
    data := map[string]interface{}{
        "email": u.Email,
        "id": u.ID,
        "logout": logoutURL,
        "uploadUrl": GetUploadURL(c, "/staff/dashboard"),
    }
    if r.Method == "POST" {
        f, v := UploadImage(c, r)
        if v != nil {
            p := &Post{
                Title: string(v["title"][0]),
                Description: string(v["description"][0]),
                Student: Student{
                    Name: string(v["name"][0]),
                    Tagline: string(v["tagline"][0]),
                },
            }
            if f == nil {
                p.PostImage = "null"
            } else {
                p.PostImage = f.BlobKey
            }
            SubmitPost(c, p)
            data["newPost"] = p
        }
    }
    if IsCampaignStaff(u.Email) {
        page := mustache.RenderFile(GetPath("dash.html"), data)
        fmt.Fprint(w, page)
    } else {
        fmt.Fprint(w, "This page is restricted to campaign staff only.")
    }
}

func ImageServeHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    ServeImage(w, vars["blobKey"])
}