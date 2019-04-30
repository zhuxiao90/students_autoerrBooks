package main

import (
	"context"
	"fmt"
	"github.com/zhuxiao90/students_autoerrBooks/AllTest"
	"google.golang.org/grpc"
	"net"
)
const post  = "127.0.0.1:18881"
//对象要和protocol内定的服务一样
type service struct {}
//实现RPC SayHello 接口
func (this *service)SayHello(ctx context.Context,in *AllTest.HelloRequest)(*AllTest.HelloResponse,error)  {
return &AllTest.HelloResponse{Message: "hello"+in.Name},nil
}
func main()  {
//	创建网络
ln,err:=net.Listen("tcp",post)
	if err!=nil {
		fmt.Println("网络异常",err)
	}

	//	创建一个grpc的句柄
	srv:=grpc.NewServer()
//	将srv注册到grpc服务中
    AllTest.RegisterHelloServiceServer(srv,&service{})
//	监听
err=srv.Serve(ln)
	if err!=nil {
		fmt.Println("网络启动异常",err)
	}

}