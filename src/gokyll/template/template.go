package template

import "html/template"
import "bytes"

func RenderHtml (path string, config interface{}) []byte {
	var out bytes.Buffer
	t :=  template.Must(template.New("page_renderer").ParseFiles(path, "example/_templates/base.html"))
	_ = t.ExecuteTemplate(&out, "base", config)
	return out.Bytes()
}

