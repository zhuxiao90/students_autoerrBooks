package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"net/http"
	"study/students_autoerrBooks/IhomeWeb/handler"
)

func main() {
	// create new web service
	// 创建web服务
	service := web.NewService(
		web.Name("go.micro.web.IhomeWeb"),
		web.Version("latest"),
		web.Address(":8999"),
	)

	// initialise service
	// 初始化服务
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	rou := httprouter.New()
	//  静态映射页面
	rou.NotFound = http.FileServer(http.Dir("html"))
	// register html handler
	// 注册服务
	service.Handle("/", rou)

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// run service
	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
