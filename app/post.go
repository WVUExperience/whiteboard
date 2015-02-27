package main

import (
    "appengine"
    "appengine/datastore"
    "github.com/kennygrant/sanitize"
    "math/rand"
    "strings"
    "time"
)

type Post struct {
    Title, Description, Path, Category string
    Votes Votes
    Student Student
    PostImage appengine.BlobKey `datastore:",noindex"`
}

type Votes struct {
    Count int
    Voters []string
    HasVoted bool
}

type Student struct {
    Name, Tagline string
}

func SubmitPost(c appengine.Context, p *Post) {
    p.Path = GetSlug(p.Title)
    p.Votes.Count = 0
    key := datastore.NewKey(c, "Post", p.Path, 0, nil)
    datastore.Put(c, key, p)
}

func GetSlug(s string) string {
    s = sanitize.Path(s)
    s = strings.Replace(s, ".", "", -1)
    if strings.Count(s, "-") > 5 {
        letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
        b := make([]rune, 5)
        rand.Seed(time.Now().UTC().UnixNano())
        for i := range b {
            b[i] = letters[rand.Intn(len(letters))]
        }
        temp, n := "", 4
        for _, v := range strings.Split(s, "-") {
            temp += v + "-"
            if n == 0 {
                temp += string(b)
                break;
            }
            n--
        }
        s = temp
    }
    return s
}

func (p *Post) HasVoted(email string) bool {
    for _, v := range p.Votes.Voters {
        if email == v {
            return true
        }
    }
    return false
}

func (p *Post) SubmitVote(c appengine.Context, email string) {
    p.Votes.Count++;
    p.Votes.Voters = append(p.Votes.Voters, email)
    key := datastore.NewKey(c, "Post", p.Path, 0, nil)
    datastore.Put(c, key, p)
}

func GetPost(c appengine.Context, slug string) *Post {
    var posts []*Post
    q := datastore.NewQuery("Post").Filter("Path =", slug).Limit(1)
    q.GetAll(c, &posts)
    if len(posts) > 0 {
        return posts[0]
    } else {
        return nil
    }
}

func GetAllPosts(c appengine.Context) []Post {
    var posts []Post
    q := datastore.NewQuery("Post")
    q.GetAll(c, &posts)
    return posts
}