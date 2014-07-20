package main

import "path/filepath"
import "os"
import "io/ioutil"
import "fmt"

func getHtmlFilesInDir(dirname string) []string{
    files, _ := ioutil.ReadDir(dirname)
    htmls := make([]string, 0)
    for _, f := range files {
        if !f.IsDir() {
            if filepath.Ext(f.Name()) == ".html" {
                htmls = append(htmls, f.Name())
            }
        }
    }
    return htmls
}

func help() {
    fmt.Println(`usage: gokyll directory`)
}

func main() {
    if len(os.Args) < 2 {
        help()
        return;
    }
    siteDir := os.Args[1]

    htmlFiles := getHtmlFilesInDir(siteDir)
    os.Mkdir(siteDir + "/_site", 0776)
    fmt.Println("These are the html files found:")
    for _, v := range htmlFiles {
        fmt.Println(v)
    }
}


