package app

import (
	"net/http"
)

// Mount mounts handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))
	mux.Handle("/upload/", http.StripPrefix("/upload",http.FileServer(http.Dir("upload")))

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/edit", adminEdit)
	adminMux.HandleFunc("/create", adminCreate)

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
