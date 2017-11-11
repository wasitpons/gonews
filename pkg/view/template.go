package view

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpIndex       = parseTemplate("template/root.html", "template/index.html")
	tpAdminLogin  = parseTemplate("template/root.html", "template/admin/admin.html")
	tpNews        = parseTemplate("template/root.html", "template/news.html")
	tpAdminList   = parseTemplate("template/root.html", "template/admin/list.html")
	tpAdminCreate = parseTemplate("template/root.html", "template/admin/create.html")
	tpAdminEdit   = parseTemplate("template/root.html", "template/admin/edit.html")
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
