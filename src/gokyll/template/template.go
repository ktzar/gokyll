package template

import "html/template"
import "bytes"
import "fmt"

type TemplateData struct {
	Site interface{}
	Data map[string]interface{}
}

func RenderHtml (siteDir string, file string, config interface{}, data map[string]interface{}) []byte {
	var out bytes.Buffer
	path := siteDir + "/" + file
	templateFile := siteDir + "/_templates/base.html"
	templateData := TemplateData{config, data}
	t :=  template.Must(template.New("page_renderer").ParseFiles(path, templateFile))
	err := t.ExecuteTemplate(&out, "base", templateData)
	if err != nil {
		fmt.Println(err)
	}
	return out.Bytes()
}

