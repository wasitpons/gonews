package app

import "net/http"
import "github.com/wasitpons/gonews/pkg/view"

func adminLogin(w http.ResponseWriter, r *http.Request) {
	view.AdminLogin(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {

}

func adminEdit(w http.ResponseWriter, r *http.Request) {

}
