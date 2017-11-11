package view

import (
	"net/http"

	"github.com/wasitpons/gonews/pkg/model"
)

// List IndexData
type IndexData struct {
	List []*model.News
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

// Admin List Handler
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

// Admin Craete Handler
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// Admin Edit Handler
func AdminEdit(w http.ResponseWriter, data interface{}) {
	render(tpAdminEdit, w, data)
}
