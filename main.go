package main

import (
    "strings"
    "path/filepath"
    "io/ioutil"
    "os"
    "github.com/russross/blackfriday"
    "github.com/flosch/pongo2"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func main() {
    files, _ := ioutil.ReadDir("./blog/")
    public_html, err := exists("./public_html/")
    check(err)
    if !public_html {
      os.Mkdir("./public_html/", 0755)
    }

    for _, filename := range files {
        // Ignore drafts
        if strings.HasPrefix(filename.Name(), "draft") {
            continue
        }

        filecontent, err := ioutil.ReadFile("./blog/" + filename.Name())
        check(err)

        tpl, err := pongo2.FromString("{% extends \"base.html\" %}{% block content %}"+string(blackfriday.MarkdownCommon(filecontent))+"{% endblock %}")
        check(err)

        f, err := tpl.Execute(pongo2.Context{})
        check(err)

        finalfilename := strings.TrimSuffix(filename.Name(), filepath.Ext(filename.Name()))
        ioutil.WriteFile(finalfilename + ".html", []byte(f), 0644)
    }
}
