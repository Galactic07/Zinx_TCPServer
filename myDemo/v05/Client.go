package main

import (
	"fmt"
	"io"
	"net"
	"src/zinx/znet"
	"time"
)

// 模拟客户端
func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)
	//1 直接链接远程服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err,exit!", err)
		return
	}
	//2 链接调用write 写数据
	for {
		//发送封包的message消息 MsgID:0
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("Zinxv0.5 client Test Message")))
		if err != nil {
			fmt.Println("client Pack err,exit!", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println(" Write err", err)
			return
		}
		//服务器一个回复一个message数据，MsgID：1 pingpingping

		//先读取流中的head部分 得到ID和dataLen

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head err", err)
			break
		}
		//将二进制的head拆包到msg 结构体中
		msgHead, err := dp.UnPack(binaryHead)
		if err != nil {
			fmt.Println("client unpack err", err)
			break
		}
		if msgHead.GetMsgLen() > 0 {
			//再根据dataLen进行第二次读取，将data读出来
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data err", err)
				return
			}
			fmt.Println("-->Recv Server Msg:ID=", msg.Id, "len =", msg.GetMsgLen(), "data=", string(msg.Data))
		}

		//cpu阻塞
		time.Sleep(1 * time.Second)
	}

}
