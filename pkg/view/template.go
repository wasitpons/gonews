package view

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpIndex      = parseTemplate("template/root.html", "template/index.html")
	tpAdminLogin = parseTemplate("template/root.html", "template/admin.html")
	tpNews       = parseTemplate("template/root.html", "template/news.html")
)

func parseTemplate(files ...string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t

}
func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		log.Print(err)
	}

}

// Index render
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}

// Admin render
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// News render
func News(w http.ResponseWriter, data interface{}) {
	render(tpNews, w, data)
}
