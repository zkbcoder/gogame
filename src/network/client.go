// ClientNet 客户端网络类
package network

import (
	"fmt"
	"net"
	"time"

	"github.com/astaxie/beego/logs"
)

type ClientNet struct {
	ServerAddr string   // 服务端IP:Port
	conn       net.Conn // 链接
	Conn       net.Conn // 链接
	StartTime  time.Time
	RoleName   string
	Log        *logs.BeeLogger
}

func (this *ClientNet) SendMsg(msg IMsg) {
	//	fmt.Println(this.conn)
	this.conn.Write(msg.Packet().Bytes())
}

func (this *ClientNet) Run() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", this.ServerAddr)
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		return
	}

	this.conn = conn // 保存链接
	this.Conn = conn
	go this.reader()
	fmt.Println("conn ok", conn, this.conn, this.RoleName)
}

func (this *ClientNet) Close() {
	this.conn.Close()
}

func (this *ClientNet) dealMsg(buf []byte) ([]byte, bool) {
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

func (this *ClientNet) reader() {
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
