package app

import (
	"net/http"

	"github.com/wasitpons/gonews/pkg/view"
)

func newsView(w http.ResponseWriter, r *http.Request) {
	view.News(w, nil)
}
