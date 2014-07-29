package files

import "fmt"
import "os"
import "io/ioutil"
import "encoding/json"
import "os/exec"
import "path/filepath"
import "strings"
import "template"

type Page struct {
	Title string
	File string
}

type Config struct {
	Title string
	Pages []Page
}

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
	os.Mkdir(siteDir + string(filepath.Separator) + "_site", 0776)
}

func ProcessFile(siteDir string, file string) {

	config := readConfig(siteDir)
	fmt.Println("Config", config)

	fmt.Println("Processing "+file)
	data, _:= ioutil.ReadFile(siteDir + "/" + file)
	data = template.RenderHtml(siteDir + "/" + file, config)
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

func readConfig (siteDir string) (store Config) {
	data, err := ioutil.ReadFile(siteDir + "/config.json")
	if err != nil {
		fmt.Println("Cannot read config.json", err)
	}
	fmt.Println(string(data))
	err = json.Unmarshal(data, &store)
	if err != nil {
		fmt.Println("Cannot parse config.json", err)
	}
	return store
}

