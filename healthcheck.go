package main

import (
	"healthcheck/config"
	"healthcheck/work"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	config := config.Get()
	err := config.Make("conf.json")
	if err != nil {
		os.Exit(1)
	}
	go work.New(*config)
	r := mux.NewRouter()
	r.HandleFunc("/fails", work.GetErrorsHandler)
	http.ListenAndServe(":8080", r)
}
