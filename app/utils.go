package main

import (
    "os"
    "path"
)

func GetPath(templateName string) string {
    dir := path.Join(os.Getenv("PWD"), "templates")
    tmpl := path.Join(dir, templateName)
    return tmpl
}