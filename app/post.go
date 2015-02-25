package main

import (
    "appengine"
    "appengine/datastore"
    "github.com/kennygrant/sanitize"
)

type Post struct {
    Title, Description, Path string
    Votes int
    Student Student
    PostImage appengine.BlobKey `datastore:",noindex"`
}

type Student struct {
    Name, Tagline string
}

func SubmitPost(c appengine.Context, p *Post) {
    p.Path = sanitize.Path(p.Title)
    p.Votes = 0
    key := datastore.NewKey(c, "Post", p.Path, 0, nil)
    datastore.Put(c, key, p)
}

func GetAllPosts(c appengine.Context) []Post {
    var posts []Post
    q := datastore.NewQuery("Post")
    q.GetAll(c, &posts)
    return posts
}