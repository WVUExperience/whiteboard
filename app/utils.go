package main

import (
    "appengine"
    "appengine/blobstore"
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