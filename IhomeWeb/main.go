package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"net/http"
	"study/students_autoerrBooks/IhomeWeb/handler"
	_ "study/students_autoerrBooks/IhomeWeb/models"
)

func main() {
	// create new web service
	// 创建web服务
	service := web.NewService(
		web.Name("go.micro.web.IhomeWeb"),
		web.Version("latest"),
		web.Address(":8080"),
	)

	// initialise service
	// 初始化服务
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	rou := httprouter.New()
	//  静态映射页面
	rou.NotFound = http.FileServer(http.Dir("html"))
	rou.GET("/api/v1.0/areas",handler.GetArea)
	//先写出来的来的来嗯个服务
	//获取session
	rou.GET("/api/v1.0/session",handler.GetSession)
	//获取 index
	rou.GET("api/v1.0/house/index",handler.GetIndex)
	// register html handler
	// 注册服务
	service.Handle("/", rou)
	// run service
	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
