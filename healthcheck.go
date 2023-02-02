package main

import (
	"healthcheck/config"
	"healthcheck/work"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config := config.Get()
	config.Make("conf.json")
	go work.New(*config)
	r := mux.NewRouter()
	r.HandleFunc("/fails", work.GetErrorsHandler)
	http.ListenAndServe(":8080", r)
}
