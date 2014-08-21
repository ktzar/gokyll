package files

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"os/exec"
	"path/filepath"
	"strings"
	"bytes"
	"html/template"
	"time"
)

var OUTPUT_DIR = "generated"

type TemplateData struct {
	Site Config
	Data map[string]interface{}
	File string
}

func (data TemplateData) Year() (string) {
	return fmt.Sprintf("%d", time.Now().Year())
}

func (data TemplateData) PageTitle() (string) {
	for _, page := range data.Site.Pages {
		if page.File == strings.Replace(data.File, ".html", "", 1) {
			return page.Title
		}
	}
	return ""
}

func (data TemplateData) SiteTitle() (string) {
	return data.Site.Title
}

type Page struct {
	Title string
	File string
}

type Config struct {
	Title string
	Pages []Page
}

// Return the static directories in the site
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

// Return the html files that are in the site directory
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

// Create the output directory of the site if it doesn't exist
func MakeSiteDir(siteDir string) {
	os.Mkdir(OUTPUT_DIR, 0776)
}

// Generate output html for one of the files in the site
func ProcessFile(siteDir string, file string) {
	fmt.Println("Processing "+file)

	config := readConfig(siteDir)
	siteData := readData(siteDir)

	data, _:= ioutil.ReadFile(siteDir + "/" + file)
	data = RenderHtml(siteDir, file, config, siteData)
	path := OUTPUT_DIR + "/" + file
	ioutil.WriteFile(path, data, 0644)
}

// Copy all the directories that are not prefixed with _ to the output directory
func CopyDirectoryToSite(dir string, siteDir string) {
	origin := siteDir + string(filepath.Separator) + dir
	target := OUTPUT_DIR
	cmd := exec.Command("cp", "-rf", origin, target)
	cmd.Run()
}

func RenderHtml (siteDir string, file string, config Config, data map[string]interface{}) []byte {
	var out bytes.Buffer
	path := siteDir + "/" + file
	templateFile := siteDir + "/_templates/base.html"
	templateData := TemplateData{config, data, file}
	t :=  template.Must(template.New("page_renderer").ParseFiles(path, templateFile))
	err := t.ExecuteTemplate(&out, "base", templateData)
	if err != nil {
		fmt.Println(err)
	}
	return out.Bytes()
}

func readData (siteDir string) (map[string]interface{}) {
	output := map[string] interface{} {}

	files, _ := ioutil.ReadDir(siteDir + "/_data")
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".json") {
			dataName := strings.Split(f.Name(), ".")[0]
			binaryData, err := ioutil.ReadFile(siteDir + "/_data/"+f.Name())
			if err != nil {
				fmt.Println("Cannot read config.json", err)
			}
			var data interface{}
			json.Unmarshal(binaryData, &data)
			output[strings.Title(dataName)] = data
		}
	}
	return output
}

func readConfig (siteDir string) (store Config) {
	data, err := ioutil.ReadFile(siteDir + "/config.json")
	if err != nil {
		fmt.Println("Cannot read config.json", err)
	} else {
		err = json.Unmarshal(data, &store)
		if err != nil {
			fmt.Println("Cannot parse config.json", err)
		}
	}
	return store
}

