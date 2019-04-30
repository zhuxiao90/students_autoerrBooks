package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"net"
	"net/http"
	"net/rpc"
)

type Zhu int

func (this *Zhu)Getinfo(argType int,replyType *int)error  {
	beego.Info(argType)
	*replyType=1+argType
	return nil
}
func main()  {
//	new一个新对象
pd:=new(Zhu)
//注册服务
//register在默认服务中注册并公布，接收服务 pd对象的方法
rpc.Register(pd)
rpc.HandleHTTP()
//建立网络监听
ln,err:=net.Listen("tcp","127.0.0.1:10086")
	if err!=nil {
		beego.Info("lianjie shibai")
	}
beego.Info("正在监听10086")
//service接收监听器上传的http链接
http.Serve(ln,nil)

}