package main

import (
	"fmt"
	"src/zinx/ziface"
	"src/zinx/znet"
)

/*
基于Zinx框架来开发的服务端应用程序
*/

// PingRoute test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle...")
	//先读取客户端的数据，再回写ping...ping...ping

	fmt.Println("recv from client: msgID=",
		request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping!!"))
	if err != nil {
		fmt.Println(err)
	}
}

// hello Zinx test 自定义路由
type HelloZinxRouter struct {
	znet.BaseRouter
}

// Test Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle...")
	//先读取客户端的数据，再回写ping...ping...ping

	fmt.Println("recv from client: msgID=",
		request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("Hello Welcome to Zinx!!"))
	if err != nil {
		fmt.Println(err)
	}
}

// 创建连接之后执行钩子函数
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("=====>DoConnectionBegin begin...")
	if err := conn.SendMsg(202, []byte("DoConnection  BEGIN")); err != nil {
		fmt.Println(err)
	}

}

// 链接断开之前需要执行的函数
func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("=====>DoConnectionLost is Called ...")
	fmt.Println("conn ID=", conn.GetConnID(), "is Lost...")
}

func main() {
	//	1、创建一个Server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.8]")

	//2 注册链接Hook钩子函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//3 给当前zinx框架添加自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//4、启动Server
	s.Serve()

}
