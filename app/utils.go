package main

import (
    "appengine"
    "appengine/blobstore"
    "encoding/json"
    "net/http"
    "net/url"
    "os"
    "path"
)

func GetPath(templateName string) string {
    dir := path.Join(os.Getenv("PWD"), "templates")
    tmpl := path.Join(dir, templateName)
    return tmpl
}

func WriteJSON(w http.ResponseWriter, data map[string]interface{}) {
    json.NewEncoder(w).Encode(data)
}

func UploadImage(c appengine.Context, r *http.Request) (*blobstore.BlobInfo, url.Values) {
    blobs, others, err := blobstore.ParseUpload(r)
    if err != nil {
        c.Errorf(err.Error())
        return nil, others
    }
    file := blobs["image"]
    if len(file) == 0 {
        return nil, others
    } else {
        return file[0], others
    }
}

func GetUploadURL(c appengine.Context, path string) string {
    uri, _ := blobstore.UploadURL(c, path, nil)
    return uri.String()
}

func ServeImage(w http.ResponseWriter, key string) {
    blobstore.Send(w, appengine.BlobKey(key))
}