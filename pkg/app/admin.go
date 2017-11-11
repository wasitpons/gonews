package app

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/wasitpons/gonews/pkg/model"
	"github.com/wasitpons/gonews/pkg/view"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	view.AdminLogin(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {
	view.AdminList(w, nil)
}

func adminEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().GET("id")
	n, err := model.GetNews(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		n.Title = r.FormValue("title")
		n.Detail = r.FormValue("detail")
		if file, fileHandler, err := r.FromFile("image"); err == nil {
			defer file.Close()
			fileName := time.Now().Format(time.RFC3339) + "-" + fileHandler.fileName
			fp, err := os.Create("upload/" + fileName)
			if err != nil {
				io.Copy(fp, file)
			}
			fp.Close()
			n.Image = "/upload/" + fileName
		}
		err := model.EditNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "admin/list", http.StatusSeeOther)
	}
	view.AdminEdit(w, nil)
}

func adminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		n := model.News{
			Title:  r.FormValue("header"),
			Detail: r.FormValue("detail"),
		}

		if file, fileHandler, err := r.FromFile("image"); err == nil {
			defer file.Close()
			fileName := time.Now().Format(time.RFC3339) + "-" + fileHandler.fileName
			fp, err := os.Create("upload/" + fileName)
			if err != nil {
				io.Copy(fp, file)
			}
			fp.Close()
			n.Image = "/upload/" + fileName
		}

		err := model.CreateNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "admin/list", http.StatusSeeOther)
		return
	}
	view.AdminCreate(w, nil)
}
