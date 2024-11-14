package main

import (
	"net/http"
	"web_core/framework"
)

func main() {
	core := framework.NewCore()
	registerRoute(core)
	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8888",
	}
	server.ListenAndServe()
}
