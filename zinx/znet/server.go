package znet

import (
	"fmt"
	"net"
	"src/zinx/utils"
	"src/zinx/ziface"
)

// 实体层
// iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器名称
	Name string
	//	服务器绑定的ip版本
	IPVersion string
	//服务器监听的IP
	IP string
	//服务器监听的端口
	Port int
	//当前server的消息管理模块，用来绑定MsgID和对应处理业务API关系
	MsgHandler ziface.IMsgHandle
	//该server的链接管理器
	ConnMgr ziface.IConnManager
	//该Server创建链接之后自动调用Hook函数--OnConnStart
	OnConnStart func(conn ziface.IConnection)
	//该Server创建链接之后自动调用Hook函数--OnConnStop
	OnConnStop func(conn ziface.IConnection)
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name : %s,listenner at IP:%s,Port:%d is starting\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version:%s,MaxConn:%d,MaxPacketSize:%d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	go func() {
		//0 开启消息队列及Worker工作池
		s.MsgHandler.StartWorkerPool()

		//1 获取一个TCP的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
			return
		}
		//2 监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start Zinx server success!", s.Name, "succ,Listening...")
		var cid uint32
		cid = 0

		//3 阻塞的等待客户端链接，处理客户端链接业务（读写）
		for {
			//如果有客户端链接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//设置最大链接个数的判断，如果超过最大链接，那么则关闭此新的链接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				//todo 给客户端相应一个超出最大链接的错误包
				fmt.Println(" Too many connection Maxconn = ", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			//将处理新链接的业务方法 和conn 进行绑定 得到我们的链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			//启动当前的链接业务处理
			go dealConn.Start()
		}
	}()

}

// 停止服务器
func (s *Server) Stop() {

	//将一些服务器的资源、状态或者一些已经开辟的链接信息 进行停止或者回收
	fmt.Println("[STOP] Zinx Server name", s.Name)
	s.ConnMgr.ClearConn()

}
func (s *Server) Serve() {
	//启动Server的服务功能
	s.Start()

	//todo 做一些启动服务器之后的业务
	//阻塞状态
	select {}
}

// 路由功能：给当前的服务注册一个路由方法，供客户端的链接处理使用
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router Success! ")
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

// 初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

// 注册OnConnStart钩子函数
func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// 注册OnConnStop钩子函数
func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// 调用OnConnStart钩子函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("----> Call OnConnStart() ")
		s.OnConnStart(conn)
	}
}

// 调用OnConnStart钩子函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("----> Call OnConnStop() ")
		s.OnConnStop(conn)
	}
}
