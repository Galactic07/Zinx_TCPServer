package main

import (
	"src/zinx/znet"
)

/*
基于Zinx框架来开发的服务端应用程序
*/

func main() {
	//	1、创建一个Server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.1]")
	//2、启动Server
	s.Serve()

}
