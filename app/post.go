package main

import (
    "appengine"
    "appengine/datastore"
    "github.com/kennygrant/sanitize"
)

type Post struct {
    Title, Description, Path string
    Student Student
    PostImage appengine.BlobKey `datastore:",noindex"`
}

type Student struct {
    Name, Tagline string
}

func SubmitPost(c appengine.Context, p *Post) {
    p.Path = sanitize.Path(p.Title)
    key := datastore.NewKey(c, "Post", p.Path, 0, nil)
    datastore.Put(c, key, p)
}