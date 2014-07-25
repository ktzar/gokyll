package template

import "html/template"
import "io"
import "fmt"

type Out struct {
	html []byte
	io.Writer
}

func (o *Out) Write(p []byte) (n int, err error) {
	fmt.Println(p)
	o.html = p
	return len(p), nil
}


func RenderHtml (tpl []byte) []byte {
	var out = new(Out)
	t, _ :=  template.New("page_renderer").Parse(string(tpl))
	_ = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	return out.html
}

