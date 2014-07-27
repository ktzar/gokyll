package template

import "html/template"
import "bytes"

func RenderHtml (path string) []byte {
	var out bytes.Buffer
	t :=  template.Must(template.New("page_renderer").ParseFiles(path, "example/_templates/base.html"))
	_ = t.ExecuteTemplate(&out, "base", "MMM")
	return out.Bytes()
}

