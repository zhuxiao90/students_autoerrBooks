package main

import (
	"github.com/astaxie/beego"
	"net/rpc"
)

func main()  {
//	rpc网络链接
cli,err:=rpc.DialHTTP("tcp","127.0.0.1:10086")
	if err !=nil {
		beego.Info("网络连接失败")
	}
var val int
	//远程调用函数（被调用的方法，传入的参数 ，返回的参数）
	err =cli.Call("Zhu.Getinfo",321,&val)
	if err!=nil{
		beego.Info("打call失败")
	}
	beego.Info("返回结果",val)
}
