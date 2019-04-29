package main

import ("context"
"google.golang.org/grpc"
pt "github.com/zhuxiao90/students_autoerrBooks/goGrpctest/grpcproto"
	"net"
)
const post  = "127.0.0.1:18881"
//对象要和protocol内定的服务一样
type service struct {}
//实现RPC SayHello 接口
func (this *service)SayHello(ctx context.Context,in *pt.HelloRequest)(*pt.HelloResponse,error)  {
return &pt.HelloResponse{Message:"hello"+in.Name},nil
}
func main()  {
//	创建网络
ln,err:=net.Listen("tcp",post)



}