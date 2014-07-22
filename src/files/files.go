package files

import "fmt"
import "io/ioutil"
import "path/filepath"

func GetHtmlFilesInDir(dirname string) []string{
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

func Help() {
    fmt.Println(`usage: gokyll directory`)
}
