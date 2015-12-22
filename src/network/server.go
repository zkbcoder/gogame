// ServerNet 服务的类
package network

//import (
//	"fmt"
//	"net"
//)

type ServerNet struct {
	NetConn
	Port string // 端口格式 :9988
}

//func (this *ServerNet) Run() {
//	netListen, err := net.Listen("tcp", this.Port)
//	if nil != err {
//		fmt.Println(err)
//		return
//	}
//	defer netListen.Close()

//	for {
//		conn, err := netListen.Accept()
//		if err != nil {
//			fmt.Println("Accept Err ", err)
//			continue
//		}

//		fmt.Println(conn.RemoteAddr().String(), "tcp connet success!")
//		go this.handleConnection(conn)
//	}
//}

//func (this *ServerNet) handleConnection(conn net.Conn) {
//	// 创建一个临时缓冲区，用来存储被截断的数据
//	tmpBuffer := make([]byte, 0)

//	buffer := make([]byte, 1024)
//	for {
//		n, err := conn.Read(buffer)
//		if err != nil {
//			fmt.Println(conn.RemoteAddr().String(), "connect err:", err)
//			return
//		}

//		tmpBuffer = append(tmpBuffer, buffer[:n]...)
//		fmt.Println("tmp=", tmpBuffer)
//		var bRet bool
//		for {
//			tmpBuffer, bRet = this.dealMsg(tmpBuffer)
//			if bRet {
//				break
//			}
//		}
//	}
//}

//func (this *ServerNet) dealMsg(buf []byte) ([]byte, bool) {
//	if len(buf) == 0 {
//		return buf, true
//	}
//	// 先读取包大小
//	var uLen uint16
//	UnPacket(buf, &uLen)
//	if int(uLen) > len(buf) {
//		return buf, true
//	}

//	// 开始解包
//	var msgId uint16
//	UnPacket(buf[2:], &msgId)

//	// TODO...
//	msg := MsgFactoryMgr.GetMsg(msgId)
//	if nil != msg {
//		// 解包
//		msg.UnPacket(buf)
//		fmt.Println("msg=", msg)
//		// 回调
//		// MsgFactoryMgr.CallBack(msgId, this, msg)
//	} else {
//		fmt.Println("No Deal Msg=%d msg=", msgId, msg)
//	}

//	buf = buf[uLen:]
//	return buf, false
//}
