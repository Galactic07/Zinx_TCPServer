package main

import (
	"fmt"
	"net"
	"time"
)

// 模拟客户端
func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)
	//1 直接链接远程服务器
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err,exit!", err)
		return
	}
	for {
		_, err := conn.Write([]byte("Hello Zinx v0.2.."))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err", err)
			return
		}
		fmt.Printf("server call back:%s,cnt=%d\n", buf, cnt)

		//cpu阻塞
		time.Sleep(1 * time.Second)
	}
	//2 链接调用write 写数据

}
