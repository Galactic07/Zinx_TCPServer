package utils

import (
	"encoding/json"
	"os"
	"src/zinx/ziface"
)

/*
存储一切有关Zinx框架的全局参数，供其他模块使用
一些参数是可以通过zinx.json由用户进行配置
*/
type GlobalObj struct {
	//Server
	TcpServer ziface.IServer //当前Zinx全局的Server对象
	Host      string         //当前服务器主机监听的IP
	TcpPort   int            //当前服务器主机监听的端口号
	Name      string         //当前服务器的名称

	//Zinx
	Version          string //当前Zinx的版本号
	MaxConn          int    //当前服务器主机允许的最大链接数
	MaxPacketSize    uint32 //当前Zinx框架数据包的最大值
	WorkerPoolSize   uint32 //当前业务工作worker池的Goroutine数量
	MaxWorkerTaskLen uint32 //Zinx框架用户最多开辟多少个Worker（限定条件）

}

// 定义一个全局的对外Globalobj
var GlobalObject *GlobalObj

// 从zinx.json去加载用于自定义的参数
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	//json文件数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}

}

// 提供一个init方法，初始化当前的GlobalObject
func init() {
	//如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerAPP",
		Version:          "V0.10",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,   //Worker工作池的队列的个数
		MaxWorkerTaskLen: 1024, //每个worker对应的消息队列的任务的数量的最大值
	}
	//应该尝试从conf/zinx.json去加载一些用户自定义的参数
	GlobalObject.Reload()
}
