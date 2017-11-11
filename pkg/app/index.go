package app

import (
	"net/http"

	"github.com/wasitpons/gonews/pkg/model"

	"github.com/wasitpons/gonews/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	list, _ := model.ListNews()
	view.Index(w, &view.IndexData{
		List: list,
	})
}
