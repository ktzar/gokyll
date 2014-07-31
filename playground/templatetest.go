package main

import "io"
import "io/ioutil"
import "fmt"
import "html/template"

type StdoutWriter struct {
	io.Writer
}

func (self StdoutWriter) Write(p []byte) (n int, err error) {
	fmt.Print((string(p)))
	return len(p), nil
}

type Page struct {
	Title string
	Author string
	Content string
	Date string
}

func (self Page) SayHi() (out string) {
	return "Hello I am "+self.Author
}

func main() {
	page := Page{
		Title:"This is a title",
		Author:"Miguel L Gonzalez",
		Content:"This is the content",
		Date:"2014-07-07",
	}
	templ, _ := ioutil.ReadFile(`templatetest.html`)
	t, _ := template.New("foo").Parse(string(templ))
	t.Execute(new(StdoutWriter), page)
}
