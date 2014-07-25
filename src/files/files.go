package files

import "fmt"
import "os"
import "io/ioutil"
import "path/filepath"
import "template"

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

func MakeSiteDir(siteDir string) {
    os.Mkdir(siteDir + "/_site", 0776)
}

func ProcessFile(siteDir string, file string) {
	fmt.Println("Processing "+file)
	data, _:= ioutil.ReadFile(siteDir + "/" + file)
	data = template.RenderHtml(data)
	path := siteDir + "/_site/" + file
	fmt.Println("Writing in "+ path)
	ioutil.WriteFile(path, data, 0644)
}
