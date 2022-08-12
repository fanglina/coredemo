package main

import (
	"coredemo/framework"
	"coredemo/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()

	//注册中间件
	core.Use(middleware.Cost())
	core.Use(middleware.Recovery())

	//注册路由
	registerRouter(core)
	server := &http.Server{
		Addr:    ":8080",
		Handler: core,
	}
	server.ListenAndServe()
}
