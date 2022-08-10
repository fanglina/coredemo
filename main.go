package main

import (
	"coredemo/framework"
	"net/http"
)

func main()  {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           framework.NewCore(),
	}
	server.ListenAndServe()
}