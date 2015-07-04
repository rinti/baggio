package main

import (
  // "fmt"
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
  file, err := ioutil.ReadFile("test.md")
  check(err)

  tpl, err := pongo2.FromString("{% extends \"base.html\" %}{% block content %}"+string(blackfriday.MarkdownCommon(file))+"{% endblock %}")
  check(err)

  f, err := tpl.Execute(pongo2.Context{})
  check(err)

  // file, err := tpl.Execute(pongo2.Context{})
  // check(err)

  ioutil.WriteFile("index.html", []byte(f), 0644)
}
