package main

import (
	"context"
	"fmt"
	"github.com/zhuxiao90/students_autoerrBooks/AllTest"
	"google.golang.org/grpc"
)

const post  = "127.0.0.1:18881"

func main()  {
//	客户端链接服务器
conn,err:=grpc.Dial(post,grpc.WithInsecure())
	if err!=nil {
		fmt.Println("连接服务器失败",err)
	}

	defer conn.Close()

	//获得grpc句柄
	c:= AllTest.NewHelloServiceClient(conn)
//	远程调用sayhello
r1,err:=c.SayHello(context.Background(),&AllTest.HelloRequest{Name: "xiao"})
	if err != nil {
		fmt.Println("cloud not get Hello server ..", err)
		return
	}
	fmt.Println("HelloServer resp: ", r1.Message)
}