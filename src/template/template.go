package template

//import "html/template"
import "fmt"
import "io/ioutil"

func RenderHtml (template string) []byte {
	return []byte(template)
}

func ProcessHtml(file string) {
    fmt.Printf("Processing %s\n", file)
       data, _:= ioutil.ReadFile(file)
       ioutil.WriteFile("_site/" + file, data, 0644)
}

