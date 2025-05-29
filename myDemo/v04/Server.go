package main

import (
	"fmt"
	"src/zinx/ziface"
	"src/zinx/znet"
)

/*
基于Zinx框架来开发的服务端应用程序
*/
//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("Call back before ping error:")
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("Call back ping...ping...ping error:")
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping...\n"))
	if err != nil {
		fmt.Println("Call back After  ping error:")
	}
}

func main() {
	//	1、创建一个Server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.4]")

	//2 给当前zinx框架添加一个自定义的router
	s.AddRouter(&PingRouter{})

	//3、启动Server
	s.Serve()

}
