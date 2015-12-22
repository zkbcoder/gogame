// socket
package network

import (
	"fmt"
	"net"
)

type NetConn struct {
	conn  net.Conn // 链接
	Child interface{}
}

func (this *NetConn) SendMsg(msg IMsg) {
	this.conn.Write(msg.Packet().Bytes())
}

func (this *NetConn) Close() {
	this.conn.Close()
}

func (this *NetConn) ClientRun(addr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		return
	}

	this.conn = conn
	go this.reader()
	fmt.Println("conn ok", conn, this.conn)
}

func (this *NetConn) ServerRun(port string) {
	netListen, err := net.Listen("tcp", port)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer netListen.Close()

	for {
		var clietConn NetConn
		conn, err := netListen.Accept()
		if err != nil {
			fmt.Println("Accept Err ", err)
			continue
		}

		clietConn.conn = conn
		fmt.Println(conn.RemoteAddr().String(), "tcp connet success!")
		go clietConn.reader()
		// go this.handleConnection(conn)
	}
}

func (this *NetConn) reader() {
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)

	buffer := make([]byte, 1024)
	for {
		n, err := this.conn.Read(buffer)
		if err != nil {
			fmt.Println(this.conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		tmpBuffer = append(tmpBuffer, buffer[:n]...)
		var bRet bool
		for {
			tmpBuffer, bRet = this.dealMsg(tmpBuffer)
			if bRet {
				break
			}
		}
	}
}

func (this *NetConn) dealMsg(buf []byte) ([]byte, bool) {
	if len(buf) == 0 {
		return buf, true
	}
	// 先读取包大小
	var uLen uint16
	// fmt.Println(buf, "--", uLen)
	UnPacket(buf, &uLen)
	//fmt.Println(buf, "--", uLen, len(buf))
	if int(uLen) > len(buf) {
		return buf, true
	}

	// 开始解包
	var msgId uint16
	UnPacket(buf[2:], &msgId)

	// TODO...
	msg := MsgFactoryMgr.GetMsg(msgId)
	if nil != msg {
		// 解包
		// fmt.Println("msgid=", msgId)
		msg.UnPacket(buf)
		// 回调
		// fmt.Println("client", &msg)
		MsgFactoryMgr.CallBack(msgId, this, msg)
	} else {
		//fmt.Println("No Deal Msg=%d", msgId)
	}

	buf = buf[uLen:]
	return buf, false
}
