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
    u := user.Current(c)
    posts := GetAllPosts(c)
    if u != nil {
        for i, v := range posts {
            if v.HasVoted(u.Email) {
                posts[i].Votes.HasVoted = true
            }
        }
    }
    data := map[string]interface{}{
        "posts": posts,
        "user": GetEmbeddedUser(u, c),
    }
    page := mustache.RenderFileInLayout(GetPath("index.html"), GetPath("layout.html"), data)
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
    u := user.Current(c)
    vars := mux.Vars(r)
    p := GetPost(c, vars["slug"])
    data := map[string]interface{}{
        "post": p,
        "user": GetEmbeddedUser(u, c),
    }
    if u != nil {
        p.Votes.HasVoted = p.HasVoted(u.Email)
    }
    page := mustache.RenderFileInLayout(GetPath("issue.html"), GetPath("layout.html"), data)
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
                Category: GetCategory(string(v["category"][0])),
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
        page := mustache.RenderFileInLayout(GetPath("dash.html"), GetPath("layout.html"), data)
        fmt.Fprint(w, page)
    } else {
        fmt.Fprint(w, "This page is restricted to campaign staff only.")
    }
}

func ImageServeHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    ServeImage(w, vars["blobKey"])
}