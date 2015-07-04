package main

import (
    "strings"
    "path/filepath"
    "io/ioutil"
    "github.com/russross/blackfriday"
    "github.com/flosch/pongo2"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    files, _ := ioutil.ReadDir("./blog/")

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
