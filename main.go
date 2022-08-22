package main

import (
	"context"
	"github.com/gohade/hade/app/provider/demo"
	"github.com/gohade/hade/framework/gin"
	"github.com/gohade/hade/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()

	//绑定具体服务
	core.Bind(&demo.DemoServiceProvider{})

	//注册中间件
	core.Use(middleware.Cost())
	core.Use(gin.Recovery())

	//注册路由
	registerRouter(core)
	server := &http.Server{
		Addr:    ":8080",
		Handler: core,
	}

	//启动服务的gorouting
	go func() {
		server.ListenAndServe()
	}()

	//当前的gorouting 等待信号liang
	quit := make(chan os.Signal)
	//监控信号量
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM )
	//阻塞当前的信号
	<-quit

	//调用server.shoutdown graceful结束
	timeoutContext, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutContext); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
