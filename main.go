package main

import (
    "strings"
    "path/filepath"
    "io/ioutil"
    "os"
    "regexp"
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
    public_html, err := exists("./public_html/")
    check(err)
    if !public_html {
      os.Mkdir("./public_html/", 0755)
      os.Mkdir("./public_html/blog/", 0755)
      os.Mkdir("./public_html/assets/", 0755)
      ioutil.WriteFile("./public_html/assets/styles.css", []byte(""), 0644)
    }

    archive := make([]map[string]string, 0)

    files, _ := ioutil.ReadDir("./blog/")
    for _, filename := range files {
        // Ignore drafts
        if strings.HasPrefix(filename.Name(), "draft") {
            continue
        }

        filecontent, err := ioutil.ReadFile("./blog/" + filename.Name())
        check(err)

        // Read the metadata
        r, _ := regexp.Compile("(?m)^Title: (.*)$")
        title := r.FindStringSubmatch(string(filecontent))[1]
        filecontent = []byte(r.ReplaceAllString(string(filecontent), ""))

        r, _ = regexp.Compile("(?m)^Published: (.*)$")
        published := r.FindStringSubmatch(string(filecontent))[1]
        filecontent = []byte(r.ReplaceAllString(string(filecontent), ""))

        tpl, err := pongo2.FromFile("detail.html")
        check(err)

        f, err := tpl.Execute(pongo2.Context{"title": title, "published": published, "content": string(blackfriday.MarkdownCommon(filecontent))})
        check(err)

        finalfilename := strings.TrimSuffix(filename.Name(), filepath.Ext(filename.Name()))
        ioutil.WriteFile("./public_html/blog/" + finalfilename + ".html", []byte(f), 0644)

        m := make(map[string]string)
        m["url"] = "./blog/" + finalfilename + ".html"
        m["title"] = title

        archive = append(archive, m)
    }

    tpl, err := pongo2.FromFile("index.html")
    check(err)

    f, err := tpl.Execute(pongo2.Context{"items": archive})
    check(err)

    ioutil.WriteFile("./public_html/index.html", []byte(f), 0644)
}
