package files

import "fmt"
import "os"
import "io/ioutil"
import "os/exec"
import "path/filepath"
import "strings"
import "template"

func GetSiteDirs(dirname string) []string{
    files, _ := ioutil.ReadDir(dirname)
    dirs := make([]string, 0)
    for _, f := range files {
        if f.IsDir() && !strings.HasPrefix(f.Name(), "_") {
			dirs = append(dirs, f.Name())
        }
    }
    return dirs
}


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
	data = template.RenderHtml(siteDir + "/" + file)
	fmt.Println("Got "+string(data))
	path := siteDir + "/_site/" + file
	fmt.Println("Writing in "+ path)
	ioutil.WriteFile(path, data, 0644)
}

func CopyDirectoryToSite(dir string, siteDir string) {
	origin := siteDir + string(filepath.Separator) + dir
	target := siteDir + string(filepath.Separator) + "_site"
	cmd := exec.Command("cp", "-rf", origin, target)
	cmd.Run()
}
