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
  files, _ := ioutil.ReadDir("./blog/")

  for _, filename := range files {
    filecontent, err := ioutil.ReadFile("./blog/" + filename.Name())
    check(err)

    tpl, err := pongo2.FromString("{% extends \"base.html\" %}{% block content %}"+string(blackfriday.MarkdownCommon(filecontent))+"{% endblock %}")
    check(err)

    f, err := tpl.Execute(pongo2.Context{})
    check(err)

    ioutil.WriteFile(filename.Name(), []byte(f), 0644)
  }

}
