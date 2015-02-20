package main

import (
    "appengine"
    "appengine/datastore"
    "github.com/kennygrant/sanitize"
)

type Post struct {
    Title, Description string
    Student Student
    PostImage appengine.BlobKey `datastore:",noindex"`
}

type Student struct {
    Name, Tagline string
}

func SubmitPost(c appengine.Context, p *Post) {
    key := datastore.NewKey(c, "Post", sanitize.Path(p.Title), 0, nil)
    datastore.Put(c, key, p)
}