package ziface

// 抽象层
type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve()
	//路由功能：当前的服务器注册一个路由方法，供客户端的链接进行处理使用
	AddRouter(msgID uint32, router IRouter)
	//获取当前Server的链接管理器
	GetConnMgr() IConnManager
	//注册OnConnStart钩子函数
	SetOnConnStart(func(connection IConnection))
	//注册OnConnStop钩子函数
	SetOnConnStop(func(connection IConnection))
	//调用OnConnStart钩子函数
	CallOnConnStart(connection IConnection)
	//调用OnConnStart钩子函数
	CallOnConnStop(connection IConnection)
}
