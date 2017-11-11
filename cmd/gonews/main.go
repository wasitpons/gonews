package main

import (
	"log"
	"net/http"

	"github.com/wasitpons/gonews/pkg/app"
	"github.com/wasitpons/gonews/pkg/model"
)

const port = ":8080"
const mongoURL = "mongodb://127.0.0.1:27017"

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model: %w", err)
	}
	http.ListenAndServe(port, mux)
}
